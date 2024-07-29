// Package console does this
package console

import "fmt"

type Format int

const (
	Reset Format = iota
	Bold
	Dim
	Italic
	Underlined
	Blink
	Reverse       = 7
	Hidden        = 8
	StrikeThrough = 9

	EndBold = 22
	EndDim
	EndItalic
	EndUnderlined
	EndBlink
	EndReverse
	EndHidden
	EndStrikeThrough = 29
) // https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters

type Colour int

const (
	ForegroundBlack Colour = iota + 30
	ForegroundRed
	ForegroundGreen
	ForegroundYellow
	ForegroundBlue
	ForegroundMagenta
	ForegroundCyan
	ForegroundWhite
)

type BackgroundColour int

const (
	BackgroundBlack BackgroundColour = iota
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

const NewLine = "\n"

func PrintFormat(params ...any) {
	var output string

	for _, param := range params {
		switch param.(type) {
		case Format:
			output += fmt.Sprintf("\033[%dm", param)
		case Colour:
			output += fmt.Sprintf("\033[%dm", param)
		case BackgroundColour:
			output += fmt.Sprintf("\033[48;5;%dm", param)
		case string:
			output += param.(string)
		default:
			output += fmt.Sprintf("%v", param)
		}
	}

	fmt.Print(output)
}

// Clear empties the console of all text
func Clear() {
	fmt.Print("\033[H\033[2J")
}

// WriteLine writes a line of text to the console with a newline character at the end
func WriteLine(message ...string) {
	fmt.Println(message)
}

// Write appends text to the console
func Write(message string) {
	fmt.Print(message)
}

// ReadLine reads a line of text from the console
func ReadLine() string {
	var value string
	fmt.Scanln(&value)
	return value
}

// Read reads a line of text from the console
func Read() string {
	var value string
	fmt.Scan(&value)
	return value
}

func SetColor(color string) {
	fmt.Print("\033[" + color + "m")
}

func ResetColor() {
	fmt.Print("\033[0m")
}

func SetBackgroundColor(color string) {
	fmt.Print("\033[48;5;" + color + "m")
}

func ResetBackgroundColor() {
	fmt.Print("\033[49m")
}

func SetBold() {
	fmt.Print("\033[1m")
}

func ResetBold() {
	fmt.Print("\033[22m")
}

func SetUnderline() {
	fmt.Print("\033[4m")
}

func ResetUnderline() {
	fmt.Print("\033[24m")
}

func SetItalic() {
	fmt.Print("\033[3m")
}

func ResetItalic() {
	fmt.Print("\033[23m")
}

func SetStrikeThrough() {
	fmt.Print("\033[9m")
}

func ResetStrikeThrough() {
	fmt.Print("\033[29m")
}

func SetInverse() {
	fmt.Print("\033[7m")
}

func ResetInverse() {
	fmt.Print("\033[27m")
}

func SetHidden() {
	fmt.Print("\033[8m")
}

func ResetHidden() {
	fmt.Print("\033[28m")
}

func SetFramed() {
	fmt.Print("\033[51m")
}

func ResetFramed() {
	fmt.Print("\033[54m")
}

func SetEncircled() {
	fmt.Print("\033[52m")
}

func ResetEncircled() {
	fmt.Print("\033[54m")
}

func SetOverlined() {
	fmt.Print("\033[53m")
}

func ResetOverlined() {
	fmt.Print("\033[55m")
}

func SetColor256(color string) {
	fmt.Print("\033[38;5;" + color + "m")
}

func SetBackgroundColor256(color string) {
	fmt.Print("\033[48;5;" + color + "m")
}

func SetColorRGB(r, g, b string) {
	fmt.Print("\033[38;2;" + r + ";" + g + ";" + b + "m")
}

func SetBackgroundColorRGB(r, g, b string) {
	fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m")
}

func SetColor256RGB(r, g, b string) {
	fmt.Print("\033[38;5;" + r + ";" + g + ";" + b + "m")
}

func SetBackgroundColor256RGB(r, g, b string) {
	fmt.Print("\033[48;5;" + r + ";" + g + ";" + b + "m")
}
