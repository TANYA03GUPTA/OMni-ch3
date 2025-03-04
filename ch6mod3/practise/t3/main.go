package main

import (
	"fmt"
)

func printNum(ch chan int){
	for i:= 1;i<=3;i++{
		ch <- i+2
	}
	close(ch)
}
func senddata(ch chan int, n1 int){
	ch <- n1
}

func main(){
	ch := make(chan int)

	go printNum(ch)
	go senddata(ch,8)
	go senddata(ch,10)

	for nm := range ch{
		fmt.Println(nm)
	}

} 