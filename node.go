package jcheck

import "fmt"

type valueType int

const (
	typeInvalid valueType = iota
	typeString
	typeNumber
	typeObject
	typeArray
	typeTrue
	typeFalse
	typeNull
)

func (t valueType) String() string {
	switch t {
	case typeString:
		return "String"
	case typeNumber:
		return "Number"
	case typeObject:
		return "Object"
	case typeArray:
		return "Array"
	case typeTrue:
		return "True"
	case typeFalse:
		return "False"
	case typeNull:
		return "Null"
	default:
		return ""
	}
}

type node struct {
	path     string
	typ      valueType
	str      string
	num      float64
	arrayLen int

	hasPermitRule bool

	parent   *node
	children []*node
}

func (n *node) String() string {
	switch n.typ {
	case typeString:
		return fmt.Sprintf("%s=%q", n.path, n.str)
	case typeNumber:
		return fmt.Sprintf("%s=%v", n.path, n.num)
	case typeObject, typeNull:
		return fmt.Sprintf("%s=(%s)", n.path, n.typ.String())
	case typeArray:
		return fmt.Sprintf("%s=(%s[%d])", n.path, n.typ.String(), n.arrayLen)
	case typeTrue, typeFalse:
		return fmt.Sprintf("%s=%s", n.path, n.typ.String())
	default:
		return ""
	}
}

func (n *node) forEachNode(f func(*node)) {
	f(n)
	if n.children != nil {
		for _, c := range n.children {
			c.forEachNode(f)
		}
	}
}

func (n *node) forEachNodeReverse(f func(*node)) {
	if n.children != nil {
		for _, c := range n.children {
			c.forEachNode(f)
		}
	}
	f(n)
}

func scanNodeMap(parent *node, path string, obj map[string]interface{}) []*node {
	var nodes []*node
	for k, v := range obj {
		var p string

		if path != "" {
			p = fmt.Sprintf("%s.%s", path, k)
		} else {
			p = k
		}

		nodes = append(nodes, scanNodeValue(parent, p, v))
	}
	return nodes
}

func scanNodeArray(parent *node, path string, obj []interface{}) []*node {
	var nodes []*node
	for i, v := range obj {
		nodes = append(nodes, scanNodeValue(parent, fmt.Sprintf("%s.%d", path, i), v))
	}
	return nodes
}

func scanNodeValue(parent *node, path string, value interface{}) *node {
	switch v := value.(type) {
	case string:
		return &node{
			path:   path,
			typ:    typeString,
			str:    v,
			parent: parent,
		}
	case float64:
		return &node{
			path:   path,
			typ:    typeNumber,
			num:    v,
			parent: parent,
		}
	case bool:
		var t valueType
		if v {
			t = typeTrue
		} else {
			t = typeFalse
		}
		return &node{
			path:   path,
			typ:    t,
			parent: parent,
		}
	case nil:
		return &node{
			path:   path,
			typ:    typeNull,
			parent: parent,
		}
	case map[string]interface{}:
		n := &node{
			path:   path,
			typ:    typeObject,
			parent: parent,
		}
		n.children = scanNodeMap(n, path, v)
		return n
	case []interface{}:
		n := &node{
			path:     path,
			typ:      typeArray,
			arrayLen: len(v),
			parent:   parent,
		}
		n.children = scanNodeArray(n, path, v)
		return n
	}
	panic("oops")
}
