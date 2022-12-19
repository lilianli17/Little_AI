package rsg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"sort"
	"strings"
)

type NGramModel struct {
	gram        int
	totalGram   [][]string
	contextGram [][]string
}

// MarkovModel interface declares a simple n-gram language model that
// can be used to generate random text resembling a source document
type MarkovModel interface {
	Tokenization(int, string) []string

	updateNGram([]string)

	GetProb([]string, string) float64

	GetRandomToken([]string) string

	GetRandomText(int) string

	NewNGram(int, string)
}

// helper function: identifying punctuations
// (only the most commonly used in English)
func IsPunct(s string) bool {
	switch s {
	case ".", ":", "!", "?", "'", ",", "-", "\"", ";",
		"(", ")", "[", "]", "{", "}":
		return true
	}
	return false
}

// Generates a list of tokens from the input string of text
func Tokenization(gram int, sen string) []string {
	sen = strings.TrimSpace(sen)
	tokenList := []string{}

	for a := 1; a < gram; a++ {
		tokenList = append(tokenList, "<start>")
	}

	var word string = ""
	for _, w := range sen {
		if w != ' ' && w != '\n' && !IsPunct(string(w)) {
			word = word + string(w)
		} else {
			if word != "" {
				tokenList = append(tokenList, word)
				word = ""
			}
			if IsPunct(string(w)) {
				tokenList = append(tokenList, string(w))
			}
		}

	}
	if word != "" {
		tokenList = append(tokenList, word)
	}
	tokenList = append(tokenList, "<end>")

	return tokenList
}

// Computes the n-grams for the input sentence and
// updates the internal totalGram and contextGram
func (m *NGramModel) updateNGram(sen string) {
	// get token list of sentence
	tokenList := Tokenization(m.gram, sen)
	// a list of all n-grams: context + current token
	var totalGram [][]string
	// n−1 words preceding the current token: only context
	var contextGram [][]string

	// update totalGram and contetGram in the model
	for i := 0; i < len(tokenList)-m.gram+1; i++ {
		totalGram = append(totalGram, tokenList[i:i+m.gram])
		contextGram = append(contextGram, tokenList[i:i+m.gram-1])
	}

	m.totalGram = append(m.totalGram, totalGram...)
	m.contextGram = append(m.contextGram, contextGram...)
}

// Computes the probability of that token occuring, given the preceding context
func (m *NGramModel) GetProb(context []string, token string) float64 {
	var total float64 = 0.0
	var counter float64 = 0.0

	for i := 0; i < len(m.contextGram); i++ {
		if reflect.DeepEqual(m.contextGram[i], context) {
			total = total + 1
			if strings.Compare(m.totalGram[i][m.gram-1], token) == 0 {
				counter = counter + 1
			}
		}
	}
	if total == 0 {
		return 0.0
	}

	return counter / total

}

// Return a random token according to the probability distribution
// determined by the given context
func (m *NGramModel) GetRandomToken(context []string) string {
	r := rand.Float64()
	counter := 0.0
	tokens := make(map[string]float64)
	t_list := []string{}

	for i := 0; i < len(m.totalGram); i++ {
		if reflect.DeepEqual(m.contextGram[i], context) {
			counter += 1
			if val, ok := tokens[m.totalGram[i][m.gram-1]]; ok {
				tokens[m.totalGram[i][m.gram-1]] = val + 1
			} else {
				tokens[m.totalGram[i][m.gram-1]] = 1
				t_list = append(t_list, m.totalGram[i][m.gram-1])
			}
		}
	}

	sort.Strings(t_list)
	prob := 1 / counter
	sum_prob := 0.0
	for _, word := range t_list {
		sum_prob += tokens[word] * prob
		if sum_prob >= r {
			return word
		}

	}
	return "<>"
}

// Returns a string of space-separated tokens chosen at random using GetRandomToken
func (m *NGramModel) GetRandomText(length int) string {
	result := ""
	// the (n−1)-list filled with "<start>"
	c_context := []string{}
	for j := 0; j < m.gram-1; j++ {
		c_context = append(c_context, "<start>")
	}

	for i := 0; i < length; i++ {
		cur := m.GetRandomToken(c_context)
		result += cur
		result += " "
		if cur == "<end>" {
			// reinitialize the context list when reaching the end of sentence
			c_context = []string{}
			for j := 0; j < m.gram-1; j++ {
				c_context = append(c_context, "<start>")
			}
		} else {
			if m.gram != 1 {
				c_context = append(c_context, cur)
				c_context = c_context[1:]
			}
		}
	}
	// exempting space at the end
	return result[:len(result)-1]
}

// Loads the text at the given path and creates an n-gram model from the resulting data
func NewNGram(gram_n int, path string) NGramModel {
	model := NGramModel{gram: gram_n}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)
	// separate whole text file by sentence
	sen_list := strings.Split(fileContent, ".")
	// update the current model by sentence
	for _, sen := range sen_list {
		model.updateNGram(sen)
	}

	return model
}
