package main

import "fmt"

func main() {
	// https://leetcode.com/problems/integer-to-roman/description/
	// Input: num = 1994
	// Output:"MCMXCIV"
	num := 1994
	fmt.Println(intToRoman(num))

}
func intToRoman(num int) string {
	// Definir o mapa de valores romanos
	romanMap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	roman := ""
	for value, symbol := range romanMap {
		// Verificar quantas vezes o símbolo pode ser subtraído do número atual
		if num >= value {
			// Adicionar o símbolo ao resultado
			roman += symbol
			// Subtrair o valor do número atual
			num -= value
		}
	}

	return roman
}
