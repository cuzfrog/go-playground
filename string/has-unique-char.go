package string

func hasUniqueChar(s *string) bool {
	cnt := make(map[rune]int)
	for _, v := range *s {
		if cnt[v] > 0 {
			return false
		}
		cnt[v]++
	}
	return true
}

func hasUniqueChar2(s *string) bool {
	mark := uint64(0)
	for _, v := range *s {
		if v < 'A' || v > 'z' {
			panic("only support A-z")
		}
		offset := v - 'A'
		if mark>>offset&1 > 0 {
			return false
		}
		mark = mark | 1<<offset
	}
	return true
}
