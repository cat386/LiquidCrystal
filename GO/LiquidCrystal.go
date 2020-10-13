package LiquidCrystal

import (
	"fmt"
	"strconv"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

const (
	// commands
	LCD_CLEARDISPLAY   = 0x01
	LCD_RETURNHOME     = 0x02
	LCD_ENTRYMODESET   = 0x04
	LCD_DISPLAYCONTROL = 0x08
	LCD_CURSORSHIFT    = 0x10
	LCD_FUNCTIONSET    = 0x20
	LCD_SETCGRAMADDR   = 0x40
	LCD_SETDDRAMADDR   = 0x80

	// flags for display entry mode
	LCD_ENTRYRIGHT          = 0x00
	LCD_ENTRYLEFT           = 0x02
	LCD_ENTRYSHIFTINCREMENT = 0x01
	LCD_ENTRYSHIFTDECREMENT = 0x00

	// flags for display on/off control
	LCD_DISPLAYON  = 0x04
	LCD_DISPLAYOFF = 0x00
	LCD_CURSORON   = 0x02
	LCD_CURSOROFF  = 0x00
	LCD_BLINKON    = 0x01
	LCD_BLINKOFF   = 0x00

	// flags for display/cursor shift
	LCD_DISPLAYMOVE = 0x08
	LCD_CURSORMOVE  = 0x00
	LCD_MOVERIGHT   = 0x04
	LCD_MOVELEFT    = 0x00

	// flags for function set
	LCD_8BITMODE = 0x10
	LCD_4BITMODE = 0x00
	LCD_2LINE    = 0x08
	LCD_1LINE    = 0x00
	LCD_5x10DOTS = 0x04
	LCD_5x8DOTS  = 0x00
)

type DataPin struct {
	//data        string //???? - хз
	values      uint8
	hardwarePin gpio.PinIO
}

type Datas struct {
	//name string
	rs_pin     DataPin //uint8
	rw_pin     DataPin //uint8
	enable_pin DataPin //uint8

	data_pins [8]DataPin //uint8

	displayFunction uint8
	displayControl  uint8
	displayMode     uint8
	initialized     uint8
	numLines        uint8
	row_offsets     [4]uint8
}

var (
	Pins Datas
)


// Init(fourbitmode,rs,rw,enable,d0,d1,d2,d3,d4,d5,d6,d7 uint8)
// init display
func Init(fourbitmode, rs, rw, enable, d0, d1, d2, d3, d4, d5, d6, d7 uint8) {

	_ ,err := host.Init()
	if err != nil{
		fmt.Println("error, in host.init")
	}else{
		Pins.rs_pin.values          := rs
		Pins.rs_pin.hardwarePin     := gpioreg.ByName(strconv.Itoa(rs)) 

		Pins.rw_pin.values          := rw
		Pins.rw_pin.hardwarePin     := gpioreg.ByName(strconv.Itoa(rw)) 

		Pins.enable_pin.values      := enable
		Pins.enable_pin.hardwarePin := gpioreg.ByName(strconv.Itoa(enable)) 

		if d0 > 0{
			Pins.data_pins[0].values      := d0
			Pins.data_pins[0].hardwarePin := gpioreg.ByName(strconv.Itoa(d0)) 
		}

		if d1 > 0{
			Pins.data_pins[1].values      := d1
			Pins.data_pins[1].hardwarePin := gpioreg.ByName(strconv.Itoa(d1)) 
		}

		if d2 > 0{
			Pins.data_pins[2].values      := d2
			Pins.data_pins[2].hardwarePin := gpioreg.ByName(strconv.Itoa(d2)) 
		}

		if d3 > 0{
			Pins.data_pins[3].values      := d3
			Pins.data_pins[3].hardwarePin := gpioreg.ByName(strconv.Itoa(d3)) 
		}

		if d4 > 0{
			Pins.data_pins[4].values      := d4
			Pins.data_pins[4].hardwarePin := gpioreg.ByName(strconv.Itoa(d4)) 
		}

		if d5 > 0{
			Pins.data_pins[5].values      := d5
			Pins.data_pins[5].hardwarePin := gpioreg.ByName(strconv.Itoa(d5)) 
		}

		if d6 > 0{
			Pins.data_pins[6].values      := d6
			Pins.data_pins[6].hardwarePin := gpioreg.ByName(strconv.Itoa(d6)) 
		}

		if d7 > 0{
			Pins.data_pins[7].values      := d7
			Pins.data_pins[7].hardwarePin := gpioreg.ByName(strconv.Itoa(d7)) 
		}

		

		// register pins gpio.PinIO

		if fourbitmode == 0 {
			Pins.displayFunction := LCD_4BITMODE | LCD_1LINE | LCD_5x8DOTS
		} else {
			Pins.displayFunction := LCD_8BITMODE | LCD_1LINE | LCD_5x8DOTS
		}
	}
	
}

//func Begin(cols, lines, dotsize uint8)
// Begin init work
func Begin(cols, lines, dotsize uint8) {

	if lines > 1 {
		displayFunction |= LCD_2LINE
	}

	numLines = lines

	// next
	//setRowOffcets(0x00, 0x40, 0x00+cols, 0x40+cols)

	if dotsize != LCD_5x8DOTS && lines == 1 {
		displayFunction |= LCD_5x10DOTS
	}

	// next
	//pinMode(rs_pin, OUTPUT)

	//if rw_pin != 255 {
	//	pinMode(Rw_pin, OUTPUT)
	//}

	//pinMode(enable_pin, OUTPUT)

	//for
}
