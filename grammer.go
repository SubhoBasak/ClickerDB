package main

const (
	_stateStart = iota
	_stateEntry
	_state0
	_state1
	_state2
	_state3
	_state4
)

func SyntaxChecker() (_SyntaxChecker func(int) (string, string)) {
	curState := _stateStart

	_SyntaxChecker = func(token int) (string, string) {
		switch curState {
		case _stateStart:
			{
				if token == OPERATOR_TOKEN {
					curState = _stateEntry
					return EMPTY, EMPTY
				}
				return EMPTY, "E:Must start with an operator"
			}
		case _stateEntry:
			{
				if token == TYPE_TOKEN {
					curState = _state0
				} else if token == KEY_TOKEN {
					curState = _state3
				} else if token == END_TOKEN {
					return ACTION0, EMPTY
				} else {
					return EMPTY, "E:Syntax error1"
				}
				return EMPTY, EMPTY
			}
		case _state0:
			{
				if token == KEY_TOKEN {
					curState = _state1
				} else if token == END_TOKEN {
					return ACTION3, EMPTY
				} else {
					return EMPTY, "E:Syntax error2"
				}
				return EMPTY, EMPTY
			}
		case _state1:
			{
				if token == STRING_TOKEN {
					curState = _state2
				} else {
					return EMPTY, "E:Syntax error3"
				}
				return EMPTY, EMPTY
			}
		case _state3:
			{
				if token == STRING_TOKEN {
					curState = _state4
				} else if token == END_TOKEN {
					return ACTION1, EMPTY
				} else {
					return EMPTY, "E:Syntax error4"
				}
				return EMPTY, EMPTY
			}
		case _state4:
			{
				if token == END_TOKEN {
					return ACTION2, EMPTY
				} else {
					return EMPTY, "E:Syntax error5"
				}
			}
		case _state2:
			return ACTION4, EMPTY
		default:
			return EMPTY, "E:Syntax error6"
		}
	}

	return
}
