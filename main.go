package main

import (
	Console "github.com/PaulFMcGinley/GoPaul/packages"
	"strconv"
)

func main() {

	for {
		Console.Clear()
		Console.WriteLine("1. Write")
		Console.WriteLine("2. Read")
		Console.WriteLine("3. Set Color")
		Console.WriteLine("4. Reset Color")
		Console.WriteLine("5. Set Background Color")
		Console.WriteLine("6. Reset Background Color")
		Console.WriteLine("7. Set Bold")
		Console.WriteLine("8. Reset Bold")
		Console.WriteLine("9. Set Underline")
		Console.WriteLine("10. Reset Underline")
		Console.WriteLine("11. Exit")
		Console.WriteLine("Choose an option: ")
		option := Console.ReadLine()

		switch option {
		case "1":
			Console.Clear()
			Console.WriteLine("Write: ")
			message := Console.ReadLine()
			Console.Write(message)
			Console.ReadLine()
		case "2":
			Console.Clear()
			Console.WriteLine("Read: ")
			message := Console.Read()
			Console.WriteLine("You wrote: " + message)
			Console.ReadLine()
		case "3":
			Console.Clear()
			Console.WriteLine("Set Color: ")
			color := Console.ReadLine()
			Console.SetColor(color)
			Console.WriteLine("This is a colored text")
			Console.ResetColor()
			Console.ReadLine()
		case "4":
			Console.Clear()
			Console.WriteLine("Reset Color")
			Console.ResetColor()
			Console.ReadLine()
		case "5":
			Console.Clear()

			for i := 0; i < 256; i++ {
				Console.SetBackgroundColor(strconv.Itoa(i))
				Console.Write(strconv.Itoa(i))
				Console.WriteLine("This is a colored background text")
				Console.ResetBackgroundColor()
			}
			//Console.WriteLine("Set Background Color: ")
			//color := Console.ReadLine()
			//Console.SetBackgroundColor(color)
			//Console.WriteLine("This is a colored background text")
			//Console.ResetBackgroundColor()
			Console.ReadLine()
		case "6":
			Console.Clear()
			Console.WriteLine("Reset Background Color")
			Console.ResetBackgroundColor()
			Console.ReadLine()
		case "7":
			Console.Clear()
			Console.WriteLine("Set Bold")
			Console.SetBold()
			Console.WriteLine("This is a bold text")
			Console.ResetBold()
			Console.ReadLine()
		case "8":
			Console.Clear()
			Console.WriteLine("Reset Bold")
			Console.ResetBold()
			Console.ReadLine()
		case "9":
			Console.Clear()
			Console.WriteLine("Set Underline")
			Console.SetUnderline()
			Console.WriteLine("This is an underline text")
			Console.ResetUnderline()
			Console.ReadLine()
		case "10":
			Console.Clear()
			Console.WriteLine("Reset Underline")
			Console.ResetUnderline()
			Console.ReadLine()
		case "11":
			Console.Clear()
			Console.WriteLine("Goodbye!")
			return
		default:

		}
	}
}
