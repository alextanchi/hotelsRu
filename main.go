package main

import "fmt"

func main() {
	//fmt.Println(Task1FormatComputers(99)) // Тестовый ввод для 1-го задания

	//arr := []int{10, 20, 50} // Тестовый ввод для 2-го задания
	//fmt.Println(Task2CommonDivisor(arr))

	//min, max := 11, 21 // Тестовый ввод для 2-го задания
	//fmt.Println(Task3FindPrimesInRange(min, max))

	Task4PrintMultiplicationTable(5)

}

func Task1FormatComputers(n int) string { //задание 1
	count := fmt.Sprintf("%d", n)
	var ending string
	switch n % 10 {
	case 2, 3, 4:
		ending = "компьютера"
	case 5, 6, 7, 8, 9, 0:
		ending = "компьютеров"
	case 1:
		ending = "компьютер"
	}
	if n >= 11 && n <= 19 {
		ending = "компьютеров"
	}
	return count + " " + ending
}

func Task2CommonDivisor(nums []int) []int { //задание 2
	if len(nums) == 0 {
		return []int{}
	}
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	commonDivisors := []int{}
	for i := 1; i <= minNum; i++ {
		isCommonDivisor := true
		for _, num := range nums {
			if num%i != 0 {
				isCommonDivisor = false
				break
			}
		}
		if isCommonDivisor {
			commonDivisors = append(commonDivisors, i)
		}
	}
	return commonDivisors
}

func Task3FindPrimesInRange(min, max int) []int { //задание 3
	if min <= 1 { //проверка на простое число - если это так, она устанавливает min равным 2, потому что 1 не является простым числом.
		min = 2
	}
	primes := make([]int, 0)
	for i := min; i <= max; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func Task4PrintMultiplicationTable(n int) { //задание 4
	// Создаем таблицу умножения
	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	// Выводим таблицу в консоль
	for _, row := range table {
		for _, num := range row {
			fmt.Printf("%2d ", num)
		}
		fmt.Println()
	}
}
