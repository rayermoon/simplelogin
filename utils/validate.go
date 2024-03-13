package utils

import (
	"html"
	"regexp"
)

func IsValidEmail(email string) bool {
	// Regular expression pattern for validating email addresses
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern
	regex := regexp.MustCompile(pattern)

	// Match the email against the pattern
	return regex.MatchString(email)
}

func StripXSS(input string) string {
	// Remove any <script> tags and their contents
	re := regexp.MustCompile(`<script[^>]*>.*?</script>`)
	input = re.ReplaceAllString(input, "")

	// Escape HTML entities to prevent XSS attacks
	input = html.EscapeString(input)

	return input
}
