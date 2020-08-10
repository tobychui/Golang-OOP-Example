package printer

import (
	"fmt"
)

//Create the structure of the "object", in this case, a Printer that contains the name of the user
type Printer struct{
	name string
}

//The function that the main.go call to create a new printer
func NewPrinter(printerName string) *Printer{
	return &Printer{
		name: printerName,
	}
}

//A function inside the printer object that use printer's properties to do something
//In this example, it prints the name of the printer user
func (p Printer)PrintHelloWorld(){
	fmt.Println( getHelloWorld() + " " + p.name)
}


