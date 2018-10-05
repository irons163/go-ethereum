package main

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	strRes := "Cdog"
	bytes := Encode("dog")
	str := string(bytes)
	if str != strRes {
		t.Error(fmt.Sprintf("Expected %q, got %q", strRes, str))
	}
	dec,_ := Decode(bytes, 0)
	fmt.Printf("raw: %v encoded: %q == %v\n", dec, str, "dog")


	sliceRes := "\x83CdogCgodCcat"
	strs := []string{"dog", "god", "cat"}
	bytes = Encode(strs)
	slice := string(bytes)
	if slice != sliceRes {
		t.Error(fmt.Sprintf("Expected %q, got %q", sliceRes, slice))
	}

	dec,_ = Decode(bytes, 0)
	fmt.Printf("raw: %v encoded: %q == %v\n", dec, slice, strs)
}