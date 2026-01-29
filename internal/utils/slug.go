package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// GenerateSlug generates a URL-friendly slug from a string
func GenerateSlug(text string) string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Remove accents/diacritics
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	text, _, _ = transform.String(t, text)

	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	text = reg.ReplaceAllString(text, "-")

	// Remove leading and trailing hyphens
	text = strings.Trim(text, "-")

	// Replace multiple hyphens with single hyphen
	reg = regexp.MustCompile(`-+`)
	text = reg.ReplaceAllString(text, "-")

	return text
}

// GenerateUniqueSlug generates a unique slug by appending a number if necessary
func GenerateUniqueSlug(text string, existingSlugs []string) string {
	baseSlug := GenerateSlug(text)
	slug := baseSlug
	counter := 1

	// Check if slug exists in the list
	for contains(existingSlugs, slug) {
		slug = baseSlug + "-" + string(rune(counter+'0'))
		counter++
	}

	return slug
}

// contains checks if a string exists in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
