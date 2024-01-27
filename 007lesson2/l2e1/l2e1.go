package main

import "fmt"

// объявляем и инициализируем массив целых знаковых чисел
var array = []int{2, 6, 1, 8, 4, 6, -23, -100, 345, 76, -12}
var mini = 0
var maxi = 0
var minv = 0
var maxv = 0

func main() {
	fmt.Printf("Исходный массив = %v\n", array)
	fmt.Printf("Размер массива = %d значений\n", len(array))

	fmt.Println("-----------------------------------")
	for index, value := range array {
		fmt.Printf("Индекс = %d, Значение = %d\n", index, value)
	}
	fmt.Println("-----------------------------------")

	for index := 0; index < len(array); index++ {
		if array[index] < minv {
			minv = array[index]
			mini = index
		}
		if array[index] > maxv {
			maxv = array[index]
			maxi = index
		}
	}
	fmt.Printf("Индекс минимального значения = %d\n", mini)
	fmt.Printf("Индекс максимального значения = %d\n", maxi)
}
