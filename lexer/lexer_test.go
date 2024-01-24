package lexer

import (
	"testing"

	"github.com/justinpaulosolo/interpreter-go/token"
)

func TestNextToekn(t *testing.T) {
	t.Run("basic tokens", func(t *testing.T) {
		input := `=+(){},;`

		tests := []struct {
			expectedType   token.TokenType
			expecedLiteral string
		}{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i,
					tt.expectedType,
					tok.Type)
			}

			if tok.Literal != tt.expecedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i,
					tt.expecedLiteral,
					tok.Literal)
			}
		}

	})

	t.Run("advance tokens", func(t *testing.T) {
		input := `let five = 5;
		let ten = 10;
		let add = fn(x, y) {
		  x + y;
		}
		let result = add(five, ten);
		`

		tests := []struct {
			expectedType   token.TokenType
			expecedLiteral string
		}{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},

			// let add = fn(x, y) {
			//   x + y;
			// }
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},

			//let result = add(five, ten);
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},

			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i,
					tt.expectedType,
					tok.Type)
			}

			if tok.Literal != tt.expecedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i,
					tt.expecedLiteral,
					tok.Literal)
			}
		}
	})
}
