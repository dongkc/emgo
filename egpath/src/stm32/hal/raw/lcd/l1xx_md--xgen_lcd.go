// +build l1xx_md

package lcd

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type LCD_Periph struct {
	CR  CR
	FCR FCR
	SR  SR
	CLR CLR
	_   uint32
	RAM [16]RAM
}

func (p *LCD_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var LCD = (*LCD_Periph)(unsafe.Pointer(uintptr(mmap.LCD_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) LCDEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(LCDEN)}}
}

func (p *LCD_Periph) VSEL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(VSEL)}}
}

func (p *LCD_Periph) DUTY() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DUTY)}}
}

func (p *LCD_Periph) BIAS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(BIAS)}}
}

func (p *LCD_Periph) MUX_SEG() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(MUX_SEG)}}
}

type FCR_Bits uint32

func (b FCR_Bits) Field(mask FCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FCR_Bits) J(v int) FCR_Bits {
	return FCR_Bits(bits.Make32(v, uint32(mask)))
}

type FCR struct{ mmio.U32 }

func (r *FCR) Bits(mask FCR_Bits) FCR_Bits { return FCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FCR) StoreBits(mask, b FCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FCR) SetBits(mask FCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *FCR) ClearBits(mask FCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *FCR) Load() FCR_Bits              { return FCR_Bits(r.U32.Load()) }
func (r *FCR) Store(b FCR_Bits)            { r.U32.Store(uint32(b)) }

type FCR_Mask struct{ mmio.UM32 }

func (rm FCR_Mask) Load() FCR_Bits   { return FCR_Bits(rm.UM32.Load()) }
func (rm FCR_Mask) Store(b FCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) HD() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(HD)}}
}

func (p *LCD_Periph) SOFIE() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(SOFIE)}}
}

func (p *LCD_Periph) UDDIE() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(UDDIE)}}
}

func (p *LCD_Periph) PON() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(PON)}}
}

func (p *LCD_Periph) DEAD() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(DEAD)}}
}

func (p *LCD_Periph) CC() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(CC)}}
}

func (p *LCD_Periph) BLINKF() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(BLINKF)}}
}

func (p *LCD_Periph) BLINK() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(BLINK)}}
}

func (p *LCD_Periph) DIV() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(DIV)}}
}

func (p *LCD_Periph) PS() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(PS)}}
}

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) ENS() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(ENS)}}
}

func (p *LCD_Periph) SOF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(SOF)}}
}

func (p *LCD_Periph) UDR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(UDR)}}
}

func (p *LCD_Periph) UDD() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(UDD)}}
}

func (p *LCD_Periph) RDY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(RDY)}}
}

func (p *LCD_Periph) FCRSR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FCRSR)}}
}

type CLR_Bits uint32

func (b CLR_Bits) Field(mask CLR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLR_Bits) J(v int) CLR_Bits {
	return CLR_Bits(bits.Make32(v, uint32(mask)))
}

type CLR struct{ mmio.U32 }

func (r *CLR) Bits(mask CLR_Bits) CLR_Bits { return CLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CLR) StoreBits(mask, b CLR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CLR) SetBits(mask CLR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CLR) ClearBits(mask CLR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CLR) Load() CLR_Bits              { return CLR_Bits(r.U32.Load()) }
func (r *CLR) Store(b CLR_Bits)            { r.U32.Store(uint32(b)) }

type CLR_Mask struct{ mmio.UM32 }

func (rm CLR_Mask) Load() CLR_Bits   { return CLR_Bits(rm.UM32.Load()) }
func (rm CLR_Mask) Store(b CLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) SOFC() CLR_Mask {
	return CLR_Mask{mmio.UM32{&p.CLR.U32, uint32(SOFC)}}
}

func (p *LCD_Periph) UDDC() CLR_Mask {
	return CLR_Mask{mmio.UM32{&p.CLR.U32, uint32(UDDC)}}
}

type RAM_Bits uint32

func (b RAM_Bits) Field(mask RAM_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RAM_Bits) J(v int) RAM_Bits {
	return RAM_Bits(bits.Make32(v, uint32(mask)))
}

type RAM struct{ mmio.U32 }

func (r *RAM) Bits(mask RAM_Bits) RAM_Bits { return RAM_Bits(r.U32.Bits(uint32(mask))) }
func (r *RAM) StoreBits(mask, b RAM_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAM) SetBits(mask RAM_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *RAM) ClearBits(mask RAM_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAM) Load() RAM_Bits              { return RAM_Bits(r.U32.Load()) }
func (r *RAM) Store(b RAM_Bits)            { r.U32.Store(uint32(b)) }

type RAM_Mask struct{ mmio.UM32 }

func (rm RAM_Mask) Load() RAM_Bits   { return RAM_Bits(rm.UM32.Load()) }
func (rm RAM_Mask) Store(b RAM_Bits) { rm.UM32.Store(uint32(b)) }

func (p *LCD_Periph) SEGMENT_DATA(n int) RAM_Mask {
	return RAM_Mask{mmio.UM32{&p.RAM[n].U32, uint32(SEGMENT_DATA)}}
}
