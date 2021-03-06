package i2c

import (
	"stm32/hal/dma"

	"stm32/hal/raw/i2c"
)

// StopMode determines auto-stop behavior of master connection. If auto-stop
// mode is enabled for read (ASRD) or write (ASWR) then ay read/write operation
// is finished by sending stop condition on the I2C bus and leaves connection
// inactive. This mode improves ability to sharing I2C bus between multiple
// tasks but at the same time can degrade performance. It is not  recommended to
// disable auto-stop mode for read operations.
type StopMode byte

const (
	NOAS StopMode = 0      // Manual mode (use SetStopRead, StopWrite).
	ASRD StopMode = 1 << 1 // Any read is finished by sending stop condition.
	ASWR StopMode = 1 << 2 // Any write is finished by sending stop condition.

	stoprd StopMode = 1 << 0
)

type Error int16

const (
	BusErr   Error = 1 << 0
	ArbLost  Error = 1 << 1
	AckFail  Error = 1 << 2
	Overrun  Error = 1 << 3
	PECErr   Error = 1 << 4
	Timeout  Error = 1 << 6
	SMBAlert Error = 1 << 7

	BadEvent    Error = 1 << 8
	SoftTimeout Error = 1 << 9
	BelatedStop Error = 1 << 10
	ActiveRead  Error = 1 << 11 // Write when active read transaction.
	DMAErr      Error = 1 << 12
)

func (e Error) Error() string {
	switch e {
	case BadEvent:
		return "I2C bad event"
	case SoftTimeout:
		return "I2C soft timeout"
	case BelatedStop:
		return "I2C belated stop"
	case ActiveRead:
		return "I2C active read"
	case DMAErr:
		return "I2C DMA error"
	default:
		return "I2C hw error"
	}
}

func getError(sr1 i2c.SR1_Bits) Error {
	return Error(sr1 >> 8)
}

const dmaErrMask = dma.ErrAll &^ dma.ErrFIFO // Ignore FIFO error.

const byteTimeout = 100e6 // 100 ms
