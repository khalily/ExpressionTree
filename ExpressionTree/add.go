package ExpressionTree

import (
// "fmt"
)

type Add struct {
	OperateNode
}

func (a Add) Caculate() int {
	return a.left.Caculate() + a.right.Caculate()
}
