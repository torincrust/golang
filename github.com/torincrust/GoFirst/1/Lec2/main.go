package main

import (
	"fmt"
	//"math"
)

func main() {
	/* // простейший вывод в консоль printl вывод на новую строку +\n
	fmt.Println("hello world")
	fmt.Println("second line")
	// функция принт простой вывод без переносов
	fmt.Print("first")
	fmt.Print("sec")
	fmt.Print("third")

	// форматированный вывод Printf - стандартный вывод с флагами форматирвоания на место флага форматирования %s помещается аргумент
	fmt.Printf("\nhello my name is %s\nMy age is %d\n", "bob", 33)
	// типо дз  */
	/* fmt.Println("Привет! Я новый студент!")
	fmt.Println("Это моя вторая программа! Я рад, что учусь здесь!!") */

	// декларирование переменных ********************************************************

	/* 	var age int
	   	fmt.Println("my age is:", age)
	   	age = 33
	   	fmt.Println("age after init:", age) */

	/* 	// дикларирование и инициализация пользовательских значений
	   	var height int = 123
	   	fmt.Println("my height:", height)

	   	//в чем полустрогость типизации
	   	var uid = 1234
	   	fmt.Println("my uid:", uid)
	   	fmt.Printf("%T\n", uid) // скажет какой тип данных в переменной

	   	// Присваивание переменных одного типа множественный случай

	   	var frst, scnd int = 20, 30
	   	fmt.Println(frst, scnd)
	   	fmt.Printf("FirstVar:%d, SecondVar:%d\n", frst, scnd)

	   	//декларирование блок переменных
	   	var (
	   		name string = "bob"
	   		age  int    = 31

	   		marriege bool = false
	   	)
	   	fmt.Printf("name: %s\nAge: %d\nMarried: %t\n", name, age, marriege)
	   	var a, b = 30, "bob"
	   	a = 300
	   	println(a, b)
	   	//var a = 300 // повторное декларирование переменной невозможно приводит к ошибки

	   	// короткая декларация (короткое обьявление)

	   	count := 10
	   	fmt.Println("Count:", count)
	   	count++
	   	fmt.Println("Count:", count)

	   	// множественные присваивания через :=

	   	aArg, bArg := 10, 30
	   	fmt.Println(aArg, bArg)
	   	// aArg, bArg := 20, "robik" // повторная передекларация приведет к ошибке
	   	// fmt.Println(aArg, bArg)

	// исключения

	bArg, cArg := 300, 400
	fmt.Println(bArg, cArg, aArg) // если хоть одна новая переменная то можно передекларировать */

	// пример
	/* width, length := 20.5, 30.5
	fmt.Printf("min dumensional of rectangle is: %.3f\n ", math.Min(width, length)) // %.3f подстановка дробного числа с точностью 3 знака после запятой */

	// задача
	word1, word2, word3, word4 := "имя", "твое", "мне", "знакомо"

	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word1, word3)

}
