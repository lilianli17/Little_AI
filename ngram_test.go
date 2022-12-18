package ngram

import (
	"testing"
	"reflect"
)



func TestTokenization(t *testing.T) {
	sen := "  This is an example. "
    actualList := Tokenize(1, sen)
    expectedlist := []string{"<s>", "This", "is", "an", "example", ".", "</s>"}
    if reflect.DeepEqual(actualList, expectedlist){
        t.Errorf("Expected String(%s) is not same as"+
         " actual string (%s)", expectedlist, actualList)
    }
}