package util

import "testing"

func TestShouldReturnCorrect(t *testing.T) {
	var convertedTime = StringToTime("2024-05-10T11:00:00")

	if convertedTime.Year() != 2024 {
		t.Errorf("Should be returned 2024")
	}

	if convertedTime.Month() != 05 {
		t.Errorf("Should be returned 05")
	}

	if convertedTime.Hour() != 11 {
		t.Errorf("Should be returned 11")
	}

}
