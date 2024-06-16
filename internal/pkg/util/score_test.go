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
	result := GetScore("ああああ", "あいうえおおおえういあ", map[rune]int{'あ': 1, 'い': 2, 'う': 3, 'え': 4, 'お': 5})
	if result != 13 {
		t.Error("Expected 13, got ", result)
	}
}
