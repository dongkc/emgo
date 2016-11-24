// +build cortexm0 cortexm3 cortexm4 cortexm4f cortexm7f cortexm7d

package noos

import (
	"arch/cortexm"
)

//emgo:noinline
func nmiHandler() {
	for {
		cortexm.BKPT(0)
	}
}

//emgo:export
func faultHandler()
