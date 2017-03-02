package utils

import (
  "encoding/hex"
  "crypto/md5"
)

func ArrayContainsString(list []string, a string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func Md5Sum(input string) string {
    hasher := md5.New()
    hasher.Write([]byte(input))
    return hex.EncodeToString(hasher.Sum(nil))
}
