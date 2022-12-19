package randomSentenceGenerator

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func TestTokenization1(t *testing.T) {
	sen := "  This is an example. "
	actualList := Tokenization(1, sen)
	expectedlist := []string{"This", "is", "an", "example", ".", "<end>"}
	if !reflect.DeepEqual(actualList, expectedlist) {
		t.Errorf("not the same")
	}
}

func TestTokenization2(t *testing.T) {
	sen := "  This is an example. "
	actualList := Tokenization(2, sen)
	// fmt.Printf("%v", actualList)
	expectedList := []string{"<start>", "This", "is", "an", "example", ".", "<end>"}
	if !reflect.DeepEqual(actualList, expectedList) {
		t.Errorf("not the same")
	}
}

func TestTokenization3(t *testing.T) {
	sen := "  This is an example. "
	actualList := Tokenization(3, sen)
	// fmt.Printf("%v", actualList)
	expectedlist := []string{"<start>", "<start>", "This", "is", "an", "example", ".", "<end>"}
	if !reflect.DeepEqual(actualList, expectedlist) {
		t.Errorf("not the same")
	}
}

func TestUpdateNGram1(t *testing.T) {
	model := NGramModel{gram: 1}
	sen := "a b c"
	model.updateNGram(sen)
	actualContext := [][]string{{}, {}, {}, {}}
	expectedlist := model.contextGram
	if !reflect.DeepEqual(actualContext, expectedlist) {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedlist, actualContext)
	}

	actualTotal := model.totalGram
	expectedTotal := [][]string{{"a"}, {"b"}, {"c"}, {"<end>"}}
	if !reflect.DeepEqual(actualTotal, expectedTotal) {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedTotal, actualTotal)
	}
}

func TestUpdateNGram2(t *testing.T) {
	model := NGramModel{gram: 2}
	sen := "a b c"
	model.updateNGram(sen)
	actualContext := model.contextGram
	expectedlist := [][]string{{"<start>"}, {"a"}, {"b"}, {"c"}}
	if !reflect.DeepEqual(actualContext, expectedlist) {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedlist, actualContext)
	}

	actualTotal := model.totalGram
	expectedTotal := [][]string{{"<start>", "a"}, {"a", "b"}, {"b", "c"}, {"c", "<end>"}}
	if !reflect.DeepEqual(actualTotal, expectedTotal) {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedTotal, actualTotal)
	}

}

func TestUpdateNGram3(t *testing.T) {
	model := NGramModel{gram: 3}
	sen := "a b c"
	model.updateNGram(sen)
	actualContext := model.contextGram
	expectedlist := [][]string{{"<start>", "<start>"}, {"<start>", "a"}, {"a", "b"}, {"b", "c"}}
	if !reflect.DeepEqual(actualContext, expectedlist) {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedlist, actualContext)
	}

	actualTotal := model.totalGram
	expectedTotal := [][]string{{"<start>", "<start>", "a"}, {"<start>", "a", "b"}, {"a", "b", "c"}, {"b", "c", "<end>"}}
	if !reflect.DeepEqual(actualTotal, expectedTotal) {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedTotal, actualTotal)
	}

}

func TestProb1(t *testing.T) {
	model := NGramModel{gram: 1}
	sen1 := "a b c d"
	model.updateNGram(sen1)
	sen2 := "a b a b"
	model.updateNGram(sen2)

	actualProb := model.GetProb([]string{}, "a")
	expectedProb := 0.3

	if actualProb != expectedProb {
		t.Errorf("Expected num(%f) is not same as"+
			" actual num (%f)", actualProb, expectedProb)
	}

	actualProb1 := model.GetProb([]string{}, "c")
	expectedProb1 := 0.1

	if actualProb1 != expectedProb1 {
		t.Errorf("Expected num(%f) is not same as"+
			" actual num (%f)", actualProb1, expectedProb1)
	}

}

func TestProb2(t *testing.T) {
	model := NGramModel{gram: 2}
	sen1 := "a b c d"
	model.updateNGram(sen1)
	sen2 := "a b a b"
	model.updateNGram(sen2)

	if !reflect.DeepEqual(model.contextGram[2], []string{"b"}) {
		t.Errorf("not the same %v", model.contextGram[0])
	}

	actualProb1 := model.GetProb([]string{"b"}, "c")
	expectedProb1 := 1.0 / 3.0

	if actualProb1 != expectedProb1 {
		t.Errorf("Expected num(%f) is not same as"+
			" actual num (%f)", expectedProb1, actualProb1)
	}

	actualProb2 := model.GetProb([]string{"y"}, "c")
	expectedProb2 := 0.0

	if actualProb2 != expectedProb2 {
		t.Errorf("Expected num(%f) is not same as"+
			" actual num (%f)", actualProb2, expectedProb2)
	}

}

func TestRandomText(t *testing.T) {
	model := NGramModel{gram: 2}
	sen1 := "a b c d"
	model.updateNGram(sen1)
	sen2 := "a b a b"
	model.updateNGram(sen2)

	rand.Seed(2)

	result := model.GetRandomText(6)
	expectedList := []string{"a", "b", "<end>", "a", "b", "c"}
	expectedString := strings.Join(expectedList, " ")
	if result != expectedString {
		t.Errorf("Expected list (%v)is not same as"+
			" actual list (%v)", expectedString, result)
	}
}

func TestCreateModel(t *testing.T) {
	var model NGramModel
	model = NewNGram(5, "frankenstein.txt")
	result := model.GetRandomText(3)
	expectedString := "It was morning"
	if result != expectedString {
		t.Errorf("Expected string is not same as"+
			" actual string (%v)", result)
	}
}
