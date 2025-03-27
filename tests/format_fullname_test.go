package main

import "testing"

func BenchmarkFormatFullName(b *testing.B) {
	testCases := []string{
		"РОМАН\tРЯПИСОВ\tЮРЬЕВИЧ",       // Russian (Cyrillic)
		"ИВАН    ИВАНОВ  ИВАН ИВАНОВИЧ", // Russian with multiple spaces
		"   ПЕТР ПЕТРОВИЧ  ",            // Russian with leading/trailing spaces
		"JOHN DOE",                      // English
		"maría de la cruz",              // Spanish with lowercase and prepositions
		"jean-luc picard",               // French with hyphenated name
		"ÖMER ŞIMŞEK",                   // Turkish with special characters
		"ŁUKASZ KOWALSKI",               // Polish
		"πρόδρομος παπαδόπουλος",        // Greek
	}

	for i := 0; i < b.N; i++ {
		for _, name := range testCases {
			formatName(name)
		}
	}
}
