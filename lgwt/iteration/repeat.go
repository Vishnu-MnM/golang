package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// Range over integer loops (go version 1.22 and above, using strings.)
// func Repeat(character string) string {
// 		return strings.Repeat(character, 5)
// 	}

// From go 1.22, the version mentioned below works, but range over int had some changes.
// for i := 0; i < 5; i++ {
// 	repeated = repeated + character
// }
// There is even better version where we don't have to copy the entire string for every iteration.
// var repeated string
// for range 5 {
// 	repeated += character
// }

