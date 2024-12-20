package parser_test

import (
	"fmt"
	"github.com/rostis232/parcelstrackingservice/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	testCases := []struct {
		testID       int
		trackNumbers []string
		isError      bool
	}{
		{
			testID: 1,
			trackNumbers: []string{
				"4PX3001521662170CN",
				"LZ134747205CN",
				"4PX3001505903961CN",
			},
			isError: false,
		},
	}

	for _, testCase := range testCases {
		body, err := parser.ParsePage(testCase.trackNumbers)
		if testCase.isError {
			assert.Error(t, err, testCase.testID)
		} else {
			assert.NoError(t, err, testCase.testID)
			assert.NotEmpty(t, body, testCase.testID)
			fmt.Println(body)
		}
	}
}
