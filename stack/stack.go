package stack

type Element struct {
	Value interface{}
}

type Stack struct {
	Elements []Element
	l        int
}

func NewStack() *Stack {
	return &Stack{make([]Element, 0), 0}
}

func (st *Stack) Push(v interface{}) {
	st.Elements = append(st.Elements, Element{v})
	st.l++
}

func (st Stack) Len() int {
	return st.l
}

func (st Stack) IsEmpty() bool {
	if st.l == 0 {
		return true
	}
	return false
}

func (st Stack) Top() Element {
	e := st.Elements[st.l-1]
	return e
}

func (st *Stack) Pop() Element {
	e := st.Top()
	st.Elements = st.Elements[:st.l-1]
	st.l = st.l - 1
	return e
}
