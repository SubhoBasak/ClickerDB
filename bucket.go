package main

import (
	"fmt"
	"strconv"
)

// storage buckets
var blMap = make(map[string]bool)
var i8Map = make(map[string]int8)
var i16Map = make(map[string]int16)
var i32Map = make(map[string]int32)
var stMap = make(map[string]string)
var f32Map = make(map[string]float32)
var f64Map = make(map[string]float64)
var typeMap = make(map[string]string)

// CRUD operations
func Read(k *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		return fmt.Sprintf("%d:%d", RESP_INT, i8Map[*k])
	case INT16:
		return fmt.Sprintf("%d:%d", RESP_INT16, i16Map[*k])
	case INT32:
		return fmt.Sprintf("%d:%d", RESP_INT32, i32Map[*k])
	case FLOAT:
		return fmt.Sprintf("%d:%f", RESP_FLOAT, f32Map[*k])
	case FLOAT64:
		return fmt.Sprintf("%d:%f", RESP_FLOAT64, f64Map[*k])
	case BOOL:
		return fmt.Sprintf("%d:%t", RESP_BOOL, blMap[*k])
	default:
		return fmt.Sprintf("%d:%s", RESP_STRING, stMap[*k])
	}
}

func Write(t *string, k *string, v *string) string {
	_, ok := typeMap[*k]
	if ok {
		return MSG_KEY_CONFLICT
	} else {
		typeMap[*k] = *t
	}

	switch *t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		i8Map[*k] = int8(i)
		return fmt.Sprint(RESP_OK)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		i16Map[*k] = int16(i)
		return fmt.Sprint(RESP_OK)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		i32Map[*k] = int32(i)
		return fmt.Sprint(RESP_OK)
	case FLOAT:
		f, err := strconv.ParseFloat(*v, 32)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT_INVALID
		}
		f32Map[*k] = float32(f)
		return fmt.Sprint(RESP_OK)
	case FLOAT64:
		f, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT64_INVALID
		}
		f64Map[*k] = f
		return fmt.Sprint(RESP_OK)
	case BOOL:
		b, err := strconv.ParseBool(*v)
		if err != nil {
			logger.Println(err)
			return MSG_BOOL_INVALID
		}
		blMap[*k] = b
		return fmt.Sprint(RESP_OK)
	default:
		stMap[*k] = *v
		return fmt.Sprint(RESP_OK)
	}
}

func Put(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		i8Map[*k] = int8(i)
		return fmt.Sprint(RESP_OK)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		i16Map[*k] = int16(i)
		return fmt.Sprint(RESP_OK)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		i32Map[*k] = int32(i)
		return fmt.Sprint(RESP_OK)
	case FLOAT:
		f, err := strconv.ParseFloat(*v, 32)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT_INVALID
		}
		f32Map[*k] = float32(f)
		return fmt.Sprint(RESP_OK)
	case FLOAT64:
		f, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT64_INVALID
		}
		f64Map[*k] = f
		return fmt.Sprint(RESP_OK)
	case BOOL:
		b, err := strconv.ParseBool(*v)
		if err != nil {
			logger.Println(err)
			return MSG_BOOL_INVALID
		}
		blMap[*k] = b
		return fmt.Sprint(RESP_OK)
	default:
		stMap[*k] = *v
		return fmt.Sprint(RESP_OK)
	}
}

func Delete(k *string) {
	t, ok := typeMap[*k]
	if !ok {
		return
	}

	switch t {
	case INT:
		delete(i8Map, *k)
		return
	case INT16:
		delete(i16Map, *k)
		return
	case INT32:
		delete(i32Map, *k)
		return
	case FLOAT:
		delete(f32Map, *k)
		return
	case FLOAT64:
		delete(f64Map, *k)
		return
	case BOOL:
		delete(blMap, *k)
		return
	default:
		delete(stMap, *k)
	}
}

func AllType(t *string) *string {
	output := EMPTY

	switch *t {
	case INT:
		for k, v := range i8Map {
			output += fmt.Sprintf("%s:%d;", k, v)
		}
		return &output // TODO - multiple output
	case INT16:
		for k, v := range i16Map {
			output += fmt.Sprintf("%s:%d;", k, v)
		}
		return &output
	case INT32:
		for k, v := range i32Map {
			output += fmt.Sprintf("%s:%d;", k, v)
		}
		return &output
	case FLOAT:
		for k, v := range f32Map {
			output += fmt.Sprintf("%s:%f;", k, v)
		}
		return &output
	case FLOAT64:
		for k, v := range f64Map {
			output += fmt.Sprintf("%s:%f;", k, v)
		}
		return &output
	case BOOL:
		for k, v := range blMap {
			output += fmt.Sprintf("%s:%t;", k, v)
		}
		return &output
	default:
		for k, v := range stMap {
			output += fmt.Sprintf(`%s:"%s";`, k, v)
		}
		return &output
	}
}

func Clear() {
	blMap = make(map[string]bool)
	i8Map = make(map[string]int8)
	i16Map = make(map[string]int16)
	i32Map = make(map[string]int32)
	stMap = make(map[string]string)
	f32Map = make(map[string]float32)
	f64Map = make(map[string]float64)
	typeMap = make(map[string]string)
}

// Arithmetic operations
func Add(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		tmp := i8Map[*k] + int8(i)
		i8Map[*k] = tmp
		return fmt.Sprintf("%d:%d", RESP_INT, tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		tmp := i16Map[*k] + int16(i)
		i16Map[*k] = tmp
		return fmt.Sprintf("%d:%d", RESP_INT16, tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		tmp := i32Map[*k] + int32(i)
		i32Map[*k] = tmp
		return fmt.Sprintf("%d:%d", RESP_INT32, tmp)
	case FLOAT:
		f, err := strconv.ParseFloat(*v, 32)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT_INVALID
		}
		tmp := f32Map[*k] + float32(f)
		f32Map[*k] = tmp
		return fmt.Sprintf("%d:%f", RESP_FLOAT, tmp)
	case FLOAT64:
		f, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT64_INVALID
		}
		tmp := f64Map[*k] + f
		f64Map[*k] = tmp
		return fmt.Sprintf("%d:%f", RESP_FLOAT64, tmp)
	default:
		return fmt.Sprintf("%d:Can't perform ADD on type %s", RESP_TYPE_ERROR, t)
	}
}

func Sub(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return fmt.Sprintf("%d:Key not found", RESP_KEY_ERROR)
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		tmp := i8Map[*k] - int8(i)
		i8Map[*k] = tmp
		return fmt.Sprintf("%d:%d", RESP_INT, tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		tmp := i16Map[*k] - int16(i)
		i16Map[*k] = tmp
		return fmt.Sprintf("%d:%d", RESP_INT16, tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		tmp := i32Map[*k] - int32(i)
		i32Map[*k] = tmp
		return fmt.Sprintf("%d:%d", RESP_INT32, tmp)
	case FLOAT:
		f, err := strconv.ParseFloat(*v, 32)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT_INVALID
		}
		tmp := f32Map[*k] - float32(f)
		f32Map[*k] = tmp
		return fmt.Sprintf("%d:%f", RESP_FLOAT, tmp)
	case FLOAT64:
		f, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT64_INVALID
		}
		tmp := f64Map[*k] - f
		f64Map[*k] = tmp
		return fmt.Sprintf("%d:%f", RESP_FLOAT64, tmp)
	default:
		return fmt.Sprintf("%d:Can't perform SUB on type %s", RESP_TYPE_ERROR, t)
	}
}

func Div(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		} else if i == 0 {
			return MSG_DIV_ZERO
		}
		tmp := i8Map[*k] / int8(i)
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		} else if i == 0 {
			return MSG_DIV_ZERO
		}
		tmp := i16Map[*k] / int16(i)
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		} else if i == 0 {
			return MSG_DIV_ZERO
		}
		tmp := i32Map[*k] / int32(i)
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT:
		f, err := strconv.ParseFloat(*v, 32)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT_INVALID
		} else if f == 0 {
			return MSG_DIV_ZERO
		}
		tmp := f32Map[*k] / float32(f)
		f32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT64:
		f, err := strconv.ParseFloat(*v, 64)
		if err != nil {
			logger.Println(err)
			return MSG_FLOAT64_INVALID
		} else if f == 0 {
			return MSG_DIV_ZERO
		}
		tmp := f64Map[*k] / float64(f)
		f64Map[*k] = tmp
		return fmt.Sprint(f64Map[*k])
	default:
		return fmt.Sprintf("E:Can't perform DIV on type %s", t)
	}
}

func Mul(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err == nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		tmp := i8Map[*k] * int8(i)
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err == nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		tmp := i16Map[*k] * int16(i)
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err == nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		tmp := i32Map[*k] * int32(i)
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT:
		f, err := strconv.ParseFloat(*v, 32)
		if err == nil {
			logger.Println(err)
			return MSG_FLOAT_INVALID
		}
		tmp := f32Map[*k] * float32(f)
		f32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT64:
		f, err := strconv.ParseFloat(*v, 64)
		if err == nil {
			logger.Println(err)
			return MSG_FLOAT64_INVALID
		}
		tmp := f64Map[*k] * float64(f)
		f64Map[*k] = tmp
		return fmt.Sprint(tmp)
	default:
		return fmt.Sprintf("E:Can't perform MUL on type %s", t)
	}
}

func Inc(k *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		tmp := i8Map[*k] + 1
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		tmp := i16Map[*k] + 1
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		tmp := i32Map[*k] + 1
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT:
		tmp := f32Map[*k] + 1
		f32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT64:
		tmp := f64Map[*k] + 1
		f64Map[*k] = tmp
		return fmt.Sprint(tmp)
	default:
		return fmt.Sprintf("E:Can't perform INC on type %s", t)
	}
}

func Dec(k *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		tmp := i8Map[*k] - 1
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		tmp := i16Map[*k] - 1
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		tmp := i32Map[*k] - 1
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT:
		tmp := f32Map[*k] - 1
		f32Map[*k] = tmp
		return fmt.Sprint(tmp)
	case FLOAT64:
		tmp := f64Map[*k] - 1
		f64Map[*k] = tmp
		return fmt.Sprint(tmp)
	default:
		return fmt.Sprintf("E:Can't perform DEC on type %s", t)
	}
}

// Bitwise operations
func And(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		tmp := i8Map[*k] & int8(i)
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		tmp := i16Map[*k] & int16(i)
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		tmp := i32Map[*k] & int32(i)
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	default:
		return fmt.Sprintf("E:Can't perform AND on type %s", t)
	}
}

func Or(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		tmp := i8Map[*k] | int8(i)
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		tmp := i16Map[*k] | int16(i)
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		tmp := i32Map[*k] | int32(i)
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	default:
		return fmt.Sprintf("E:Can't perform OR on type %s", t)
	}
}

func Xor(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return MSG_KEY_NOT_FOUND
	}

	switch t {
	case INT:
		i, err := strconv.ParseInt(*v, 10, 8)
		if err != nil {
			logger.Println(err)
			return MSG_INT_INVALID
		}
		tmp := i8Map[*k] ^ int8(i)
		i8Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT16:
		i, err := strconv.ParseInt(*v, 10, 16)
		if err != nil {
			logger.Println(err)
			return MSG_INT16_INVALID
		}
		tmp := i16Map[*k] ^ int16(i)
		i16Map[*k] = tmp
		return fmt.Sprint(tmp)
	case INT32:
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			logger.Println(err)
			return MSG_INT32_INVALID
		}
		tmp := i32Map[*k] ^ int32(i)
		i32Map[*k] = tmp
		return fmt.Sprint(tmp)
	default:
		return fmt.Sprintf("E:Can't perform XOR on type %s", t)
	}
}
