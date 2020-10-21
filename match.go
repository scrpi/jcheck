package jcheck

import (
	"strings"
	"unicode/utf8"
)

// scanChunk gets the next segment in pattern, which is a non-starhash string
// possibly preceded by a '*' or '#'.
func scanChunk(pattern string) (starhash rune, chunk, rest string) {
	for len(pattern) > 0 && (pattern[0] == '*' || pattern[0] == '#') {
		starhash = rune(pattern[0])
		pattern = pattern[1:]
	}

	var i int
	for i = 0; i < len(pattern); i++ {
		if pattern[i] == '*' || pattern[i] == '#' {
			break
		}
	}

	return starhash, pattern[0:i], pattern[i:]
}

// matchChunk tests whether chunk matches the beginning of the string s. If so
// it returns the remainder of s (excluding the chunk prefix).
func matchChunk(chunk, s string) (rest string, ok bool) {
	for len(chunk) > 0 {
		if len(s) == 0 {
			return "", false
		}
		switch chunk[0] {
		case '?':
			if s[0] == '.' {
				// '?' does not match a path separator
				return "", false
			}
			_, n := utf8.DecodeRuneInString(s)
			s = s[n:]
			chunk = chunk[1:]
		default:
			if chunk[0] != s[0] {
				return "", false
			}
			s = s[1:]
			chunk = chunk[1:]
		}
	}
	return s, true
}

// match determines whether the path matches the pattern provided.
//
// The pattern syntax is as follows:
//
//  pattern:
//      { term }
//
//  term:
//      '#'     matches any sequence of characters
//      '*'     matches any sequence of non-separator ('.') characters
//      '?'     matches any single non-separator ('.') character
//      c       matches character c
//
// match requires pattern to match the entire of path, not just a substring.
//
// Example:
//
//      match("array.#.field", "array.10.field")        // true
//      match("array.#.field", "array.10.object.field") // false
//      match("array.*.field", "array.10.object.field") // true
//
func match(pattern, path string) bool {
Pattern:
	for len(pattern) > 0 {
		var starhash rune
		var chunk string
		starhash, chunk, pattern = scanChunk(pattern)

		if (starhash == '*' || starhash == '#') && chunk == "" {
			switch starhash {
			case '#':
				// Trailing '#'. Matches the rest of the path.
				return true
			case '*':
				// Trailing '*'. Matches the rest of the path not including dot
				// separator.
				return !strings.Contains(path, ".")
			}
		}

		// Look for match at current position
		t, ok := matchChunk(chunk, path)

		if ok && (len(t) == 0 || len(pattern) > 0) {
			path = t
			continue
		}

		if starhash == '*' || starhash == '#' {
			// Look for match skipping i+1 bytes
			for i := 0; i < len(path); i++ {
				if starhash == '*' && path[i] == '.' {
					// '#' does not match dot separator in path.
					return false
				}
				t, ok := matchChunk(chunk, path[i+1:])
				if ok {
					path = t
					continue Pattern
				}
			}
		}
		return false
	}
	return len(path) == 0
}
