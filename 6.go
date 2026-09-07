package main

import (
	"fmt"
	"math"
)

func main() {
	var i float64
	var a float64
	var y int
	fmt.Println("Введите число i")
	fmt.Scan(&i)
	fmt.Println("Введите число a")
	fmt.Scan(&a)
	fmt.Println("Введите число y")
	fmt.Scan(&y)
	var s = i * math.Pow((1+a/100), float64(y))
	fmt.Println("Итоговая сумма через", y, " лет", s, "рублей")
}