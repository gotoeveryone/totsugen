package main

import (
	"math"
	"strings"
	"testing"
)

func TestHasMultibyte(t *testing.T) {
	patterns := map[string]bool{
		"a": false,
		"あ": true,
		"𡌛": true,
	}
	for value, result := range patterns {
		r := hasMultibyte([]rune(value)[0])
		if r != result {
			t.Errorf("Test failed: return value is not match, actual: %t, expected: %t", r, result)
		}
	}
}

func TestGetStringCount(t *testing.T) {
	patterns := map[string]int{
		"hoge123":    7,
		"突然の死":       8,
		"突然のhoge1":   11,
		"has-space ": 10,
		"サロゲートペア(𡌛)": 18,
	}
	for value, count := range patterns {
		c := getStringCount(value)
		if c != count {
			t.Errorf("Test failed: return value is not match, actual: %d, expected: %d", c, count)
		}
	}
}

func TestGetDisplaytext(t *testing.T) {
	patterns := map[string]int{
		"hoge123":    7,
		"突然の死":       8,
		"突然のhoge1":   11,
		"has-space ": 10,
		"サロゲートペア(𡌛)": 18,
	}
	for value, count := range patterns {
		text := getDisplayText(value, count)
		if !strings.Contains(text, value) {
			t.Errorf("Test failed: return value is not contains %s, actual: %s", value, text)
		}
		padLen := math.Floor(float64(count) / 2)
		repeats := []string{"人", "Y^"}
		for _, repeat := range repeats {
			r := strings.Repeat(repeat, int(padLen)+2)
			if !strings.Contains(text, r) {
				t.Errorf("Test failed: too little repetition in the return value %s, actual: %s", r, text)
			}
		}
	}
}
