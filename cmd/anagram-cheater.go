package main

import (
	"anagram-cheater/internal/dictionary_mng"
	"anagram-cheater/internal/file_mng"
	"flag"
	"fmt"
	"os"
)

func main() {

	var anagram = flag.String("letters", "", "Letter to anagram")
	var dictionary = flag.String("dictionary", "", "Dictionary to use")
	var length = flag.Int("length", 0, "Length of the word to find")
	flag.Parse()

	if *anagram == "" || *dictionary == "" || *length == 0{
		fmt.Println("Usage: " + os.Args[0] + " --help")
		fmt.Println("Example: " + os.Args[0] + " --letters abcdefg --length 3 --dictionary ./italian_dictionary.txt")
		os.Exit(1)
	}


	lines := file_mng.ReadFileByLines(*dictionary)
	wordlist := dictionary_mng.Sort(lines)

	found := dictionary_mng.FindWords(wordlist, *anagram, *length)

	for _, word := range found{
		fmt.Println(word)
	}
}
