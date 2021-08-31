package reddit

//GetSymbols get all symbols
func GetStockSymbols(s string) (m []string) {
	res := symbolRegex.FindAllStringSubmatch(s, -1)
	for _, v := range res {
		if len(v) == 3 {
			m = append(m, v[2])
		}
	}
	return
}
