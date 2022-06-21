package main

import "fmt"

const (
	ACTION0 = "a0"
	ACTION1 = "a1"
	ACTION2 = "a2"
	ACTION3 = "a3"
	ACTION4 = "a4"
)

func Action0(operator string) string {
	switch operator {
	case X:
		Clear()
		return "O:Ok"
	default:
		return fmt.Sprintf("E:Unexpected %s", operator)
	}
}

func Action1(operator string, key string) string {
	switch operator {
	case R:
		return Read(&key)
	case D:
		Delete(&key)
		return "O:Ok"
	case INC:
		return Inc(&key)
	case DEC:
		return Dec(&key)
	default:
		return fmt.Sprintf("E:Unexpected %s", operator)
	}
}

func Action2(operator string, key string, val string) string {
	switch operator {
	case P:
		return Put(&key, &val)
	case ADD:
		return Add(&key, &val)
	case SUB:
		return Sub(&key, &val)
	case DIV:
		return Div(&key, &val)
	case MUL:
		return Mul(&key, &val)
	case AND:
		return And(&key, &val)
	case OR:
		return Or(&key, &val)
	case XOR:
		return Xor(&key, &val)
	default:
		return fmt.Sprintf("E:Unexpected %s", operator)
	}
}

func Action3(operator string, t string) string {
	switch operator {
	case A:
		return *AllType(&t)
	default:
		return fmt.Sprintf("E:Unexpected %s", operator)
	}
}

func Action4(operator string, t string, key string, val string) string {
	switch operator {
	case W:
		return Write(&t, &key, &val)
	default:
		return fmt.Sprintf("E:Unexpected %s", operator)
	}
}
