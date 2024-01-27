package main

import "fmt"

var Sc = float64(10000) // текущее расстояние между людьми
var Sp = Sc             // предыдущее расстояние между людьми
var Smin = float64(10)  //расстояние, на котором будет считаться,
// что два человека встретились
var Tcnt = float64(0) //время для встречи собаки и человека
// в текущей итерации
var Man = int(2)    //к какому человеку бежит собака
var V1 = float64(1) //скорость первого человека
var V2 = float64(2) //скорость второго человека
var V3 = float64(5) //скорость лучшего друга человека

var Cnt = int(0) //счетчик забегов собаки

func main() {
	for Sc > Smin {
		if Man == 2 {
			Tcnt = Sc / (V2 + V3)
			Man = 1
		} else {
			Tcnt = Sc / (V1 + V3)
			Man = 2
		}
		Sp = Sc
		Cnt++
		Sc = Sp - (V1*Tcnt + V2*Tcnt)
	}
	fmt.Printf("Собака метнулась %d раз\n", Cnt)
}
