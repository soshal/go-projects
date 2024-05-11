package main

import "fmt"

func main(){
	var inflationrate float64
	var currentprice float64
	var futureprice float64
	var years float64

	fmt.Println("Enter the current price of the item: ")
	fmt.Scanln(&currentprice)
	fmt.Println("Enter the number of years: ")
	fmt.Scanln(&years)
	fmt.Println("Enter the inflation rate: ")
	fmt.Scanln(&inflationrate)
	

	inflationrate = inflationrate/100
	futureprice = currentprice * (1 + inflationrate) * years
	
	fmt.Println("The future price of the item is: ", futureprice)

}



