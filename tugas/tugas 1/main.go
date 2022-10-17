package main

import (
	"fmt"
)

var angka = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var vocal = []string{"a", "i", "u", "e", "o"}
var text string

func main() {

	text = "b"
	if isVocal(text) {
		fmt.Println("vocal")
	} else {
		if isAngka(text) {
			fmt.Println("error bukan huruf")
		} else {
			fmt.Println("konsonan")
		}
	}

}

func isAngka(param string) bool {
	apaAngka := false
	for i := 0; i < len(angka); i++ {
		if param == angka[i] {
			apaAngka = true

		}
	}
	return (apaAngka)
}

func isVocal(param string) bool {
	apaVocal := false
	for i := 0; i < len(vocal); i++ {
		if param == vocal[i] {
			apaVocal = true

		}
	}
	return (apaVocal)
}

// func ispolindrone(input string) bool {
// 	input = strings.ToLower(input)
// 	apaPolindrone := true
// 	for i := 0; i < len(input)/2; i++ {
// 		awal := string(input[i])
// 		akhir := string(input[(len(input) - 1 - i)])

// 		if awal != akhir {
// 			return apaPolindrone == false
// 		}

// 	}
// 	return true
// }
