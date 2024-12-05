package main

import "fmt"

func main(){
	var height , weight float64;

	fmt.Print("Enter your weight :");
	fmt.Scanln(&weight);
	fmt.Print("Enter your height :");
	fmt.Scanln(&height);

	bmi := weight / (height * height );
	fmt.Printf("your bmi is %f", bmi );
	if bmi < 18.5 {
		fmt.Print("your bmi is underweight")
	}else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Print("your bmi is normalweight")
	}else if bmi >29.9 && bmi < 29.9 {
		fmt.Print("your bmi is overweight")
	}else{
		fmt.Print("yor are obset!")
	}
 }