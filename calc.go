package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	roman = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	arabic = map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}

	table = [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
)

func ScanCalc() string { // Никак не мог найти способ передать строку с пробелами через консоль (она просто передавала все до первого пробела). Нашел в интернете единственный способ, причину ошибки не нашел
	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Что Вам нужно посчитать?")
	in.Scan()
	return in.Text()
}

func arab2roman(arab int) string { // Также код из интернета, так как не мог самостоятельно придумать адекватный способ перевода арабских чисел в римские, а писать map со 100 элементами в формате "ключ-значение" полная жесть :)

	var (
		roman = ""
		digit = 1000
	)
	for i := 3; i >= 0; i-- {
		d := arab / digit
		roman += table[i][d]
		arab %= digit
		digit /= 10
	}

	return roman
}

func main() {

	numbers := strings.Fields(ScanCalc())

	if len(numbers) >= 4 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if len(numbers) <= 2 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
	} else {

		if _, ok := roman[numbers[0]]; ok {

			if _, ok := roman[numbers[2]]; ok {

				number1 := roman[numbers[0]]
				number2 := roman[numbers[2]]

				var res int

				if numbers[1] == "+" {
					res = number1 + number2
				} else if numbers[1] == "-" {
					res = number1 - number2
				} else if numbers[1] == "*" {
					res = number1 * number2
				} else if numbers[1] == "/" {
					res = number1 / number2
				}

				if res <= 0 {
					fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.") // момент про исключения непонял - написал через условие
				} else {
					res1 := arab2roman(res)
					fmt.Println(res1)
				}

			} else if _, ok := arabic[numbers[2]]; ok {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			}

		}

		if _, ok := arabic[numbers[0]]; ok { //К сожалению, не получилось по какой-то причине передать срез как аргумент в функцию, поэтому есть повторяющийся код

			if _, ok := arabic[numbers[2]]; ok {

				number1 := arabic[numbers[0]]
				number2 := arabic[numbers[2]]

				var res int

				if numbers[1] == "+" {
					res = number1 + number2
				} else if numbers[1] == "-" {
					res = number1 - number2
				} else if numbers[1] == "*" {
					res = number1 * number2
				} else if numbers[1] == "/" {
					res = number1 / number2
				}

				fmt.Print(res)

			} else if _, ok := roman[numbers[2]]; ok {
				fmt.Print("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			}

		}
	}

}
