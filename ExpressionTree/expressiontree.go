package ExpressionTree

import (
	"expression-tree/stack"
	"fmt"
	"regexp"
	"strconv"
)

type Node interface {
	Caculate() int
	getName() string
}

type NumberNode struct {
	num int
}

type BinaryTree interface {
	setLeft(node Node)
	setRight(node Node)
	PrefixExpression() string
	SuffixExpression() string
	NifixExpression() string
}

type OperateNode struct {
	operate string
	left    Node
	right   Node
}

func (on OperateNode) getName() string {
	return on.operate + " "
}

func (on *OperateNode) setLeft(node Node) {
	on.left = node
}

func (on *OperateNode) setRight(node Node) {
	on.right = node
}

func (on OperateNode) PrefixExpression() string {
	var leftExpression string
	var rightExpression string
	if leftNode, ok := on.left.(BinaryTree); ok {
		leftExpression = leftNode.PrefixExpression()
	} else {
		leftExpression = on.left.getName()
	}

	if rightNode, ok := on.right.(BinaryTree); ok {
		rightExpression = rightNode.PrefixExpression()
	} else {
		rightExpression = on.right.getName()
	}
	return fmt.Sprint(on.getName(), leftExpression, rightExpression)
}

func (on *OperateNode) SuffixExpression() string {
	var leftExpression string
	var rightExpression string
	if leftNode, ok := on.left.(BinaryTree); ok {
		leftExpression = leftNode.SuffixExpression()
	} else {
		leftExpression = on.left.getName()
	}

	if rightNode, ok := on.right.(BinaryTree); ok {
		rightExpression = rightNode.SuffixExpression()
	} else {
		rightExpression = on.right.getName()
	}

	return fmt.Sprint(leftExpression, rightExpression, on.getName())
}

func (on OperateNode) NifixExpression() string {
	var leftExpression string
	var rightExpression string
	if leftNode, ok := on.left.(BinaryTree); ok {
		leftExpression = leftNode.NifixExpression()
	} else {
		leftExpression = on.left.getName()
	}

	if rightNode, ok := on.right.(BinaryTree); ok {
		rightExpression = rightNode.NifixExpression()
	} else {
		rightExpression = on.right.getName()
	}

	return fmt.Sprint(leftExpression, on.getName(), rightExpression)
}

func (nn NumberNode) Caculate() int {
	return nn.num
}

func (nn NumberNode) getName() string {
	name := strconv.Itoa(nn.num)
	return name + " "
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
