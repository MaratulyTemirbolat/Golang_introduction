package main // Объявлем имя пакета. Пакеты позволяют дробить наше приложение на небольшие участки.
// Main - говорит компилятору, что данный пакет является точкой входа. все начинается с функции main

import (
	"fmt"
	"reflect"
	"errors"
	) // импортируем другие пакеты, которые можем использовать

// С помощью import можно импортировать пакеты 3-х типов: 1) пакеты из стандартной библиотеки Go
// 2) другие пакеты нашего приложения, которые мы сами описали
// 3) сторонние пакеты, которые находятся в сети

/* Переменные
   Типы данных (Скорее всего примитивные ток имелись ввиду): Числа, строки, байты и bool и т.д.
   Go - язык со статической типизацией ( Каждая переменная имеет фиксированный тип)
   Способы инициализации:
	   1. Инициализации "var" после чего идет название переменной (var varName = "Hello!")
	   2. Название переменной := значение ( тоже самое, что и 1-ый пункт) (Пример: varName := "Hello!")
		   Вообще : используется только для инициализа
	   3. Случай, когда мы не можем сразу указать значение переменной, но сделаем это позже.
		   Делается следующее: var varName string
		   Затем уже где-нибудь в коде дальше varName = 'Hello!'
		   Если создать переменную, но не без указания её значения, то ей будет присвоено нулевое значение
		   Пример: int --> 0, string --> "", float32 --> 0, bool --> false и т.д


	Создание констант:
		Синтаксис: const myVariable ...
			Вместо ... может быть ( data type = value  || = value)



	Функции:
		Каждая программа на Go должна иметь функцию main. Иначе программа будет выдавать ошибку.
		Функция main не имеет ни аргументов, ни возвращающий значений.
		Функции в Go можно создавать в разных местах, нет необходимости создавать её выше её вызова.


	Указатели:
		Указатель - тоже переменная, только она хранит адрес на участок в памяти, а не само значение
		Бывают Указатели со значение nil (Это указатели, которые не указывают ни на какую область в памяти 
			и не ссылаются ни на какую переменную)

	Встроенный тип "error":
		Он может иметь значение nil или содержать объект ошибки

	Массивы:
		Это последовательность фиксированной длины из 0 или более элементов
		Синтаксис: arrayName := [size]dataType{elementA,elementB,...,elementN}
		len(arrayName) -> длина массива

	for:
		Синтаксис: 
		1) Если просто индексом 
			for indexName := range arrayName{

			}
			каждый текущий индекс в variableName

		2) Если индекс и значение под ним 
			for indexName. valueName := range arrayName{

			}

		3) Стандартный цикл for из Си:
			for i := 0; i <= 10; i++{

			} 

	Слайсы (Срезы): Как я понял, это как list в Python
		Это почти те же массивы, однако их длина не фиксирована, а динамическая
		Синтаксис: sliceName := []dataType{elementA,elementB,...,elementN}
		Функции: sliceName = append(sliceName,elementC,elementD,...)

	Map:
		Это набор элементов со структурой ключ + значение. Как dict в Python
		Синтаксис: mapName := map[keyDataType]valueDataType{keyA:valueA,keyB:valueB,...}

		Итерация по map:
			for keyName, valueName := range mapName {

			} 

		Удаление: 
			delete(mapName,keyName)

	Структуры и их методы:
		Структуры помогают объявить собственные типы и описывать необходимые поля
		Синтаксис ниже


*/

// Структура
type contact struct{
	firstName string
	lastName string
	phoneNumber string
	email string
	address string
	dateOfBirth string
}
// Метод для структуры
func (c contact) printInfo(){ // Перед именем метода мы указываем структуру в круглых скобка, для кого мы объявляем метод
	// и так называемый reciever. Reciever - это копия объекта к которой мы можем обращаться внутри метода
	// Cуществует 2 типа reciever(а). 1) Мы можем передавать значение в роли reciever(а), как мы делаем в printInfo
	// 2) Также мы можем отправлять указатель в роли ресивера. Пример setName
	fmt.Printf("Имя: %s\nФамилия: %s\nНомер телефона: %s\nEmail: %s\nАдрес: %s\n",
		c.firstName,c.lastName,c.phoneNumber,c.email,c.address)
}

func (c *contact) setName(name string){
	c.firstName = name
}

func main() {
	// Initialization

	// Variant A
	var messageOne = "Привет, Go! It's first variant.\n"
	fmt.Println(messageOne)

	// Variant B
	messageTwo := "Hello, Go! It's second variant.\n"
	fmt.Println(messageTwo)

	// Variant C
	var messageThree string
	messageThree = "Hello, Go! It's third variant"
	fmt.Println(messageThree)

	// Different initializations
	var name string
	var age int
	var weight float32
	var isAdult bool

	fmt.Println("Необъявленное значение переменной \"name\":", name)
	fmt.Println("Необъявленное значение переменной \"age\":", age)
	fmt.Println("Необъявленное значение переменной \"weight\":", weight)
	fmt.Println("Необъявленное значение переменной \"isAdult\":", isAdult)

	// Constants possible initialization

	const myConstantA int = 50
	const myConstantB = "Trial constant B"
	fmt.Println("\nМоя созданная константа A:", myConstantA)
	fmt.Println("Моя созданная константа B:", myConstantB)

	printPersonInfo("Темирболат", 21, 90.2)

	
	fmt.Println(isPerson(15))

	// Pointers
	name = "Temirbolat"
	var pName *string = &name // Pointer
	fmt.Println("")
	fmt.Println(reflect.TypeOf(pName))
	fmt.Println(pName)
	fmt.Println(*pName) // Dereferencing или разыменование

	*pName = "Asanali" // Меняем также значенеи оригинала на который ссылаемся
	fmt.Println(name)
	name = "Antony" // у pName тоже поменяется, так как он на него ссылается. Короче, хоть одно меняется, то второе тоже
	fmt.Println(*pName)

	var notSetName string = "Temirbolat"
	fmt.Println("\nName before set is:",notSetName)
	setNewName(&notSetName)
	fmt.Println("Name after set is:",notSetName)


	// Встроенный тип "error":
	change_a, err_a := buyWine(20,110) 
	if err_a != nil{
		fmt.Println("\nПокупка неуспешна:",err_a.Error(), reflect.TypeOf(err_a))
	} else{
		fmt.Printf("\nВаша сдача - %d. Хорошего дня!",change_a)
		fmt.Print(reflect.TypeOf(err_a),"\n")
	}

	change_b, err_b := buyWine(17,110) 
	if err_b != nil{
		fmt.Println("Покупка неуспешна:",err_b.Error(),reflect.TypeOf(err_b))
	} else{
		fmt.Printf("Ваша сдача - %d. Хорошего дня!",change_b)
		fmt.Print(reflect.TypeOf(err_b),"\n")
	}

	change_c, err_c := buyWine(20,90) 
	if err_c != nil{
		fmt.Println("Покупка неуспешна:",err_c.Error(),reflect.TypeOf(err_c))
	} else{
		fmt.Printf("Ваша сдача - %d. Хорошего дня!",change_c)
		fmt.Print(reflect.TypeOf(err_c),"\n")
	}


	fmt.Println("")
	// Создание экземпляра структуры на основе выше шаблона
	var structName contact = contact{
		firstName: "Вася",
		lastName: "Пупкин",
		phoneNumber: "8778",
		email: "snake@mail.ru",
		address: "Tole Bi 59",
		dateOfBirth: "31.01.2001",
	}

	structName.setName("Петя")
	structName.printInfo()
}

func printPersonInfo(name string, age int, weight float32) {
	fmt.Printf("Имя: %s \nВозраст: %d \nВес: %.1f \n ", name, age, weight) // Вывести форматированный текст
	// Для форматирования используется символ %,
	// а после него указывается буква(обозначает тип переменной, которую хотим подставить в строку)
	// s -> string, d -> int (целые числа), f - float
	// После того, как указали строку, то через запятую указываются значения, которые хотим подставить в данную строку
	// .1 означает вывод 1 числа после запятой
}


func isPerson(age int) bool {
	const adultAge int = 18
	return age >= adultAge
}

func setNewName(name *string){ // Передаем указатель
	*name = "New Temirbolat"
}

func buyWine(age int,moneyAmount int) (int,error){
	const winePrice int = 100
	if age >= 18 && moneyAmount >= winePrice{
		return moneyAmount - winePrice, nil
	}else if age < 18 {
		return moneyAmount, errors.New("Иди пей вишневый сок, сынуля")
	}else{
		return moneyAmount, errors.New("У вас недостаточно средств")
	}

}