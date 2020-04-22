package main

import (
	"fmt"
	"strings"
)

func to_pig_latin(s string) string {
	const suff = "ay"
	//var punctuation = [...]string {",", ".", ";"}

	fields := strings.Fields(s)

	var punct byte
	for i, word := range fields {
		if word[len(word)-1] == ',' {
			punct = word[len(word)-1]
			word = strings.Replace(word, ",", "", -1)
		} else {
			punct = 0
		}
		fields[i] =  word[1:] + string(word[0]) + suff + string(punct)
	}

	s = strings.Join(fields, " ")

	return s
}

func main() {
	text := `Pig Latin is a language game or argot in which words in English are altered, usually by adding a fabricated suffix or by moving the onset or initial consonant or consonant cluster of a word to the end of the word and adding a vocalic syllable to create such a suffix. The objective is to conceal the words from others not familiar with the rules. The reference to Latin is a deliberate misnomer; Pig Latin is simply a form of argot or jargon unrelated to Latin, and the name is used for its English connotations as a strange and foreign sounding language. It is most often used by young children as a fun way to confuse people unfamiliar with Pig Latin. `

	result := to_pig_latin(text)
	fmt.Println(result)
}