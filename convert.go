package cnv

import (
	"reflect"
	"strconv"
	"time"
)

// Converter interface for custom type conversion.
type Converter interface {
	Parse(s string) error
}

// ParseBool parse string boolean.
func ParseBool(s string, p *bool) error {
	b, err := strconv.ParseBool(s)
	*p = b
	return err
}

// ParseInt parse string integer.
func ParseInt(s string, p *int) error {
	i, err := strconv.ParseInt(s, 10, 0)
	*p = int(i)
	return err
}

// ParseInt8 parse string 8 bit integer.
func ParseInt8(s string, p *int8) error {
	i, err := strconv.ParseInt(s, 10, 8)
	*p = int8(i)
	return err
}

// ParseInt16 parse string 16 bit integer.
func ParseInt16(s string, p *int16) error {
	i, err := strconv.ParseInt(s, 10, 16)
	*p = int16(i)
	return err
}

// ParseInt32 parse string 32 bit integer.
func ParseInt32(s string, p *int32) error {
	i, err := strconv.ParseInt(s, 10, 32)
	*p = int32(i)
	return err
}

// ParseInt64 parse string 64 bit integer.
func ParseInt64(s string, p *int64) error {
	i, err := strconv.ParseInt(s, 10, 64)
	*p = int64(i)
	return err
}

// ParseUint parse string unsigned integer.
func ParseUint(s string, p *uint) error {
	i, err := strconv.ParseUint(s, 10, 0)
	*p = uint(i)
	return err
}

// ParseUint8 parse string 8 bit unsigned integer.
func ParseUint8(s string, p *uint8) error {
	i, err := strconv.ParseUint(s, 10, 8)
	*p = uint8(i)
	return err
}

// ParseUint16 parse string 16 bit unsigned integer.
func ParseUint16(s string, p *uint16) error {
	i, err := strconv.ParseUint(s, 10, 16)
	*p = uint16(i)
	return err
}

// ParseUint32 parse string 32 bit unsigned integer.
func ParseUint32(s string, p *uint32) error {
	i, err := strconv.ParseUint(s, 10, 32)
	*p = uint32(i)
	return err
}

// ParseUint64 parse string 64 bit unsigned integer.
func ParseUint64(s string, p *uint64) error {
	i, err := strconv.ParseUint(s, 10, 64)
	*p = uint64(i)
	return err
}

// ParseFloat32 parse string 32 bit float.
func ParseFloat32(s string, p *float32) error {
	f, err := strconv.ParseFloat(s, 32)
	*p = float32(f)
	return err
}

// ParseFloat64 parse string 64 bit float.
func ParseFloat64(s string, p *float64) error {
	f, err := strconv.ParseFloat(s, 64)
	*p = float64(f)
	return err
}

// ParseTime parse string time.
func ParseTime(s string, p *time.Time) error {
	t, err := time.Parse(time.RFC3339, s)
	*p = t
	return err
}

// Parse string to Go type or Converter interface.
func Parse(s string, p interface{}) error {
	if reflect.ValueOf(p).Kind() != reflect.Ptr {
		return ErrNotAPointer
	}

	v := reflect.Indirect(reflect.ValueOf(p))

	switch v.Kind() {
	case reflect.Bool:
		return ParseBool(s, p.(*bool))
	case reflect.Int:
		return ParseInt(s, p.(*int))
	case reflect.Int8:
		return ParseInt8(s, p.(*int8))
	case reflect.Int16:
		return ParseInt16(s, p.(*int16))
	case reflect.Int32:
		return ParseInt32(s, p.(*int32))
	case reflect.Int64:
		return ParseInt64(s, p.(*int64))
	case reflect.Uint:
		return ParseUint(s, p.(*uint))
	case reflect.Uint8:
		return ParseUint8(s, p.(*uint8))
	case reflect.Uint16:
		return ParseUint16(s, p.(*uint16))
	case reflect.Uint32:
		return ParseUint32(s, p.(*uint32))
	case reflect.Uint64:
		return ParseUint64(s, p.(*uint64))
	case reflect.Float32:
		return ParseFloat32(s, p.(*float32))
	case reflect.Float64:
		return ParseFloat64(s, p.(*float64))
	case reflect.String:
		*p.(*string) = s
		return nil
	case reflect.Struct:
		if v.Type().String() == "time.Time" {
			return ParseTime(s, p.(*time.Time))
		}

		if c, ok := p.(Converter); ok {
			return c.Parse(s)
		}
	}

	return ErrUnsupportedType
}
