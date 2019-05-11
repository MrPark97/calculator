package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
	"math"
)

type table map[string]float64

var (
	endline  = "\r\n"
	id_table = make(map[string]float64)
)

func checkId(id string) (string, error) {
	if match, _ := regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9_]*", id); match {
		return id, nil
	} else {
		return "", &wrongIdError{"wrong id syntax"}
	}

	return "", nil
}

func calculate(s string, local_table *table) (result float64, err error) {
	for index, runeValue := range s {
		if runeValue == '=' {
			if id, err := checkId(strings.Trim(s[:index], " ")); err != nil {
				return 0.0, err
			} else {
				if result, err = calculate(strings.Trim(s[index+1:], " "), local_table); err != nil {
					return 0.0, err
				} else {
					(*local_table)[id] = result
					return result, err
				}
			}
		}
	}

	parenthesesLevel := 0
	for _, runeValue := range s {
		if runeValue == '(' {
			parenthesesLevel++
		} else if runeValue == ')' {
			parenthesesLevel--
			if parenthesesLevel < 0 {
				return 0.0, &parenthesesError{"par error"}
			}
 		}
	}

	if parenthesesLevel != 0 {
		return 0.0, &parenthesesError{"par error"}
	}

	if s[0] == '(' && s[len(s)-1] == ')' {
		result, err = calculate(s[1:len(s)-1], local_table)
		return result, err
	}

	for index := len(s)-1; index>=0; index-- {
		runeValue := s[index]
		if runeValue == '(' {
			parenthesesLevel++
		} else if runeValue == ')' {
			parenthesesLevel--
		}

		if parenthesesLevel == 0 && (runeValue == '+' || runeValue == '-') {
			left_result, err := calculate(strings.Trim(s[:index], " "), local_table)
			if err != nil {
				return 0.0, err
			}

			right_result, err := calculate(strings.Trim(s[index+1:], " "), local_table)
			if err != nil {
				return 0.0, err
			}

			if(runeValue == '+') {
				return left_result + right_result, nil
			} else {
				return left_result - right_result, nil
			}
		}
	}

	for index := len(s)-1; index>=0; index-- {
		runeValue := s[index]
		if runeValue == '(' {
			parenthesesLevel++
		} else if runeValue == ')' {
			parenthesesLevel--
		}
		if parenthesesLevel == 0 && (runeValue == '*' || runeValue == '/') {
			left_result, err := calculate(strings.Trim(s[:index], " "), local_table)
			if err != nil {
				return 0.0, err
			}

			right_result, err := calculate(strings.Trim(s[index+1:], " "), local_table)
			if err != nil {
				return 0.0, err
			}


			if(runeValue == '*') {
				return left_result * right_result, nil
			} else {
				if right_result == 0.0 {
					return 0.0, &zeroDivisionError{"division by zero error"}
				}
				return left_result / right_result, nil
			}
		}
	}

	for index := len(s)-1; index>=0; index-- {
		runeValue := s[index]
		if runeValue == '(' {
			parenthesesLevel++
		} else if runeValue == ')' {
			parenthesesLevel--
		}
		if parenthesesLevel == 0 && runeValue == '^' {
			left_result, err := calculate(strings.Trim(s[:index], " "), local_table)
			if err != nil {
				return 0.0, err
			}

			right_result, err := calculate(strings.Trim(s[index+1:], " "), local_table)
			if err != nil {
				return 0.0, err
			}
			return  math.Pow(left_result, right_result), nil
		}
	}

	if f, err := strconv.ParseFloat(s, 64); err != nil {
		if _, ok := id_table[s]; ok {
			return id_table[s], nil;
		} else {
			return 0.0, &presentationError{"presentation error"}
		}
	} else {
		return f, nil
	}

	return 0.0, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmdString = strings.TrimSuffix(cmdString, endline)
		local_table := make(table)
		if result, err := calculate(cmdString, &local_table); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			for k, v := range local_table {
				id_table[k] = v
			}
			fmt.Fprintln(os.Stdout, result)
		}
	}
}
