package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1. Программа, которая создает карту с именами людей и их возрастами
func createAndDisplayMap() {
	people := map[string]int{
		"Arkadii": 30,
		"Steve":   25,
		"Oleg":    40,
	}

	people["Ivan"] = 35

	fmt.Println("Люди и их возраст:")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}
}

// 2. Функция, которая принимает карту и возвращает средний возраст
func averageAge(people map[string]int) float64 {
	totalAge := 0
	for _, age := range people {
		totalAge += age
	}
	return float64(totalAge) / float64(len(people))
}

// 3. Программа, которая удаляет запись из карты по заданному имени
func deletePerson(people map[string]int, name string) {
	delete(people, name)
	fmt.Printf("Человек с именем %s удален.\n", name)
}

// 4. Программа, которая считывает строку с ввода и выводит её в верхнем регистре
func toUpperCase() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	upper := strings.ToUpper(strings.TrimSpace(input))
	fmt.Printf("Строка в верхнем регистре: %s\n", upper)
}

// 5. Программа, которая считывает несколько чисел и выводит их сумму
func sumNumbers() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через пробел: ")
	input, _ := reader.ReadString('\n')
	numbers := strings.Fields(input)

	sum := 0
	for _, number := range numbers {
		num, _ := strconv.Atoi(number)
		sum += num
	}

	fmt.Printf("Сумма чисел: %d\n", sum)
}

// 6. Программа, которая считывает массив целых чисел и выводит их в обратном порядке
func reverseArray() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите массив чисел через пробел: ")
	input, _ := reader.ReadString('\n')
	numbers := strings.Fields(input)

	fmt.Print("Числа в обратном порядке: ")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Printf("%s ", numbers[i])
	}
	fmt.Println()
}

func main() {
	// 1. Создать карту с именами и возрастами
	createAndDisplayMap()

	// 2. Пример расчета среднего возраста
	people := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 40,
	}
	avg := averageAge(people)
	fmt.Printf("Средний возраст: %.2f\n", avg)

	// 3. Удаление записи по имени
	deletePerson(people, "Oleg")

	// 4. Перевод строки в верхний регистр
	toUpperCase()

	// 5. Ввод нескольких чисел и вывод их суммы
	sumNumbers()

	// 6. Ввод массива чисел и вывод в обратном порядке
	reverseArray()
}
