package ExpressionTree

type Division struct {
	OperateNode
}

func (d *Division) isZeroOfRigtht() bool {
	return d.right.Caculate() == 0
}

func (d Division) Caculate() int {
	if d.isZeroOfRigtht() {
		panic("div zero")
	}

	return d.left.Caculate() / d.right.Caculate()
}
