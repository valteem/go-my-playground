package main

import (
    "fmt"
)

type Stack []string

func (s *Stack) isEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) push(value string) {
    *s = append(*s, value)
}

func (s *Stack) pop() (string, bool) {
    if s.isEmpty() {
        return "", false
    } else {
        elementIndex := len(*s) - 1
        element := (*s)[elementIndex]
        *s = (*s)[:elementIndex]
        return element, true
    }
}

func isOperator(character string) bool {
    switch character {
    case "+", "-", "*", "/":
        return true
    default:
        return false
    }

}

func input() {
    var stack Stack
    fmt.Print("Please input the equation without spaces: \n")
    input := "ABC/-AK/L-*"


    for _, character := range input {
        valueCheck := isOperator(string(character))
        if valueCheck {
            operand1 := stack[len(stack)-1]
            stack.pop()
            operand2 := stack[len(stack)-1]
            stack.pop()


            temp := string(character) + string(operand2) + string(operand1)
            stack.push(temp)

        } else {
            stack.push(string(character))
        }
    }

    fmt.Print(stack)

}

func main() {
    input()
}
