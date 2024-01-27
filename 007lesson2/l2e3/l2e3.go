package main

import "fmt"

// объявляем и инициализируем массив целых знаковых чисел
var array = []int{2, 6, 1, -101, 4, 6, -23, -100, 345, 76, -120}
var mini = 0 //индекс минимального элемента
var maxi = 0 //индекс максимального элемента
var minv = 0 //значение минимального элемента
var maxv = 0 //значение максимального элемента
var sum = 0  //начальное значение искомой суммы

func main() {
	fmt.Printf("Исходный массив = %v\n", array)              //печатаем исходный массив
	fmt.Printf("Размер массива = %d значений\n", len(array)) //печатаем длину массива

	fmt.Println("-----------------------------------") //печатаем разделитель
	for index, value := range array {                  //проходим по всем элементам массива
		//выводим таблицу соответствия индексов и значений элементов:
		fmt.Printf("Индекс = %d, Значение = %d\n", index, value)
	}
	fmt.Println("-----------------------------------") //печатаем разделитель

	for index := 0; index < len(array); index++ { //проходим по всем элементам массива
		if array[index] < minv { //если элемент массива меньше сохраненного минимального значения, то
			minv = array[index] //минимальное значение равно этому значению элемента массива
			mini = index        //а, индекс минимального элемента равен индексу этого элемента массива
		}
		if array[index] > maxv { //если элемент массива больше сохраненного максимального значения, то
			maxv = array[index] //максимальное значение равно этому значению элемента массива
			maxi = index        //а, индекс максимального элемента равен индексу этого элемента массива
		}
	}
	fmt.Printf("Индекс минимального значения = %d\n", mini)  //печатаем индекс минимального элемента
	fmt.Printf("Индекс максимального значения = %d\n", maxi) //печатаем индекс максимального элемента

	//проходим по элементам массива от элемента с минимальным значением до элемента с максимальным значением
	//сами минимальный и максимальный элементы в подсчет НЕ включены
	if mini < maxi { //если минимальное значение массива находится левее максимального значения, то
		for index := mini + 1; index < maxi; index++ { //идем по массиву от индекса минимального элемента
			sum = sum + array[index] //до индекса максимального элемента, и считаем сумму
		}
	} else { //если минимальное значение массива находится правее максимального значения, то
		for index := maxi + 1; index < mini; index++ { //идем от индекса максимального элемента
			sum = sum + array[index] //до индекса минимального элемента, и считаем сумму
		}
	}
	fmt.Printf("Сумма между минимальным и максимальным элементами = %d\n", sum) //печатаем результат
}
