// this file was generated by gomacro command: import _b "go/scanner"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"go/scanner"
)

// reflection: allow interpreted code to import "go/scanner"
func init() {
	Packages["go/scanner"] = Package{
	Binds: map[string]Value{
		"PrintError":	ValueOf(scanner.PrintError),
		"ScanComments":	ValueOf(scanner.ScanComments),
	},
	Types: map[string]Type{
		"Error":	TypeOf((*scanner.Error)(nil)).Elem(),
		"ErrorHandler":	TypeOf((*scanner.ErrorHandler)(nil)).Elem(),
		"ErrorList":	TypeOf((*scanner.ErrorList)(nil)).Elem(),
		"Mode":	TypeOf((*scanner.Mode)(nil)).Elem(),
		"Scanner":	TypeOf((*scanner.Scanner)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Untypeds: map[string]string{
	},
	Wrappers: map[string][]string{
	} }
}
