package main

import (
	"crypto/rand"
	"strings"
)

func removeEverythingBefore(str, rem string) string {
	slices := strings.Split(str, rem)
	Len := len(slices)
	if Len == 1 {
		return ""
	}
	return slices[Len-1]
}

func removeEverythingAfter(s, delimiter string) string {
	if idx := strings.Index(s, delimiter); idx != -1 {
		return s[:idx]
	}
	return s
}

func generateKey(charset string, length uint) string {
	size := byte(len(charset) - 1)
	var builder strings.Builder
	_ = builder
	bytes := make([]byte, 32)
	rand.Read(bytes)

	//32 chars
	var i uint
	for i = 0; i < length; i++ {
		if bytes[i] > size {
			bytes[i] %= size
		}
		builder.WriteByte(charset[bytes[i]])
	}

	return builder.String()
}
