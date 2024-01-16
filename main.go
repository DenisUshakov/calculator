package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanInt = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
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
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	8:   "VIII",
	7:   "VII",
	6:   "VI",
	5:   "V",
	4:   "IV",
	3:   "III",
	2:   "II",
	1:   "I",
}

var intMain = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			return
		} else {
			s := strings.ReplaceAll(text, " ", "")
			calculate(strings.ToUpper(strings.TrimSpace(s)))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
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
		for _, elem := range intMain {
			v := min(elem, romanResult)
			val, ok := intRoman[v]
			if ok {
				romanNum += val
				romanResult -= v
				break
			}
		}
	}

	fmt.Println(romanNum)
}
