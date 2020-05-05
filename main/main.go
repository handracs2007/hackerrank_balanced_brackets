// https://www.hackerrank.com/challenges/balanced-brackets/problem

package main

import (
	"errors"
	"fmt"
)

type RuneStack struct {
	elements []rune
}

func NewRuneStack() *RuneStack {
	var runeStack = &RuneStack{
		elements: make([]rune, 0),
	}

	return runeStack
}

func (runeStack *RuneStack) push(element rune) {
	runeStack.elements = append(runeStack.elements, element)
}

func (runeStack *RuneStack) pop() (rune, error) {
	var length = len(runeStack.elements)

	if length == 0 {
		return -1, errors.New("stack is empty")
	}

	var element = runeStack.elements[length-1]
	runeStack.elements = runeStack.elements[:length-1]

	return element, nil
}

func (runeStack *RuneStack) isEmpty() bool {
	return len(runeStack.elements) == 0
}

func isBalanced(s string) string {
	var runeStack = NewRuneStack()

	for _, bracket := range s {
		switch bracket {
		case '{', '[', '(':
			runeStack.push(bracket)

		case '}', ']', ')':
			var openBracket, err = runeStack.pop()

			if err == nil {
				if (bracket == '}' && openBracket != '{') ||
					(bracket == ']' && openBracket != '[') ||
					(bracket == ')' && openBracket != '(') {
					return "NO"
				}
			} else {
				return "NO"
			}
		}
	}

	if runeStack.isEmpty() {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(isBalanced("{[()]}"))
	fmt.Println(isBalanced("{[(])}"))
	fmt.Println(isBalanced("{{[[(())]]}}"))
}
