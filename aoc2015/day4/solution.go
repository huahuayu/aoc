package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func MD5SaltFinder(str string, leadingZero int) int {
	stringToMatch := buildZeroString(leadingZero)
	for i := 0; ; i++ {
		hash := MD5(str + fmt.Sprint(i))
		if stringToMatch == hash[:leadingZero] {
			return i
		}
	}
}

func MD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func buildZeroString(length int) string {
	return strings.Repeat("0", length)
}
