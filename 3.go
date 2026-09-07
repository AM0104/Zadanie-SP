package main

import "fmt"

func main() {
	var x = 5000
	var f = 256
	var sum = x / f
	var res = x % f
	fmt.Println("Поместится файлов", sum)
	fmt.Println("Свободного места останеться", res)
}


	
	