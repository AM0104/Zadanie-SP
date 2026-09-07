package main
import (
	"fmt"
 	"math"
)
func main(){
	var r float64
	fmt.Println("Введите радиус окр")
	fmt.Scan(&r)
	var d = 2 * math.Pi * r
	var p = math.Pi * math.Pow(r, 2)
	fmt.Println("Длинна равна:", d)
	fmt.Println("Площадь равна:", p)
}
