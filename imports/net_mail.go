// this file was generated by gomacro command: import _b "net/mail"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/mail"
)

// reflection: allow interpreted code to import "net/mail"
func init() {
	Packages["net/mail"] = Package{
	Binds: map[string]Value{
		"ErrHeaderNotPresent":	ValueOf(&mail.ErrHeaderNotPresent).Elem(),
		"ParseAddress":	ValueOf(mail.ParseAddress),
		"ParseAddressList":	ValueOf(mail.ParseAddressList),
		"ParseDate":	ValueOf(mail.ParseDate),
		"ReadMessage":	ValueOf(mail.ReadMessage),
	},
	Types: map[string]Type{
		"Address":	TypeOf((*mail.Address)(nil)).Elem(),
		"AddressParser":	TypeOf((*mail.AddressParser)(nil)).Elem(),
		"Header":	TypeOf((*mail.Header)(nil)).Elem(),
		"Message":	TypeOf((*mail.Message)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Untypeds: map[string]string{
	},
	Wrappers: map[string][]string{
	} }
}
