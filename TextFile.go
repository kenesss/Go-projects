package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	filePath := "sample.txt"

	wordCounts, err := countWords(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printTopWords(wordCounts, 5)
}

func countWords(filePath string) (map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			wordCounts[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordCounts, nil
}

func printTopWords(wordCounts map[string]int, topN int) {
	type wordPair struct {
		Word  string
		Count int
	}

	var pairs []wordPair
	for word, count := range wordCounts {
		pairs = append(pairs, wordPair{word, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	fmt.Printf("\nTop %d most frequent words:\n", topN)
	for i := 0; i < topN && i < len(pairs); i++ {
		fmt.Printf("%d. %s: %d\n", i+1, pairs[i].Word, pairs[i].Count)
	}
}
