func topKFrequent(nums []int, k int) []int {
	count:=make(map[int]int)
	res:=[]int{}

	for i:=0;i<len(nums);i++{
		count[nums[i]]++
	}
	for j:=0;j<k;j++{
		best:=0
		max:=0

		for num,c:=range count{
			c++
			if c >max{
				max=c
				best=num
			}
		}
		res=append(res,best)
		delete(count, best)

	}
	return res
}
