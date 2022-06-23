package main

const (
	INVALID_TOKEN = iota
	OPERATOR_TOKEN
	TYPE_TOKEN
	KEY_TOKEN
	STRING_TOKEN
	NUMBER_TOKEN
	END_TOKEN
)

func Lex(s *string) int {
	switch *s {
	case W:
		return OPERATOR_TOKEN
	case R:
		return OPERATOR_TOKEN
	case P:
		return OPERATOR_TOKEN
	case D:
		return OPERATOR_TOKEN
	case A:
		return OPERATOR_TOKEN
	case X:
		return OPERATOR_TOKEN
	case ADD:
		return OPERATOR_TOKEN
	case SUB:
		return OPERATOR_TOKEN
	case MUL:
		return OPERATOR_TOKEN
	case DIV:
		return OPERATOR_TOKEN
	case INC:
		return OPERATOR_TOKEN
	case DEC:
		return OPERATOR_TOKEN
	case AND:
		return OPERATOR_TOKEN
	case OR:
		return OPERATOR_TOKEN
	case XOR:
		return OPERATOR_TOKEN
	case LSHIFT:
		return OPERATOR_TOKEN
	case RSHIFT:
		return OPERATOR_TOKEN
	case INT:
		return TYPE_TOKEN
	case INT16:
		return TYPE_TOKEN
	case INT32:
		return TYPE_TOKEN
	case FLOAT:
		return TYPE_TOKEN
	case FLOAT64:
		return TYPE_TOKEN
	case STRING:
		return TYPE_TOKEN
	case BOOL:
		return TYPE_TOKEN
	default:
		{
			logger.Print(*s)
			c := (*s)[0]
			if c > 47 && c < 58 || c == '+' || c == '-' {
				// check if the string starts with 0-9 or + or -
				dot := true
				for _, c := range (*s)[1:] {
					// check if the following characters are not 0-9 or .
					if c == '.' && dot {
						// this is to ensure that . does not appear more than one time
						dot = false
					} else if c < 48 && c > 57 {
						return INVALID_TOKEN
					}
				}
				return NUMBER_TOKEN
			} else if (c > 64 && c < 92) || (c > 96 && c < 123) || c == '_' {
				// check if the string starts with a-z or A-Z or _ or not
				for _, c := range (*s)[1:] {
					// check if the following characters are not 0-9 and a-z and A-Z and _
					if (c < 48 && c > 57) && (c < 65 && c > 90) && (c < 97 && c > 122) && c != '_' {
						return INVALID_TOKEN
					}
				}
				return KEY_TOKEN
			}
			return INVALID_TOKEN
		}
	}
}
