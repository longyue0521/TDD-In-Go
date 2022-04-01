package parser_test

import (
	"testing"

	"github.com/longyue0521/TDD-In-Go/args/parser"
	"github.com/stretchr/testify/assert"
)

func TestBoolOptionParser(t *testing.T) {
	testcases := map[string]struct {
		options   []string
		option    string
		expected  interface{}
		assertion assert.ErrorAssertionFunc
	}{
		"should not accept extra argument for bool option": {
			options:  []string{"-l", "t"},
			option:   "l",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrTooManyArguments)
			},
		},
		"should not accept more extra arguments for bool option": {
			options:  []string{"-l", "t", "f"},
			option:   "l",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrTooManyArguments)
			},
		},
		"should get default value if bool option not present": {
			options:  []string{},
			option:   "l",
			expected: (interface{})(false),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		"should set value to true if bool option present": {
			options:  []string{"-l"},
			option:   "l",
			expected: (interface{})(true),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, err := parser.BoolOptionParser().Parse(tt.options, tt.option)
			assert.Equal(t, tt.expected, actual)
			tt.assertion(t, err)
		})
	}
}

func TestIntOptionParser(t *testing.T) {
	testcases := map[string]struct {
		options   []string
		option    string
		expected  interface{}
		assertion assert.ErrorAssertionFunc
	}{
		"should not accept extra argument for single value option": {
			options:  []string{"-p", "8080", "8081"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrTooManyArguments)
			},
		},
		"should not missing argument for single value option": {
			options:  []string{"-p"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrMissingArgument)
			},
		},
		"should not missing argument for single value option but with another option": {
			options:  []string{"-p", "-l"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrMissingArgument)
			},
		},
		"should set default value if single value option present": {
			options:  []string{},
			option:   "p",
			expected: (interface{})(0),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		"should parse value if single value option present": {
			options:  []string{"-p", "9080"},
			option:   "p",
			expected: (interface{})(9080),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, err := parser.IntOptionParser().Parse(tt.options, tt.option)
			tt.assertion(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
