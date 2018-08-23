package main

import (
	"log"
	"strconv"
	"strings"
)

func ExprEvaluate(expression string) int {
	var lastNumber bool
	ops := NewStack()
	vals := NewStack()

	for _, c := range expression {
		char := strings.Trim(string(c), " ")

		if char == "" {
			lastNumber = false
			continue
		} else if char == "(" {
			lastNumber = false
			continue
		} else if char == "+" {
			lastNumber = false
			ops.Push(char)
		} else if char == "-" {
			lastNumber = false
			ops.Push(char)
		} else if char == "*" {
			lastNumber = false
			ops.Push(char)
		} else if char == "/" {
			lastNumber = false
			ops.Push(char)
		} else if char == ")" {
			lastNumber = false
			var val int
			if v, err := vals.Pop(); err != nil {
				log.Fatal("values stack is empty")
			} else {
				val, _ = strconv.Atoi(v.(string))
			}
			var op string
			if o, err := ops.Pop(); err != nil {
				log.Fatal("operators stack is empty")
			} else {
				op = o.(string)
			}

			if op == "+" {
				if v, err := vals.Pop(); err != nil {
					log.Fatal("values stack is empty")
				} else {
					v2, _ := strconv.Atoi(v.(string))
					val = val + v2
				}
			} else if op == "-" {
				if v, err := vals.Pop(); err != nil {
					log.Fatal("values stack is empty")
				} else {
					v2, _ := strconv.Atoi(v.(string))
					val = val - v2
				}
			} else if op == "*" {
				if v, err := vals.Pop(); err != nil {
					log.Fatal("values stack is empty")
				} else {
					v2, _ := strconv.Atoi(v.(string))
					val = val * v2
				}
			} else if op == "/" {
				if v, err := vals.Pop(); err != nil {
					log.Fatal("values stack is empty")
				} else {
					v2, _ := strconv.Atoi(v.(string))
					val = val / v2
				}
			}

			vals.Push(strconv.Itoa(val))
		} else {
			if lastNumber {
				var val int
				if v, err := vals.Pop(); err != nil {
					log.Fatal("values stack is empty")
				} else {
					v2, _ := strconv.Atoi(v.(string))
					v3, _ := strconv.Atoi(char)
					val = v2*10 + v3
				}
				vals.Push(strconv.Itoa(val))
			} else {
				vals.Push(char)
			}
			lastNumber = true
		}
	}

	if v, err := vals.Pop(); err == nil {
		result, _ := strconv.Atoi(v.(string))
		return result
	}
	return -1
}
