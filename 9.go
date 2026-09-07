package main
import "fmt"
func main(){
	var sum float64
    fmt.Println("Введите сумму покупки ")
    fmt.Scan(&sum)
    res := sum * 0.8
    fmt.Println("Сумма со скидкой", res)
}