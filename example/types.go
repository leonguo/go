package main

import (
	"fmt"
	"math"
	"strconv"
	"unsafe"
)

func main() {
	var v_bool bool
	var v_byte byte
	var v_rune rune
	var v_int8 int8
	var v_int16 int16
	var v_int32 int32
	var v_int64 int64
	var v_uint8 uint8
	var v_uint16 uint16
	var v_uint32 uint32
	var v_uint64 uint64

	var v_float32 float32
	var v_float64 float64

	fmt.Printf("%-20s%-20s%-20s%-50s\n", "Type", "Sizeof", "DefaultValue", "Description")
	fmt.Printf("%-20s%-20d%-20t%-50s\n", "Bool", unsafe.Sizeof(v_bool), v_bool, "")
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "byte", unsafe.Sizeof(v_byte), v_byte, "0 -"+strconv.Itoa(math.MaxInt8))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "rune", unsafe.Sizeof(v_rune), v_rune, strconv.Itoa(math.MinInt32)+" - "+strconv.Itoa(math.MaxInt32))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "int8", unsafe.Sizeof(v_int8), v_int8, strconv.Itoa(math.MinInt8)+" - "+strconv.Itoa(math.MaxInt8))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "int16", unsafe.Sizeof(v_int16), v_int16, strconv.Itoa(math.MinInt16)+" - "+strconv.Itoa(math.MaxInt16))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "int32", unsafe.Sizeof(v_int32), v_int32, strconv.Itoa(math.MinInt32)+" - "+strconv.Itoa(math.MaxInt32))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "int64", unsafe.Sizeof(v_int64), v_int64, strconv.Itoa(math.MinInt64)+" - "+strconv.Itoa(math.MaxInt64))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "uint8", unsafe.Sizeof(v_uint8), v_uint8, "0 - "+strconv.Itoa(math.MaxUint8))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "uint16", unsafe.Sizeof(v_uint16), v_uint16, "0 - "+strconv.Itoa(math.MaxUint16))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "uint32", unsafe.Sizeof(v_uint32), v_uint32, "0 - "+strconv.Itoa(math.MaxUint32))
	fmt.Printf("%-20s%-20d%-20d%-50s\n", "uint64", unsafe.Sizeof(v_uint64), v_uint64, "0 - "+strconv.FormatUint(math.MaxUint64, 10))

	fmt.Printf("%-20s%-20d%-20f%-50s\n", "float32", unsafe.Sizeof(v_float32), v_float32, "0 - "+strconv.FormatFloat(math.MaxFloat32, 'f', 6, 32))
	fmt.Printf("%-20s%-20d%-20f%-50s\n", "float64", unsafe.Sizeof(v_float64), v_float64, "")

}
