// this file was generated by gomacro command: import _b "compress/zlib"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"compress/zlib"
	"io"
)

// reflection: allow interpreted code to import "compress/zlib"
func init() {
	Packages["compress/zlib"] = Package{
	Binds: map[string]Value{
		"BestCompression":	ValueOf(zlib.BestCompression),
		"BestSpeed":	ValueOf(zlib.BestSpeed),
		"DefaultCompression":	ValueOf(zlib.DefaultCompression),
		"ErrChecksum":	ValueOf(&zlib.ErrChecksum).Elem(),
		"ErrDictionary":	ValueOf(&zlib.ErrDictionary).Elem(),
		"ErrHeader":	ValueOf(&zlib.ErrHeader).Elem(),
		"HuffmanOnly":	ValueOf(zlib.HuffmanOnly),
		"NewReader":	ValueOf(zlib.NewReader),
		"NewReaderDict":	ValueOf(zlib.NewReaderDict),
		"NewWriter":	ValueOf(zlib.NewWriter),
		"NewWriterLevel":	ValueOf(zlib.NewWriterLevel),
		"NewWriterLevelDict":	ValueOf(zlib.NewWriterLevelDict),
		"NoCompression":	ValueOf(zlib.NoCompression),
	},
	Types: map[string]Type{
		"Resetter":	TypeOf((*zlib.Resetter)(nil)).Elem(),
		"Writer":	TypeOf((*zlib.Writer)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"Resetter":	TypeOf((*Resetter_compress_zlib)(nil)).Elem(),
	},
	Untypeds: map[string]string{
		"BestCompression":	"int:9",
		"BestSpeed":	"int:1",
		"DefaultCompression":	"int:-1",
		"HuffmanOnly":	"int:-2",
		"NoCompression":	"int:0",
	},
	Wrappers: map[string][]string{
	} }
}

// --------------- proxy for compress/zlib.Resetter ---------------
type Resetter_compress_zlib struct {
	Object	interface{}
	Reset_	func(r io.Reader, dict []byte) error
}
func (Proxy *Resetter_compress_zlib) Reset(r io.Reader, dict []byte) error {
	return Proxy.Reset_(r, dict)
}
