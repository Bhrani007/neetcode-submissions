func groupAnagrams(s []string) [][]string {
	groups:=make(map[[26]int][]string)
	keys:=[][26]int{}

	for i:=0;i<len(s);i++{
		arr:=[]rune(s[i])
		var count [26]int

		for j:=0;j<len(arr);j++{
			count[arr[j]-'a']++
		}
		_,found:=groups[count]
		if !found{
			keys=append(keys, count)
		}
		groups[count]=append(groups[count], s[i])
	}
	res:=[][]string{}

	for i:=0;i<len(keys);i++{
		res=append(res, groups[keys[i]])
	}
	return res
}
