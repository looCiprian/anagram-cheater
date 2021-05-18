package dictionary_mng

import (
	"sort"
	"strings"
)

type dictionary struct {
	word string
	sortedWord string
}

func Sort(dictionaryLines []string) []dictionary{

	var wordlist []dictionary

	for _, line := range dictionaryLines{
		wordlist = append(wordlist, dictionary{word: line, sortedWord: sortString(line)})
	}
	return wordlist
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func FindWords(wordlist []dictionary, anagram string, length int) []string {

	anagramSorted := sortString(strings.ToLower(anagram))

	var found []string

	for k:=0; k < len(wordlist); k++ {

		if len(wordlist[k].word) != length {
			continue
		}
		i := 0
		j := 0
		tempWord := ""
		for  {
			if tempWord == wordlist[k].sortedWord{
				found = append(found, wordlist[k].word)
				break
			}
			if len(anagramSorted) == i {
				break
			}
			if len(wordlist[k].sortedWord) == j {
				break
			}
			if anagramSorted[i] == wordlist[k].sortedWord[j] {
				tempWord = tempWord + string(anagramSorted[i])
				i++
				j++
				continue
			}
			if anagramSorted[i] > wordlist[k].sortedWord[j] {
				j++
				continue
			}
			if anagramSorted[i] < wordlist[k].sortedWord[j] {
				i++
				continue
			}
		}
	}
	return found
}