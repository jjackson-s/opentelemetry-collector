// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configschema

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"reflect"
	"strings"

	"golang.org/x/mod/modfile"
)

// DefaultSrcRoot is the default root of the collector repo, relative to the
// current working directory. Can be used to create a DirResolver.
const DefaultSrcRoot = "."

// DefaultModule is the module prefix of otelcol. Can be used to create a
// DirResolver.
const DefaultModule = "go.opentelemetry.io/collector"

// DirResolver is used to resolve the base directory of a given reflect.Type.
type DirResolver struct {
	SrcRoot    string
	ModuleName string
}

// NewDefaultDirResolver creates a DirResolver with a default SrcRoot and
// ModuleName, suitable for using this package's API using otelcol with an
// executable running from the otelcol's source root (not tests).
func NewDefaultDirResolver() DirResolver {
	return NewDirResolver(DefaultSrcRoot, DefaultModule)
}

// NewDirResolver creates a DirResolver with a custom SrcRoot and ModuleName.
// Useful for testing and for using this package's API from a repository other
// than otelcol (e.g. contrib).
func NewDirResolver(srcRoot string, moduleName string) DirResolver {
	return DirResolver{
		SrcRoot:    srcRoot,
		ModuleName: moduleName,
	}
}

// PackageDir accepts a Type and returns its package dir.
// If package is not found locally, searches for package location through go.mod.
// Attempts to search through an absolute path or GOPATH
func (dr DirResolver) PackageDir(t reflect.Type) string {
	pkg := strings.TrimPrefix(t.PkgPath(), dr.ModuleName+"/")
	dir := localPackageDir(dr, pkg)
	_, err := os.ReadDir(dir)
	if err != nil {
		if dir, err = externalPackageDir(dr, pkg); err != nil {
			panic(fmt.Sprintf("Could not find the pkg: %s", pkg))
		}
	}
	return dir
}

// localPackageDir returns the path to a local package.
func localPackageDir(dr DirResolver, pkg string) string {
	return path.Join(dr.SrcRoot, pkg)
}

// externalPackageDir returns the path to an external package.
func externalPackageDir(dr DirResolver, pkg string) (string, error) {
	line, version, err := grepMod(path.Join(dr.SrcRoot, "go.mod"), pkg)
	if err != nil {
		return "", err
	}
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	pkgPath := buildExternalPath(dr, goPath, pkg, line, version)
	if _, err = os.ReadDir(pkgPath); err != nil {
		return "", err
	}
	return pkgPath, nil
}

// grepMod returns the line within go.mod associated with the package
// we are looking for.
func grepMod(goModPath string, pkg string) (string, string, error) {
	file, err := os.ReadFile(goModPath)
	if err != nil {
		return "", "", err
	}
	modContents, err := modfile.Parse(goModPath, file, nil)
	if err != nil {
		return "", "", err
	}
	var defaultModulePath, defaultModuleVersion string
	for _, line := range modContents.Replace {
		switch {
		case strings.Compare(line.Old.Path, DefaultModule) == 0:
			defaultModulePath, defaultModuleVersion = line.New.Path, line.New.Version
		case strings.Contains(line.Old.Path, pkg):
			return line.New.Path, line.New.Version, nil
		}
	}
	for _, line := range modContents.Require {
		switch {
		case defaultModulePath == "" && strings.Compare(line.Mod.Path, DefaultModule) == 0:
			defaultModulePath, defaultModuleVersion = line.Mod.Path, line.Mod.Version
		case strings.Contains(line.Mod.Path, pkg):
			return line.Mod.Path, line.Mod.Version, nil
		}
	}
	isDefaultModPkg := strings.Contains(pkg, DefaultModule)
	if isDefaultModPkg && defaultModulePath != "" {
		return defaultModulePath, defaultModuleVersion, nil
	}
	return "", "", nil
}

// buildExternalPath builds a path to a package that is not local to directory.
func buildExternalPath(dr DirResolver, goPath, pkg, line, v string) string {
	switch {
	case strings.HasPrefix(line, "./"):
		return path.Join(dr.SrcRoot, line)
	case strings.HasPrefix(line, "../"):
		return path.Join(dr.SrcRoot, line, strings.TrimPrefix(pkg, DefaultModule))
	default:
		return path.Join(goPath, "pkg", "mod", line+"@"+v)
	}
}
