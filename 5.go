package main
import "fmt", "math"
func main(){
	var r int
	fmt.Println("Введите радиус окр")
	fmt.Scan(&r)
	var d = 2 * math.PI * r
	var p = math.PI * math.Pow(r, 2)
	fmt.Println("Длинна равна:", d)
	fmt.Println("Площадь равна:", p)
}
