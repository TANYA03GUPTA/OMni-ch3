package main

import (
	"fmt"
)


func main(){
	var size int 
	fmt.Print("Enter the size of the array: ")
	_, err := fmt.Scan(&size)
	if err != nil || size<=0 {
		fmt.Println("Invalid input")
	}
	num := make([]int, size)
	for i := 0;i<size;i++ {
		fmt.Printf("Enter the %d element: ", i+1)
		_, err = fmt.Scan(&num[i])
		if err != nil{
			fmt.Println("Invalid input")
			return
		}
	}
	//create channle
	ch := make(chan []int)
	ch2 := make(chan string)

	go func(){
	for k:= 0;k<size;k++{
		var s1 string 
		fmt.Printf("enter %d word ," , k+1)
		_, err := fmt.Scan(&s1)
		if err != nil{
			fmt.Println("Invalid input")
			return
		}
		ch2 <- s1
	}
	close(ch2)
}()
  var words []string
  for wd := range ch2{
	words = append(words, wd)
  }
  fmt.Println("Words :", words)
  

	//sned data
	go func(){
		ch <- num
		close(ch)
	}()

	res := <- ch
	fmt.Println("recieved", res)

	ch3 := make(chan string, 3)
	ch3 <- "highalands"
	ch3 <- "fauji"

	fmt.Println(<- ch3)
	fmt.Println(<- ch3)

}
//bufered channels 
//letbus send dat w/o an immediate reciever 