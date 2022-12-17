package ngram

import (
	"reflect"
	"strings"
	"Math/rand"
	"sort"
	"io/ioutil"
	"fmt"
)

type NGramModel struct {
	gram    int
	totalGram [][]string
	contextGram [][]string

}

// MarkovModel declares a simple language model that
// can be used to generate random text resembling a source document
type MarkovModel interface {
	Tokenize(int, string) []string

	updateNGram([]string)

	GetProb([]string, string) float64

	GetRandomToken([]string) string

	GetRandomText(int) string

	NewNGram(int, string)
}

// helper function: identifying punctuations
func IsPunct(s string) bool {
	switch s {
	case ".", ":", "!", "?", "'", ",", "-", "\"", ";", "(", ")", "[", "]", "{", "}", "\n":
		return true
	}
	return false
}

func Tokenize(gram int, sen string) []string {
	sen = strings.TrimSpace(sen)
	tokenList := []string{}

	for a := 0; a < gram-1; a++ {
		tokenList = append(tokenList, "<s>")
	}
	var word string = ""
	for _, w := range sen {
		if w != ' ' && w != '\n'  && !IsPunct(string(w)) {
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
	for b := 0; b < gram-1; b++ {
		tokenList = append(tokenList, "</s>")
	}
	return tokenList
}

func (m *NGramModel)updateNGram(sen string) {
	// get token list of sentence
	tokenList := Tokenize(m.gram, sen)
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

func (m *NGramModel)GetProb(context []string, token string) float64 {
	var total float64 = 0.0
	var counter float64 = 0.0
	
	for i := 0; i < len(m.totalGram); i++ {
		if reflect.DeepEqual(m.contextGram[i], context) {
			total += 1
			if m.totalGram[i][m.gram - 1] == token {
				counter += 1
			}
		}
	}
	if total == 0 {return 0.0}
	return counter / total

}

// helper function: check if slice s contains string e
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

// returns a random token according to the probability distribution 
// determined by the given context.
func (m *NGramModel)GetRandomToken(context []string) string{
	r := rand.Float64()
	var counter float64 = 0.0
	tokens := make(map[string]float64)
	t_list := []string{}

	for i := 0; i < len(m.totalGram); i++ {
		if reflect.DeepEqual(m.contextGram[i], context) {
			counter += 1
			if val, ok := tokens[m.totalGram[i][m.gram - 1]]; ok {
				counter += 1
				temp := val + 1
				tokens[m.totalGram[i][m.gram - 1]] = temp
			} else {
				tokens[m.totalGram[i][m.gram - 1]] = 1
			}

			if !contains(t_list, m.totalGram[i][m.gram - 1]) {
				t_list = append(t_list, m.totalGram[i][m.gram - 1])
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

func (m *NGramModel) GetRandomText(length int) string {
	result := ""
	// the (n−1)-list filled with "<s>"
	c_context := []string{}
	for j := 0; j < m.gram-1; j++ {
		c_context = append(c_context, "<s>")
	}

	for i := 0; i < length; i++ {
		cur := m.GetRandomToken(c_context)
		result += cur
		result += " "
		if cur == "</s>" {
			// reinitialize the context list when reaching the end of sentence
			c_context = []string{}
			for j := 0; j < m.gram-1; j++ {
				c_context = append(c_context, "<s>")
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

func NewNGram(gram_n int, path string) NGramModel {
	model := NGramModel{gram: gram_n}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	fileContent := string(file)
	// separate whole text file by sentence
	sen_list := strings.Split(fileContent, ".")
	for _, sen := range sen_list {
		model.updateNGram(sen)
	}

	return model
}
