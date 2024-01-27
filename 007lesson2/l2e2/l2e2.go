package main

import "fmt"

// объявляем и инициализируем массив целых знаковых чисел
var array = []int{2, 6, 1, -101, 4, 6, -23, -100, 345, 76, -12}
var mirrorarray []int

func main() {
	fmt.Printf("Исходный массив = %v\n", array)
	fmt.Printf("Размер массива = %d значений\n", len(array))

	fmt.Println("-----------------------------------")
	for index, value := range array {
		fmt.Printf("Индекс = %d, Значение = %d\n", index, value)
	}
	fmt.Println("-----------------------------------")

	for index := len(array) - 1; index >= 0; index-- {
		mirrorarray = append(mirrorarray, array[index])
	}
	fmt.Printf("Перевернутый массив = %v\n", mirrorarray)
}
