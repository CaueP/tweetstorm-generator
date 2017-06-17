// Author: Cauê Garcia Polimanti
// Date: 17/06/2017
package main

import (
	"reflect"
	"testing"
)

type TestData struct {
	Input  string
	Output []string
}

// Test for Tweetstorm's convertTextIntoTweets function
func TestTweetstorm(t *testing.T) {

	// Testing data array
	// Case Test 1: 140 characters-long text
	// Case Test 2: 141 characters-long text
	// Case Test 3: 280 characters-long text
	testData := []TestData{
		TestData{"Este texto contem 140 caracteres. Este texto contem 140 caracteres. Este texto contem 140 caracteres. Este texto contem 140 caracteres. Este", []string{"Este texto contem 140 caracteres. Este texto contem 140 caracteres. Este texto contem 140 caracteres. Este texto contem 140 caracteres. Este"}},
		TestData{"Este texto contem 141 caracteres. Este texto contem 141 caracteres. Este texto contem 141 caracteres. Este texto contem 141 caracteres. To141", []string{"1/2 Este texto contem 141 caracteres. Este texto contem 141 caracteres. Este texto contem 141 caracteres. Este texto contem 141 caracteres. ", "2/2 To141"}},
		TestData{"Este é um texto exemplo de entrada que será convertido para Tweets separados com até 140 caracteres cada um. Esta parte é só para ultrapassar os 140 caracteres. Este é um texto exemplo de entrada que será convertido para Tweets separados com até 140 caracteres cada um. Este é um ", []string{"1/3 Este é um texto exemplo de entrada que será convertido para Tweets separados com até 140 caracteres cada um. Esta parte é só para u", "2/3 ltrapassar os 140 caracteres. Este é um texto exemplo de entrada que será convertido para Tweets separados com até 140 caracteres cad", "3/3 a um. Este é um "}},
	}

	for _, v := range testData {
		got := convertTextIntoTweets(v.Input)

		if !reflect.DeepEqual(got, v.Output) {
			t.Errorf("convertTextIntoTweets(%q) == %q, want %q", v.Input, got, v.Output)
		}
	}
}
