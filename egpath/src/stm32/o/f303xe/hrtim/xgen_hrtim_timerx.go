package hrtim

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type HRTIM_Timerx_Periph struct {
	TIMxCR   TIMxCR
	TIMxISR  TIMxISR
	TIMxICR  TIMxICR
	TIMxDIER TIMxDIER
	CNTxR    CNTxR
	PERxR    PERxR
	REPxR    REPxR
	CMP1xR   CMP1xR
	CMP1CxR  CMP1CxR
	CMP2xR   CMP2xR
	CMP3xR   CMP3xR
	CMP4xR   CMP4xR
	CPT1xR   CPT1xR
	CPT2xR   CPT2xR
	DTxR     DTxR
	SETx1R   SETx1R
	RSTx1R   RSTx1R
	SETx2R   SETx2R
	RSTx2R   RSTx2R
	EEFxR1   EEFxR1
	EEFxR2   EEFxR2
	RSTxR    RSTxR
	CHPxR    CHPxR
	CPT1xCR  CPT1xCR
	CPT2xCR  CPT2xCR
	OUTxR    OUTxR
	FLTxR    FLTxR
}

func (p *HRTIM_Timerx_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TIMxCR_Bits uint32

func (b TIMxCR_Bits) Field(mask TIMxCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMxCR_Bits) J(v int) TIMxCR_Bits {
	return TIMxCR_Bits(bits.Make32(v, uint32(mask)))
}

type TIMxCR struct{ mmio.U32 }

func (r *TIMxCR) Bits(mask TIMxCR_Bits) TIMxCR_Bits { return TIMxCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TIMxCR) StoreBits(mask, b TIMxCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIMxCR) SetBits(mask TIMxCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *TIMxCR) ClearBits(mask TIMxCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *TIMxCR) Load() TIMxCR_Bits                 { return TIMxCR_Bits(r.U32.Load()) }
func (r *TIMxCR) Store(b TIMxCR_Bits)               { r.U32.Store(uint32(b)) }

type TIMxCR_Mask struct{ mmio.UM32 }

func (rm TIMxCR_Mask) Load() TIMxCR_Bits   { return TIMxCR_Bits(rm.UM32.Load()) }
func (rm TIMxCR_Mask) Store(b TIMxCR_Bits) { rm.UM32.Store(uint32(b)) }

type TIMxISR_Bits uint32

func (b TIMxISR_Bits) Field(mask TIMxISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMxISR_Bits) J(v int) TIMxISR_Bits {
	return TIMxISR_Bits(bits.Make32(v, uint32(mask)))
}

type TIMxISR struct{ mmio.U32 }

func (r *TIMxISR) Bits(mask TIMxISR_Bits) TIMxISR_Bits { return TIMxISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TIMxISR) StoreBits(mask, b TIMxISR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIMxISR) SetBits(mask TIMxISR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *TIMxISR) ClearBits(mask TIMxISR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *TIMxISR) Load() TIMxISR_Bits                  { return TIMxISR_Bits(r.U32.Load()) }
func (r *TIMxISR) Store(b TIMxISR_Bits)                { r.U32.Store(uint32(b)) }

type TIMxISR_Mask struct{ mmio.UM32 }

func (rm TIMxISR_Mask) Load() TIMxISR_Bits   { return TIMxISR_Bits(rm.UM32.Load()) }
func (rm TIMxISR_Mask) Store(b TIMxISR_Bits) { rm.UM32.Store(uint32(b)) }

type TIMxICR_Bits uint32

func (b TIMxICR_Bits) Field(mask TIMxICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMxICR_Bits) J(v int) TIMxICR_Bits {
	return TIMxICR_Bits(bits.Make32(v, uint32(mask)))
}

type TIMxICR struct{ mmio.U32 }

func (r *TIMxICR) Bits(mask TIMxICR_Bits) TIMxICR_Bits { return TIMxICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TIMxICR) StoreBits(mask, b TIMxICR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIMxICR) SetBits(mask TIMxICR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *TIMxICR) ClearBits(mask TIMxICR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *TIMxICR) Load() TIMxICR_Bits                  { return TIMxICR_Bits(r.U32.Load()) }
func (r *TIMxICR) Store(b TIMxICR_Bits)                { r.U32.Store(uint32(b)) }

type TIMxICR_Mask struct{ mmio.UM32 }

func (rm TIMxICR_Mask) Load() TIMxICR_Bits   { return TIMxICR_Bits(rm.UM32.Load()) }
func (rm TIMxICR_Mask) Store(b TIMxICR_Bits) { rm.UM32.Store(uint32(b)) }

type TIMxDIER_Bits uint32

func (b TIMxDIER_Bits) Field(mask TIMxDIER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMxDIER_Bits) J(v int) TIMxDIER_Bits {
	return TIMxDIER_Bits(bits.Make32(v, uint32(mask)))
}

type TIMxDIER struct{ mmio.U32 }

func (r *TIMxDIER) Bits(mask TIMxDIER_Bits) TIMxDIER_Bits {
	return TIMxDIER_Bits(r.U32.Bits(uint32(mask)))
}
func (r *TIMxDIER) StoreBits(mask, b TIMxDIER_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIMxDIER) SetBits(mask TIMxDIER_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *TIMxDIER) ClearBits(mask TIMxDIER_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *TIMxDIER) Load() TIMxDIER_Bits             { return TIMxDIER_Bits(r.U32.Load()) }
func (r *TIMxDIER) Store(b TIMxDIER_Bits)           { r.U32.Store(uint32(b)) }

type TIMxDIER_Mask struct{ mmio.UM32 }

func (rm TIMxDIER_Mask) Load() TIMxDIER_Bits   { return TIMxDIER_Bits(rm.UM32.Load()) }
func (rm TIMxDIER_Mask) Store(b TIMxDIER_Bits) { rm.UM32.Store(uint32(b)) }

type CNTxR_Bits uint32

func (b CNTxR_Bits) Field(mask CNTxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CNTxR_Bits) J(v int) CNTxR_Bits {
	return CNTxR_Bits(bits.Make32(v, uint32(mask)))
}

type CNTxR struct{ mmio.U32 }

func (r *CNTxR) Bits(mask CNTxR_Bits) CNTxR_Bits { return CNTxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CNTxR) StoreBits(mask, b CNTxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CNTxR) SetBits(mask CNTxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CNTxR) ClearBits(mask CNTxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CNTxR) Load() CNTxR_Bits                { return CNTxR_Bits(r.U32.Load()) }
func (r *CNTxR) Store(b CNTxR_Bits)              { r.U32.Store(uint32(b)) }

type CNTxR_Mask struct{ mmio.UM32 }

func (rm CNTxR_Mask) Load() CNTxR_Bits   { return CNTxR_Bits(rm.UM32.Load()) }
func (rm CNTxR_Mask) Store(b CNTxR_Bits) { rm.UM32.Store(uint32(b)) }

type PERxR_Bits uint32

func (b PERxR_Bits) Field(mask PERxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PERxR_Bits) J(v int) PERxR_Bits {
	return PERxR_Bits(bits.Make32(v, uint32(mask)))
}

type PERxR struct{ mmio.U32 }

func (r *PERxR) Bits(mask PERxR_Bits) PERxR_Bits { return PERxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PERxR) StoreBits(mask, b PERxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PERxR) SetBits(mask PERxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PERxR) ClearBits(mask PERxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PERxR) Load() PERxR_Bits                { return PERxR_Bits(r.U32.Load()) }
func (r *PERxR) Store(b PERxR_Bits)              { r.U32.Store(uint32(b)) }

type PERxR_Mask struct{ mmio.UM32 }

func (rm PERxR_Mask) Load() PERxR_Bits   { return PERxR_Bits(rm.UM32.Load()) }
func (rm PERxR_Mask) Store(b PERxR_Bits) { rm.UM32.Store(uint32(b)) }

type REPxR_Bits uint32

func (b REPxR_Bits) Field(mask REPxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask REPxR_Bits) J(v int) REPxR_Bits {
	return REPxR_Bits(bits.Make32(v, uint32(mask)))
}

type REPxR struct{ mmio.U32 }

func (r *REPxR) Bits(mask REPxR_Bits) REPxR_Bits { return REPxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *REPxR) StoreBits(mask, b REPxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REPxR) SetBits(mask REPxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *REPxR) ClearBits(mask REPxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *REPxR) Load() REPxR_Bits                { return REPxR_Bits(r.U32.Load()) }
func (r *REPxR) Store(b REPxR_Bits)              { r.U32.Store(uint32(b)) }

type REPxR_Mask struct{ mmio.UM32 }

func (rm REPxR_Mask) Load() REPxR_Bits   { return REPxR_Bits(rm.UM32.Load()) }
func (rm REPxR_Mask) Store(b REPxR_Bits) { rm.UM32.Store(uint32(b)) }

type CMP1xR_Bits uint32

func (b CMP1xR_Bits) Field(mask CMP1xR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMP1xR_Bits) J(v int) CMP1xR_Bits {
	return CMP1xR_Bits(bits.Make32(v, uint32(mask)))
}

type CMP1xR struct{ mmio.U32 }

func (r *CMP1xR) Bits(mask CMP1xR_Bits) CMP1xR_Bits { return CMP1xR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMP1xR) StoreBits(mask, b CMP1xR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMP1xR) SetBits(mask CMP1xR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CMP1xR) ClearBits(mask CMP1xR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CMP1xR) Load() CMP1xR_Bits                 { return CMP1xR_Bits(r.U32.Load()) }
func (r *CMP1xR) Store(b CMP1xR_Bits)               { r.U32.Store(uint32(b)) }

type CMP1xR_Mask struct{ mmio.UM32 }

func (rm CMP1xR_Mask) Load() CMP1xR_Bits   { return CMP1xR_Bits(rm.UM32.Load()) }
func (rm CMP1xR_Mask) Store(b CMP1xR_Bits) { rm.UM32.Store(uint32(b)) }

type CMP1CxR_Bits uint32

func (b CMP1CxR_Bits) Field(mask CMP1CxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMP1CxR_Bits) J(v int) CMP1CxR_Bits {
	return CMP1CxR_Bits(bits.Make32(v, uint32(mask)))
}

type CMP1CxR struct{ mmio.U32 }

func (r *CMP1CxR) Bits(mask CMP1CxR_Bits) CMP1CxR_Bits { return CMP1CxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMP1CxR) StoreBits(mask, b CMP1CxR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMP1CxR) SetBits(mask CMP1CxR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CMP1CxR) ClearBits(mask CMP1CxR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CMP1CxR) Load() CMP1CxR_Bits                  { return CMP1CxR_Bits(r.U32.Load()) }
func (r *CMP1CxR) Store(b CMP1CxR_Bits)                { r.U32.Store(uint32(b)) }

type CMP1CxR_Mask struct{ mmio.UM32 }

func (rm CMP1CxR_Mask) Load() CMP1CxR_Bits   { return CMP1CxR_Bits(rm.UM32.Load()) }
func (rm CMP1CxR_Mask) Store(b CMP1CxR_Bits) { rm.UM32.Store(uint32(b)) }

type CMP2xR_Bits uint32

func (b CMP2xR_Bits) Field(mask CMP2xR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMP2xR_Bits) J(v int) CMP2xR_Bits {
	return CMP2xR_Bits(bits.Make32(v, uint32(mask)))
}

type CMP2xR struct{ mmio.U32 }

func (r *CMP2xR) Bits(mask CMP2xR_Bits) CMP2xR_Bits { return CMP2xR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMP2xR) StoreBits(mask, b CMP2xR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMP2xR) SetBits(mask CMP2xR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CMP2xR) ClearBits(mask CMP2xR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CMP2xR) Load() CMP2xR_Bits                 { return CMP2xR_Bits(r.U32.Load()) }
func (r *CMP2xR) Store(b CMP2xR_Bits)               { r.U32.Store(uint32(b)) }

type CMP2xR_Mask struct{ mmio.UM32 }

func (rm CMP2xR_Mask) Load() CMP2xR_Bits   { return CMP2xR_Bits(rm.UM32.Load()) }
func (rm CMP2xR_Mask) Store(b CMP2xR_Bits) { rm.UM32.Store(uint32(b)) }

type CMP3xR_Bits uint32

func (b CMP3xR_Bits) Field(mask CMP3xR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMP3xR_Bits) J(v int) CMP3xR_Bits {
	return CMP3xR_Bits(bits.Make32(v, uint32(mask)))
}

type CMP3xR struct{ mmio.U32 }

func (r *CMP3xR) Bits(mask CMP3xR_Bits) CMP3xR_Bits { return CMP3xR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMP3xR) StoreBits(mask, b CMP3xR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMP3xR) SetBits(mask CMP3xR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CMP3xR) ClearBits(mask CMP3xR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CMP3xR) Load() CMP3xR_Bits                 { return CMP3xR_Bits(r.U32.Load()) }
func (r *CMP3xR) Store(b CMP3xR_Bits)               { r.U32.Store(uint32(b)) }

type CMP3xR_Mask struct{ mmio.UM32 }

func (rm CMP3xR_Mask) Load() CMP3xR_Bits   { return CMP3xR_Bits(rm.UM32.Load()) }
func (rm CMP3xR_Mask) Store(b CMP3xR_Bits) { rm.UM32.Store(uint32(b)) }

type CMP4xR_Bits uint32

func (b CMP4xR_Bits) Field(mask CMP4xR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMP4xR_Bits) J(v int) CMP4xR_Bits {
	return CMP4xR_Bits(bits.Make32(v, uint32(mask)))
}

type CMP4xR struct{ mmio.U32 }

func (r *CMP4xR) Bits(mask CMP4xR_Bits) CMP4xR_Bits { return CMP4xR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMP4xR) StoreBits(mask, b CMP4xR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMP4xR) SetBits(mask CMP4xR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CMP4xR) ClearBits(mask CMP4xR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CMP4xR) Load() CMP4xR_Bits                 { return CMP4xR_Bits(r.U32.Load()) }
func (r *CMP4xR) Store(b CMP4xR_Bits)               { r.U32.Store(uint32(b)) }

type CMP4xR_Mask struct{ mmio.UM32 }

func (rm CMP4xR_Mask) Load() CMP4xR_Bits   { return CMP4xR_Bits(rm.UM32.Load()) }
func (rm CMP4xR_Mask) Store(b CMP4xR_Bits) { rm.UM32.Store(uint32(b)) }

type CPT1xR_Bits uint32

func (b CPT1xR_Bits) Field(mask CPT1xR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPT1xR_Bits) J(v int) CPT1xR_Bits {
	return CPT1xR_Bits(bits.Make32(v, uint32(mask)))
}

type CPT1xR struct{ mmio.U32 }

func (r *CPT1xR) Bits(mask CPT1xR_Bits) CPT1xR_Bits { return CPT1xR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CPT1xR) StoreBits(mask, b CPT1xR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CPT1xR) SetBits(mask CPT1xR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CPT1xR) ClearBits(mask CPT1xR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CPT1xR) Load() CPT1xR_Bits                 { return CPT1xR_Bits(r.U32.Load()) }
func (r *CPT1xR) Store(b CPT1xR_Bits)               { r.U32.Store(uint32(b)) }

type CPT1xR_Mask struct{ mmio.UM32 }

func (rm CPT1xR_Mask) Load() CPT1xR_Bits   { return CPT1xR_Bits(rm.UM32.Load()) }
func (rm CPT1xR_Mask) Store(b CPT1xR_Bits) { rm.UM32.Store(uint32(b)) }

type CPT2xR_Bits uint32

func (b CPT2xR_Bits) Field(mask CPT2xR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPT2xR_Bits) J(v int) CPT2xR_Bits {
	return CPT2xR_Bits(bits.Make32(v, uint32(mask)))
}

type CPT2xR struct{ mmio.U32 }

func (r *CPT2xR) Bits(mask CPT2xR_Bits) CPT2xR_Bits { return CPT2xR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CPT2xR) StoreBits(mask, b CPT2xR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CPT2xR) SetBits(mask CPT2xR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CPT2xR) ClearBits(mask CPT2xR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CPT2xR) Load() CPT2xR_Bits                 { return CPT2xR_Bits(r.U32.Load()) }
func (r *CPT2xR) Store(b CPT2xR_Bits)               { r.U32.Store(uint32(b)) }

type CPT2xR_Mask struct{ mmio.UM32 }

func (rm CPT2xR_Mask) Load() CPT2xR_Bits   { return CPT2xR_Bits(rm.UM32.Load()) }
func (rm CPT2xR_Mask) Store(b CPT2xR_Bits) { rm.UM32.Store(uint32(b)) }

type DTxR_Bits uint32

func (b DTxR_Bits) Field(mask DTxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTxR_Bits) J(v int) DTxR_Bits {
	return DTxR_Bits(bits.Make32(v, uint32(mask)))
}

type DTxR struct{ mmio.U32 }

func (r *DTxR) Bits(mask DTxR_Bits) DTxR_Bits { return DTxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DTxR) StoreBits(mask, b DTxR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DTxR) SetBits(mask DTxR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DTxR) ClearBits(mask DTxR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DTxR) Load() DTxR_Bits               { return DTxR_Bits(r.U32.Load()) }
func (r *DTxR) Store(b DTxR_Bits)             { r.U32.Store(uint32(b)) }

type DTxR_Mask struct{ mmio.UM32 }

func (rm DTxR_Mask) Load() DTxR_Bits   { return DTxR_Bits(rm.UM32.Load()) }
func (rm DTxR_Mask) Store(b DTxR_Bits) { rm.UM32.Store(uint32(b)) }

type SETx1R_Bits uint32

func (b SETx1R_Bits) Field(mask SETx1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SETx1R_Bits) J(v int) SETx1R_Bits {
	return SETx1R_Bits(bits.Make32(v, uint32(mask)))
}

type SETx1R struct{ mmio.U32 }

func (r *SETx1R) Bits(mask SETx1R_Bits) SETx1R_Bits { return SETx1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *SETx1R) StoreBits(mask, b SETx1R_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SETx1R) SetBits(mask SETx1R_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *SETx1R) ClearBits(mask SETx1R_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *SETx1R) Load() SETx1R_Bits                 { return SETx1R_Bits(r.U32.Load()) }
func (r *SETx1R) Store(b SETx1R_Bits)               { r.U32.Store(uint32(b)) }

type SETx1R_Mask struct{ mmio.UM32 }

func (rm SETx1R_Mask) Load() SETx1R_Bits   { return SETx1R_Bits(rm.UM32.Load()) }
func (rm SETx1R_Mask) Store(b SETx1R_Bits) { rm.UM32.Store(uint32(b)) }

type RSTx1R_Bits uint32

func (b RSTx1R_Bits) Field(mask RSTx1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RSTx1R_Bits) J(v int) RSTx1R_Bits {
	return RSTx1R_Bits(bits.Make32(v, uint32(mask)))
}

type RSTx1R struct{ mmio.U32 }

func (r *RSTx1R) Bits(mask RSTx1R_Bits) RSTx1R_Bits { return RSTx1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *RSTx1R) StoreBits(mask, b RSTx1R_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSTx1R) SetBits(mask RSTx1R_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *RSTx1R) ClearBits(mask RSTx1R_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *RSTx1R) Load() RSTx1R_Bits                 { return RSTx1R_Bits(r.U32.Load()) }
func (r *RSTx1R) Store(b RSTx1R_Bits)               { r.U32.Store(uint32(b)) }

type RSTx1R_Mask struct{ mmio.UM32 }

func (rm RSTx1R_Mask) Load() RSTx1R_Bits   { return RSTx1R_Bits(rm.UM32.Load()) }
func (rm RSTx1R_Mask) Store(b RSTx1R_Bits) { rm.UM32.Store(uint32(b)) }

type SETx2R_Bits uint32

func (b SETx2R_Bits) Field(mask SETx2R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SETx2R_Bits) J(v int) SETx2R_Bits {
	return SETx2R_Bits(bits.Make32(v, uint32(mask)))
}

type SETx2R struct{ mmio.U32 }

func (r *SETx2R) Bits(mask SETx2R_Bits) SETx2R_Bits { return SETx2R_Bits(r.U32.Bits(uint32(mask))) }
func (r *SETx2R) StoreBits(mask, b SETx2R_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SETx2R) SetBits(mask SETx2R_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *SETx2R) ClearBits(mask SETx2R_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *SETx2R) Load() SETx2R_Bits                 { return SETx2R_Bits(r.U32.Load()) }
func (r *SETx2R) Store(b SETx2R_Bits)               { r.U32.Store(uint32(b)) }

type SETx2R_Mask struct{ mmio.UM32 }

func (rm SETx2R_Mask) Load() SETx2R_Bits   { return SETx2R_Bits(rm.UM32.Load()) }
func (rm SETx2R_Mask) Store(b SETx2R_Bits) { rm.UM32.Store(uint32(b)) }

type RSTx2R_Bits uint32

func (b RSTx2R_Bits) Field(mask RSTx2R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RSTx2R_Bits) J(v int) RSTx2R_Bits {
	return RSTx2R_Bits(bits.Make32(v, uint32(mask)))
}

type RSTx2R struct{ mmio.U32 }

func (r *RSTx2R) Bits(mask RSTx2R_Bits) RSTx2R_Bits { return RSTx2R_Bits(r.U32.Bits(uint32(mask))) }
func (r *RSTx2R) StoreBits(mask, b RSTx2R_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSTx2R) SetBits(mask RSTx2R_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *RSTx2R) ClearBits(mask RSTx2R_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *RSTx2R) Load() RSTx2R_Bits                 { return RSTx2R_Bits(r.U32.Load()) }
func (r *RSTx2R) Store(b RSTx2R_Bits)               { r.U32.Store(uint32(b)) }

type RSTx2R_Mask struct{ mmio.UM32 }

func (rm RSTx2R_Mask) Load() RSTx2R_Bits   { return RSTx2R_Bits(rm.UM32.Load()) }
func (rm RSTx2R_Mask) Store(b RSTx2R_Bits) { rm.UM32.Store(uint32(b)) }

type EEFxR1_Bits uint32

func (b EEFxR1_Bits) Field(mask EEFxR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EEFxR1_Bits) J(v int) EEFxR1_Bits {
	return EEFxR1_Bits(bits.Make32(v, uint32(mask)))
}

type EEFxR1 struct{ mmio.U32 }

func (r *EEFxR1) Bits(mask EEFxR1_Bits) EEFxR1_Bits { return EEFxR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *EEFxR1) StoreBits(mask, b EEFxR1_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EEFxR1) SetBits(mask EEFxR1_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *EEFxR1) ClearBits(mask EEFxR1_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *EEFxR1) Load() EEFxR1_Bits                 { return EEFxR1_Bits(r.U32.Load()) }
func (r *EEFxR1) Store(b EEFxR1_Bits)               { r.U32.Store(uint32(b)) }

type EEFxR1_Mask struct{ mmio.UM32 }

func (rm EEFxR1_Mask) Load() EEFxR1_Bits   { return EEFxR1_Bits(rm.UM32.Load()) }
func (rm EEFxR1_Mask) Store(b EEFxR1_Bits) { rm.UM32.Store(uint32(b)) }

type EEFxR2_Bits uint32

func (b EEFxR2_Bits) Field(mask EEFxR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EEFxR2_Bits) J(v int) EEFxR2_Bits {
	return EEFxR2_Bits(bits.Make32(v, uint32(mask)))
}

type EEFxR2 struct{ mmio.U32 }

func (r *EEFxR2) Bits(mask EEFxR2_Bits) EEFxR2_Bits { return EEFxR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *EEFxR2) StoreBits(mask, b EEFxR2_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *EEFxR2) SetBits(mask EEFxR2_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *EEFxR2) ClearBits(mask EEFxR2_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *EEFxR2) Load() EEFxR2_Bits                 { return EEFxR2_Bits(r.U32.Load()) }
func (r *EEFxR2) Store(b EEFxR2_Bits)               { r.U32.Store(uint32(b)) }

type EEFxR2_Mask struct{ mmio.UM32 }

func (rm EEFxR2_Mask) Load() EEFxR2_Bits   { return EEFxR2_Bits(rm.UM32.Load()) }
func (rm EEFxR2_Mask) Store(b EEFxR2_Bits) { rm.UM32.Store(uint32(b)) }

type RSTxR_Bits uint32

func (b RSTxR_Bits) Field(mask RSTxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RSTxR_Bits) J(v int) RSTxR_Bits {
	return RSTxR_Bits(bits.Make32(v, uint32(mask)))
}

type RSTxR struct{ mmio.U32 }

func (r *RSTxR) Bits(mask RSTxR_Bits) RSTxR_Bits { return RSTxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RSTxR) StoreBits(mask, b RSTxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSTxR) SetBits(mask RSTxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *RSTxR) ClearBits(mask RSTxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *RSTxR) Load() RSTxR_Bits                { return RSTxR_Bits(r.U32.Load()) }
func (r *RSTxR) Store(b RSTxR_Bits)              { r.U32.Store(uint32(b)) }

type RSTxR_Mask struct{ mmio.UM32 }

func (rm RSTxR_Mask) Load() RSTxR_Bits   { return RSTxR_Bits(rm.UM32.Load()) }
func (rm RSTxR_Mask) Store(b RSTxR_Bits) { rm.UM32.Store(uint32(b)) }

type CHPxR_Bits uint32

func (b CHPxR_Bits) Field(mask CHPxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CHPxR_Bits) J(v int) CHPxR_Bits {
	return CHPxR_Bits(bits.Make32(v, uint32(mask)))
}

type CHPxR struct{ mmio.U32 }

func (r *CHPxR) Bits(mask CHPxR_Bits) CHPxR_Bits { return CHPxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CHPxR) StoreBits(mask, b CHPxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CHPxR) SetBits(mask CHPxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CHPxR) ClearBits(mask CHPxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CHPxR) Load() CHPxR_Bits                { return CHPxR_Bits(r.U32.Load()) }
func (r *CHPxR) Store(b CHPxR_Bits)              { r.U32.Store(uint32(b)) }

type CHPxR_Mask struct{ mmio.UM32 }

func (rm CHPxR_Mask) Load() CHPxR_Bits   { return CHPxR_Bits(rm.UM32.Load()) }
func (rm CHPxR_Mask) Store(b CHPxR_Bits) { rm.UM32.Store(uint32(b)) }

type CPT1xCR_Bits uint32

func (b CPT1xCR_Bits) Field(mask CPT1xCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPT1xCR_Bits) J(v int) CPT1xCR_Bits {
	return CPT1xCR_Bits(bits.Make32(v, uint32(mask)))
}

type CPT1xCR struct{ mmio.U32 }

func (r *CPT1xCR) Bits(mask CPT1xCR_Bits) CPT1xCR_Bits { return CPT1xCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CPT1xCR) StoreBits(mask, b CPT1xCR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CPT1xCR) SetBits(mask CPT1xCR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CPT1xCR) ClearBits(mask CPT1xCR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CPT1xCR) Load() CPT1xCR_Bits                  { return CPT1xCR_Bits(r.U32.Load()) }
func (r *CPT1xCR) Store(b CPT1xCR_Bits)                { r.U32.Store(uint32(b)) }

type CPT1xCR_Mask struct{ mmio.UM32 }

func (rm CPT1xCR_Mask) Load() CPT1xCR_Bits   { return CPT1xCR_Bits(rm.UM32.Load()) }
func (rm CPT1xCR_Mask) Store(b CPT1xCR_Bits) { rm.UM32.Store(uint32(b)) }

type CPT2xCR_Bits uint32

func (b CPT2xCR_Bits) Field(mask CPT2xCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPT2xCR_Bits) J(v int) CPT2xCR_Bits {
	return CPT2xCR_Bits(bits.Make32(v, uint32(mask)))
}

type CPT2xCR struct{ mmio.U32 }

func (r *CPT2xCR) Bits(mask CPT2xCR_Bits) CPT2xCR_Bits { return CPT2xCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CPT2xCR) StoreBits(mask, b CPT2xCR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CPT2xCR) SetBits(mask CPT2xCR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *CPT2xCR) ClearBits(mask CPT2xCR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *CPT2xCR) Load() CPT2xCR_Bits                  { return CPT2xCR_Bits(r.U32.Load()) }
func (r *CPT2xCR) Store(b CPT2xCR_Bits)                { r.U32.Store(uint32(b)) }

type CPT2xCR_Mask struct{ mmio.UM32 }

func (rm CPT2xCR_Mask) Load() CPT2xCR_Bits   { return CPT2xCR_Bits(rm.UM32.Load()) }
func (rm CPT2xCR_Mask) Store(b CPT2xCR_Bits) { rm.UM32.Store(uint32(b)) }

type OUTxR_Bits uint32

func (b OUTxR_Bits) Field(mask OUTxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OUTxR_Bits) J(v int) OUTxR_Bits {
	return OUTxR_Bits(bits.Make32(v, uint32(mask)))
}

type OUTxR struct{ mmio.U32 }

func (r *OUTxR) Bits(mask OUTxR_Bits) OUTxR_Bits { return OUTxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OUTxR) StoreBits(mask, b OUTxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OUTxR) SetBits(mask OUTxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *OUTxR) ClearBits(mask OUTxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *OUTxR) Load() OUTxR_Bits                { return OUTxR_Bits(r.U32.Load()) }
func (r *OUTxR) Store(b OUTxR_Bits)              { r.U32.Store(uint32(b)) }

type OUTxR_Mask struct{ mmio.UM32 }

func (rm OUTxR_Mask) Load() OUTxR_Bits   { return OUTxR_Bits(rm.UM32.Load()) }
func (rm OUTxR_Mask) Store(b OUTxR_Bits) { rm.UM32.Store(uint32(b)) }

type FLTxR_Bits uint32

func (b FLTxR_Bits) Field(mask FLTxR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FLTxR_Bits) J(v int) FLTxR_Bits {
	return FLTxR_Bits(bits.Make32(v, uint32(mask)))
}

type FLTxR struct{ mmio.U32 }

func (r *FLTxR) Bits(mask FLTxR_Bits) FLTxR_Bits { return FLTxR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FLTxR) StoreBits(mask, b FLTxR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FLTxR) SetBits(mask FLTxR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *FLTxR) ClearBits(mask FLTxR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *FLTxR) Load() FLTxR_Bits                { return FLTxR_Bits(r.U32.Load()) }
func (r *FLTxR) Store(b FLTxR_Bits)              { r.U32.Store(uint32(b)) }

type FLTxR_Mask struct{ mmio.UM32 }

func (rm FLTxR_Mask) Load() FLTxR_Bits   { return FLTxR_Bits(rm.UM32.Load()) }
func (rm FLTxR_Mask) Store(b FLTxR_Bits) { rm.UM32.Store(uint32(b)) }
