package parser_test

import (
	"testing"

	"github.com/longyue0521/TDD-In-Go/args/parser"
	"github.com/stretchr/testify/assert"
)

func TestTestUnaryOptionParser_BoolOption(t *testing.T) {
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
			// 注意:并发问题
			tt := tt
			// 利用多核,并行运行
			t.Parallel()

			actual, err := parser.BoolOptionParser().Parse(tt.options, tt.option)
			assert.Equal(t, tt.expected, actual)
			tt.assertion(t, err)
		})
	}
}

func TestUnaryOptionParser_IntOption(t *testing.T) {
	testcases := map[string]struct {
		options   []string
		option    string
		expected  interface{}
		assertion assert.ErrorAssertionFunc
	}{
		"should not accept extra argument for int option": {
			options:  []string{"-p", "8080", "8081"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrTooManyArguments)
			},
		},
		"should not missing argument for int option": {
			options:  []string{"-p"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrMissingArgument)
			},
		},
		"should not missing argument for int option but with another option": {
			options:  []string{"-p", "-l"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrMissingArgument)
			},
		},
		"should set default value if int option not present": {
			options:  []string{},
			option:   "p",
			expected: (interface{})(0),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		"should parse value if int option present": {
			options:  []string{"-p", "9080"},
			option:   "p",
			expected: (interface{})(9080),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		"should not parse illegal value if int option present": {
			options:  []string{"-p", "9x8y"},
			option:   "p",
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrIllegalValue)
			},
		},
	}

	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			// 注意:并发问题
			tt := tt
			// 利用多核,并行运行
			t.Parallel()

			actual, err := parser.IntOptionParser().Parse(tt.options, tt.option)
			tt.assertion(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListOptionParser_StringListOption(t *testing.T) {
	testcases := map[string]struct {
		options     []string
		option      string
		parseValues func(s ...string) ([]string, error)
		expected    interface{}
		assertion   assert.ErrorAssertionFunc
	}{

		"should set default list value if string list option not present": {
			options: []string{"this", "is", "list"},
			option:  "g",
			parseValues: func(s ...string) ([]string, error) {
				return s, nil
			},
			expected: (interface{})([]string{}),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		"should parse list value if string list option present": {
			options: []string{"-g", "this", "is", "list"},
			option:  "g",
			parseValues: func(s ...string) ([]string, error) {
				return s, nil
			},
			expected: (interface{})([]string{"this", "is", "list"}),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		"should have at least one argument for string list option present": {
			options: []string{"-g"},
			option:  "g",
			parseValues: func(s ...string) ([]string, error) {
				return s, nil
			},
			expected: (interface{})(nil),
			assertion: func(tt assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, parser.ErrAtLeastOneArgument)
			},
		},
		"should parse special list value if string list option present": {
			options: []string{"-g", "number", "-1", "-l2", "--list"},
			option:  "g",
			parseValues: func(s ...string) ([]string, error) {
				return s, nil
			},
			expected:  (interface{})([]string{"number", "-1", "-l2"}),
			assertion: assert.NoError,
		},
	}
	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			// 注意:并发问题
			tt := tt
			// 利用多核,并行运行
			t.Parallel()

			actual, err := parser.StringListParser(tt.parseValues).Parse(tt.options, tt.option)
			tt.assertion(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
