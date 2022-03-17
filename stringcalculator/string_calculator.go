package stringcalculator

import (
	"regexp"
	"strconv"
	"strings"
)

type StringCalculator struct {
}

func NewStringCalculator() *StringCalculator {
	return &StringCalculator{}
}

func (s *StringCalculator) Add(template string) int {
	sum := 0
	delimiter, numbers := s.parseTemplate(template)
	for _, number := range strings.Split(strings.ReplaceAll(numbers, `\n`, delimiter), delimiter) {
		num, _ := strconv.Atoi(number)
		sum += num
	}
	return sum
}

func (s *StringCalculator) parseTemplate(template string) (delimiter string, numbers string) {
	delimiter, numbers = ",", template
	reg, err := regexp.Compile(`^//[\D]+\\n`)
	if err != nil {
		return
	}
	templateBytes := []byte(template)
	delimiterHeaderIndexes := reg.FindIndex(templateBytes)
	if len(delimiterHeaderIndexes) == 0 {
		return
	}
	return s.parseTemplateBytes(delimiterHeaderIndexes, templateBytes)
}

func (s *StringCalculator) parseTemplateBytes(delimiterHeaderIndexes []int, templateBytes []byte) (string, string) {
	delimiterHeaderStartIndex, delimiterHeaderEndIndex := delimiterHeaderIndexes[0], delimiterHeaderIndexes[1]
	delimiterStartIndex, delimiterEndIndex := delimiterHeaderStartIndex+len(`//`), delimiterHeaderEndIndex-len(`\n`)
	return string(templateBytes[delimiterStartIndex:delimiterEndIndex]), string(templateBytes[delimiterHeaderEndIndex:])
}
