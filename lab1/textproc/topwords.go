// Find the top K most common words in a text document.
// Input path: location of the document, K top words
// Output: Slice of top K words
// For this excercise, word is defined as characters separated by a whitespace

// Note: You should use `checkError` to handle potential errors.

package textproc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func topWords(path string, K int) []WordCount {
	// open the file and check for error
	file, err := os.Open(path)
	checkError(err)
	
	defer file.Close()

	wordMap := make(map[string]int)

	//Read the file and split lines to words and increment the count
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")
		for _, word := range words {
			wordMap[word]++
		}
	}

	//check if there is any arrors in the scanner
	err = scanner.Err()
	checkError(err)

	//convert to slice and sort the wordcount
	var wc []WordCount
	for k, v := range wordMap {
		wc = append(wc, WordCount{k, v})
	}
	sortWordCounts(wc)
	
	//sort the Top K
	return wc[:K]
}

//--------------- DO NOT MODIFY----------------!

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

// Method to convert struct to string format
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.

func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
