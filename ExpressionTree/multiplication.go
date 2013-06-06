package ExpressionTree

type Mutiplication struct {
	OperateNode
}

func (m Mutiplication) Caculate() int {
	return m.left.Caculate() * m.right.Caculate()
}
