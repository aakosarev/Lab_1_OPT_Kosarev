package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//var infix14 = "(l - k - j - h*g + f/d/s)/(a*p*o + i + u - y + t)"

var infix14 = "(a+b*c)"

type Stack []string

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Push(str string) {
	*st = append(*st, str)
}

func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1
		*st = (*st)[:index]
		return true
	}
}

func (st *Stack) Top() string {
	if st.IsEmpty() {
		return ""
	} else {
		index := len(*st) - 1
		element := (*st)[index]
		return element
	}
}

func IsLetter(le string) bool {
	if le >= "a" && le <= "z" {
		return true
	}
	return false
}

func IsNumber(nu string) bool {
	if nu >= "0" && nu <= "9" {
		return true
	}
	return false
}

func IsBracket(br string) bool {
	if br == "(" || br <= ")" {
		return true
	}
	return false
}

func IsOperation(op string) bool {
	if op == "+" || op == "-" || op == "*" || op == "/" {
		return true
	}
	return false
}

func Priority(op string) int {
	if (op == "/") || (op == "*") {
		return 2
	} else if (op == "+") || (op == "-") {
		return 1
	} else {
		return -1
	}
}

func IsCorrectBrackets(infix string) bool {
	count := 0
	for char := range infix {
		if char == '(' {
			count += 1
		}
		if char == ')' {
			count -= 1
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

func IsCorrect(infix string) bool {
	if !IsCorrectBrackets(infix) {
		return false
	}
	if (IsOperation(string(infix[0])) && infix[0] != '-') || IsOperation(string(infix[len(infix)-1])) {
		return false
	}
	for i := 0; i < len(infix); i++ {
		if !IsLetter(string(infix[i])) && !IsOperation(string(infix[i])) && !IsNumber(string(infix[i])) && !IsBracket(string(infix[i])) {
			return false
		}
	}
	for i := 0; i < len(infix)-1; i++ {
		if infix[i] == '(' {
			if !IsLetter(string(infix[i+1])) && !IsNumber(string(infix[i+1])) && infix[i+1] != '-' && infix[i+1] != '(' {
				return false
			}
		}
		if infix[i] == ')' {
			if !IsOperation(string(infix[i+1])) && !IsNumber(string(infix[i+1])) && infix[i+1] != '-' && infix[i+1] != ')' {
				return false
			}
		}
		if IsOperation(string(infix[i])) {
			if IsOperation(string(infix[i+1])) {
				return false
			}
			if infix[i+1] == ')' {
				return false
			}
		}
		if IsLetter(string(infix[i])) {
			if IsLetter(string(infix[i+1])) {
				return false
			}
			if infix[i+1] == '(' {
				return false
			}
		}
	}
	return true
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Parse(infix string) []string {
	separators := []string{"+", "-", "*", "/", "(", ")"}
	var infixTokens []string
	token := ""
	for _, char := range infix {
		if string(char) != " " {
			if Contains(separators, string(char)) {
				if token != "" {
					infixTokens = append(infixTokens, token)
				}
				infixTokens = append(infixTokens, string(char))
				token = ""
			} else {
				token += string(char)
			}
		}
	}
	if token != "" {
		infixTokens = append(infixTokens, token)
	}
	return infixTokens
}

func InfixToPostfix(infix string) []string {
	infixTokens := Parse(infix)
	var postfixTokens []string
	var sta Stack

	for _, token := range infixTokens {
		if IsLetter(token) || IsNumber(token) {
			postfixTokens = append(postfixTokens, token)
		} else if token == "(" {
			sta.Push(token)
		} else if token == ")" {
			for sta.Top() != "(" {
				postfixTokens = append(postfixTokens, sta.Top())
				sta.Pop()
			}
			sta.Pop()
		} else {
			for !sta.IsEmpty() && Priority(token) <= Priority(sta.Top()) {
				postfixTokens = append(postfixTokens, sta.Top())
				sta.Pop()
			}
			sta.Push(token)
		}
	}
	for !sta.IsEmpty() {
		postfixTokens = append(postfixTokens, sta.Top())
		sta.Pop()
	}
	return postfixTokens
}

func CalculateRPN(postfixTokens []string) (int, error) {
	var sta Stack
	var value int
	for _, token := range postfixTokens {
		switch token {
		case "+", "-", "*", "/":
			right, _ := strconv.Atoi(sta.Top())
			sta.Pop()
			left, _ := strconv.Atoi(sta.Top())
			sta.Pop()

			switch token {
			case "+":
				value = left + right
			case "-":
				value = left - right
			case "*":
				value = left * right
			case "/":
				if right == 0 {
					return 0, errors.New("ошибка деления на 0")
				}
				value = left / right
			}
		default:
			value, _ = strconv.Atoi(token)
		}
		sta.Push(strconv.Itoa(value))
	}
	result, _ := strconv.Atoi(sta[0])

	return result, nil
}

func Check(l, k, j, h, g, f, d, s, a, p, o, i, u, y, t int) int {
	return (l - k - j - h*g + f/d/s) / (a*p*o + i + u - y + t)
}

func ShowMenu() {
	fmt.Print("\n\n")
	fmt.Println("Выберите нужный пункт:")
	fmt.Println("1: Решить задачу c выражением из индивидуального варианта")
	fmt.Println("2: Ввести своё выражение и решить задачу")
	fmt.Println("3: Выйти из программы")
	fmt.Print("\n")
	fmt.Print("Ввод:  ")
}

func main() {
	var selection string

	for selection != "3" {

		ShowMenu()

		_, err := fmt.Scanf("%s\n", &selection)
		if err != nil {
			log.Fatal(err)
		}

		switch selection {

		case "1":
			fmt.Printf("Инфиксная запись   :    %s\n", infix14)
			infix14 = strings.Join(strings.Fields(infix14), "")
			infix14 = strings.ToLower(infix14)

			m := make(map[string]string)
			var valueList []string

			if IsCorrect(infix14) {
				postfix := InfixToPostfix(infix14)
				fmt.Printf("Постфиксная запись :    %s\n", strings.Join(postfix, " "))

				for _, token := range postfix {
					if IsLetter(token) {
						_, ok := m[token]
						if !ok { // тут бахать бесконечный цикл для правильного ввода
							fmt.Print(token, ": ")
							var num string
							for {
								_, err := fmt.Scanf("%s\n", &num)
								if err != nil {
									log.Fatal(err)
								}
								if _, err := strconv.Atoi(num); err != nil {
									fmt.Println("Некорректный ввод. Необходимо ввести число!")
									fmt.Print(token, ": ")
								} else {
									break
								}
							}
							m[token] = num
							valueList = append(valueList, m[token])
						} else {
							valueList = append(valueList, m[token])
						}
					} else {
						valueList = append(valueList, token)
					}
				}
				result, err := CalculateRPN(valueList)
				if err != nil {
					fmt.Println("Ошибка: ", err.Error())
				}
				fmt.Println("Результат: ", result)
				fmt.Println("Результат средствами языка: ", Check(1, 1, 1, 1000, 1000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1))

			} else {
				fmt.Println("Некорректный ввод")
			}

		case "2":
			fmt.Println("Введите выражение в инфиксной форме: ")
			fmt.Printf("Инфиксная запись   :    %s\n", infix14)
			infix14 = strings.Join(strings.Fields(infix14), "")
			infix14 = strings.ToLower(infix14)

			m := make(map[string]string)
			var valueList []string

			if IsCorrect(infix14) {
				postfix := InfixToPostfix(infix14)
				fmt.Printf("Постфиксная запись :    %s\n", strings.Join(postfix, " "))

				for _, token := range postfix {
					if IsLetter(token) {
						_, ok := m[token]
						if !ok { // тут бахать бесконечный цикл для правильного ввода
							fmt.Print(token, ": ")
							var num string
							fmt.Scanf("%s\n", &num)
							m[token] = num
							valueList = append(valueList, m[token])
						} else {
							valueList = append(valueList, m[token])
						}
					} else {
						valueList = append(valueList, token)
					}
				}
				result, err := CalculateRPN(valueList)
				if err != nil {
					fmt.Println("Ошибка: ", err.Error())
				}
				fmt.Println("Результат: ", result)
				fmt.Println("Результат средствами языка: ", Check(1, 1, 1, 1000, 1000, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1))

			} else {
				fmt.Println("Некорректный ввод")
			}

		case "3":
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Введено некорректное значение! Введите еще раз:")
		}
	}
}
