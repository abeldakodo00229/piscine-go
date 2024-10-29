package piscine

func ToLower(s string) string {
	Tab := []rune(s)
	for i := range Tab {
		if Tab[i] >= 'A' && Tab[i] <= 'Z' {
			Tab[i] = Tab[i] + 32
		}
	}
	return string(Tab)
}
