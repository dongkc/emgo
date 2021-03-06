package ili9341

import (
	"image"
)

type Display struct {
	dci    DCI
	width  uint16
	height uint16
	swapWH bool
}

// MakeDisplay returns initialised Display value.
func MakeDisplay(dci DCI) Display {
	return Display{
		dci:    dci,
		width:  240,
		height: 320,
	}
}

// NewDisplay works like MakeDisplay but returns a pointer to the heap allocated
// variable.
func NewDisplay(dci DCI) *Display {
	d := new(Display)
	*d = MakeDisplay(dci)
	return d
}

// DCI allows to direct access to the internal DCI.
func (d *Display) DCI() DCI {
	return d.dci
}

// Err returns and clears internal error variable.
func (d *Display) Err() error {
	return d.dci.Err()
}

// Bounds returns the bounds of the display
func (d *Display) Bounds() image.Rectangle {
	if d.swapWH {
		return image.Rectangle{Max: image.Pt(int(d.height), int(d.width))}
	}
	return image.Rectangle{Max: image.Pt(int(d.width), int(d.height))}
}

// SetWordSize changes the data word size.
func (d *Display) SetWordSize(size int) {
	d.dci.SetWordSize(size)
}

func (d *Display) Area(r image.Rectangle) Area {
	a := Area{disp: d, rect: r.Canon()}
	a.updateBounds()
	return a
}

func (d *Display) NewArea(r image.Rectangle) *Area {
	a := new(Area)
	*a = d.Area(r)
	return a
}
