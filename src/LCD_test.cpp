5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26

// Подключение LCD на базе HD44780 к ATmega16 (LM016L LCD 16x2)
// сайт http://micro-pi.ru

#define F_CPU 8000000UL
#include <avr/io.h>
#include <util/delay.h>
#include <string.h>
#include "LCD.h"

int main(void) {
  _delay_ms(100);
  lcdInit();
  lcdClear();
  lcdSetDisplay(LCD_DISPLAY_ON);
  lcdSetCursor(LCD_CURSOR_OFF);

  char text[17];
  strcpy(text, "  Hello World!  ");
  lcdGotoXY(0, 0);
  lcdPuts(text);
  strcpy(text, "site:micro-pi.ru");
  lcdGotoXY(1, 0);
  lcdPuts(text);

  while (1);
}
