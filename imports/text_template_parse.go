// this file was generated by gomacro command: import _b "text/template/parse"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"text/template/parse"
)

// reflection: allow interpreted code to import "text/template/parse"
func init() {
	Packages["text/template/parse"] = Package{
	Binds: map[string]Value{
		"IsEmptyTree":	ValueOf(parse.IsEmptyTree),
		"New":	ValueOf(parse.New),
		"NewIdentifier":	ValueOf(parse.NewIdentifier),
		"NodeAction":	ValueOf(parse.NodeAction),
		"NodeBool":	ValueOf(parse.NodeBool),
		"NodeChain":	ValueOf(parse.NodeChain),
		"NodeCommand":	ValueOf(parse.NodeCommand),
		"NodeDot":	ValueOf(parse.NodeDot),
		"NodeField":	ValueOf(parse.NodeField),
		"NodeIdentifier":	ValueOf(parse.NodeIdentifier),
		"NodeIf":	ValueOf(parse.NodeIf),
		"NodeList":	ValueOf(parse.NodeList),
		"NodeNil":	ValueOf(parse.NodeNil),
		"NodeNumber":	ValueOf(parse.NodeNumber),
		"NodePipe":	ValueOf(parse.NodePipe),
		"NodeRange":	ValueOf(parse.NodeRange),
		"NodeString":	ValueOf(parse.NodeString),
		"NodeTemplate":	ValueOf(parse.NodeTemplate),
		"NodeText":	ValueOf(parse.NodeText),
		"NodeVariable":	ValueOf(parse.NodeVariable),
		"NodeWith":	ValueOf(parse.NodeWith),
		"Parse":	ValueOf(parse.Parse),
	},
	Types: map[string]Type{
		"ActionNode":	TypeOf((*parse.ActionNode)(nil)).Elem(),
		"BoolNode":	TypeOf((*parse.BoolNode)(nil)).Elem(),
		"BranchNode":	TypeOf((*parse.BranchNode)(nil)).Elem(),
		"ChainNode":	TypeOf((*parse.ChainNode)(nil)).Elem(),
		"CommandNode":	TypeOf((*parse.CommandNode)(nil)).Elem(),
		"DotNode":	TypeOf((*parse.DotNode)(nil)).Elem(),
		"FieldNode":	TypeOf((*parse.FieldNode)(nil)).Elem(),
		"IdentifierNode":	TypeOf((*parse.IdentifierNode)(nil)).Elem(),
		"IfNode":	TypeOf((*parse.IfNode)(nil)).Elem(),
		"ListNode":	TypeOf((*parse.ListNode)(nil)).Elem(),
		"NilNode":	TypeOf((*parse.NilNode)(nil)).Elem(),
		"Node":	TypeOf((*parse.Node)(nil)).Elem(),
		"NodeType":	TypeOf((*parse.NodeType)(nil)).Elem(),
		"NumberNode":	TypeOf((*parse.NumberNode)(nil)).Elem(),
		"PipeNode":	TypeOf((*parse.PipeNode)(nil)).Elem(),
		"Pos":	TypeOf((*parse.Pos)(nil)).Elem(),
		"RangeNode":	TypeOf((*parse.RangeNode)(nil)).Elem(),
		"StringNode":	TypeOf((*parse.StringNode)(nil)).Elem(),
		"TemplateNode":	TypeOf((*parse.TemplateNode)(nil)).Elem(),
		"TextNode":	TypeOf((*parse.TextNode)(nil)).Elem(),
		"Tree":	TypeOf((*parse.Tree)(nil)).Elem(),
		"VariableNode":	TypeOf((*parse.VariableNode)(nil)).Elem(),
		"WithNode":	TypeOf((*parse.WithNode)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	},
	Untypeds: map[string]string{
	},
	Wrappers: map[string][]string{
		"ActionNode":	[]string{"Position","Type",},
		"BoolNode":	[]string{"Position","Type",},
		"BranchNode":	[]string{"Position","Type",},
		"ChainNode":	[]string{"Position","Type",},
		"CommandNode":	[]string{"Position","Type",},
		"DotNode":	[]string{"Position",},
		"FieldNode":	[]string{"Position","Type",},
		"IdentifierNode":	[]string{"Position","Type",},
		"IfNode":	[]string{"Position","String","Type",},
		"ListNode":	[]string{"Position","Type",},
		"NilNode":	[]string{"Position",},
		"NumberNode":	[]string{"Position","Type",},
		"PipeNode":	[]string{"Position","Type",},
		"RangeNode":	[]string{"Position","String","Type",},
		"StringNode":	[]string{"Position","Type",},
		"TemplateNode":	[]string{"Position","Type",},
		"TextNode":	[]string{"Position","Type",},
		"VariableNode":	[]string{"Position","Type",},
		"WithNode":	[]string{"Position","String","Type",},
	} }
}
