// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import "fmt"

func main() {
	s := "abcabcbb" // Output: 3
	ss := "pwwkew"  // Output: 3
	// Explanation: The answer is "wke", with the length of 3.
	// Notice that the answer must be a substring,
	// "pwke" is a subsequence and not a substring.
	Output := lengthOfLongestSubstring(s)
	fmt.Println(Output)
	Output = lengthOfLongestSubstring(ss)
	fmt.Println(Output)
}

func lengthOfLongestSubstring(s string) int {
	// maxLength armazena o comprimento da substring mais longa sem caracteres repetidos
	maxLength := 0
	// start indica o índice inicial da substring atual sem caracteres repetidos
	start := 0
	// lastSeen é um mapa que armazena o último índice em que cada caractere foi visto
	lastSeen := make(map[string]bool)

	// Percorrer a string s
	for indexChar := 0; indexChar < len(s); indexChar++ {

		currentChar := string(s[indexChar]) // currentChar armazena o caractere atual da string s

		// Verificar se o caractere atual já foi visto antes
		if lastPos, found := lastSeen[currentChar]; found && lastPos {

			// Se o caractere atual foi visto anteriormente...
			for string(s[start]) != currentChar { // percorre a string a partir do início até encontrar uma ocorrência do caractere atual
				delete(lastSeen, string(s[start])) // Remove os caracteres vistos anteriores do mapa
				start++                            // Avança o índice de início para frente
			}

			start++ // Avança o índice de início para frente novamente para evitar a repetição do caractere atual
		}

		lastSeen[currentChar] = true // Marca o índice do caractere atual como visto no mapa

		// Atualiza maxLength com o comprimento da substring atual, se necessário
		if indexChar-start+1 > maxLength {
			maxLength = indexChar - start + 1
		}
	}
	// Retorna o comprimento da substring mais longa sem caracteres repetidos
	return maxLength
}
