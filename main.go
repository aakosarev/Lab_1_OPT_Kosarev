package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	for i := 0; i < len(infix); i++ {
		if string(infix[i]) == "(" {
			count += 1
		}
		if string(infix[i]) == ")" {
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
	fmt.Println("1: Решить задачу")
	fmt.Println("2: Выйти из программы")
	fmt.Print("\n")
	fmt.Print("Ввод:  ")
}

func ShowTaskMenu() {
	fmt.Print("\n")
	fmt.Println("Выберите задачу:")
	fmt.Println("1) Решить задачу с выражением из индивидуального варианта")
	fmt.Println("2) Решить задачу с выражением, вводимым пользователем")
	fmt.Print("\n")
	fmt.Print("Ввод:  ")
}

func main() {
	var selection string
	for selection != "2" {
		ShowMenu()
		_, err := fmt.Scanf("%s\n", &selection)
		if err != nil {
			log.Fatal(err)
		}
		switch selection {
		case "1":
			var task string
			var infix string
			for flag := true; flag; {
				ShowTaskMenu()
				_, err := fmt.Scanf("%s\n", &task)
				if err != nil {
					log.Fatal(err)
				}
				switch task {
				case "1":
					infix = "(a+b*c))"
					flag = false
				case "2":
					fmt.Print("Введите выражение в инфиксной форме: ")
					reader := bufio.NewReader(os.Stdin)
					infix, _ = reader.ReadString('\n')
					if err != nil {
						log.Fatal(err)
					}
					infix = strings.TrimSuffix(infix, "\n")
					flag = false
				default:
					fmt.Println("Введено некорректное значение! Введите еще раз:")
				}
			}
			fmt.Printf("\nИнфиксная запись   :    %s\n", infix)
			infix = strings.Join(strings.Fields(infix), "")
			infix = strings.ToLower(infix)
			m := make(map[string]string)
			var valueList []string
			if IsCorrect(infix) {
				fmt.Println("КОРРЕКТНО!")
				postfix := InfixToPostfix(infix)
				fmt.Printf("Постфиксная запись :    %s\n", strings.Join(postfix, " "))
				for _, token := range postfix {
					if IsLetter(token) {
						_, ok := m[token]
						if !ok {
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
				fmt.Println("Некорректное выражение")
			}
		case "2":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Введено некорректное значение! Введите еще раз:")
		}
	}
}
