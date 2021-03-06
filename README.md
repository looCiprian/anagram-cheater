# Anagram Cheater
Anagram cheater allows you to find all anagram words from a set of letters a length and a dictionary.

Can be useful in some games :)

# How to use
## From Source
````
$ git clone https://github.com/looCiprian/anagram-cheater.git
$ cd anagram-cheater
$ go run cmd/anagram-cheater.go --help
$ go run cmd/anagram-cheater.go --letters abcdefg --length 3 --dictionary ./italian_dictionary.txt
````
Binaries available [here](https://github.com/looCiprian/anagram-cheater/releases/tag/v1.0)

Tested on:
- Windows
- macOS
- Linux

## Docker
````
$ git clone https://github.com/looCiprian/anagram-cheater.git
$ cd anagram-cheater
$ docker build -t anagram-cheater .
$ docker run -it --name anagram anagram-cheater
@ anagram-cheater --help
````
For dictionaries, you can mount volume on /dictionary inside docker
````
$ docker run -it --name anagram -v full/path/to/your/dictionary/:/dictionary anagram-cheater
@ anagram-cheater --letters abcdefg --length 3 --dictionary /dictionary/your_dictionary.txt
````

# Dictionaries
[Italian words](https://github.com/napolux/paroleitaliane/blob/master/paroleitaliane/60000_parole_italiane.txt)

# Complexity
- d = words in dictionary
- len(dWord) = length of a word in the dictionary
- len(anagram) = numbers of letters in  "--letters" option

Sort complexity = O(d * len(dWord)log(len(dWord)))

Search complexity = O(d * (len(anagram) + len(dWord)))

Final complexity = O(d * len(dWord)log(len(dWord))) + O(d * (len(anagram) + len(dWord)))