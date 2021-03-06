// +build l476xx

package exti

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type EXTI_Periph struct {
	IMR1   IMR1
	EMR1   EMR1
	RTSR1  RTSR1
	FTSR1  FTSR1
	SWIER1 SWIER1
	PR1    PR1
	_      [2]uint32
	IMR2   IMR2
	EMR2   EMR2
	RTSR2  RTSR2
	FTSR2  FTSR2
	SWIER2 SWIER2
	PR2    PR2
}

func (p *EXTI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var EXTI = (*EXTI_Periph)(unsafe.Pointer(uintptr(mmap.EXTI_BASE)))

type IMR1_Bits uint32

func (b IMR1_Bits) Field(mask IMR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMR1_Bits) J(v int) IMR1_Bits {
	return IMR1_Bits(bits.Make32(v, uint32(mask)))
}

type IMR1 struct{ mmio.U32 }

func (r *IMR1) Bits(mask IMR1_Bits) IMR1_Bits { return IMR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMR1) StoreBits(mask, b IMR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMR1) SetBits(mask IMR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *IMR1) ClearBits(mask IMR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *IMR1) Load() IMR1_Bits               { return IMR1_Bits(r.U32.Load()) }
func (r *IMR1) Store(b IMR1_Bits)             { r.U32.Store(uint32(b)) }

type IMR1_Mask struct{ mmio.UM32 }

func (rm IMR1_Mask) Load() IMR1_Bits   { return IMR1_Bits(rm.UM32.Load()) }
func (rm IMR1_Mask) Store(b IMR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) IM0() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM0)}}
}

func (p *EXTI_Periph) IM1() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM1)}}
}

func (p *EXTI_Periph) IM2() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM2)}}
}

func (p *EXTI_Periph) IM3() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM3)}}
}

func (p *EXTI_Periph) IM4() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM4)}}
}

func (p *EXTI_Periph) IM5() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM5)}}
}

func (p *EXTI_Periph) IM6() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM6)}}
}

func (p *EXTI_Periph) IM7() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM7)}}
}

func (p *EXTI_Periph) IM8() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM8)}}
}

func (p *EXTI_Periph) IM9() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM9)}}
}

func (p *EXTI_Periph) IM10() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM10)}}
}

func (p *EXTI_Periph) IM11() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM11)}}
}

func (p *EXTI_Periph) IM12() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM12)}}
}

func (p *EXTI_Periph) IM13() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM13)}}
}

func (p *EXTI_Periph) IM14() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM14)}}
}

func (p *EXTI_Periph) IM15() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM15)}}
}

func (p *EXTI_Periph) IM16() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM16)}}
}

func (p *EXTI_Periph) IM17() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM17)}}
}

func (p *EXTI_Periph) IM18() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM18)}}
}

func (p *EXTI_Periph) IM19() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM19)}}
}

func (p *EXTI_Periph) IM20() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM20)}}
}

func (p *EXTI_Periph) IM21() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM21)}}
}

func (p *EXTI_Periph) IM22() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM22)}}
}

func (p *EXTI_Periph) IM23() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM23)}}
}

func (p *EXTI_Periph) IM24() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM24)}}
}

func (p *EXTI_Periph) IM25() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM25)}}
}

func (p *EXTI_Periph) IM26() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM26)}}
}

func (p *EXTI_Periph) IM27() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM27)}}
}

func (p *EXTI_Periph) IM28() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM28)}}
}

func (p *EXTI_Periph) IM29() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM29)}}
}

func (p *EXTI_Periph) IM30() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM30)}}
}

func (p *EXTI_Periph) IM31() IMR1_Mask {
	return IMR1_Mask{mmio.UM32{&p.IMR1.U32, uint32(IM31)}}
}

type EMR1_Bits uint32

func (b EMR1_Bits) Field(mask EMR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EMR1_Bits) J(v int) EMR1_Bits {
	return EMR1_Bits(bits.Make32(v, uint32(mask)))
}

type EMR1 struct{ mmio.U32 }

func (r *EMR1) Bits(mask EMR1_Bits) EMR1_Bits { return EMR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *EMR1) StoreBits(mask, b EMR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EMR1) SetBits(mask EMR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *EMR1) ClearBits(mask EMR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *EMR1) Load() EMR1_Bits               { return EMR1_Bits(r.U32.Load()) }
func (r *EMR1) Store(b EMR1_Bits)             { r.U32.Store(uint32(b)) }

type EMR1_Mask struct{ mmio.UM32 }

func (rm EMR1_Mask) Load() EMR1_Bits   { return EMR1_Bits(rm.UM32.Load()) }
func (rm EMR1_Mask) Store(b EMR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) EM0() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM0)}}
}

func (p *EXTI_Periph) EM1() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM1)}}
}

func (p *EXTI_Periph) EM2() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM2)}}
}

func (p *EXTI_Periph) EM3() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM3)}}
}

func (p *EXTI_Periph) EM4() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM4)}}
}

func (p *EXTI_Periph) EM5() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM5)}}
}

func (p *EXTI_Periph) EM6() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM6)}}
}

func (p *EXTI_Periph) EM7() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM7)}}
}

func (p *EXTI_Periph) EM8() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM8)}}
}

func (p *EXTI_Periph) EM9() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM9)}}
}

func (p *EXTI_Periph) EM10() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM10)}}
}

func (p *EXTI_Periph) EM11() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM11)}}
}

func (p *EXTI_Periph) EM12() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM12)}}
}

func (p *EXTI_Periph) EM13() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM13)}}
}

func (p *EXTI_Periph) EM14() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM14)}}
}

func (p *EXTI_Periph) EM15() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM15)}}
}

func (p *EXTI_Periph) EM16() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM16)}}
}

func (p *EXTI_Periph) EM17() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM17)}}
}

func (p *EXTI_Periph) EM18() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM18)}}
}

func (p *EXTI_Periph) EM19() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM19)}}
}

func (p *EXTI_Periph) EM20() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM20)}}
}

func (p *EXTI_Periph) EM21() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM21)}}
}

func (p *EXTI_Periph) EM22() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM22)}}
}

func (p *EXTI_Periph) EM23() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM23)}}
}

func (p *EXTI_Periph) EM24() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM24)}}
}

func (p *EXTI_Periph) EM25() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM25)}}
}

func (p *EXTI_Periph) EM26() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM26)}}
}

func (p *EXTI_Periph) EM27() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM27)}}
}

func (p *EXTI_Periph) EM28() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM28)}}
}

func (p *EXTI_Periph) EM29() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM29)}}
}

func (p *EXTI_Periph) EM30() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM30)}}
}

func (p *EXTI_Periph) EM31() EMR1_Mask {
	return EMR1_Mask{mmio.UM32{&p.EMR1.U32, uint32(EM31)}}
}

type RTSR1_Bits uint32

func (b RTSR1_Bits) Field(mask RTSR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RTSR1_Bits) J(v int) RTSR1_Bits {
	return RTSR1_Bits(bits.Make32(v, uint32(mask)))
}

type RTSR1 struct{ mmio.U32 }

func (r *RTSR1) Bits(mask RTSR1_Bits) RTSR1_Bits { return RTSR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *RTSR1) StoreBits(mask, b RTSR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSR1) SetBits(mask RTSR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *RTSR1) ClearBits(mask RTSR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *RTSR1) Load() RTSR1_Bits                { return RTSR1_Bits(r.U32.Load()) }
func (r *RTSR1) Store(b RTSR1_Bits)              { r.U32.Store(uint32(b)) }

type RTSR1_Mask struct{ mmio.UM32 }

func (rm RTSR1_Mask) Load() RTSR1_Bits   { return RTSR1_Bits(rm.UM32.Load()) }
func (rm RTSR1_Mask) Store(b RTSR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) RT0() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT0)}}
}

func (p *EXTI_Periph) RT1() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT1)}}
}

func (p *EXTI_Periph) RT2() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT2)}}
}

func (p *EXTI_Periph) RT3() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT3)}}
}

func (p *EXTI_Periph) RT4() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT4)}}
}

func (p *EXTI_Periph) RT5() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT5)}}
}

func (p *EXTI_Periph) RT6() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT6)}}
}

func (p *EXTI_Periph) RT7() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT7)}}
}

func (p *EXTI_Periph) RT8() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT8)}}
}

func (p *EXTI_Periph) RT9() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT9)}}
}

func (p *EXTI_Periph) RT10() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT10)}}
}

func (p *EXTI_Periph) RT11() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT11)}}
}

func (p *EXTI_Periph) RT12() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT12)}}
}

func (p *EXTI_Periph) RT13() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT13)}}
}

func (p *EXTI_Periph) RT14() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT14)}}
}

func (p *EXTI_Periph) RT15() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT15)}}
}

func (p *EXTI_Periph) RT16() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT16)}}
}

func (p *EXTI_Periph) RT18() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT18)}}
}

func (p *EXTI_Periph) RT19() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT19)}}
}

func (p *EXTI_Periph) RT20() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT20)}}
}

func (p *EXTI_Periph) RT21() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT21)}}
}

func (p *EXTI_Periph) RT22() RTSR1_Mask {
	return RTSR1_Mask{mmio.UM32{&p.RTSR1.U32, uint32(RT22)}}
}

type FTSR1_Bits uint32

func (b FTSR1_Bits) Field(mask FTSR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FTSR1_Bits) J(v int) FTSR1_Bits {
	return FTSR1_Bits(bits.Make32(v, uint32(mask)))
}

type FTSR1 struct{ mmio.U32 }

func (r *FTSR1) Bits(mask FTSR1_Bits) FTSR1_Bits { return FTSR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *FTSR1) StoreBits(mask, b FTSR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FTSR1) SetBits(mask FTSR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *FTSR1) ClearBits(mask FTSR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *FTSR1) Load() FTSR1_Bits                { return FTSR1_Bits(r.U32.Load()) }
func (r *FTSR1) Store(b FTSR1_Bits)              { r.U32.Store(uint32(b)) }

type FTSR1_Mask struct{ mmio.UM32 }

func (rm FTSR1_Mask) Load() FTSR1_Bits   { return FTSR1_Bits(rm.UM32.Load()) }
func (rm FTSR1_Mask) Store(b FTSR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) FT0() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT0)}}
}

func (p *EXTI_Periph) FT1() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT1)}}
}

func (p *EXTI_Periph) FT2() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT2)}}
}

func (p *EXTI_Periph) FT3() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT3)}}
}

func (p *EXTI_Periph) FT4() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT4)}}
}

func (p *EXTI_Periph) FT5() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT5)}}
}

func (p *EXTI_Periph) FT6() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT6)}}
}

func (p *EXTI_Periph) FT7() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT7)}}
}

func (p *EXTI_Periph) FT8() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT8)}}
}

func (p *EXTI_Periph) FT9() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT9)}}
}

func (p *EXTI_Periph) FT10() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT10)}}
}

func (p *EXTI_Periph) FT11() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT11)}}
}

func (p *EXTI_Periph) FT12() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT12)}}
}

func (p *EXTI_Periph) FT13() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT13)}}
}

func (p *EXTI_Periph) FT14() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT14)}}
}

func (p *EXTI_Periph) FT15() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT15)}}
}

func (p *EXTI_Periph) FT16() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT16)}}
}

func (p *EXTI_Periph) FT18() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT18)}}
}

func (p *EXTI_Periph) FT19() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT19)}}
}

func (p *EXTI_Periph) FT20() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT20)}}
}

func (p *EXTI_Periph) FT21() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT21)}}
}

func (p *EXTI_Periph) FT22() FTSR1_Mask {
	return FTSR1_Mask{mmio.UM32{&p.FTSR1.U32, uint32(FT22)}}
}

type SWIER1_Bits uint32

func (b SWIER1_Bits) Field(mask SWIER1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWIER1_Bits) J(v int) SWIER1_Bits {
	return SWIER1_Bits(bits.Make32(v, uint32(mask)))
}

type SWIER1 struct{ mmio.U32 }

func (r *SWIER1) Bits(mask SWIER1_Bits) SWIER1_Bits { return SWIER1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SWIER1) StoreBits(mask, b SWIER1_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SWIER1) SetBits(mask SWIER1_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *SWIER1) ClearBits(mask SWIER1_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *SWIER1) Load() SWIER1_Bits                 { return SWIER1_Bits(r.U32.Load()) }
func (r *SWIER1) Store(b SWIER1_Bits)               { r.U32.Store(uint32(b)) }

type SWIER1_Mask struct{ mmio.UM32 }

func (rm SWIER1_Mask) Load() SWIER1_Bits   { return SWIER1_Bits(rm.UM32.Load()) }
func (rm SWIER1_Mask) Store(b SWIER1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) SWI0() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI0)}}
}

func (p *EXTI_Periph) SWI1() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI1)}}
}

func (p *EXTI_Periph) SWI2() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI2)}}
}

func (p *EXTI_Periph) SWI3() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI3)}}
}

func (p *EXTI_Periph) SWI4() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI4)}}
}

func (p *EXTI_Periph) SWI5() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI5)}}
}

func (p *EXTI_Periph) SWI6() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI6)}}
}

func (p *EXTI_Periph) SWI7() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI7)}}
}

func (p *EXTI_Periph) SWI8() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI8)}}
}

func (p *EXTI_Periph) SWI9() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI9)}}
}

func (p *EXTI_Periph) SWI10() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI10)}}
}

func (p *EXTI_Periph) SWI11() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI11)}}
}

func (p *EXTI_Periph) SWI12() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI12)}}
}

func (p *EXTI_Periph) SWI13() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI13)}}
}

func (p *EXTI_Periph) SWI14() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI14)}}
}

func (p *EXTI_Periph) SWI15() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI15)}}
}

func (p *EXTI_Periph) SWI16() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI16)}}
}

func (p *EXTI_Periph) SWI18() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI18)}}
}

func (p *EXTI_Periph) SWI19() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI19)}}
}

func (p *EXTI_Periph) SWI20() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI20)}}
}

func (p *EXTI_Periph) SWI21() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI21)}}
}

func (p *EXTI_Periph) SWI22() SWIER1_Mask {
	return SWIER1_Mask{mmio.UM32{&p.SWIER1.U32, uint32(SWI22)}}
}

type PR1_Bits uint32

func (b PR1_Bits) Field(mask PR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PR1_Bits) J(v int) PR1_Bits {
	return PR1_Bits(bits.Make32(v, uint32(mask)))
}

type PR1 struct{ mmio.U32 }

func (r *PR1) Bits(mask PR1_Bits) PR1_Bits { return PR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *PR1) StoreBits(mask, b PR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PR1) SetBits(mask PR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *PR1) ClearBits(mask PR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *PR1) Load() PR1_Bits              { return PR1_Bits(r.U32.Load()) }
func (r *PR1) Store(b PR1_Bits)            { r.U32.Store(uint32(b)) }

type PR1_Mask struct{ mmio.UM32 }

func (rm PR1_Mask) Load() PR1_Bits   { return PR1_Bits(rm.UM32.Load()) }
func (rm PR1_Mask) Store(b PR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) PIF0() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF0)}}
}

func (p *EXTI_Periph) PIF1() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF1)}}
}

func (p *EXTI_Periph) PIF2() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF2)}}
}

func (p *EXTI_Periph) PIF3() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF3)}}
}

func (p *EXTI_Periph) PIF4() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF4)}}
}

func (p *EXTI_Periph) PIF5() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF5)}}
}

func (p *EXTI_Periph) PIF6() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF6)}}
}

func (p *EXTI_Periph) PIF7() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF7)}}
}

func (p *EXTI_Periph) PIF8() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF8)}}
}

func (p *EXTI_Periph) PIF9() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF9)}}
}

func (p *EXTI_Periph) PIF10() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF10)}}
}

func (p *EXTI_Periph) PIF11() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF11)}}
}

func (p *EXTI_Periph) PIF12() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF12)}}
}

func (p *EXTI_Periph) PIF13() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF13)}}
}

func (p *EXTI_Periph) PIF14() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF14)}}
}

func (p *EXTI_Periph) PIF15() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF15)}}
}

func (p *EXTI_Periph) PIF16() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF16)}}
}

func (p *EXTI_Periph) PIF18() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF18)}}
}

func (p *EXTI_Periph) PIF19() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF19)}}
}

func (p *EXTI_Periph) PIF20() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF20)}}
}

func (p *EXTI_Periph) PIF21() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF21)}}
}

func (p *EXTI_Periph) PIF22() PR1_Mask {
	return PR1_Mask{mmio.UM32{&p.PR1.U32, uint32(PIF22)}}
}

type IMR2_Bits uint32

func (b IMR2_Bits) Field(mask IMR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMR2_Bits) J(v int) IMR2_Bits {
	return IMR2_Bits(bits.Make32(v, uint32(mask)))
}

type IMR2 struct{ mmio.U32 }

func (r *IMR2) Bits(mask IMR2_Bits) IMR2_Bits { return IMR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMR2) StoreBits(mask, b IMR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMR2) SetBits(mask IMR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *IMR2) ClearBits(mask IMR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *IMR2) Load() IMR2_Bits               { return IMR2_Bits(r.U32.Load()) }
func (r *IMR2) Store(b IMR2_Bits)             { r.U32.Store(uint32(b)) }

type IMR2_Mask struct{ mmio.UM32 }

func (rm IMR2_Mask) Load() IMR2_Bits   { return IMR2_Bits(rm.UM32.Load()) }
func (rm IMR2_Mask) Store(b IMR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) IM32() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM32)}}
}

func (p *EXTI_Periph) IM33() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM33)}}
}

func (p *EXTI_Periph) IM34() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM34)}}
}

func (p *EXTI_Periph) IM35() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM35)}}
}

func (p *EXTI_Periph) IM36() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM36)}}
}

func (p *EXTI_Periph) IM37() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM37)}}
}

func (p *EXTI_Periph) IM38() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM38)}}
}

func (p *EXTI_Periph) IM39() IMR2_Mask {
	return IMR2_Mask{mmio.UM32{&p.IMR2.U32, uint32(IM39)}}
}

type EMR2_Bits uint32

func (b EMR2_Bits) Field(mask EMR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EMR2_Bits) J(v int) EMR2_Bits {
	return EMR2_Bits(bits.Make32(v, uint32(mask)))
}

type EMR2 struct{ mmio.U32 }

func (r *EMR2) Bits(mask EMR2_Bits) EMR2_Bits { return EMR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *EMR2) StoreBits(mask, b EMR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EMR2) SetBits(mask EMR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *EMR2) ClearBits(mask EMR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *EMR2) Load() EMR2_Bits               { return EMR2_Bits(r.U32.Load()) }
func (r *EMR2) Store(b EMR2_Bits)             { r.U32.Store(uint32(b)) }

type EMR2_Mask struct{ mmio.UM32 }

func (rm EMR2_Mask) Load() EMR2_Bits   { return EMR2_Bits(rm.UM32.Load()) }
func (rm EMR2_Mask) Store(b EMR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) EM32() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM32)}}
}

func (p *EXTI_Periph) EM33() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM33)}}
}

func (p *EXTI_Periph) EM34() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM34)}}
}

func (p *EXTI_Periph) EM35() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM35)}}
}

func (p *EXTI_Periph) EM36() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM36)}}
}

func (p *EXTI_Periph) EM37() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM37)}}
}

func (p *EXTI_Periph) EM38() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM38)}}
}

func (p *EXTI_Periph) EM39() EMR2_Mask {
	return EMR2_Mask{mmio.UM32{&p.EMR2.U32, uint32(EM39)}}
}

type RTSR2_Bits uint32

func (b RTSR2_Bits) Field(mask RTSR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RTSR2_Bits) J(v int) RTSR2_Bits {
	return RTSR2_Bits(bits.Make32(v, uint32(mask)))
}

type RTSR2 struct{ mmio.U32 }

func (r *RTSR2) Bits(mask RTSR2_Bits) RTSR2_Bits { return RTSR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *RTSR2) StoreBits(mask, b RTSR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTSR2) SetBits(mask RTSR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *RTSR2) ClearBits(mask RTSR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *RTSR2) Load() RTSR2_Bits                { return RTSR2_Bits(r.U32.Load()) }
func (r *RTSR2) Store(b RTSR2_Bits)              { r.U32.Store(uint32(b)) }

type RTSR2_Mask struct{ mmio.UM32 }

func (rm RTSR2_Mask) Load() RTSR2_Bits   { return RTSR2_Bits(rm.UM32.Load()) }
func (rm RTSR2_Mask) Store(b RTSR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) RT35() RTSR2_Mask {
	return RTSR2_Mask{mmio.UM32{&p.RTSR2.U32, uint32(RT35)}}
}

func (p *EXTI_Periph) RT36() RTSR2_Mask {
	return RTSR2_Mask{mmio.UM32{&p.RTSR2.U32, uint32(RT36)}}
}

func (p *EXTI_Periph) RT37() RTSR2_Mask {
	return RTSR2_Mask{mmio.UM32{&p.RTSR2.U32, uint32(RT37)}}
}

func (p *EXTI_Periph) RT38() RTSR2_Mask {
	return RTSR2_Mask{mmio.UM32{&p.RTSR2.U32, uint32(RT38)}}
}

type FTSR2_Bits uint32

func (b FTSR2_Bits) Field(mask FTSR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FTSR2_Bits) J(v int) FTSR2_Bits {
	return FTSR2_Bits(bits.Make32(v, uint32(mask)))
}

type FTSR2 struct{ mmio.U32 }

func (r *FTSR2) Bits(mask FTSR2_Bits) FTSR2_Bits { return FTSR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *FTSR2) StoreBits(mask, b FTSR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FTSR2) SetBits(mask FTSR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *FTSR2) ClearBits(mask FTSR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *FTSR2) Load() FTSR2_Bits                { return FTSR2_Bits(r.U32.Load()) }
func (r *FTSR2) Store(b FTSR2_Bits)              { r.U32.Store(uint32(b)) }

type FTSR2_Mask struct{ mmio.UM32 }

func (rm FTSR2_Mask) Load() FTSR2_Bits   { return FTSR2_Bits(rm.UM32.Load()) }
func (rm FTSR2_Mask) Store(b FTSR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) FT35() FTSR2_Mask {
	return FTSR2_Mask{mmio.UM32{&p.FTSR2.U32, uint32(FT35)}}
}

func (p *EXTI_Periph) FT36() FTSR2_Mask {
	return FTSR2_Mask{mmio.UM32{&p.FTSR2.U32, uint32(FT36)}}
}

func (p *EXTI_Periph) FT37() FTSR2_Mask {
	return FTSR2_Mask{mmio.UM32{&p.FTSR2.U32, uint32(FT37)}}
}

func (p *EXTI_Periph) FT38() FTSR2_Mask {
	return FTSR2_Mask{mmio.UM32{&p.FTSR2.U32, uint32(FT38)}}
}

type SWIER2_Bits uint32

func (b SWIER2_Bits) Field(mask SWIER2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWIER2_Bits) J(v int) SWIER2_Bits {
	return SWIER2_Bits(bits.Make32(v, uint32(mask)))
}

type SWIER2 struct{ mmio.U32 }

func (r *SWIER2) Bits(mask SWIER2_Bits) SWIER2_Bits { return SWIER2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SWIER2) StoreBits(mask, b SWIER2_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SWIER2) SetBits(mask SWIER2_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *SWIER2) ClearBits(mask SWIER2_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *SWIER2) Load() SWIER2_Bits                 { return SWIER2_Bits(r.U32.Load()) }
func (r *SWIER2) Store(b SWIER2_Bits)               { r.U32.Store(uint32(b)) }

type SWIER2_Mask struct{ mmio.UM32 }

func (rm SWIER2_Mask) Load() SWIER2_Bits   { return SWIER2_Bits(rm.UM32.Load()) }
func (rm SWIER2_Mask) Store(b SWIER2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) SWI35() SWIER2_Mask {
	return SWIER2_Mask{mmio.UM32{&p.SWIER2.U32, uint32(SWI35)}}
}

func (p *EXTI_Periph) SWI36() SWIER2_Mask {
	return SWIER2_Mask{mmio.UM32{&p.SWIER2.U32, uint32(SWI36)}}
}

func (p *EXTI_Periph) SWI37() SWIER2_Mask {
	return SWIER2_Mask{mmio.UM32{&p.SWIER2.U32, uint32(SWI37)}}
}

func (p *EXTI_Periph) SWI38() SWIER2_Mask {
	return SWIER2_Mask{mmio.UM32{&p.SWIER2.U32, uint32(SWI38)}}
}

type PR2_Bits uint32

func (b PR2_Bits) Field(mask PR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PR2_Bits) J(v int) PR2_Bits {
	return PR2_Bits(bits.Make32(v, uint32(mask)))
}

type PR2 struct{ mmio.U32 }

func (r *PR2) Bits(mask PR2_Bits) PR2_Bits { return PR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *PR2) StoreBits(mask, b PR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PR2) SetBits(mask PR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *PR2) ClearBits(mask PR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *PR2) Load() PR2_Bits              { return PR2_Bits(r.U32.Load()) }
func (r *PR2) Store(b PR2_Bits)            { r.U32.Store(uint32(b)) }

type PR2_Mask struct{ mmio.UM32 }

func (rm PR2_Mask) Load() PR2_Bits   { return PR2_Bits(rm.UM32.Load()) }
func (rm PR2_Mask) Store(b PR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) PIF35() PR2_Mask {
	return PR2_Mask{mmio.UM32{&p.PR2.U32, uint32(PIF35)}}
}

func (p *EXTI_Periph) PIF36() PR2_Mask {
	return PR2_Mask{mmio.UM32{&p.PR2.U32, uint32(PIF36)}}
}

func (p *EXTI_Periph) PIF37() PR2_Mask {
	return PR2_Mask{mmio.UM32{&p.PR2.U32, uint32(PIF37)}}
}

func (p *EXTI_Periph) PIF38() PR2_Mask {
	return PR2_Mask{mmio.UM32{&p.PR2.U32, uint32(PIF38)}}
}
