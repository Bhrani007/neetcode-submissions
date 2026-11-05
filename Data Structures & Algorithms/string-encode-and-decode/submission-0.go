type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res:=""
	for i:=0;i<len(strs);i++{
		arr:=[]rune(strs[i])
		res=res+strconv.Itoa(len(arr))+"#"

		for j:=0;j<len(arr);j++{
			res=res+string(arr[j])
		}

	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	num:=0
	arr:=encoded
	isFound:=false
	res:=[]string{}

	for i:=0;i<len(arr);i++{
		if arr[i]>='0'&& arr[i]<='9'{
			num = num*10 + int(arr[i]-'0')
		}
		if string(arr[i])=="#"{
			isFound=true
			i++
		}
		if isFound{
			word:=""
			for j:=0;j<num;j++{
				word=word+string(arr[i])
				i++
			}
			res=append(res, word)
			i--
			isFound=false
			num=0
		}

	}
	return res
}
