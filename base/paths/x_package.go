// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/base/paths"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package paths

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/base/paths"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/base/paths"] = imports.Package{
		Binds: map[string]r.Value{
			"DirName":                  r.ValueOf(DirName),
			"FileName":                 r.ValueOf(FileName),
			"GetImportsSrcDir":         r.ValueOf(GetImportsSrcDir),
			"GoSrcDir":                 r.ValueOf(&GoSrcDir).Elem(),
			"RemoveFinalSlash":         r.ValueOf(RemoveFinalSlash),
			"Subdir":                   r.ValueOf(Subdir),
			"SymbolFromImportsPackage": r.ValueOf(&SymbolFromImportsPackage).Elem(),
			"UserHomeDir":              r.ValueOf(UserHomeDir),
		},
	}
}
