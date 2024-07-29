package main

import (
	"github.com/PaulMcGinley/GoPaul/console"
	"strconv"
)

func main() {

	for {
		console.Clear()
		console.WriteLine("1. Write")
		console.WriteLine("2. Read")
		console.WriteLine("3. Set Color")
		console.WriteLine("4. Reset Color")
		console.WriteLine("5. Set Background Color")
		console.WriteLine("6. Reset Background Color")
		console.WriteLine("7. Set Bold")
		console.WriteLine("8. Reset Bold")
		console.WriteLine("9. Set Underline")
		console.WriteLine("10. Reset Underline")
		console.WriteLine("11. Exit")
		console.WriteLine("Choose an option: ")
		option := console.ReadLine()

		switch option {
		case "1":
			console.Clear()
			console.WriteLine("Write: ")
			message := console.ReadLine()
			console.Write(message)
			console.ReadLine()
		case "2":
			console.Clear()
			console.WriteLine("Read: ")
			message := console.Read()
			console.WriteLine("You wrote: " + message)
			console.ReadLine()
		case "3":
			console.Clear()
			console.WriteLine("Set Color: ")
			color := console.ReadLine()
			console.SetColor(color)
			console.WriteLine("This is a colored text")
			console.ResetColor()
			console.ReadLine()
		case "4":
			console.Clear()
			console.WriteLine("Reset Color")
			console.ResetColor()
			console.ReadLine()
		case "5":
			console.Clear()

			for i := 0; i < 256; i++ {
				console.SetBackgroundColor(strconv.Itoa(i))
				console.Write(strconv.Itoa(i))
				console.WriteLine("This is a colored background text")
				console.ResetBackgroundColor()
			}
			//console.WriteLine("Set Background Color: ")
			//color := console.ReadLine()
			//console.SetBackgroundColor(color)
			//console.WriteLine("This is a colored background text")
			//console.ResetBackgroundColor()
			console.ReadLine()
		case "6":
			console.Clear()
			console.WriteLine("Reset Background Color")
			console.ResetBackgroundColor()
			console.ReadLine()
		case "7":
			console.Clear()
			console.WriteLine("Set Bold")
			console.SetBold()
			console.WriteLine("This is a bold text")
			console.ResetBold()
			console.ReadLine()
		case "8":
			console.Clear()
			console.WriteLine("Reset Bold")
			console.ResetBold()
			console.ReadLine()
		case "9":
			console.Clear()
			console.WriteLine("Set Underline")
			console.SetUnderline()
			console.WriteLine("This is an underline text")
			console.ResetUnderline()
			console.ReadLine()
		case "10":
			console.Clear()
			console.WriteLine("Reset Underline")
			console.ResetUnderline()
			console.ReadLine()
		case "11":
			console.Clear()
			console.WriteLine("Goodbye!")
			return
		default:

		}
	}
}
