package webx

import (
	"fmt"
	"reflect"
	"strconv"
)

func toInt64(val string) (int64, error) {
	return strconv.ParseInt(val, 10, 64)
}

// -----------------------------------------------------------------------------

func toInt32(val string) (int32, error) {
	n, err := strconv.ParseInt(val, 10, 32)
	return int32(n), err
}

// -----------------------------------------------------------------------------

func toInt16(val string) (int16, error) {
	n, err := strconv.ParseInt(val, 10, 16)
	return int16(n), err
}

// -----------------------------------------------------------------------------

func toInt8(val string) (int8, error) {
	n, err := strconv.ParseInt(val, 10, 8)
	return int8(n), err
}

// -----------------------------------------------------------------------------

func toInt(val string) (int, error) {
	return strconv.Atoi(val)
}

// -----------------------------------------------------------------------------

func toUint64(val string) (uint64, error) {
	return strconv.ParseUint(val, 10, 64)
}

// -----------------------------------------------------------------------------

func toUint32(val string) (uint32, error) {
	n, err := strconv.ParseUint(val, 10, 32)
	return uint32(n), err
}

// -----------------------------------------------------------------------------

func toUint16(val string) (uint16, error) {
	n, err := strconv.ParseUint(val, 10, 16)
	return uint16(n), err
}

// -----------------------------------------------------------------------------

func toUint8(val string) (uint8, error) {
	n, err := strconv.ParseUint(val, 10, 8)
	return uint8(n), err
}

// -----------------------------------------------------------------------------

func toUint(val string) (uint, error) {
	n, err := strconv.ParseUint(val, 10, 0)
	return uint(n), err
}

// -----------------------------------------------------------------------------

func toFloat64(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

// -----------------------------------------------------------------------------

func toFloat32(val string) (float32, error) {
	n, err := strconv.ParseFloat(val, 32)
	return float32(n), err
}

// -----------------------------------------------------------------------------

func toBool(val string) (bool, error) {
	return strconv.ParseBool(val)
}

// -----------------------------------------------------------------------------

func toComplex128(val string) (complex128, error) {
	return strconv.ParseComplex128(val, 128)
}

// -----------------------------------------------------------------------------

func toComplex64(val string) (complex64, error) {
	n, err := strconv.ParseComplex(val, 64)
	return complex64(n), err
}

// -----------------------------------------------------------------------------

func cast(val string, typ reflect.Type) (reflect.Value, error) {
	var v interface{}
	var err error

	switch typ.Kind() {
	case reflect.String:
		return reflect.ValueOf(val), nil
	case reflect.Bool:
		v, err = toBool(val)
	case reflect.Int:
		v, err = toInt(val)
	case reflect.Int8:
		v, err = toInt8(val)
	case reflect.Int16:
		v, err = toInt16(val)
	case reflect.Int32:
		v, err = toInt32(val)
	case reflect.Int64:
		v, err = toInt64(val)
	case reflect.Uint:
		v, err = toUint(val)
	case reflect.Uint8:
		v, err = toUint8(val)
	case reflect.Uint16:
		v, err = toUint16(val)
	case reflect.Uint32:
		v, err = toUint32(val)
	case reflect.Uint64:
		v, err = toUint64(val)
	case reflect.Float32:
		v, err = toFloat32(val)
	case reflect.Float64:
		v, err = toFloat64(val)
	case reflect.Complex64:
		v, err = toComplex64(val)
	case reflect.Complex128:
		v, err = toComplex128(val)
	default:
		err = fmt.Errorf("unsupported type of %v", typ)
	}

	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(v), err
}
