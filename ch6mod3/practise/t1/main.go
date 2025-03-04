package main 


import ("fmt")

func add(a int, b int)int {
	return a + b
}

func main(){
	res := add(4,5)
	fmt.Println("Result of the addition is : ",res)

	var r1,r2 int
	fmt.Print("enter the first no:")
	_,err := fmt.Scan(&r1)
	if err!=nil{
		fmt.Println("Invalid input")
		return
	}
	fmt.Print("enter the second no:")
	_, er2 := fmt.Scan(&r2)
	if er2 != nil{
		fmt.Print("Invalid")
		return
	}
	fmt.Println("Result of the addition is : ",r1+r2)
}