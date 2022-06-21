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
			if token == OPERATOR_TOKEN {
				curState = _stateEntry
				return EMPTY, EMPTY
			}
			return EMPTY, "E:Query must starts with an operator"
		case _stateEntry:
			if token == TYPE_TOKEN {
				curState = _state0
			} else if token == KEY_TOKEN {
				curState = _state3
			} else if token == END_TOKEN {
				return ACTION0, EMPTY
			} else {
				return EMPTY, "E:Expecting type or key or ; after the operator"
			}
			return EMPTY, EMPTY
		case _state0:
			if token == KEY_TOKEN {
				curState = _state1
			} else if token == END_TOKEN {
				return ACTION3, EMPTY
			} else {
				return EMPTY, "E:Expection key or ; after the type"
			}
			return EMPTY, EMPTY
		case _state1:
			if token == STRING_TOKEN {
				curState = _state2
			} else if token == NUMBER_TOKEN {
				curState = _state2
			} else {
				return EMPTY, "E:Expecting a value after the key"
			}
			return EMPTY, EMPTY
		case _state2:
			if token == END_TOKEN {
				return ACTION4, EMPTY
			} else {
				return EMPTY, "E:Expecting ; after the value"
			}
		case _state3:
			if token == STRING_TOKEN {
				curState = _state4
			} else if token == NUMBER_TOKEN {
				curState = _state4
			} else if token == END_TOKEN {
				return ACTION1, EMPTY
			} else {
				return EMPTY, "E:Expecting a value or ; after the key"
			}
			return EMPTY, EMPTY
		case _state4:
			if token == END_TOKEN {
				return ACTION2, EMPTY
			} else {
				return EMPTY, "E:Expecting ; after the value"
			}
		default:
			return EMPTY, "E:Syntax error"
		}
	}

	return
}
