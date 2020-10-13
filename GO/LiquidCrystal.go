package LiquidCrystal

import (
	"periph.io/x/periph/conn/gpio"
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
	data        string
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

/*
	rs_pin          uint8
	rw_pin          uint8
	enable_pin      uint8
	data_pins       [8]uint8
	displayFunction uint8
	displayControl  uint8
	displayMode     uint8
	initialized     uint8
	numLines        uint8
	row_offsets     [4]uint8
	//-----------------------
	names [8]namePin
)
*/

// Init(fourbitmode,rs,rw,enable,d0,d1,d2,d3,d4,d5,d6,d7 uint8)
// init display
func Init(fourbitmode, rs, rw, enable, d0, d1, d2, d3, d4, d5, d6, d7 uint8) {
	/*rs_pin = rs
	rw_pin = rw
	enable_pin = enable
	data_pins[0] = d0
	data_pins[1] = d1
	data_pins[2] = d2
	data_pins[3] = d3
	data_pins[4] = d4
	data_pins[5] = d5
	data_pins[6] = d6
	data_pins[7] = d7
	*/
	// register pins gpio.PinIO

	if fourbitmode == 0 {
		displayFunction = LCD_4BITMODE | LCD_1LINE | LCD_5x8DOTS
	} else {
		displayFunction = LCD_8BITMODE | LCD_1LINE | LCD_5x8DOTS
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
