package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func Uitoa(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func Sha256Hex(data []byte) string {
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}

func Sha256Bin(data []byte) []byte {
	hash := sha256.Sum256(data)

	return hash[:]
}

// Helper function for comparing slices
func CompareIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Returns the amount of nibbles that match each other from 0 ...
func MatchingNibbleLength(a, b []int) int { // Nibble是少量的意思
	i := 0
	for CompareIntSlice(a[:i+1], b[:i+1]) && i < len(b) { // 比對前i個內容是否相同
		i+=1
	}

	//fmt.Println(a, b, i-1)

	return i // 返回前i個前綴相同，如i==2等於前兩個值相同
}