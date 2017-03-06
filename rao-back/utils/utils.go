package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func ArrayContainsString(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SplitString(input string, size int) []string {
	slices := []string{}
	count := 0
	lastIndex := 0
	for i, _ := range input {
		count++
		if count%(size+1) == 0 {
			slices = append(slices, input[lastIndex:i])
			lastIndex = i
		}
	}
	return slices
}

func Md5Sum(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
