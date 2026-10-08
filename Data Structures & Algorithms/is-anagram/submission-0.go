func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false 
	}
	myChars:=make (map[rune]int)
	arr1:=[]rune(s)
	arr2:=[]rune(t)
	

	for i:=0;i<len(arr1);i++{
		myChars[arr1[i]]++
	}

	for i:=0;i<len(arr2);i++{
		myChars[arr2[i]]--
		if myChars[arr2[i]]<0{
			return false
		}
	}
	return true
}
