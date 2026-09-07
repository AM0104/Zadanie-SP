package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Введите число a ")
	fmt.Scan(&a)
	fmt.Print("Введите число b ")
	fmt.Scan(&b)
	var del = a / b
	var c = math.Round((del))
	var o = math.Floor(float64(del))
	fmt.Println("Результат деления", del)
	fmt.Println("ближайшего целого ", c)
	fmt.Println("Округление вверх", o)
}