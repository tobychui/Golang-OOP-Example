# Golang-OOP-Example
A very basic example showcasing the use of go module and struct for OOP in Golang project


## Qestions you might want to ask

1. Why the import module is "GoOOP/demo/printer" instead of just "printer" or "Golang-OOP-Example/printer"?

	It is because this project is initialized with 
	```
	go mod init GoOOP/demo
	```

2. Why the compiled binary is named "demo.exe" instead of the folder name "Golang-OOP-Example.exe"?

	It is because this project is initialized with 
	```
	go mod init GoOOP/demo
	```

	and the go build command will use the last section of the module name as the binary executable name.


3. If I am starting a new project, how can I name my module?

	Use the go mod command as follows.

	```
	go mod init {your module name}

	/*
		Example of module names: 
		github.com/tobychui/mymodule
		mydomain.com/mymodule
		anything/you/want
		
	*/ 

	go mod tidy
	```

4. Why the module has to be initialize with ```p := printer.NewPrinter("Tim")``` instead of other module name?
	
	You can choose to import "GoOOP/demo/printer" with other names like
	```
		import (
			magic "GoOOP/demo/printer"
		)
	```
	
	and call to the imported module using 
	
	```
		p := magic.NewPrinter("Tim")
	```
	
	
## Questions?
	Feel free to open a new issue and I will add your Q&A into the README.md file.
	