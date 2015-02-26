package strreverse

func Reverse(s string) string {
	runeForm := []rune(s)

	for i, j := 0, len(runeForm)-1; i < len(runeForm)/2; i, j = i+1, j-1 {
		runeForm[i], runeForm[j] = runeForm[j], runeForm[i]
	}

	return string(runeForm)
}
