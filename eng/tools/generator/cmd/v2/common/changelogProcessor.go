// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/autorest/model"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/repo"
	"github.com/Azure/azure-sdk-for-go/eng/tools/internal/delta"
	"github.com/Azure/azure-sdk-for-go/eng/tools/internal/exports"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

const (
	sdk_tag_fetch_url = "https://api.github.com/repos/Azure/azure-sdk-for-go/git/refs/tags"
	sdk_remote_url    = "https://github.com/Azure/azure-sdk-for-go.git"
)

func GetAllVersionTags(rpName, namespaceName string) ([]string, error) {
	log.Printf("Fetching all release tags from GitHub for RP: '%s' Package: '%s' ...", rpName, namespaceName)
	client := http.Client{}
	res, err := client.Get(sdk_tag_fetch_url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := []map[string]interface{}{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	var tags []string
	for _, tag := range result {
		tagName := tag["ref"].(string)
		if strings.Contains(tagName, "sdk/resourcemanager/"+rpName+"/"+namespaceName) {
			tags = append(tags, tag["ref"].(string))
		}
	}
	sort.Sort(releaseTagsSort(tags))

	return tags, nil
}

type releaseTagsSort []string

func (t releaseTagsSort) Len() int {
	return len(t)
}

func (t releaseTagsSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t releaseTagsSort) Less(i, j int) bool {
	if t[i] > t[j] {
		return !strings.Contains(t[i], t[j])
	} else {
		return strings.Contains(t[j], t[i])
	}
}

func ContainsPreviewAPIVersion(packagePath string) (bool, error) {
	log.Printf("Judge whether contains preview API version from '%s' ...", packagePath)

	files, err := ioutil.ReadDir(packagePath)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			b, err := ioutil.ReadFile(path.Join(packagePath, file.Name()))
			if err != nil {
				return false, err
			}

			lines := strings.Split(string(b), "\n")
			for _, line := range lines {
				if strings.Contains(line, "\"api-version\"") {
					parts := strings.Split(line, "\"")
					if len(parts) == 5 && strings.Contains(parts[3], "preview") {
						return true, nil
					}
				}
			}
		}
	}

	return false, nil
}

func GetPreviousVersionTag(isCurrentPreview bool, allReleases []string) string {
	if isCurrentPreview {
		// for preview api, always compare with latest release
		return allReleases[0]
	} else {
		// for stable api, always compare with previous stable, if no stable, then latest release
		for _, release := range allReleases {
			if !strings.Contains(release, "beta") {
				return release
			}
		}
		return allReleases[0]
	}
}

func GetExportsFromTag(sdkRepo repo.SDKRepository, packagePath, tag string) (*exports.Content, error) {
	log.Printf("Get exports from specific tag '%s' ...", tag)

	// get current head branch name
	currentRef, err := sdkRepo.Head()
	if err != nil {
		return nil, err
	}

	// add package change
	err = sdkRepo.Add(packagePath)
	if err != nil {
		return nil, err
	}

	// stash current change
	err = sdkRepo.Stash()
	if err != nil {
		return nil, err
	}

	// create remote with center sdk repo
	remoteName := "release_remote"
	_, err = sdkRepo.CreateRemote(&config.RemoteConfig{Name: remoteName, URLs: []string{sdk_remote_url}})
	if err != nil {
		if err != git.ErrRemoteExists {
			return nil, err
		}
	}

	// fetch tag from remote
	err = sdkRepo.Fetch(&git.FetchOptions{RemoteName: remoteName, RefSpecs: []config.RefSpec{config.RefSpec(tag + ":" + tag)}})
	if err != nil {
		if err.Error() != "already up-to-date" {
			return nil, err
		}
	}

	// checkout to the specific tag
	err = sdkRepo.CheckoutTag(strings.TrimPrefix(tag, "ref/tags/"))
	if err != nil {
		return nil, err
	}

	// get exports
	result, err := exports.Get(packagePath)
	if err != nil {
		return nil, err
	}

	// checkout back to head branch
	err = sdkRepo.Checkout(&repo.CheckoutOptions{
		Branch: plumbing.ReferenceName(currentRef.Name()),
	})
	if err != nil {
		return nil, err
	}

	// remove remote
	err = sdkRepo.DeleteRemote(remoteName)
	if err != nil {
		return nil, err
	}

	// restore current change
	err = sdkRepo.StashPop()
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func MarshalUnmarshalFilter(changelog *model.Changelog) {
	if changelog.Modified != nil {
		if changelog.Modified.AdditiveChanges != nil {
			removeMarshalUnmarshalFunc(changelog.Modified.AdditiveChanges.Funcs)
		}
		if changelog.Modified.BreakingChanges != nil && changelog.Modified.BreakingChanges.Removed != nil {
			removeMarshalUnmarshalFunc(changelog.Modified.BreakingChanges.Removed.Funcs)
		}
	}
}

func removeMarshalUnmarshalFunc(funcs map[string]exports.Func) {
	for k := range funcs {
		if strings.HasSuffix(k, ".MarshalJSON") || strings.HasSuffix(k, ".UnmarshalJSON") {
			delete(funcs, k)
		}
	}
}

func FilterChangelog(changelog *model.Changelog, opts ...func(changelog *model.Changelog)) {
	if changelog.Modified != nil {
		for _, opt := range opts {
			opt(changelog)
		}
	}
}

func EnumFilter(changelog *model.Changelog) {
	if changelog.Modified.HasAdditiveChanges() {
		if changelog.Modified.AdditiveChanges != nil && changelog.Modified.AdditiveChanges.TypeAliases != nil {
			for typeAliases := range changelog.Modified.AdditiveChanges.TypeAliases {
				funcKeys, funcExist := searchKey(changelog.Modified.AdditiveChanges.Funcs, typeAliases, "Possible")
				if funcExist && len(funcKeys) == 1 {
					for _, f := range funcKeys {
						delete(changelog.Modified.AdditiveChanges.Funcs, f)
					}
				}
			}
		}
	}

	if changelog.Modified.HasBreakingChanges() {
		enumOperation(changelog.Modified.BreakingChanges.Removed)
	}
}

func enumOperation(content *delta.Content) {
	if content != nil && content.TypeAliases != nil {
		for typeAliases := range content.TypeAliases {
			constKeys, constExist := searchKey(content.Consts, typeAliases, "")
			funcKeys, funcExist := searchKey(content.Funcs, typeAliases, "Possible")

			if constExist && funcExist && len(funcKeys) == 1 {
				for _, c := range constKeys {
					delete(content.Consts, c)
				}
				for _, f := range funcKeys {
					delete(content.Funcs, f)
				}
			}
		}
	}
}

func searchKey[T exports.Const | exports.Func | exports.Struct](m map[string]T, key1, prefix string) ([]string, bool) {
	keys := make([]string, 0)
	for k := range m {
		if regexp.MustCompile(fmt.Sprintf("^%s%s\\w*", prefix, key1)).MatchString(k) {
			keys = append(keys, k)
		}
	}
	if len(keys) != 0 {
		return keys, true
	}
	return nil, false
}

func FuncFilter(changelog *model.Changelog) {
	if changelog.Modified.HasAdditiveChanges() {
		funcOperation(changelog.Modified.AdditiveChanges)
	}

	if changelog.Modified.HasBreakingChanges() {
		funcOperation(changelog.Modified.BreakingChanges.Removed)
	}
}

func funcOperation(content *delta.Content) {
	if content != nil && content.Funcs != nil {
		for funcName, funcValue := range content.Funcs {
			clientFunc := strings.Split(funcName, ".")
			if len(clientFunc) == 2 {
				// the last parameter
				if funcValue.Params != nil {
					ps := strings.Split(*funcValue.Params, ",")
					clientFuncOptions := ps[len(ps)-1]
					clientFuncOptions = strings.TrimLeft(strings.TrimSpace(clientFuncOptions), "*")
					if clientFuncOptions != "" && content.CompleteStructs != nil {
						delete(content.Structs, clientFuncOptions)
						for i, v := range content.CompleteStructs {
							if v == clientFuncOptions {
								content.CompleteStructs = append(content.CompleteStructs[:i],
									content.CompleteStructs[i+1:]...)
								break
							}
						}
					}
				}

				// the first return value
				if funcValue.Returns != nil {
					rs := strings.Split(*funcValue.Returns, ",")
					clientFuncResponse := rs[0]
					if strings.Contains(clientFunc[1], "runtime.Poller") {
						re := regexp.MustCompile("\\[(?P<response>.*)\\]")
						clientFuncResponse = re.FindString(clientFuncResponse)
						clientFuncResponse = re.ReplaceAllString(clientFuncResponse, "${response}")
					} else {
						clientFuncResponse = strings.TrimLeft(clientFuncResponse, "*")
					}
					if clientFuncResponse != "" && content.CompleteStructs != nil {
						delete(content.Structs, clientFuncResponse)
						for i, v := range content.CompleteStructs {
							if v == clientFuncResponse {
								content.CompleteStructs = append(content.CompleteStructs[:i],
									content.CompleteStructs[i+1:]...)
								break
							}
						}
					}
				}
			}
		}
	}
}

// LROFilter LROFilter after OperationFilter
func LROFilter(changelog *model.Changelog) {
	if changelog.Modified.HasBreakingChanges() && changelog.Modified.HasAdditiveChanges() && changelog.Modified.BreakingChanges.Removed != nil && changelog.Modified.BreakingChanges.Removed.Funcs != nil {
		removedContent := changelog.Modified.BreakingChanges.Removed
		for bFunc, v := range removedContent.Funcs {
			var beginFunc string
			clientFunc := strings.Split(bFunc, ".")
			if len(clientFunc) == 2 {
				if strings.Contains(clientFunc[1], "Begin") {
					clientFunc[1] = strings.ReplaceAll(clientFunc[1], "Being", "")
					beginFunc = fmt.Sprintf("%s.%s", clientFunc[0], clientFunc[1])
				} else {
					beginFunc = fmt.Sprintf("%s.Begin%s", clientFunc[0], clientFunc[1])
				}
				if _, ok := changelog.Modified.AdditiveChanges.Funcs[beginFunc]; ok {
					delete(changelog.Modified.AdditiveChanges.Funcs, beginFunc)
					v.ReplacedBy = &beginFunc
					removedContent.Funcs[bFunc] = v
				}
			}
		}
	}
}
