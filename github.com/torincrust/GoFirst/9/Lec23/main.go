package main

import (
	_ "Lec23/circle"
	"Lec23/rectangle"
	"fmt"
)

//1. Функци init() - данная функция отрпбатывает единожды при первом импортировании пакета
//2. Данных фукнций в пакете может быть несколько штук! (не в одном модуле, т.к. в одном модуле нельзя создать
// более одной функции с каким-то определенным именем)

//3. init() вызывает в момент инициализации пакета:
//* Данный процесс выглядит следующим образом:
// ** сначала компилятор смотрит на содержимое пакета
// ** затем компилятор смотрит на пути импорта (если что-то импортируется, компилятор уходит туда)
// ** затем компилятор инициализурет переменные уровня пакета
// ** затем компилятор запускает функцию init() для данного пакета
// ** повторяет данную процедуру для всех пакетов проекта
// ** после чего вызывается функция main()

//4. Что произойдет , если запустить go run main.go
// * Сначала смотрим в main.go на предмет синтаксических ошибок ,но ничего не инициализируется
// * Затем импорты : сначала импоритруем Lec23/rectangle
// ** Компилятор идет в rectangle
// ** Смотрим в пакет на предмет синтаксических ошибок
// ** Затем импорт fmt
// ** Затем инициализируем переменные уровня пакета
// ** Затем запускаем функцию init() пакета rectangle
// ** Затем подружаем (определяем) все имена пакета rectangle
// **Функции main тут нет, возвращаемся назад
// * Пытаемся импоритровать fmt (т.к. он уже был импортирован одним из пакетов - повторноый импорт не требуется)
// * Инициализируем переменные уровня пакета main
// * Запускаем функцию init() в main
// * Затем определяем имена (тут дополнительных имен нет, тут ничего не делаем)
// * Затем запускаем функцию main()

//5. Все импорты (вне зависимости , стандартные или пользовательские) сортируются по алфавиту
func init() {
	fmt.Println("Init function for main package!")
	fmt.Println("Name:", name, "Age:", age)
}

var (
	name string = "Bob"
	age  int    = 99
)

func main() {
	r := rectangle.New(10, 10)
	fmt.Println(r)

}
