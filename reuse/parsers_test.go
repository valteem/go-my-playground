package reuse_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

// Input from LLM, immediately goes to infinite loop ))
func ParseStringOpenCloseTokensFromLLM(inputString, tokenOpen, tokenClose string) []string {
	var result []string
	current := ""
	startIndex := strings.Index(inputString, tokenOpen) + len(tokenOpen)
	for startIndex != -1 {
		endIndex := strings.Index(inputString[startIndex:], tokenClose)
		if endIndex == -1 {
			break
		} else {
			endIndex += startIndex
		}
		current = inputString[startIndex : endIndex+len(tokenClose)]
		result = append(result, current)
		startIndex = strings.Index(inputString[endIndex+len(tokenClose):], tokenOpen) + endIndex + len(tokenClose)
	}
	return result
}

func ParseStringOpenCloseTokens(input, tokenOpen, tokenClose string) []string {

	var output []string

	if len(tokenOpen)+len(tokenClose) >= len(input) {
		return output
	}

	countOpen, countClose := countTokens(input, tokenOpen), countTokens(input, tokenClose)
	count := min(countOpen, countClose)
	if count == 0 {
		return output
	}

	// TODO:
	// - make sure that at least one tokenOpen occurs before tokenClose
	// - count all pairs where tokenOpen precedes corresponding tokenClose
	// - exclude tokenOpen-tokenClose pairs that include at least one nested pair
	// - ... and probably some more ...

	return output

}

func countTokens(input, token string) int {
	if len(token) > len(input) || len(input) == 0 {
		return 0
	}
	var indexAbs, indexRun, count int
	for {
		indexRun = strings.Index(input[indexAbs:], token)
		if indexRun != -1 {
			count++
			indexAbs += indexRun + len(token)
		} else {
			break
		}
	}
	return count
}

func TestTokenCount(t *testing.T) {

	tests := []struct {
		input string
		token string
		count int
	}{
		{"input with token and token", "token", 2},
		{"input without what we need", "token", 0},
	}

	for _, tc := range tests {
		if count := countTokens(tc.input, tc.token); count != tc.count {
			t.Errorf("toekn count for %q, %q: get %d, expect %d", tc.input, tc.token, count, tc.count)
		}
	}

}

func TestParseStringOpenClose(t *testing.T) {
	tests := []struct {
		input        string
		openToken    string
		closeToken   string
		expectedList []string
	}{
		{
			input:        "<first substring><second_substring>",
			openToken:    "<",
			closeToken:   ">",
			expectedList: []string{"first_substring", "second_substring"},
		},
		{
			input:        "<first substring>something in between<second_substring>",
			openToken:    "<",
			closeToken:   ">",
			expectedList: []string{"first_substring", "second_substring"},
		},
		{
			input:        "<first sub>string><second_substring>",
			openToken:    "<",
			closeToken:   ">",
			expectedList: []string{"first_sub", "second_substring"},
		},
		{
			input:        "<first <substring><second_substring>",
			openToken:    "<",
			closeToken:   ">",
			expectedList: []string{"substring", "second_substring"},
		},
		{
			input:        "<first substring><<<second_substring>",
			openToken:    "<",
			closeToken:   ">",
			expectedList: []string{"first_substring", "second_substring"},
		},
	}

	for _, test := range tests {
		actualList := ParseStringOpenCloseTokens(test.input, test.openToken, test.closeToken)
		if fmt.Sprint(actualList) != fmt.Sprint(test.expectedList) {
			fmt.Printf("Test failed for input '%s' with expected list %v but got %v\n", test.input, test.expectedList, actualList)
		}
	}
}

func TestParseTokens(t *testing.T) {

	tests := []struct {
		input      string
		tokenOpen  string
		tokenClose string
		count      int
	}{
		{"some <substring> and more", "<", ">", 1},
		{"some <<nested substring> and more>", "<", ">", 1},
		{"some <<nested substring> and more> and more><", "<", ">", 1},
		{"some <<nested substring> and more stuff> and <another substring>>", "<", ">", 2},
	}

	for _, tc := range tests {
		pattern := regexp.MustCompile(tc.tokenOpen + "(.*?)" + tc.tokenClose)
		occ := pattern.FindAll([]byte(tc.input), -1)
		if count := len(occ); count != tc.count {
			t.Errorf("input string %q\nwith tokens %q and %q\n: get %d, expect %d", tc.input, tc.tokenOpen, tc.tokenClose, count, tc.count)
		}
	}
}

type customHeaderKV struct {
	key   string
	value string
}

type customHeader struct {
	headers []customHeaderKV
}

func (h *customHeader) Add(key, value string) {
	h.headers = append(h.headers, customHeaderKV{key: key, value: value})
}

type customHeaderValue *customHeader

func NewCHV() customHeaderValue {
	return &customHeader{}
}

func TestStructTypeConversion(t *testing.T) {

	ch := NewCHV()

	key, value := "Header-Key", "header-value"

	// ch.Add(key, value) // ch.Add() undefined
	(*customHeader)(ch).Add(key, value) // type conversion, same as int(int32)

	if headerKey, headerValue := ch.headers[0].key, ch.headers[0].value; headerKey != key || headerValue != value {
		t.Errorf("header key/value: get %s/%s, expect %s/%s", headerKey, headerValue, key, value)
	}

}
