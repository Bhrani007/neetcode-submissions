func getConcatenation(nums []int) []int {
    res:=[]int{}
	
	for j:=0;j<2;j++{
		for i:=0;i<len(nums);i++{
			res=append(res ,nums[i])
		}
	}
	return res
}
