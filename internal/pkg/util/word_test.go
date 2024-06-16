package util

import "testing"


func TestGetFirstRune(t *testing.T) {
	result := GetFirstRune("あいうえお")
	if result != 'あ' {
		t.Error("Expected 'あ', got ", result)
	}
}

func TestGetLastRune(t *testing.T) {
	result := GetLastRune("あいうえお")
	if result != 'お' {
		t.Error("Expected 'お', got ", result)
	}
}

func TestInsert(t *testing.T) {
	result := Insert("あいうえお", 3, "か")
	if result != "あいうかえお" {
		t.Error("Expected 'あいかうえお', got ", result)
	}
	result = Insert("あいうえお", 5, "か")
	if result != "あいうえおか" {
		t.Error("Expected 'あいうえおか', got ", result)
	}
}

func TestGetLastVaildRune(t *testing.T) {
	result := GetLastValidRune("あいうえお")
	
	if result != 'お' {
		t.Error("Expected 'お', got ", result)
	}
	result = GetLastValidRune("あいうえおっ")
	if result != 'お' {
		t.Error("Expected 'お', got ", result)
	}
}

func TestSiritoriCheck(t *testing.T) {
	result := SiritoriCheck("あいうえお", "かきくけこ")
	if result != false {
		t.Error("Expected false, got ", result)
	}
	result = SiritoriCheck("あいうえお", "おかきくけこ")
	if result != true {
		t.Error("Expected true, got ", result)
	}
	result = SiritoriCheck("こうきゅうしゃ", "しんか")
	if result != true {
		t.Error("Expected true, got ", result)
	}
}

func TestIsValid(t *testing.T) {
	result := IsValid("あいうえお")
	if result != true {
		t.Error("Expected true, got ", result)
	}
	result = IsValid("ゃい")
	if result != false {
		t.Error("Expected false, got ", result)
	}
}

func TestCount(t *testing.T) {
	result := Count("あいうえおしゃしん")
	m := map[rune]int{
		'あ': 1, 'い': 1, 'う': 1, 'え': 1, 'お': 1, 'し': 2, 'ゃ': 1, 'ん': 1,
	}
	for k, v := range result {
		if m[k] != v {
			t.Error("Expected ", m, ", got ", result)
		}
	}
}