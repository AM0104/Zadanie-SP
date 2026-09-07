package main

import "fmt"

func main() {
	var f float64
	fmt.Println("Введите значение для f")
	fmt.Scan(&f)
	var c = (f - 32) * 5 / 9
	fmt.Println("Температура равна", c)
}
