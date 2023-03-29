package main

import (

	"example.com/customer"
	
)

func main() {
	a := customer.App{}
	a.Initialize()
	a.Run()
}