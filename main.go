package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var morseCode = map[string]string{
	// Russian
	"А": ".-", "Б": "-...", "В": ".--", "Г": "--.", "Д": "-..", "Е": ".",
	"Ж": "...-", "З": "--..", "И": "..", "Й": ".---", "К": "-.-", "Л": ".-..",
	"М": "--", "Н": "-.", "О": "---", "П": ".--.", "Р": ".-.", "С": "...",
	"Т": "-", "У": "..-", "Ф": "..-.", "Х": "....", "Ц": "-.-.", "Ч": "---.",
	"Ш": "----", "Щ": "--.-", "Ъ": "--.--", "Ы": "-.--", "Ь": "-..-",
	"Э": "..-..", "Ю": "..--", "Я": ".-.-",

	// English
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".", "F": "..-.",
	"G": "--.", "H": "....", "I": "..", "J": ".---", "K": "-.-", "L": ".-..",
	"M": "--", "N": "-.", "O": "---", "P": ".--.", "Q": "--.-", "R": ".-.",
	"S": "...", "T": "-", "U": "..-", "V": "...-", "W": ".--", "X": "-..-",
	"Y": "-.--", "Z": "--..",

	// Numbers
	"0": "-----", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
	"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",

	// Special symbols
	".": ".-.-.-", ",": "--..--", "?": "..--..", "'": ".----.", "!": "-.-.--",
	"/": "-..-.", "(": "-.--.", ")": "-.--.-", "&": ".-...", ":": "---...",
	";": "-.-.-.", "=": "-...-", "+": ".-.-.", "-": "-....-", "_": "..--.-",
	"\"": ".-..-.", "$": "...-..-", "@": ".--.-.", "¿": "..-.-", "¡": "--...-",
}

var reverseMorseCode = map[string]string{}

func main() {
	var input string
	var output string

	for key, value := range morseCode {
		reverseMorseCode[value] = key
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Write word or sentence: ")
	input, _ = reader.ReadString('\n')

	input = strings.ToUpper(input)

	// output = coder(input)
	output = decoder(input)

	fmt.Println(output)
}

func coder(str string) string {
	var result string

	words := strings.Fields(str)

	for _, word := range words {
		runes := []rune(word)

		for _, rune := range runes {
			element, _ := morseCode[string(rune)]
			result += element + " "
		}
	}

	return result
}

func decoder(str string) string {
	var result string

	words := strings.Split(str, " ")

	for _, word := range words {
		// runes := []rune(word)
		element, _ := reverseMorseCode[string(word)]
		result += element + " "

		// for _, rune := range runes {

		// }
	}

	return result
}
