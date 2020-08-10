package main

import (
	"fmt"
	"GoOOP/demo/printer"
)

//Program entry point for this demo

func main(){
	//Printing Hello World without OOP
	fmt.Println("Hello World without OOP")

	//Creating a new object call printer and call printer to print hello world
	p := printer.NewPrinter("Tim")

	//The type of p is *printer.Printer, where the first printer
	//is the module name and the next Printer is the (exported) struct inside the printer module

	//Call to the "PrintHelloWorld" function for the printer object
	p.PrintHelloWorld();

}