package main

import (
	"fmt"
	"strconv"
	"strings"
)

var romanInt = map[string]int{
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var intRoman = map[int]string{
	10: "X",
	9:  "IX",
	8:  "VIII",
	7:  "VII",
	6:  "VI",
	5:  "V",
	4:  "IV",
	3:  "III",
	2:  "II",
	1:  "I",
}

var a, b *int

var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

const (
	LOW  = "Ошибка. Cтрока не является математической операцией."
	HIGH = "Ошибка. Формат математической операции " +
		"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Ошибка. Используются одновременно разные системы счисления."
	DIV   = "Ошибка. В римской системе нет отрицательных чисел."
	ZERO  = "Ошибка. В римской системе нет числа 0."
	RANGE = "Ошибка. Калькулятор умеет работать только с арабскими целыми " +
		"числами или римскими цифрами от 1 до 10 включительно"
)

func main() {
	fmt.Println("Welcome to calculator, enter the string:")

	var text string
	fmt.Scan(&text)
	s := strings.ReplaceAll(text, " ", "")
	calculate(strings.ToUpper(strings.TrimSpace(s)))
}

func calculate(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}

	switch {
	case len(operator) > 1:
		panic(HIGH)
	case len(operator) < 1:
		panic(LOW)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(SCALE)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11 // должно выполняться условие число > 0 и число < 11, второе аналогично
		if val, ok := operators[operator]; ok && errCheck {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(RANGE)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := romanInt[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(RANGE)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}

func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(ZERO)
	} else if romanResult < 0 {
		panic(DIV)
	}

	for romanResult > 0 {
		for i := 10; i > 0; i -= 1 {
			v := min(i, romanResult)
			val, ok := intRoman[v]
			if ok {
				romanNum += val
				romanResult -= v
				break
			} else {
				panic("not found roman")
			}
		}
	}

	fmt.Println(romanNum)
}
