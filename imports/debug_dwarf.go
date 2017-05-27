// this file was generated by gomacro command: import _b "debug/dwarf"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"debug/dwarf"
)

// reflection: allow interpreted code to import "debug/dwarf"
func init() {
	Packages["debug/dwarf"] = Package{
	Binds: map[string]Value{
		"AttrAbstractOrigin":	ValueOf(dwarf.AttrAbstractOrigin),
		"AttrAccessibility":	ValueOf(dwarf.AttrAccessibility),
		"AttrAddrClass":	ValueOf(dwarf.AttrAddrClass),
		"AttrAllocated":	ValueOf(dwarf.AttrAllocated),
		"AttrArtificial":	ValueOf(dwarf.AttrArtificial),
		"AttrAssociated":	ValueOf(dwarf.AttrAssociated),
		"AttrBaseTypes":	ValueOf(dwarf.AttrBaseTypes),
		"AttrBitOffset":	ValueOf(dwarf.AttrBitOffset),
		"AttrBitSize":	ValueOf(dwarf.AttrBitSize),
		"AttrByteSize":	ValueOf(dwarf.AttrByteSize),
		"AttrCallColumn":	ValueOf(dwarf.AttrCallColumn),
		"AttrCallFile":	ValueOf(dwarf.AttrCallFile),
		"AttrCallLine":	ValueOf(dwarf.AttrCallLine),
		"AttrCalling":	ValueOf(dwarf.AttrCalling),
		"AttrCommonRef":	ValueOf(dwarf.AttrCommonRef),
		"AttrCompDir":	ValueOf(dwarf.AttrCompDir),
		"AttrConstValue":	ValueOf(dwarf.AttrConstValue),
		"AttrContainingType":	ValueOf(dwarf.AttrContainingType),
		"AttrCount":	ValueOf(dwarf.AttrCount),
		"AttrDataLocation":	ValueOf(dwarf.AttrDataLocation),
		"AttrDataMemberLoc":	ValueOf(dwarf.AttrDataMemberLoc),
		"AttrDeclColumn":	ValueOf(dwarf.AttrDeclColumn),
		"AttrDeclFile":	ValueOf(dwarf.AttrDeclFile),
		"AttrDeclLine":	ValueOf(dwarf.AttrDeclLine),
		"AttrDeclaration":	ValueOf(dwarf.AttrDeclaration),
		"AttrDefaultValue":	ValueOf(dwarf.AttrDefaultValue),
		"AttrDescription":	ValueOf(dwarf.AttrDescription),
		"AttrDiscr":	ValueOf(dwarf.AttrDiscr),
		"AttrDiscrList":	ValueOf(dwarf.AttrDiscrList),
		"AttrDiscrValue":	ValueOf(dwarf.AttrDiscrValue),
		"AttrEncoding":	ValueOf(dwarf.AttrEncoding),
		"AttrEntrypc":	ValueOf(dwarf.AttrEntrypc),
		"AttrExtension":	ValueOf(dwarf.AttrExtension),
		"AttrExternal":	ValueOf(dwarf.AttrExternal),
		"AttrFrameBase":	ValueOf(dwarf.AttrFrameBase),
		"AttrFriend":	ValueOf(dwarf.AttrFriend),
		"AttrHighpc":	ValueOf(dwarf.AttrHighpc),
		"AttrIdentifierCase":	ValueOf(dwarf.AttrIdentifierCase),
		"AttrImport":	ValueOf(dwarf.AttrImport),
		"AttrInline":	ValueOf(dwarf.AttrInline),
		"AttrIsOptional":	ValueOf(dwarf.AttrIsOptional),
		"AttrLanguage":	ValueOf(dwarf.AttrLanguage),
		"AttrLocation":	ValueOf(dwarf.AttrLocation),
		"AttrLowerBound":	ValueOf(dwarf.AttrLowerBound),
		"AttrLowpc":	ValueOf(dwarf.AttrLowpc),
		"AttrMacroInfo":	ValueOf(dwarf.AttrMacroInfo),
		"AttrName":	ValueOf(dwarf.AttrName),
		"AttrNamelistItem":	ValueOf(dwarf.AttrNamelistItem),
		"AttrOrdering":	ValueOf(dwarf.AttrOrdering),
		"AttrPriority":	ValueOf(dwarf.AttrPriority),
		"AttrProducer":	ValueOf(dwarf.AttrProducer),
		"AttrPrototyped":	ValueOf(dwarf.AttrPrototyped),
		"AttrRanges":	ValueOf(dwarf.AttrRanges),
		"AttrReturnAddr":	ValueOf(dwarf.AttrReturnAddr),
		"AttrSegment":	ValueOf(dwarf.AttrSegment),
		"AttrSibling":	ValueOf(dwarf.AttrSibling),
		"AttrSpecification":	ValueOf(dwarf.AttrSpecification),
		"AttrStartScope":	ValueOf(dwarf.AttrStartScope),
		"AttrStaticLink":	ValueOf(dwarf.AttrStaticLink),
		"AttrStmtList":	ValueOf(dwarf.AttrStmtList),
		"AttrStride":	ValueOf(dwarf.AttrStride),
		"AttrStrideSize":	ValueOf(dwarf.AttrStrideSize),
		"AttrStringLength":	ValueOf(dwarf.AttrStringLength),
		"AttrTrampoline":	ValueOf(dwarf.AttrTrampoline),
		"AttrType":	ValueOf(dwarf.AttrType),
		"AttrUpperBound":	ValueOf(dwarf.AttrUpperBound),
		"AttrUseLocation":	ValueOf(dwarf.AttrUseLocation),
		"AttrUseUTF8":	ValueOf(dwarf.AttrUseUTF8),
		"AttrVarParam":	ValueOf(dwarf.AttrVarParam),
		"AttrVirtuality":	ValueOf(dwarf.AttrVirtuality),
		"AttrVisibility":	ValueOf(dwarf.AttrVisibility),
		"AttrVtableElemLoc":	ValueOf(dwarf.AttrVtableElemLoc),
		"ClassAddress":	ValueOf(dwarf.ClassAddress),
		"ClassBlock":	ValueOf(dwarf.ClassBlock),
		"ClassConstant":	ValueOf(dwarf.ClassConstant),
		"ClassExprLoc":	ValueOf(dwarf.ClassExprLoc),
		"ClassFlag":	ValueOf(dwarf.ClassFlag),
		"ClassLinePtr":	ValueOf(dwarf.ClassLinePtr),
		"ClassLocListPtr":	ValueOf(dwarf.ClassLocListPtr),
		"ClassMacPtr":	ValueOf(dwarf.ClassMacPtr),
		"ClassRangeListPtr":	ValueOf(dwarf.ClassRangeListPtr),
		"ClassReference":	ValueOf(dwarf.ClassReference),
		"ClassReferenceAlt":	ValueOf(dwarf.ClassReferenceAlt),
		"ClassReferenceSig":	ValueOf(dwarf.ClassReferenceSig),
		"ClassString":	ValueOf(dwarf.ClassString),
		"ClassStringAlt":	ValueOf(dwarf.ClassStringAlt),
		"ClassUnknown":	ValueOf(dwarf.ClassUnknown),
		"ErrUnknownPC":	ValueOf(&dwarf.ErrUnknownPC).Elem(),
		"New":	ValueOf(dwarf.New),
		"TagAccessDeclaration":	ValueOf(dwarf.TagAccessDeclaration),
		"TagArrayType":	ValueOf(dwarf.TagArrayType),
		"TagBaseType":	ValueOf(dwarf.TagBaseType),
		"TagCatchDwarfBlock":	ValueOf(dwarf.TagCatchDwarfBlock),
		"TagClassType":	ValueOf(dwarf.TagClassType),
		"TagCommonDwarfBlock":	ValueOf(dwarf.TagCommonDwarfBlock),
		"TagCommonInclusion":	ValueOf(dwarf.TagCommonInclusion),
		"TagCompileUnit":	ValueOf(dwarf.TagCompileUnit),
		"TagCondition":	ValueOf(dwarf.TagCondition),
		"TagConstType":	ValueOf(dwarf.TagConstType),
		"TagConstant":	ValueOf(dwarf.TagConstant),
		"TagDwarfProcedure":	ValueOf(dwarf.TagDwarfProcedure),
		"TagEntryPoint":	ValueOf(dwarf.TagEntryPoint),
		"TagEnumerationType":	ValueOf(dwarf.TagEnumerationType),
		"TagEnumerator":	ValueOf(dwarf.TagEnumerator),
		"TagFileType":	ValueOf(dwarf.TagFileType),
		"TagFormalParameter":	ValueOf(dwarf.TagFormalParameter),
		"TagFriend":	ValueOf(dwarf.TagFriend),
		"TagImportedDeclaration":	ValueOf(dwarf.TagImportedDeclaration),
		"TagImportedModule":	ValueOf(dwarf.TagImportedModule),
		"TagImportedUnit":	ValueOf(dwarf.TagImportedUnit),
		"TagInheritance":	ValueOf(dwarf.TagInheritance),
		"TagInlinedSubroutine":	ValueOf(dwarf.TagInlinedSubroutine),
		"TagInterfaceType":	ValueOf(dwarf.TagInterfaceType),
		"TagLabel":	ValueOf(dwarf.TagLabel),
		"TagLexDwarfBlock":	ValueOf(dwarf.TagLexDwarfBlock),
		"TagMember":	ValueOf(dwarf.TagMember),
		"TagModule":	ValueOf(dwarf.TagModule),
		"TagMutableType":	ValueOf(dwarf.TagMutableType),
		"TagNamelist":	ValueOf(dwarf.TagNamelist),
		"TagNamelistItem":	ValueOf(dwarf.TagNamelistItem),
		"TagNamespace":	ValueOf(dwarf.TagNamespace),
		"TagPackedType":	ValueOf(dwarf.TagPackedType),
		"TagPartialUnit":	ValueOf(dwarf.TagPartialUnit),
		"TagPointerType":	ValueOf(dwarf.TagPointerType),
		"TagPtrToMemberType":	ValueOf(dwarf.TagPtrToMemberType),
		"TagReferenceType":	ValueOf(dwarf.TagReferenceType),
		"TagRestrictType":	ValueOf(dwarf.TagRestrictType),
		"TagRvalueReferenceType":	ValueOf(dwarf.TagRvalueReferenceType),
		"TagSetType":	ValueOf(dwarf.TagSetType),
		"TagSharedType":	ValueOf(dwarf.TagSharedType),
		"TagStringType":	ValueOf(dwarf.TagStringType),
		"TagStructType":	ValueOf(dwarf.TagStructType),
		"TagSubprogram":	ValueOf(dwarf.TagSubprogram),
		"TagSubrangeType":	ValueOf(dwarf.TagSubrangeType),
		"TagSubroutineType":	ValueOf(dwarf.TagSubroutineType),
		"TagTemplateAlias":	ValueOf(dwarf.TagTemplateAlias),
		"TagTemplateTypeParameter":	ValueOf(dwarf.TagTemplateTypeParameter),
		"TagTemplateValueParameter":	ValueOf(dwarf.TagTemplateValueParameter),
		"TagThrownType":	ValueOf(dwarf.TagThrownType),
		"TagTryDwarfBlock":	ValueOf(dwarf.TagTryDwarfBlock),
		"TagTypeUnit":	ValueOf(dwarf.TagTypeUnit),
		"TagTypedef":	ValueOf(dwarf.TagTypedef),
		"TagUnionType":	ValueOf(dwarf.TagUnionType),
		"TagUnspecifiedParameters":	ValueOf(dwarf.TagUnspecifiedParameters),
		"TagUnspecifiedType":	ValueOf(dwarf.TagUnspecifiedType),
		"TagVariable":	ValueOf(dwarf.TagVariable),
		"TagVariant":	ValueOf(dwarf.TagVariant),
		"TagVariantPart":	ValueOf(dwarf.TagVariantPart),
		"TagVolatileType":	ValueOf(dwarf.TagVolatileType),
		"TagWithStmt":	ValueOf(dwarf.TagWithStmt),
	},
	Types: map[string]Type{
		"AddrType":	TypeOf((*dwarf.AddrType)(nil)).Elem(),
		"ArrayType":	TypeOf((*dwarf.ArrayType)(nil)).Elem(),
		"Attr":	TypeOf((*dwarf.Attr)(nil)).Elem(),
		"BasicType":	TypeOf((*dwarf.BasicType)(nil)).Elem(),
		"BoolType":	TypeOf((*dwarf.BoolType)(nil)).Elem(),
		"CharType":	TypeOf((*dwarf.CharType)(nil)).Elem(),
		"Class":	TypeOf((*dwarf.Class)(nil)).Elem(),
		"CommonType":	TypeOf((*dwarf.CommonType)(nil)).Elem(),
		"ComplexType":	TypeOf((*dwarf.ComplexType)(nil)).Elem(),
		"Data":	TypeOf((*dwarf.Data)(nil)).Elem(),
		"DecodeError":	TypeOf((*dwarf.DecodeError)(nil)).Elem(),
		"DotDotDotType":	TypeOf((*dwarf.DotDotDotType)(nil)).Elem(),
		"Entry":	TypeOf((*dwarf.Entry)(nil)).Elem(),
		"EnumType":	TypeOf((*dwarf.EnumType)(nil)).Elem(),
		"EnumValue":	TypeOf((*dwarf.EnumValue)(nil)).Elem(),
		"Field":	TypeOf((*dwarf.Field)(nil)).Elem(),
		"FloatType":	TypeOf((*dwarf.FloatType)(nil)).Elem(),
		"FuncType":	TypeOf((*dwarf.FuncType)(nil)).Elem(),
		"IntType":	TypeOf((*dwarf.IntType)(nil)).Elem(),
		"LineEntry":	TypeOf((*dwarf.LineEntry)(nil)).Elem(),
		"LineFile":	TypeOf((*dwarf.LineFile)(nil)).Elem(),
		"LineReader":	TypeOf((*dwarf.LineReader)(nil)).Elem(),
		"LineReaderPos":	TypeOf((*dwarf.LineReaderPos)(nil)).Elem(),
		"Offset":	TypeOf((*dwarf.Offset)(nil)).Elem(),
		"PtrType":	TypeOf((*dwarf.PtrType)(nil)).Elem(),
		"QualType":	TypeOf((*dwarf.QualType)(nil)).Elem(),
		"Reader":	TypeOf((*dwarf.Reader)(nil)).Elem(),
		"StructField":	TypeOf((*dwarf.StructField)(nil)).Elem(),
		"StructType":	TypeOf((*dwarf.StructType)(nil)).Elem(),
		"Tag":	TypeOf((*dwarf.Tag)(nil)).Elem(),
		"Type":	TypeOf((*dwarf.Type)(nil)).Elem(),
		"TypedefType":	TypeOf((*dwarf.TypedefType)(nil)).Elem(),
		"UcharType":	TypeOf((*dwarf.UcharType)(nil)).Elem(),
		"UintType":	TypeOf((*dwarf.UintType)(nil)).Elem(),
		"UnspecifiedType":	TypeOf((*dwarf.UnspecifiedType)(nil)).Elem(),
		"VoidType":	TypeOf((*dwarf.VoidType)(nil)).Elem(),
	},
	Proxies: map[string]Type{
		"Type":	TypeOf((*Type_debug_dwarf)(nil)).Elem(),
	},
	Untypeds: map[string]string{
	},
	Wrappers: map[string][]string{
		"AddrType":	[]string{"Basic","Common","Size","String",},
		"ArrayType":	[]string{"Common",},
		"BasicType":	[]string{"Common","Size",},
		"BoolType":	[]string{"Basic","Common","Size","String",},
		"CharType":	[]string{"Basic","Common","Size","String",},
		"ComplexType":	[]string{"Basic","Common","Size","String",},
		"DotDotDotType":	[]string{"Common","Size",},
		"EnumType":	[]string{"Common","Size",},
		"FloatType":	[]string{"Basic","Common","Size","String",},
		"FuncType":	[]string{"Common","Size",},
		"IntType":	[]string{"Basic","Common","Size","String",},
		"PtrType":	[]string{"Common","Size",},
		"QualType":	[]string{"Common",},
		"StructType":	[]string{"Common","Size",},
		"TypedefType":	[]string{"Common",},
		"UcharType":	[]string{"Basic","Common","Size","String",},
		"UintType":	[]string{"Basic","Common","Size","String",},
		"UnspecifiedType":	[]string{"Basic","Common","Size","String",},
		"VoidType":	[]string{"Common","Size",},
	} }
}

// --------------- proxy for debug/dwarf.Type ---------------
type Type_debug_dwarf struct {
	Object	interface{}
	Common_	func() *dwarf.CommonType
	Size_	func() int64
	String_	func() string
}
func (Proxy *Type_debug_dwarf) Common() *dwarf.CommonType {
	return Proxy.Common_()
}
func (Proxy *Type_debug_dwarf) Size() int64 {
	return Proxy.Size_()
}
func (Proxy *Type_debug_dwarf) String() string {
	return Proxy.String_()
}
