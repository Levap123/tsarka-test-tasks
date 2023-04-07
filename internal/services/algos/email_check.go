package algos

import (
	"fmt"
	"regexp"
	"strings"
)

func (s *AlgosService) EmailCheck(str string) []string {
	allowedSymbols := "a-zA-Z.-_"
	emailPattern := fmt.Sprintf(`Email:\s*([%s]+@[%s]+\.[%s])+`, allowedSymbols, allowedSymbols, allowedSymbols)
	emailPatternRegex := regexp.MustCompile(emailPattern)

	matches := emailPatternRegex.FindAllString(str, -1)
	emails := make([]string, 0, len(matches))
	for _, match := range matches {
		buffer := strings.TrimPrefix(match, "Email:")
		buffer = strings.TrimSpace(buffer)
		emails = append(emails, buffer)
	}

	return emails
}
