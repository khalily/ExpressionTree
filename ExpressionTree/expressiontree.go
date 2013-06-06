package ExpressionTree

import (
	"expression-tree/stack"
	// "fmt"
	"regexp"
	"strconv"
)

type Node interface {
	Caculate() int
}

type NumberNode struct {
	num int
}

type BinaryTree interface {
	setLeft(node Node)
	setRight(node Node)
}

type OperateNode struct {
	operate string
	left    Node
	right   Node
}

func (o *OperateNode) setLeft(node Node) {
	o.left = node
}

func (o *OperateNode) setRight(node Node) {
	o.right = node
}

func (nn NumberNode) Caculate() int {
	return nn.num
}

func isOperate(s string) bool {
	re, _ := regexp.Compile("[+*/\\-]{1}")
	ok := re.MatchString(s)
	// fmt.Println(ok, s)
	return ok
}

func CreateOperateNode(v string) (node BinaryTree) {
	switch v {
	case "+":
		return &Add{OperateNode{"+", nil, nil}}
	case "-":
		return &Subtraction{OperateNode{"-", nil, nil}}
	case "*":
		return &Mutiplication{OperateNode{"*", nil, nil}}
	case "/":
		return &Division{OperateNode{"/", nil, nil}}
	default:
		panic("no support")
	}
	return nil
}

func CreateNumberNode(v int) *NumberNode {
	return &NumberNode{v}
}

func CreateExpressionTree(exps []string) Node {
	st := stack.NewStack()
	for _, s := range exps {
		if isOperate(s) {
			operateNode := CreateOperateNode(s)
			right, _ := st.Pop().Value.(Node)
			left, _ := st.Pop().Value.(Node)
			operateNode.setRight(right)
			operateNode.setLeft(left)
			// fmt.Println(operateNode)
			st.Push(operateNode)
		} else {
			num, _ := strconv.Atoi(s)
			numberNode := CreateNumberNode(num)
			st.Push(numberNode)
		}
	}
	if st.Len() != 1 {
		panic("CreateExpressionTree error")
	}
	node, _ := st.Pop().Value.(Node)
	// fmt.Println(node)
	return node
}
