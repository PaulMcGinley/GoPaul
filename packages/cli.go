package Console

import "fmt"

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func WriteLine(message string) {
	fmt.Println(message)
}

func Write(message string) {
	fmt.Print(message)
}

func ReadLine() string {
	var value string
	fmt.Scanln(&value)
	return value
}

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
