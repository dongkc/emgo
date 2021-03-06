package main

import (
	"fmt"
	"rtos"
	"unsafe"

	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
)

var dbg = rtos.Debug(0)

func printa(s string, gs, ga, cs, ca uintptr) {
	fmt.Fprintf(dbg, "%6s %2d %2d\t%2d %2d\n", s, gs, ga, cs, ca)
}

type S16 struct {
	i16 int16
	i8  int8
}

type S32 struct {
	i32 int32
	i16 int16
	i8  int8
}

type S64 struct {
	i64 int64
	i32 int32
	i16 int16
	i8  int8
}

func SizeByte() uintptr
func SizeInt() uintptr
func SizeInt16() uintptr
func SizeInt32() uintptr
func SizeInt64() uintptr
func SizeS16() uintptr
func SizeS32() uintptr
func SizeS64() uintptr

func AlignByte() uintptr
func AlignInt() uintptr
func AlignInt16() uintptr
func AlignInt32() uintptr
func AlignInt64() uintptr
func AlignS16() uintptr
func AlignS32() uintptr
func AlignS64() uintptr

func main() {
	system.Setup32(0)
	systick.Setup()

	dbg.WriteString("Data size/alignment\n\n")
	dbg.WriteString("Type    Emgo       C\n")
	dbg.WriteString("---------------------\n")
	var b byte
	printa(
		"byte  ",
		unsafe.Sizeof(b), unsafe.Alignof(b), SizeByte(), AlignByte(),
	)
	var i int
	printa(
		"int   ",
		unsafe.Sizeof(i), unsafe.Alignof(i), SizeInt(), AlignInt(),
	)
	var i16 int16
	printa(
		"int16 ",
		unsafe.Sizeof(i16), unsafe.Alignof(i16), SizeInt16(), AlignInt16(),
	)
	var i32 int32
	printa(
		"int32 ",
		unsafe.Sizeof(i32), unsafe.Alignof(i32), SizeInt32(), AlignInt32(),
	)
	var i64 int64
	printa(
		"int64 ",
		unsafe.Sizeof(i64), unsafe.Alignof(i64), SizeInt64(), AlignInt64(),
	)

	printa(
		"S16   ",
		unsafe.Sizeof(S16{}), unsafe.Alignof(S16{}), SizeS16(), AlignS16(),
	)
	printa(
		"S32   ",
		unsafe.Sizeof(S32{}), unsafe.Alignof(S32{}), SizeS32(), AlignS32(),
	)
	printa(
		"S64   ",
		unsafe.Sizeof(S64{}), unsafe.Alignof(S64{}), SizeS64(), AlignS64(),
	)
}
