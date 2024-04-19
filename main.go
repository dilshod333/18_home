package main

import "fmt"

func main() {
	num := 0
	ch:=make(chan int ,1)
	fmt.Print("Enter the number: ")
	fmt.Scanln(&num)
	go factorial(num,ch)

	fmt.Println(<-ch)
	
}

func factorial(num int,ch chan int) {
	fact:=1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	ch<-fact

}
