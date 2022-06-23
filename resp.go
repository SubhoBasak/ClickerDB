package main

import "fmt"

// ERRORS
const (
	RESP_SERVER_ERROR = 100 + iota
	RESP_QUERY_ERROR
	RESP_KEY_ERROR
	RESP_TYPE_ERROR
	RESP_VALUE_ERROR
	RESP_SYNTAX_ERROR
	RESP_IDENTIFIER_ERROR
	RESP_LOGICAL_ERROR
	RESP_DIV_ZERO_ERROR
)

// SUCCESS
const (
	RESP_OK = 200 + iota
	RESP_INT
	RESP_INT16
	RESP_INT32
	RESP_FLOAT
	RESP_FLOAT64
	RESP_STRING
	RESP_BOOL
	RESP_SINGLE_TYPE
	RESP_MULTIPLE_TYPE
)

var (
	MSG_OK              = fmt.Sprint(RESP_OK)
	MSG_KEY_CONFLICT    = fmt.Sprintf("%d:Key conflict", RESP_KEY_ERROR)
	MSG_KEY_NOT_EXIST   = fmt.Sprintf("%d:Key not exist", RESP_KEY_ERROR)
	MSG_KEY_NOT_FOUND   = fmt.Sprintf("%d:Key not found", RESP_KEY_ERROR)
	MSG_INT_INVALID     = fmt.Sprintf("%d:Invalid INT value", RESP_VALUE_ERROR)
	MSG_INT16_INVALID   = fmt.Sprintf("%d:Invalid INT16 value", RESP_VALUE_ERROR)
	MSG_INT32_INVALID   = fmt.Sprintf("%d:Invalid INT32 value", RESP_VALUE_ERROR)
	MSG_FLOAT_INVALID   = fmt.Sprintf("%d:Invalid FLOAT value", RESP_VALUE_ERROR)
	MSG_FLOAT64_INVALID = fmt.Sprintf("%d:Invalid FLOAT64 value", RESP_VALUE_ERROR)
	MSG_BOOL_INVALID    = fmt.Sprintf("%d:Invalid BOOL value", RESP_VALUE_ERROR)
	MSG_DIV_ZERO        = fmt.Sprintf("%d:Can't divide by zero", RESP_DIV_ZERO_ERROR)
)
