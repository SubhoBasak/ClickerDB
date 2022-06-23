package main

import "fmt"

func Execute(s *string) string {
	var keywords []string
	var strval byte = 0
	var token int

	lex := EMPTY
	ln := len(*s)
	synChcker := SyntaxChecker()

	for i := 0; i < ln; i++ {
		chr := (*s)[i]

		if (chr == SPACE || chr == TAB || chr == LINE || chr == END) && strval == 0 {
			if lex != EMPTY {
				token = Lex(&lex)
				if token == INVALID_TOKEN {
					return fmt.Sprintf("%d:Unknown keyword : %s", RESP_IDENTIFIER_ERROR, lex)
				}
				_, err := synChcker(token)
				if err == EMPTY {
					keywords = append(keywords, lex)
					lex = ""
				} else {
					return err
				}
			}
			continue
		} else if chr == SQUOTE || chr == DQUOTE {
			if strval == 0 {
				if lex == EMPTY {
					strval = chr
				} else {
					return fmt.Sprintf("%d:Unexpected %c after %s", RESP_SYNTAX_ERROR, chr, lex)
				}
			} else if strval == chr {
				_, err := synChcker(STRING_TOKEN)
				if err != EMPTY {
					return err
				}
				keywords = append(keywords, lex)
				lex = ""
				strval = 0
			}
			continue
		}

		lex += string(chr)
	}

	if lex == EMPTY {
		act, err := synChcker(END_TOKEN)
		if err != EMPTY {
			return err
		}

		// Call the action
		switch act {
		case ACTION0:
			return Action0(keywords[0])
		case ACTION1:
			return Action1(keywords[0], keywords[1])
		case ACTION2:
			return Action2(keywords[0], keywords[1], keywords[2])
		case ACTION3:
			return Action3(keywords[0], keywords[1])
		case ACTION4:
			return Action4(keywords[0], keywords[1], keywords[2], keywords[3])
		default:
			return fmt.Sprintf("%d:Invalid syntax meaning", RESP_LOGICAL_ERROR)
		}
	} else {
		if strval != 0 {
			return fmt.Sprintf("%d:Missing %c", RESP_SYNTAX_ERROR, strval)
		} else {
			logger.Printf("Parse error : %s", *s)
			return fmt.Sprintf("%d:Parse error", RESP_SYNTAX_ERROR)
		}
	}
}
