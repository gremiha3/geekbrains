package main

import "fmt"

// объявляем и инициализируем массив целых знаковых чисел
var array = []int{2, 5, 13, 7, 6, 4} //задаем исходный массив
var avarsum = 0                      //задаем начальное значение среднего арифметического.

func main() {
	fmt.Printf("Исходный массив = %v\n", array)
	fmt.Printf("Размер массива = %d значений\n", len(array))

	fmt.Println("-----------------------------------")
	for index, value := range array {
		fmt.Printf("Индекс = %d, Значение = %d\n", index, value)
	}
	fmt.Println("-----------------------------------")

	for index := 0; index < len(array); index++ {
		avarsum = avarsum + array[index]
	}
	avarsum = avarsum / len(array)
	fmt.Printf("Среднее арифметическое = %d\n", avarsum)
}
