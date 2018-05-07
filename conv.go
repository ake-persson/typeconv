package cnv

import (
	"reflect"
	"strconv"
)

type Converter interface {
	Interface() (interface{}, error)
}

func Conv(s string, kind reflect.Kind) (interface{}, error) {
	switch kind {
	case reflect.Bool:
		return strconv.ParseBool(s)

	case reflect.Int:
		i, err := strconv.ParseInt(s, 10, 0)
		return int(i), err
	case reflect.Int8:
		i, err := strconv.ParseInt(s, 10, 8)
		return int8(i), err
	case reflect.Int16:
		i, err := strconv.ParseInt(s, 10, 16)
		return int16(i), err
	case reflect.Int32:
		i, err := strconv.ParseInt(s, 10, 32)
		return int32(i), err
	case reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		return int64(i), err

	case reflect.Uint:
		i, err := strconv.ParseUint(s, 10, 0)
		return uint(i), err
	case reflect.Uint8:
		i, err := strconv.ParseInt(s, 10, 8)
		return uint8(i), err
	case reflect.Uint16:
		i, err := strconv.ParseInt(s, 10, 16)
		return uint16(i), err
	case reflect.Uint32:
		i, err := strconv.ParseInt(s, 10, 32)
		return uint32(i), err
	case reflect.Uint64:
		i, err := strconv.ParseInt(s, 10, 64)
		return uint64(i), err

	case reflect.Float32:
		f, err := strconv.ParseFloat(s, 32)
		return float32(f), err
	case reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		return float64(f), err

	case reflect.String:
		return s, nil

		//	case reflect.Struct:
		//		if va.Type().String() == "time.Time" && vb.Type().String() == "time.Time" {
		//			return va.Interface().(time.Time).Equal(vb.Interface().(time.Time)), nil
		//		}

		//	if ca, ok := va.Interface().(Comparer); ok {
		//		return ca.Eq(vb.Interface())
		//	}
	}

	return nil, ErrKindNotSupported
}
