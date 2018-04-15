// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/base"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package base

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/base"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/base"] = imports.Package{
		Binds: map[string]r.Value{
			"CMacroExpand":            r.ValueOf(CMacroExpand),
			"CMacroExpand1":           r.ValueOf(CMacroExpand1),
			"CMacroExpandCodewalk":    r.ValueOf(CMacroExpandCodewalk),
			"ConvertValue":            r.ValueOf(ConvertValue),
			"Debugf":                  r.ValueOf(Debugf),
			"DefaultImporter":         r.ValueOf(DefaultImporter),
			"DescendNestedUnquotes":   r.ValueOf(DescendNestedUnquotes),
			"DirName":                 r.ValueOf(DirName),
			"DuplicateNestedUnquotes": r.ValueOf(DuplicateNestedUnquotes),
			"Error":                   r.ValueOf(Error),
			"Errorf":                  r.ValueOf(Errorf),
			"False":                   r.ValueOf(&False).Elem(),
			"FileName":                r.ValueOf(FileName),
			"FindFirstToken":          r.ValueOf(FindFirstToken),
			"GoSrcDir":                r.ValueOf(&GoSrcDir).Elem(),
			"GomacroDir":              r.ValueOf(&GomacroDir).Elem(),
			"ImBuiltin":               r.ValueOf(ImBuiltin),
			"ImInception":             r.ValueOf(ImInception),
			"ImPlugin":                r.ValueOf(ImPlugin),
			"ImThirdParty":            r.ValueOf(ImThirdParty),
			"IsCategory":              r.ValueOf(IsCategory),
			"IsGensym":                r.ValueOf(IsGensym),
			"IsGensymEmbedded":        r.ValueOf(IsGensymEmbedded),
			"IsGensymInterface":       r.ValueOf(IsGensymInterface),
			"IsGensymPrivate":         r.ValueOf(IsGensymPrivate),
			"IsNillableKind":          r.ValueOf(IsNillableKind),
			"IsOptimizedKind":         r.ValueOf(IsOptimizedKind),
			"KindToCategory":          r.ValueOf(KindToCategory),
			"KindToType":              r.ValueOf(KindToType),
			"MakeBufReadline":         r.ValueOf(MakeBufReadline),
			"MakeQuote":               r.ValueOf(MakeQuote),
			"MakeQuote2":              r.ValueOf(MakeQuote2),
			"MakeTtyReadline":         r.ValueOf(MakeTtyReadline),
			"MarshalUntyped":          r.ValueOf(MarshalUntyped),
			"MaxInt":                  r.ValueOf(MaxInt),
			"MaxUint":                 r.ValueOf(MaxUint),
			"MaxUint16":               r.ValueOf(MaxUint16),
			"MinInt":                  r.ValueOf(MinInt),
			"MinUint":                 r.ValueOf(MinUint),
			"NewGlobals":              r.ValueOf(NewGlobals),
			"Nil":                     r.ValueOf(&Nil).Elem(),
			"None":                    r.ValueOf(&None).Elem(),
			"One":                     r.ValueOf(&One).Elem(),
			"OptCollectDeclarations":     r.ValueOf(OptCollectDeclarations),
			"OptCollectStatements":       r.ValueOf(OptCollectStatements),
			"OptDebugCallStack":          r.ValueOf(OptDebugCallStack),
			"OptDebugField":              r.ValueOf(OptDebugField),
			"OptDebugFromReflect":        r.ValueOf(OptDebugFromReflect),
			"OptDebugMacroExpand":        r.ValueOf(OptDebugMacroExpand),
			"OptDebugMethod":             r.ValueOf(OptDebugMethod),
			"OptDebugPanicRecover":       r.ValueOf(OptDebugPanicRecover),
			"OptDebugParse":              r.ValueOf(OptDebugParse),
			"OptDebugQuasiquote":         r.ValueOf(OptDebugQuasiquote),
			"OptDebugSleepOnSwitch":      r.ValueOf(OptDebugSleepOnSwitch),
			"OptFastInterpreter":         r.ValueOf(OptFastInterpreter),
			"OptMacroExpandOnly":         r.ValueOf(OptMacroExpandOnly),
			"OptPanicStackTrace":         r.ValueOf(OptPanicStackTrace),
			"OptShowCompile":             r.ValueOf(OptShowCompile),
			"OptShowEval":                r.ValueOf(OptShowEval),
			"OptShowEvalType":            r.ValueOf(OptShowEvalType),
			"OptShowMacroExpand":         r.ValueOf(OptShowMacroExpand),
			"OptShowParse":               r.ValueOf(OptShowParse),
			"OptShowPrompt":              r.ValueOf(OptShowPrompt),
			"OptShowTime":                r.ValueOf(OptShowTime),
			"OptTrapPanic":               r.ValueOf(OptTrapPanic),
			"PackValues":                 r.ValueOf(PackValues),
			"PackValuesAndTypes":         r.ValueOf(PackValuesAndTypes),
			"ParseOptions":               r.ValueOf(ParseOptions),
			"ReadBytes":                  r.ValueOf(ReadBytes),
			"ReadMultiline":              r.ValueOf(ReadMultiline),
			"ReadOptCollectAllComments":  r.ValueOf(ReadOptCollectAllComments),
			"ReadOptShowPrompt":          r.ValueOf(ReadOptShowPrompt),
			"ReadString":                 r.ValueOf(ReadString),
			"RemoveLastByte":             r.ValueOf(RemoveLastByte),
			"ShowPackageHeader":          r.ValueOf(ShowPackageHeader),
			"SigAll":                     r.ValueOf(SigAll),
			"SigDefer":                   r.ValueOf(SigDefer),
			"SigInterrupt":               r.ValueOf(SigInterrupt),
			"SigNone":                    r.ValueOf(SigNone),
			"SigReturn":                  r.ValueOf(SigReturn),
			"SimplifyAstForQuote":        r.ValueOf(SimplifyAstForQuote),
			"SimplifyNodeForQuote":       r.ValueOf(SimplifyNodeForQuote),
			"StartSignalHandler":         r.ValueOf(StartSignalHandler),
			"StopSignalHandler":          r.ValueOf(StopSignalHandler),
			"StrGensym":                  r.ValueOf(StrGensym),
			"StrGensymEmbedded":          r.ValueOf(StrGensymEmbedded),
			"StrGensymInterface":         r.ValueOf(StrGensymInterface),
			"StrGensymPrivate":           r.ValueOf(StrGensymPrivate),
			"Subdir":                     r.ValueOf(Subdir),
			"True":                       r.ValueOf(&True).Elem(),
			"TypeOfBool":                 r.ValueOf(&TypeOfBool).Elem(),
			"TypeOfByte":                 r.ValueOf(&TypeOfByte).Elem(),
			"TypeOfComplex128":           r.ValueOf(&TypeOfComplex128).Elem(),
			"TypeOfComplex64":            r.ValueOf(&TypeOfComplex64).Elem(),
			"TypeOfDeferFunc":            r.ValueOf(&TypeOfDeferFunc).Elem(),
			"TypeOfError":                r.ValueOf(&TypeOfError).Elem(),
			"TypeOfFloat32":              r.ValueOf(&TypeOfFloat32).Elem(),
			"TypeOfFloat64":              r.ValueOf(&TypeOfFloat64).Elem(),
			"TypeOfInt":                  r.ValueOf(&TypeOfInt).Elem(),
			"TypeOfInt16":                r.ValueOf(&TypeOfInt16).Elem(),
			"TypeOfInt32":                r.ValueOf(&TypeOfInt32).Elem(),
			"TypeOfInt64":                r.ValueOf(&TypeOfInt64).Elem(),
			"TypeOfInt8":                 r.ValueOf(&TypeOfInt8).Elem(),
			"TypeOfInterface":            r.ValueOf(&TypeOfInterface).Elem(),
			"TypeOfPtrBool":              r.ValueOf(&TypeOfPtrBool).Elem(),
			"TypeOfPtrComplex128":        r.ValueOf(&TypeOfPtrComplex128).Elem(),
			"TypeOfPtrComplex64":         r.ValueOf(&TypeOfPtrComplex64).Elem(),
			"TypeOfPtrFloat32":           r.ValueOf(&TypeOfPtrFloat32).Elem(),
			"TypeOfPtrFloat64":           r.ValueOf(&TypeOfPtrFloat64).Elem(),
			"TypeOfPtrInt":               r.ValueOf(&TypeOfPtrInt).Elem(),
			"TypeOfPtrInt16":             r.ValueOf(&TypeOfPtrInt16).Elem(),
			"TypeOfPtrInt32":             r.ValueOf(&TypeOfPtrInt32).Elem(),
			"TypeOfPtrInt64":             r.ValueOf(&TypeOfPtrInt64).Elem(),
			"TypeOfPtrInt8":              r.ValueOf(&TypeOfPtrInt8).Elem(),
			"TypeOfPtrString":            r.ValueOf(&TypeOfPtrString).Elem(),
			"TypeOfPtrUint":              r.ValueOf(&TypeOfPtrUint).Elem(),
			"TypeOfPtrUint16":            r.ValueOf(&TypeOfPtrUint16).Elem(),
			"TypeOfPtrUint32":            r.ValueOf(&TypeOfPtrUint32).Elem(),
			"TypeOfPtrUint64":            r.ValueOf(&TypeOfPtrUint64).Elem(),
			"TypeOfPtrUint8":             r.ValueOf(&TypeOfPtrUint8).Elem(),
			"TypeOfPtrUintptr":           r.ValueOf(&TypeOfPtrUintptr).Elem(),
			"TypeOfReflectType":          r.ValueOf(&TypeOfReflectType).Elem(),
			"TypeOfRune":                 r.ValueOf(&TypeOfRune).Elem(),
			"TypeOfString":               r.ValueOf(&TypeOfString).Elem(),
			"TypeOfUint":                 r.ValueOf(&TypeOfUint).Elem(),
			"TypeOfUint16":               r.ValueOf(&TypeOfUint16).Elem(),
			"TypeOfUint32":               r.ValueOf(&TypeOfUint32).Elem(),
			"TypeOfUint64":               r.ValueOf(&TypeOfUint64).Elem(),
			"TypeOfUint8":                r.ValueOf(&TypeOfUint8).Elem(),
			"TypeOfUintptr":              r.ValueOf(&TypeOfUintptr).Elem(),
			"UnescapeChar":               r.ValueOf(UnescapeChar),
			"UnescapeString":             r.ValueOf(UnescapeString),
			"UnmarshalUntyped":           r.ValueOf(UnmarshalUntyped),
			"UnmarshalUntypedVal":        r.ValueOf(UnmarshalUntypedVal),
			"UnpackValues":               r.ValueOf(UnpackValues),
			"UntypedKindToReflectKind":   r.ValueOf(UntypedKindToReflectKind),
			"UnwrapTrivialAst":           r.ValueOf(UnwrapTrivialAst),
			"UnwrapTrivialAstKeepBlocks": r.ValueOf(UnwrapTrivialAstKeepBlocks),
			"UnwrapTrivialNode":          r.ValueOf(UnwrapTrivialNode),
			"UserHomeDir":                r.ValueOf(UserHomeDir),
			"ValueInterface":             r.ValueOf(ValueInterface),
			"ValueType":                  r.ValueOf(ValueType),
			"Warnf":                      r.ValueOf(Warnf),
			"ZeroStrings":                r.ValueOf(&ZeroStrings).Elem(),
			"ZeroTypes":                  r.ValueOf(&ZeroTypes).Elem(),
			"ZeroValues":                 r.ValueOf(&ZeroValues).Elem(),
		},
		Types: map[string]r.Type{
			"BufReadline":      r.TypeOf((*BufReadline)(nil)).Elem(),
			"Globals":          r.TypeOf((*Globals)(nil)).Elem(),
			"ImportMode":       r.TypeOf((*ImportMode)(nil)).Elem(),
			"Importer":         r.TypeOf((*Importer)(nil)).Elem(),
			"Options":          r.TypeOf((*Options)(nil)).Elem(),
			"Output":           r.TypeOf((*Output)(nil)).Elem(),
			"PackageRef":       r.TypeOf((*PackageRef)(nil)).Elem(),
			"ReadOptions":      r.TypeOf((*ReadOptions)(nil)).Elem(),
			"Readline":         r.TypeOf((*Readline)(nil)).Elem(),
			"RuntimeError":     r.TypeOf((*RuntimeError)(nil)).Elem(),
			"Signal":           r.TypeOf((*Signal)(nil)).Elem(),
			"Signals":          r.TypeOf((*Signals)(nil)).Elem(),
			"Stringer":         r.TypeOf((*Stringer)(nil)).Elem(),
			"TtyReadline":      r.TypeOf((*TtyReadline)(nil)).Elem(),
			"TypeVisitor":      r.TypeOf((*TypeVisitor)(nil)).Elem(),
			"UntypedVal":       r.TypeOf((*UntypedVal)(nil)).Elem(),
			"WhichMacroExpand": r.TypeOf((*WhichMacroExpand)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{
			"Readline": r.TypeOf((*P_github_com_cosmos72_gomacro_base_Readline)(nil)).Elem(),
		},
		Wrappers: map[string][]string{
			"Globals":    []string{"CollectPackageImportsWithRename", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "IncLine", "IncLineBytes", "MakeRuntimeError", "Position", "Sprintf", "ToString", "WarnExtraValues", "Warnf"},
			"Output":     []string{"Copy", "Errorf", "Fprintf", "IncLine", "IncLineBytes", "MakeRuntimeError", "Position", "Sprintf", "ToString"},
			"PackageRef": []string{"LazyInit", "Merge"},
		},
	}
}

// --------------- proxy for github.com/cosmos72/gomacro/base.Readline ---------------
type P_github_com_cosmos72_gomacro_base_Readline struct {
	Object interface{}
	Read_  func(_proxy_obj_ interface{}, prompt string) ([]byte, error)
}

func (P *P_github_com_cosmos72_gomacro_base_Readline) Read(prompt string) ([]byte, error) {
	return P.Read_(P.Object, prompt)
}
