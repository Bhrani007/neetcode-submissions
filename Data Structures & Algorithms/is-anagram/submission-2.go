func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	myChar:= make(map[rune]int)

	for i:=0;i<len(s);i++{
		myChar[rune(s[i])]++
	}

	for i:=0;i<len(t);i++{
		myChar[rune(t[i])]--

		if myChar[rune(t[i])]< 0{
			return false
		}
	}
	return true 
}
