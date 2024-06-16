package util

import "testing"

func TestPalindrome(t *testing.T) {
	result := IsPalindrome("あいうえおかおえういあ")
	if result != true {
		t.Error("Expected true, got ", result)
	}
}

func TestCountSuffixRhyme(t *testing.T) {
	result := CountSuffixRhyme("やまたのおろち", "がんせきおとし")
	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}

func TestGetMode(t *testing.T) {
	result := GetMode("あいうえおおおえういあ")
	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}

func TestGetScore(t *testing.T) {
	result := GetScore("あいうえおおおえういあ")
	if result != 13 {
		t.Error("Expected 13, got ", result)
	}
}
