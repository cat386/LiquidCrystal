package main

import (
	"fmt"
	unsafe "unsafe"

	"periph.io/x/periph/conn/gpio"
)

type data struct {
	name string
	data uint8
	hPin gpio.PinIO
}

func main() {
	var ttt data
	var tta [20]data

	ttt.name = "11"
	ttt.data = 11

	const infoSize = unsafe.Sizeof(ttt)

	fmt.Println("size struct: ", infoSize)
	fmt.Println("All size for display: ", infoSize*20)
	fmt.Println("-----------------------------------")
	fmt.Println("data.name = ", unsafe.Sizeof(ttt.name))
	fmt.Println("data.data = ", unsafe.Sizeof(ttt.data))
	fmt.Println("data.hPin = ", unsafe.Sizeof(ttt.hPin))
	fmt.Println("Array size for display: ", unsafe.Sizeof(tta))
	fmt.Println("-----------------------------------")
}
