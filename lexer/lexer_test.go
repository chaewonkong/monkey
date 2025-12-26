package lexer_test

import (
	"monkey/lexer"
	"monkey/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	l := lexer.New(input)

	for _, tt := range []struct {
		expectedType    token.TokenType
		expectedLeteral string
	}{} {
		t.Run(tt.expectedLeteral, func(t *testing.T) {
			tok := l.NextToken()
			assert.Equal(t, tt.expectedType, tok.Type)
			assert.Equal(t, tt.expectedLeteral, tok.Literal)
		})
	}
}
