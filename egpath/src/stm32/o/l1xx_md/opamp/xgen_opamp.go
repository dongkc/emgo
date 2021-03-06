package opamp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type OPAMP_Periph struct {
	CSR   CSR
	OTR   OTR
	LPOTR LPOTR
}

func (p *OPAMP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OPAMP = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP_BASE)))

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *OPAMP_Periph) OPA1PD() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA1PD)}}
}

func (p *OPAMP_Periph) S3SEL1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S3SEL1)}}
}

func (p *OPAMP_Periph) S4SEL1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S4SEL1)}}
}

func (p *OPAMP_Periph) S5SEL1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S5SEL1)}}
}

func (p *OPAMP_Periph) S6SEL1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S6SEL1)}}
}

func (p *OPAMP_Periph) OPA1CAL_L() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA1CAL_L)}}
}

func (p *OPAMP_Periph) OPA1CAL_H() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA1CAL_H)}}
}

func (p *OPAMP_Periph) OPA1LPM() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA1LPM)}}
}

func (p *OPAMP_Periph) OPA2PD() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA2PD)}}
}

func (p *OPAMP_Periph) S3SEL2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S3SEL2)}}
}

func (p *OPAMP_Periph) S4SEL2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S4SEL2)}}
}

func (p *OPAMP_Periph) S5SEL2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S5SEL2)}}
}

func (p *OPAMP_Periph) S6SEL2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S6SEL2)}}
}

func (p *OPAMP_Periph) OPA2CAL_L() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA2CAL_L)}}
}

func (p *OPAMP_Periph) OPA2CAL_H() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA2CAL_H)}}
}

func (p *OPAMP_Periph) OPA2LPM() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA2LPM)}}
}

func (p *OPAMP_Periph) OPA3PD() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA3PD)}}
}

func (p *OPAMP_Periph) S3SEL3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S3SEL3)}}
}

func (p *OPAMP_Periph) S4SEL3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S4SEL3)}}
}

func (p *OPAMP_Periph) S5SEL3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S5SEL3)}}
}

func (p *OPAMP_Periph) S6SEL3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S6SEL3)}}
}

func (p *OPAMP_Periph) OPA3CAL_L() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA3CAL_L)}}
}

func (p *OPAMP_Periph) OPA3CAL_H() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA3CAL_H)}}
}

func (p *OPAMP_Periph) OPA3LPM() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA3LPM)}}
}

func (p *OPAMP_Periph) ANAWSEL1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ANAWSEL1)}}
}

func (p *OPAMP_Periph) ANAWSEL2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ANAWSEL2)}}
}

func (p *OPAMP_Periph) ANAWSEL3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(ANAWSEL3)}}
}

func (p *OPAMP_Periph) S7SEL2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(S7SEL2)}}
}

func (p *OPAMP_Periph) AOP_RANGE() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AOP_RANGE)}}
}

func (p *OPAMP_Periph) OPA1CALOUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA1CALOUT)}}
}

func (p *OPAMP_Periph) OPA2CALOUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA2CALOUT)}}
}

func (p *OPAMP_Periph) OPA3CALOUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPA3CALOUT)}}
}

type OTR_Bits uint32

func (b OTR_Bits) Field(mask OTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTR_Bits) J(v int) OTR_Bits {
	return OTR_Bits(bits.Make32(v, uint32(mask)))
}

type OTR struct{ mmio.U32 }

func (r *OTR) Bits(mask OTR_Bits) OTR_Bits { return OTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OTR) StoreBits(mask, b OTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OTR) SetBits(mask OTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OTR) ClearBits(mask OTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OTR) Load() OTR_Bits              { return OTR_Bits(r.U32.Load()) }
func (r *OTR) Store(b OTR_Bits)            { r.U32.Store(uint32(b)) }

type OTR_Mask struct{ mmio.UM32 }

func (rm OTR_Mask) Load() OTR_Bits   { return OTR_Bits(rm.UM32.Load()) }
func (rm OTR_Mask) Store(b OTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *OPAMP_Periph) AO1_OPT_OFFSET_TRIM() OTR_Mask {
	return OTR_Mask{mmio.UM32{&p.OTR.U32, uint32(AO1_OPT_OFFSET_TRIM)}}
}

func (p *OPAMP_Periph) AO2_OPT_OFFSET_TRIM() OTR_Mask {
	return OTR_Mask{mmio.UM32{&p.OTR.U32, uint32(AO2_OPT_OFFSET_TRIM)}}
}

func (p *OPAMP_Periph) AO3_OPT_OFFSET_TRIM() OTR_Mask {
	return OTR_Mask{mmio.UM32{&p.OTR.U32, uint32(AO3_OPT_OFFSET_TRIM)}}
}

func (p *OPAMP_Periph) OT_USER() OTR_Mask {
	return OTR_Mask{mmio.UM32{&p.OTR.U32, uint32(OT_USER)}}
}

type LPOTR_Bits uint32

func (b LPOTR_Bits) Field(mask LPOTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LPOTR_Bits) J(v int) LPOTR_Bits {
	return LPOTR_Bits(bits.Make32(v, uint32(mask)))
}

type LPOTR struct{ mmio.U32 }

func (r *LPOTR) Bits(mask LPOTR_Bits) LPOTR_Bits { return LPOTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LPOTR) StoreBits(mask, b LPOTR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LPOTR) SetBits(mask LPOTR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *LPOTR) ClearBits(mask LPOTR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *LPOTR) Load() LPOTR_Bits                { return LPOTR_Bits(r.U32.Load()) }
func (r *LPOTR) Store(b LPOTR_Bits)              { r.U32.Store(uint32(b)) }

type LPOTR_Mask struct{ mmio.UM32 }

func (rm LPOTR_Mask) Load() LPOTR_Bits   { return LPOTR_Bits(rm.UM32.Load()) }
func (rm LPOTR_Mask) Store(b LPOTR_Bits) { rm.UM32.Store(uint32(b)) }
