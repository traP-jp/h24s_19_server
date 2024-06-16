package util

func getLength(s string) int {
	return len([]rune(s))
}

func countPreSuffixMatch(s string) int {
	count := 0
	size := getLength(s)
	runes := []rune(s)
	for i := 0; i < size; i++ {
		if runes[i] == runes[size-1-i] {
			count++
		} else {
			break
		}
	}
	return count
}

func IsPalindrome(s string) bool {
	return countPreSuffixMatch(s) == getLength(s)
}

func CountSuffixRhyme(s1, s2 string) int {
	size1 := getLength(s1)
	size2 := getLength(s2)
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	count := 0
	for i := 0; i < size1 && i < size2; i++ {
		if Getvowel(runes1[size1-1-i]) == Getvowel(runes2[size2-1-i]) {
			count++
		} else {
			break
		}
	}
	return count
}

func GetMode(s string) int {
	runes := []rune(s)
	size := getLength(s)
	m := make(map[rune]int)
	max := 0
	for i := 0; i < size; i++ {
		m[runes[i]]++
		if m[runes[i]] > max {
			max = m[runes[i]]
		}
	}
	return max
}

func GetScoreByFrequency(s string, count map[rune]int) int {
	score := 0
	runes := []rune(s)
	size := getLength(s)
	for i := 0; i < size; i++ {
		score = score + 100/count[runes[i]]
	}
	return score
}

var coefLength = 1
var coefSuffixRhyme = 1
var coefMode = 10
var coefFrequency = 1

func GetScore(lastReading, currentReading string, count map[rune]int) int {
	return coefLength*getLength(currentReading) + coefSuffixRhyme*CountSuffixRhyme(lastReading, currentReading) + coefMode*GetMode(currentReading) + coefFrequency*GetScoreByFrequency(currentReading, count)
}
