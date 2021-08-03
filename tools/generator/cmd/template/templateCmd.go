// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/tools/generator/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Command returns the template command
func Command() *cobra.Command {
	templateCmd := &cobra.Command{
		Use:   "template (<rpName> <packageName>) | <packagePath>",
		Short: "Onboards new RP with the template",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var rpName, packageName string
			if len(args) == 1 {
				segments := strings.Split(args[0], "/")
				if len(segments) != 2 {
					return fmt.Errorf("%s is not a valid package path. Please assign the package path using `rpName/packageName` format", args[0])
				}
				rpName = segments[0]
				packageName = segments[1]
			} else {
				rpName = args[0]
				packageName = args[1]
			}

			return GeneratePackageByTemplate(rpName, packageName, ParseFlags(cmd.Flags()))
		},
	}

	BindFlags(templateCmd.Flags())
	templateCmd.MarkFlagRequired("package-title")

	return templateCmd
}

// BindFlags binds the flags to this command
func BindFlags(flagSet *pflag.FlagSet) {
	flagSet.String("go-sdk-folder", ".", "Specifies the path of root of azure-sdk-for-go")
	flagSet.String("template-path", "tools/generator/template/rpName/packageName", "Specifies the path of the template")
	flagSet.String("package-title", "", "Specifies the title of this package")
}

// ParseFlags parses the flags to a Flags struct
func ParseFlags(flagSet *pflag.FlagSet) Flags {
	return Flags{
		SDKRoot:      flags.GetString(flagSet, "go-sdk-folder"),
		TemplatePath: flags.GetString(flagSet, "template-path"),
		PackageTitle: flags.GetString(flagSet, "package-title"),
	}
}

// Flags ...
type Flags struct {
	SDKRoot      string
	TemplatePath string
	PackageTitle string
}

func GeneratePackageByTemplate(rpName, packageName string, flags Flags) error {
	root, err := filepath.Abs(flags.SDKRoot)
	if err != nil {
		return fmt.Errorf("cannot get the root of azure-sdk-for-go from '%s': %+v", flags.SDKRoot, err)
	}
	var absTemplateDir string
	if filepath.IsAbs(flags.TemplatePath) {
		absTemplateDir = flags.TemplatePath
	} else {
		absTemplateDir = filepath.Join(root, flags.TemplatePath)
	}
	fileList, err := ioutil.ReadDir(absTemplateDir)
	if err != nil {
		return fmt.Errorf("cannot read the directory '%s': %+v", absTemplateDir, err)
	}

	// build the replaceMap
	buildReplaceMap(rpName, packageName, flags.PackageTitle)

	// copy everything to destination directory
	for _, file := range fileList {
		path := filepath.Join(absTemplateDir, file.Name())
		content, err := readAndReplace(path)
		if err != nil {
			return err
		}

		dirPath := filepath.Join(root, "sdk", rpName, packageName)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory '%s': %+v", dirPath, err)
		}

		newFilePath := filepath.Join(dirPath, strings.TrimSuffix(file.Name(), FilenameSuffix))
		if err := createNewFile(newFilePath, content); err != nil {
			return err
		}
	}

	return nil
}

func buildReplaceMap(rpName, packageName, packageTitle string) {
	replaceMap = make(map[string]string)

	replaceMap[RPNameKey] = rpName
	replaceMap[PackageNameKey] = packageName
	replaceMap[PackageTitleKey] = packageTitle
}

func readAndReplace(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read from file '%s': %+v", path, err)
	}

	content := string(b)
	for k, v := range replaceMap {
		content = strings.ReplaceAll(content, k, v)
	}

	return content, nil
}

func createNewFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot create file '%s': %+v", path, err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("cannot write to file '%s': %+v", path, err)
	}

	return nil
}

var (
	replaceMap map[string]string
)

const (
	RPNameKey       = "{{rpName}}"
	PackageNameKey  = "{{packageName}}"
	PackageTitleKey = "{{PackageTitle}}"
	FilenameSuffix  = ".tpl"
)
