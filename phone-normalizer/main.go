package main

import (
	"regexp"
)

// REGEX version
func normalize(phone string) string {
	// If not caught with testing, bad stuff can happen with this MustCompile
	//re := regexp.MustCompile("[^0-9]")
	//Another way of Regexing digits double backslash is to escape backslash. Capitol D is any non-digits
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

//func normalize(phone string) string {
//	var buf bytes.Buffer
//
//	for _, ch := range phone {
//		if ch >= '0' && ch <= '9' {
//			buf.WriteRune(ch)
//		}
//	}
//	// We want these - 0123456789
//	return buf.String()
//}
