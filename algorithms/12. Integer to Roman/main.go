func intToRoman(num int) string {
	// 查表法，直接對應每個位數
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds  := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	var sb strings.Builder
	sb.WriteString(thousands[num/1000])
	sb.WriteString(hundreds[num%1000/100])
	sb.WriteString(tens[num%100/10])
	sb.WriteString(ones[num%10])

	return sb.String()
}