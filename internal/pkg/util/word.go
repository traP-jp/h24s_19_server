package util

import "strings"

//import "github.com/google/uuid"

func CountRunes(s string) int {
	return len([]rune(s))
}

func GetFirstRune(s string) rune {
	if len(s) == 0 {
		return rune(0)
	}
	return []rune(s)[0]
}

func GetLastRune(s string) rune {
	if len(s) == 0 {
		return rune(0)
	}
	return []rune(s)[CountRunes(s)-1]
}

func CheckShiritori(a string, b string) bool {
	return GetLastRune(a) == GetFirstRune(b)
}

var vowel = map[rune]rune{
	'あ': 'あ', 'い': 'い', 'う': 'う', 'え': 'え', 'お': 'お', 'か': 'あ', 'き': 'い', 'く': 'う', 'け': 'え', 'こ': 'お', 'さ': 'あ', 'し': 'い', 'す': 'う', 'せ': 'え', 'そ': 'お', 'た': 'あ', 'ち': 'い', 'つ': 'う', 'て': 'え', 'と': 'お', 'な': 'あ', 'に': 'い', 'ぬ': 'う', 'ね': 'え', 'の': 'お', 'は': 'あ', 'ひ': 'い', 'ふ': 'う', 'へ': 'え', 'ほ': 'お', 'ま': 'あ', 'み': 'い', 'む': 'う', 'め': 'え', 'も': 'お', 'や': 'あ', 'ゆ': 'う', 'よ': 'お', 'ら': 'あ', 'り': 'い', 'る': 'う', 'れ': 'え', 'ろ': 'お', 'わ': 'あ', 'を': 'お', 'ん': 'ん', 'が': 'あ', 'ぎ': 'い', 'ぐ': 'う', 'げ': 'え', 'ご': 'お', 'ざ': 'あ', 'じ': 'い', 'ず': 'う', 'ぜ': 'え', 'ぞ': 'お', 'だ': 'あ', 'ぢ': 'い', 'づ': 'う', 'で': 'え', 'ど': 'お', 'ば': 'あ', 'び': 'い', 'ぶ': 'う', 'べ': 'え', 'ぼ': 'お', 'ぱ': 'あ', 'ぴ': 'い', 'ぷ': 'う', 'ぺ': 'え', 'ぽ': 'お', 'ゃ': 'あ', 'ゅ': 'う', 'ょ': 'お', 'っ': 'っ',
}

func Getvowel(s rune) rune {
	v, ok := vowel[s]
	if !ok {
		return s
	}
	return v
}

func GetLastValidRune(s string) rune {
	runes := []rune(s)
	if len(s) == 0 {
		return rune(0)
	}
	if strings.Contains("ゃゅょっ", string(runes[len(runes)-1])) {
		return runes[len([]rune(s))-2]
	}
	return runes[len([]rune(s))-1]
}

func IsValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if strings.Contains("ゃゅょっ", string(GetFirstRune(s))) {
		return false
	}
	return true
}

func SiritoriCheck(lastReading, currentReading string) bool {
	return GetLastValidRune(lastReading) == GetFirstRune(currentReading)
}

func Insert(s string, index int, value string) string {
	return string([]rune(s)[:index]) + value + string([]rune(s)[index:])
}

func Count(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}
