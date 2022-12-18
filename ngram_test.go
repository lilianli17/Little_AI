package ngram

import (
	"testing"
	"reflect"
	"fmt"
)

func TestTokenization1(t *testing.T) {
	sen := "  This is an example. "
    actualList := Tokenize(1, sen)
    expectedlist := []string{"<s>", "This", "is", "an", "example", ".", "</s>"}
    if reflect.DeepEqual(actualList, expectedlist){
        t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expectedlist, actualList)
    }
}

func TestTokenization2(t *testing.T) {
	sen := "  This is an example. "
    actualList := Tokenize(2, sen)
	fmt.Printf("%v", actualList)
    expectedlist := []string{"<s>", "<s>", "This", "is", "an", "example", ".", "</s>", "</s>"}
    if reflect.DeepEqual(actualList, expectedlist){
        t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expectedlist, actualList)
    }
}