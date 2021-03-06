package main

import (
	"nrf5/hal/clock"
	"nrf5/hal/gpio"
	"nrf5/hal/irq"
	"nrf5/hal/rtc"
	"nrf5/hal/system"
	"nrf5/hal/system/timer/rtcst"
	"nrf5/hal/uart"
)

var (
	led gpio.Pin
	u   *uart.Periph
)

func init() {
	system.Setup(clock.XTAL, clock.XTAL, true)
	rtcst.Setup(rtc.RTC0, 1)

	p0 := gpio.P0

	led = p0.Pin(18)
	led.Setup(gpio.ModeOut)

	u = uart.UART0
	u.StorePSEL(uart.SignalRXD, p0.Pin(11))
	u.StorePSEL(uart.SignalTXD, p0.Pin(9))
	u.StoreBAUDRATE(uart.Baud115200)
	u.StoreENABLE(true)
}

func main() {
	u.Task(uart.STARTRX).Trigger()
	u.Task(uart.STARTTX).Trigger()
	rxdrdy := u.Event(uart.RXDRDY)
	txdrdy := u.Event(uart.TXDRDY)

	for i := 0; ; i++ {
		for !rxdrdy.IsSet() {
		}
		rxdrdy.Clear()
		b := u.LoadRXD()
		for i != 0 && !txdrdy.IsSet() {
		}
		txdrdy.Clear()
		u.StoreTXD(b)
		led.Store(i)
	}
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.RTC0: rtcst.ISR,
}
