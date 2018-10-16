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