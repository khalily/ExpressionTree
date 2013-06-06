package main

import (
	"expression-tree/ExpressionTree"
	"expression-tree/stack"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isOperate(s string) bool {
	re, _ := regexp.Compile("[+*/\\-\\(\\)]{1}")
	ok := re.MatchString(s)
	// fmt.Println(ok, s)
	return ok
}

func isPop(st *stack.Stack, s string) (op []string, ok bool) {
	switch string(s) {
	case "(":
		ok = false
		return
	case ")":
		ok = true
		var cur stack.Element
		for !st.IsEmpty() {
			cur = st.Top()
			if curValue, ok2 := cur.Value.(string); ok2 {
				if curValue == "(" {
					st.Pop()
					return
				}
				op = append(op, curValue)
				st.Pop()
			}
		}
	default:
		var cur stack.Element
		for !st.IsEmpty() {
			cur = st.Top()
			if curValue, ok2 := cur.Value.(string); ok2 {
				if level, ok3 := po[curValue]; ok3 && level >= po[s] {
					ok = true
					op = append(op, curValue)
					// fmt.Println(op)
					st.Pop()
				} else {
					// fmt.Println(curValue, op)
					if len(op) != 0 {
						ok = true
					} else {
						ok = false
					}
					return
				}
			}
		}
	}
	return
}

func pre2stuf(exps []string) (exps2 []string) {
	st1 := stack.NewStack()
	st2 := stack.NewStack()

	for _, exp := range exps {
		if isOperate(exp) {
			if op, ok := isPop(st1, exp); ok {
				for _, s := range op {
					st2.Push(s)
				}
			}
			if exp == ")" {
				continue
			}
			st1.Push(exp)
		} else {
			st2.Push(exp)
		}
	}

	// fmt.Print(cur.Value)
	for !st1.IsEmpty() {
		st2.Push(st1.Pop().Value)
	}

	for _, e := range st2.Elements {
		s, _ := e.Value.(string)
		exps2 = append(exps2, s)
	}

	// fmt.Println(exps2)
	return
}

func parseExp(s string) (exps []string, err error) {
	re, err := regexp.Compile("[0-9]+|[+*/\\-\\(\\)]{1}")
	if err != nil {
		fmt.Println("regexp compile error:", err)
		return
	}
	for _, exp := range re.FindAll([]byte(s), -1) {
		exps = append(exps, string(exp))
	}
	return
}

const (
	LEVEL1 = iota
	LEVEL2
	LEVEL3
)

var po = make(map[string]int)

func main() {
	po["+"] = LEVEL1
	po["-"] = LEVEL1
	po["*"] = LEVEL2
	po["/"] = LEVEL2

	a := "1 + 23 * 1 - 12 * (4 + 2 * 5)"

	a = strings.Replace(a, " ", "", -1)
	exps, err := parseExp(a)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Input Expression:\t", exps)
	exps2 := pre2stuf(exps)
	// fmt.Println(exps2)

	root := ExpressionTree.CreateExpressionTree(exps2)
	fmt.Println("Caculate Result:\t", root.Caculate())

	if btree, ok := root.(ExpressionTree.BinaryTree); ok {
		fmt.Println("Prefix Expression:\t", btree.PrefixExpression())
		fmt.Println("Nifix Expression:\t", btree.NifixExpression())
		fmt.Println("Suffix Expression:\t", btree.SuffixExpression())
	}
}
