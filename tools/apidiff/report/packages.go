// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package report

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/tools/apidiff/markdown"
)

const apiDirSuffix = "api"

// GetPackages returns all the go sdk packages under the given root directory
func GetPackages(dir string) ([]string, error) {
	var pkgDirs []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// check if leaf dir
			fi, err := ioutil.ReadDir(path)
			if err != nil {
				return err
			}
			hasSubDirs := false
			for _, f := range fi {
				// check if this is the interfaces subdir, if it is don't count it as a subdir
				if f.IsDir() && f.Name() != filepath.Base(path)+apiDirSuffix {
					hasSubDirs = true
					break
				}
			}
			if !hasSubDirs {
				pkgDirs = append(pkgDirs, path)
				// skip any dirs under us (i.e. interfaces subdir)
				return filepath.SkipDir
			}
		}
		return nil
	})
	return pkgDirs, err
}

// PkgsList contains a collection of packages
type PkgsList []string

// ModifiedPackages contains a collection of package reports, it's structured as "package path":pkgReport
type ModifiedPackages map[string]Package

// IsEmpty ...
func (m ModifiedPackages) IsEmpty() bool {
	return len(m) == 0
}

// HasBreakingChanges returns true if any package contained in has a breaking change
func (m ModifiedPackages) HasBreakingChanges() bool {
	for _, p := range m {
		if p.HasBreakingChanges() {
			return true
		}
	}
	return false
}

// HasAdditiveChanges returns true if any package contained in has an additive change
func (m ModifiedPackages) HasAdditiveChanges() bool {
	for _, p := range m {
		if p.HasAdditiveChanges() {
			return true
		}
	}
	return false
}

// CommitPkgsReport represents a collection of reports, one for each commit hash.
type CommitPkgsReport struct {
	// AffectedPackages stores the package list with key of commit hashes
	AffectedPackages map[string]PkgsList `json:"affectedPackages"`
	// BreakingChanges stores the commit hashes that contain breaking changes
	BreakingChanges []string `json:"breakingChanges,omitempty"`
	// CommitsReports stores the detailed reports with the key of commit hashes
	CommitsReports map[string]PkgsReport `json:"deltas"`
}

// IsEmpty returns true if the report contains no data.
func (c CommitPkgsReport) IsEmpty() bool {
	for _, r := range c.CommitsReports {
		if !r.IsEmpty() {
			return false
		}
	}
	return true
}

// HasBreakingChanges returns true if the report contains breaking changes.
func (c CommitPkgsReport) HasBreakingChanges() bool {
	for _, r := range c.CommitsReports {
		if r.HasBreakingChanges() {
			return true
		}
	}
	return false
}

// HasAdditiveChanges returns true if the package contains additive changes.
func (c CommitPkgsReport) HasAdditiveChanges() bool {
	for _, r := range c.CommitsReports {
		if r.HasAdditiveChanges() {
			return true
		}
	}
	return false
}

// UpdateAffectedPackages updates the collection of affected packages with the packages that were touched in the specified commit
func (c *CommitPkgsReport) UpdateAffectedPackages(commit string, r PkgsReport) {
	if c.AffectedPackages == nil {
		c.AffectedPackages = map[string]PkgsList{}
	}

	for _, pkg := range r.AddedPackages {
		c.AffectedPackages[commit] = append(c.AffectedPackages[commit], pkg)
	}

	for pkgName := range r.ModifiedPackages {
		c.AffectedPackages[commit] = append(c.AffectedPackages[commit], pkgName)
	}

	for _, pkg := range r.RemovedPackages {
		c.AffectedPackages[commit] = append(c.AffectedPackages[commit], pkg)
	}
}

// PkgsReport represents a complete report of added, removed, and modified packages
type PkgsReport struct {
	// AddedPackages stores the added packages in the report
	AddedPackages PkgsList `json:"added,omitempty"`
	// ModifiedPackages stores the details of all modified packages
	ModifiedPackages ModifiedPackages `json:"modified,omitempty"`
	// RemovedPackages stores the removed packages in the report
	RemovedPackages PkgsList `json:"removed,omitempty"`
}

// HasBreakingChanges returns true if the package report contains breaking changes
func (r PkgsReport) HasBreakingChanges() bool {
	return len(r.RemovedPackages) > 0 || (r.ModifiedPackages != nil && r.ModifiedPackages.HasBreakingChanges())
}

// HasAdditiveChanges returns true if the package report contains additive changes
func (r PkgsReport) HasAdditiveChanges() bool {
	return len(r.AddedPackages) > 0 || (r.ModifiedPackages != nil && r.ModifiedPackages.HasAdditiveChanges())
}

// IsEmpty returns true if the report contains no data
func (r PkgsReport) IsEmpty() bool {
	return len(r.AddedPackages) == 0 && len(r.ModifiedPackages) == 0 && len(r.RemovedPackages) == 0
}

// ToMarkdown writes the report to string in the markdown form
func (r *PkgsReport) ToMarkdown(withHeader bool, version, marker string) string {
	md := markdown.Writer{}
	if withHeader {
		r.writeHeader(&md, version, marker)
	}
	if r.IsEmpty() {
		return ""
	}
	r.writeAddedPackages(&md)
	r.writeRemovedPackages(&md)
	r.writeModifiedPackages(&md)
	return md.String()
}

func (r *PkgsReport) writeHeader(md *markdown.Writer, version, marker string) {
	md.WriteTitle("Release History")
	md.WriteTopLevelHeader(fmt.Sprintf("%s (%s)", version, marker))
}

func (r *PkgsReport) writeAddedPackages(md *markdown.Writer) {
	if len(r.AddedPackages) == 0 {
		return
	}
	md.WriteHeader("New Packages")
	// write list
	for _, p := range r.AddedPackages {
		md.WriteLine("- " + p)
	}
}

func (r *PkgsReport) writeModifiedPackages(md *markdown.Writer) {
	if len(r.ModifiedPackages) == 0 {
		return
	}
	md.WriteHeader("Modified Packages")
	// write list
	for n := range r.ModifiedPackages {
		md.WriteLine(fmt.Sprintf("- `%s`", n))
	}
	// write details
	for n, p := range r.ModifiedPackages {
		md.WriteHeader(fmt.Sprintf("Modified `%s`", n))
		md.WriteLine(p.ToMarkdown())
	}
}

func (r *PkgsReport) writeRemovedPackages(md *markdown.Writer) {
	if len(r.RemovedPackages) == 0 {
		return
	}
	md.WriteHeader("Removed Packages")
	// write list
	for _, p := range r.RemovedPackages {
		md.WriteLine("- " + p)
	}
	md.EmptyLine()
}
