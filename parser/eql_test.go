package parser

import (
	"strings"
	"testing"

	"github.com/sqls-server/sqls/dialect"
	"github.com/sqls-server/sqls/token"
)

func TestEQLLexer_Keywords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Kind
	}{
		{
			name:  "GET keyword",
			input: "GET",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "FROM keyword",
			input: "FROM",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "WHERE keyword",
			input: "WHERE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "AND keyword",
			input: "AND",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "OR keyword",
			input: "OR",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "ANDNOT keyword",
			input: "ANDNOT",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "GROUP BY keywords",
			input: "GROUP BY",
			expected: []token.Kind{
				token.SQLKeyword,
				token.Whitespace,
				token.SQLKeyword,
			},
		},
		{
			name:  "ORDER BY keywords",
			input: "ORDER BY",
			expected: []token.Kind{
				token.SQLKeyword,
				token.Whitespace,
				token.SQLKeyword,
			},
		},
		{
			name:  "PAGE keyword",
			input: "PAGE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "LOCK keyword",
			input: "LOCK",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "FORMAT keyword",
			input: "FORMAT",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "SUBSCRIBE keyword",
			input: "SUBSCRIBE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "UNSUBSCRIBE keyword",
			input: "UNSUBSCRIBE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "SET keyword",
			input: "SET",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "IN keyword",
			input: "IN",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "INSERT keyword",
			input: "INSERT",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "DELETE keyword",
			input: "DELETE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "TRANSACTION_START keyword",
			input: "TRANSACTION_START",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "TRANSACTION_COMMIT keyword",
			input: "TRANSACTION_COMMIT",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "TRANSACTION_ROLLBACK keyword",
			input: "TRANSACTION_ROLLBACK",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "AS keyword",
			input: "AS",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "ASC keyword",
			input: "ASC",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "DESC keyword",
			input: "DESC",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "READ keyword",
			input: "READ",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "WRITE keyword",
			input: "WRITE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
		{
			name:  "LIKE keyword",
			input: "LIKE",
			expected: []token.Kind{
				token.SQLKeyword,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := token.NewTokenizer(strings.NewReader(tt.input), &dialect.EQLDialect{})
			tokens, err := tokenizer.Tokenize()
			if err != nil {
				t.Fatalf("Tokenize() error = %v", err)
			}

			if len(tokens) != len(tt.expected) {
				t.Fatalf("Tokenize() returned wrong number of tokens: got %d, expected %d", len(tokens), len(tt.expected))
			}

			for i, tok := range tokens {
				if tok.Kind != tt.expected[i] {
					t.Errorf("Token[%d] kind wrong: got %v, expected %v, value: '%v'", i, tok.Kind, tt.expected[i], tok.Value)
				}
			}
		})
	}
}
