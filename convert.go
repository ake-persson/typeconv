package cnv

import (
	"reflect"
	"strconv"
	"time"
)

type Converter interface {
	Parse() (interface{}, error)
}

func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func ParseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 0)
	return int(i), err
}

func ParseInt8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	return int8(i), err
}

func ParseInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}

func ParseInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}

func ParseInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int64(i), err
}

func ParseUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	return uint(i), err
}

func ParseUint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	return uint8(i), err
}

func ParseUint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	return uint16(i), err
}

func ParseUint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	return uint32(i), err
}

func ParseUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	return uint64(i), err
}

func ParseFloat32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}

func ParseFloat64(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	return float64(i), err
}

func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// Parse with interface{} pointer?

func Parse(s string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Bool:
		return ParseBool(s)

	case reflect.Int:
		return ParseInt(s)
	case reflect.Int8:
		return ParseInt8(s)
	case reflect.Int16:
		return ParseInt16(s)
	case reflect.Int32:
		return ParseInt32(s)
	case reflect.Int64:
		return ParseInt64(s)

	case reflect.Uint:
		return ParseInt(s)
	case reflect.Uint8:
		return ParseInt8(s)
	case reflect.Uint16:
		return ParseInt16(s)
	case reflect.Uint32:
		return ParseInt32(s)
	case reflect.Uint64:
		return ParseInt64(s)

	case reflect.Float32:
		return ParseFloat32(s)
	case reflect.Float64:
		return ParseFloat64(s)

	case reflect.String:
		return s, nil

	case reflect.Struct:
		/*
			if kind == "time.Time" {
				return ParseTime(s)
			}
		*/

		/*
			if c, ok := s.(Converter); ok {
				return c.Parse(s)
			}
		*/
	}

	return nil, ErrKindNotSupported
}
