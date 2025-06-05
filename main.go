package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var morseCodeRus = map[string]string{
	// Russian
	"А": ".-", "Б": "-...", "В": ".--", "Г": "--.", "Д": "-..", "Е": ".",
	"Ж": "...-", "З": "--..", "И": "..", "Й": ".---", "К": "-.-", "Л": ".-..",
	"М": "--", "Н": "-.", "О": "---", "П": ".--.", "Р": ".-.", "С": "...",
	"Т": "-", "У": "..-", "Ф": "..-.", "Х": "....", "Ц": "-.-.", "Ч": "---.",
	"Ш": "----", "Щ": "--.-", "Ъ": "--.--", "Ы": "-.--", "Ь": "-..-",
	"Э": "..-..", "Ю": "..--", "Я": ".-.-",

	// Numbers
	"0": "-----", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
	"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",

	// Special symbols
	".": ".-.-.-", ",": "--..--", "?": "..--..", "'": ".----.", "!": "-.-.--",
	"/": "-..-.", "(": "-.--.", ")": "-.--.-", "&": ".-...", ":": "---...",
	";": "-.-.-.", "=": "-...-", "+": ".-.-.", "-": "-....-", "_": "..--.-",
	"\"": ".-..-.", "$": "...-..-", "@": ".--.-.", "¿": "..-.-", "¡": "--...-",
}

var morseCodeEng = map[string]string{
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
var language string

func main() {
	var input string
	var output string
	var coder_option string

	fmt.Println("1) English")
	fmt.Println("2) Russian")
	fmt.Print("Choose language: ")
	fmt.Scanln(&language)

	if language == "1" {
		for key, value := range morseCodeEng {
			reverseMorseCode[value] = key
		}
	} else if language == "2" {
		for key, value := range morseCodeRus {
			reverseMorseCode[value] = key
		}
	} else {
		fmt.Println("WRONG INPUT")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("1) Coder")
	fmt.Print("2) Decoder")
	fmt.Print("Choose option: ")

	coder_option, _ = reader.ReadString('\n')

	fmt.Printf("Write word or sentence: ")
	input, _ = reader.ReadString('\n')

	input = strings.ToUpper(input)

	if coder_option == "1" {
		output = coder(input)
	} else if coder_option == "2" {
		output = decoder(input)
	} else {
		fmt.Println("WRONG OPTION")
	}

	fmt.Println(output)
}

func coder(str string) string {
	var result string

	words := strings.Fields(str)

	for _, word := range words {
		runes := []rune(word)

		if language == "1" {
			for _, rune := range runes {
				element, _ := morseCodeEng[string(rune)]
				result += element + " "
			}
		} else {
			for _, rune := range runes {
				element, _ := morseCodeRus[string(rune)]
				result += element + " "
			}
		}
	}

	return result
}

func decoder(str string) string {
	var result string

	cleaned := strings.TrimSpace(str)
	words := strings.Fields(cleaned)
	fmt.Println(words)

	for _, word := range words {
		fmt.Println(word)
		element, _ := reverseMorseCode[word]
		fmt.Println(element)
		result += element
	}

	return result
}
