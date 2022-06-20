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
		return "E:Key not found"
	}

	switch t {
	case INT:
		return fmt.Sprint(i8Map[*k])
	case INT16:
		return fmt.Sprint(i16Map[*k])
	case INT32:
		return fmt.Sprint(i32Map[*k])
	case FLOAT:
		return fmt.Sprint(f32Map[*k])
	case FLOAT64:
		return fmt.Sprint(f64Map[*k])
	case STRING:
		return stMap[*k]
	case BOOL:
		return fmt.Sprint(blMap[*k])
	default:
		return "E:Invalid type"
	}
}

func Write(t *string, k *string, v *string) string {
	_, ok := typeMap[*k]
	if ok {
		return "E:Key already exist"
	} else {
		typeMap[*k] = *t
	}

	switch *t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] = int8(i)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] = int16(i)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid int16 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] = int32(i)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid int32 value"
			}
		}
	case FLOAT:
		{
			f, err := strconv.ParseFloat(*v, 32)
			if err == nil {
				f32Map[*k] = float32(f)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid float32 value"
			}
		}
	case FLOAT64:
		{
			f, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				f64Map[*k] = f
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid float64 value"
			}
		}
	case STRING:
		{
			stMap[*k] = *v
			return "O:Ok"
		}
	case BOOL:
		{
			b, err := strconv.ParseBool(*v)
			if err == nil {
				blMap[*k] = b
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid bool value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Put(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] = int8(i)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] = int16(i)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid int16 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] = int32(i)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid int32 value"
			}
		}
	case FLOAT:
		{
			f, err := strconv.ParseFloat(*v, 32)
			if err == nil {
				f32Map[*k] = float32(f)
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid float32 value"
			}
		}
	case FLOAT64:
		{
			f, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				f64Map[*k] = f
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid float64 value"
			}
		}
	case STRING:
		{
			stMap[*k] = *v
			return "O:Ok"
		}
	case BOOL:
		{
			b, err := strconv.ParseBool(*v)
			if err == nil {
				blMap[*k] = b
				return "O:Ok"
			} else {
				logger.Println(err)
				return "E:Invalid bool value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Delete(k *string) {
	t, ok := typeMap[*k]
	if !ok {
		return
	}

	switch t {
	case INT:
		{
			delete(i8Map, *k)
			return
		}
	case INT16:
		{
			delete(i16Map, *k)
			return
		}
	case INT32:
		{
			delete(i32Map, *k)
			return
		}
	case FLOAT:
		{
			delete(f32Map, *k)
			return
		}
	case FLOAT64:
		{
			delete(f64Map, *k)
			return
		}
	case STRING:
		{
			delete(stMap, *k)
			return
		}
	case BOOL:
		{
			delete(blMap, *k)
			return
		}
	default:
		return
	}
}

func AllType(t *string) *string {
	output := EMPTY

	switch *t {
	case INT:
		{
			for k, v := range i8Map {
				output += fmt.Sprintf("%s:%d;", k, v)
			}
		}
	case INT16:
		{
			for k, v := range i16Map {
				output += fmt.Sprintf("%s:%d;", k, v)
			}
		}
	case INT32:
		{
			for k, v := range i32Map {
				output += fmt.Sprintf("%s:%d;", k, v)
			}
		}
	case FLOAT:
		{
			for k, v := range f32Map {
				output += fmt.Sprintf("%s:%f;", k, v)
			}
		}
	case FLOAT64:
		{
			for k, v := range f64Map {
				output += fmt.Sprintf("%s:%f;", k, v)
			}
		}
	case STRING:
		{
			for k, v := range stMap {
				output += fmt.Sprintf(`%s:"%s";`, k, v)
			}
		}
	case BOOL:
		{
			for k, v := range blMap {
				output += fmt.Sprintf("%s:%t;", k, v)
			}
		}
	default:
		output = "E:Invalid type"
	}

	return &output
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
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] += int8(i)
				return fmt.Sprint(i8Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] += int16(i)
				return fmt.Sprint(i16Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] += int32(i)
				return fmt.Sprint(i32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT:
		{
			f, err := strconv.ParseFloat(*v, 32)
			if err == nil {
				f32Map[*k] += float32(f)
				return fmt.Sprint(f32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT64:
		{
			f, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				f64Map[*k] += float64(f)
				return fmt.Sprint(f64Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Sub(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] -= int8(i)
				return fmt.Sprint(i8Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] -= int16(i)
				return fmt.Sprint(i16Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] -= int32(i)
				return fmt.Sprint(i32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT:
		{
			f, err := strconv.ParseFloat(*v, 32)
			if err == nil {
				f32Map[*k] -= float32(f)
				return fmt.Sprint(f32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT64:
		{
			f, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				f64Map[*k] -= float64(f)
				return fmt.Sprint(f64Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Div(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] += int8(i)
				return fmt.Sprint(i8Map[*k])
			} else if i == 0 {
				return "E:Can't divide by zero"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] += int16(i)
				return fmt.Sprint(i16Map[*k])
			} else if i == 0 {
				return "E:Can't divide by zero"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] += int32(i)
				return fmt.Sprint(i32Map[*k])
			} else if i == 0 {
				return "E:Can't divide by zero"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT:
		{
			f, err := strconv.ParseFloat(*v, 32)
			if err == nil {
				f32Map[*k] += float32(f)
				return fmt.Sprint(f32Map[*k])
			} else if f == 0 {
				return "E:Can't divide by zero"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT64:
		{
			f, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				f64Map[*k] += float64(f)
				return fmt.Sprint(f64Map[*k])
			} else if f == 0 {
				return "E:Can't divide by zero"
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Mul(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] *= int8(i)
				return fmt.Sprint(i8Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] *= int16(i)
				return fmt.Sprint(i16Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] *= int32(i)
				return fmt.Sprint(i32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT:
		{
			f, err := strconv.ParseFloat(*v, 32)
			if err == nil {
				f32Map[*k] *= float32(f)
				return fmt.Sprint(f32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case FLOAT64:
		{
			f, err := strconv.ParseFloat(*v, 64)
			if err == nil {
				f64Map[*k] *= float64(f)
				return fmt.Sprint(f64Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Inc(k *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i8Map[*k]++
			return fmt.Sprint(i8Map[*k])
		}
	case INT16:
		{
			i16Map[*k]++
			return fmt.Sprint(i16Map[*k])
		}
	case INT32:
		{
			i32Map[*k]++
			return fmt.Sprint(i32Map[*k])
		}
	case FLOAT:
		{
			f32Map[*k]++
			return fmt.Sprint(f32Map[*k])
		}
	case FLOAT64:
		{
			f64Map[*k]++
			return fmt.Sprint(f64Map[*k])
		}
	default:
		return "E:Value is not numeric type"
	}
}

func Dec(k *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i8Map[*k]--
			return fmt.Sprint(i8Map[*k])
		}
	case INT16:
		{
			i16Map[*k]--
			return fmt.Sprint(i16Map[*k])
		}
	case INT32:
		{
			i32Map[*k]--
			return fmt.Sprint(i32Map[*k])
		}
	case FLOAT:
		{
			f32Map[*k]--
			return fmt.Sprint(f32Map[*k])
		}
	case FLOAT64:
		{
			f64Map[*k]--
			return fmt.Sprint(f64Map[*k])
		}
	default:
		return "E:Value is not numeric type"
	}
}

// Bitwise operations
func And(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] &= int8(i)
				return fmt.Sprint(i8Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] &= int16(i)
				return fmt.Sprint(i16Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] &= int32(i)
				return fmt.Sprint(i32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Or(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] |= int8(i)
				return fmt.Sprint(i8Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] |= int16(i)
				return fmt.Sprint(i16Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] |= int32(i)
				return fmt.Sprint(i32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}

func Xor(k *string, v *string) string {
	t, ok := typeMap[*k]
	if !ok {
		return "E:Key not found"
	}

	switch t {
	case INT:
		{
			i, err := strconv.ParseInt(*v, 10, 8)
			if err == nil {
				i8Map[*k] ^= int8(i)
				return fmt.Sprint(i8Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT16:
		{
			i, err := strconv.ParseInt(*v, 10, 16)
			if err == nil {
				i16Map[*k] ^= int16(i)
				return fmt.Sprint(i16Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	case INT32:
		{
			i, err := strconv.ParseInt(*v, 10, 32)
			if err == nil {
				i32Map[*k] ^= int32(i)
				return fmt.Sprint(i32Map[*k])
			} else {
				logger.Println(err)
				return "E:Invalid int8 value"
			}
		}
	default:
		return "E:Invalid type"
	}
}
