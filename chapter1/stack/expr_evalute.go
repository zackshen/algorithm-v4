package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getValueFromStack(stack *Stack) (float64, error) {
	var value interface{}
	value, err := stack.Pop()
	if err != nil {
		return 0, errors.New("values stack is empty")
	}
	val, err := strconv.ParseFloat(value.(string), 32)
	if err != nil {
		return 0, fmt.Errorf("%v parse float fail", value)
	}
	return val, nil
}

func getOpFromStack(stack *Stack) (string, error) {
	value, err := stack.Pop()
	if err != nil {
		return "", errors.New("operators stack is empty")
	}
	op, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("%v inference type failed", value)
	}
	return op, nil
}

func ExprEvaluate(expression string) float64 {
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
			val, err := getValueFromStack(vals)
			if err != nil {
				log.Panic(err)
			}
			op, err := getOpFromStack(ops)
			if err != nil {
				log.Panic(err)
			}

			if op == "+" {
				val2, err := getValueFromStack(vals)
				if err != nil {
					log.Panic(err)
				}
				val = val + val2
			} else if op == "-" {
				val2, err := getValueFromStack(vals)
				if err != nil {
					log.Panic(err)
				}
				val = val2 - val
			} else if op == "*" {
				val2, err := getValueFromStack(vals)
				if err != nil {
					log.Panic(err)
				}
				val = val2 * val
			} else if op == "/" {
				val2, err := getValueFromStack(vals)
				if err != nil {
					log.Panic(err)
				}
				if val == 0 {
					log.Panic("divide zero is illegal")
				}
				val = val2 / val
			}

			vals.Push(strconv.FormatFloat(val, 'f', 6, 64))
		} else {
			if lastNumber {
				val, err := getValueFromStack(vals)
				if err != nil {
					log.Panic(err)
				}

				val2, err := strconv.ParseFloat(char, 64)
				if err != nil {
					log.Panic(err)
				}

				val = val*10 + val2
				vals.Push(strconv.FormatFloat(val, 'f', 6, 64))
			} else {
				vals.Push(char)
			}
			lastNumber = true
		}
	}

	val, err := vals.Pop()
	if err != nil {
		log.Panic(err)
	}
	result, err := strconv.ParseFloat(val.(string), 64)
	if err != nil {
		log.Panic(err)
	}

	return result
}
