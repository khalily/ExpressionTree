package ExpressionTree

type Subtraction struct {
	OperateNode
}

func (s Subtraction) Caculate() int {
	return s.left.Caculate() - s.right.Caculate()
}
