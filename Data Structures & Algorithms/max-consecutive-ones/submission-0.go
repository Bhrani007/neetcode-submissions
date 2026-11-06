func findMaxConsecutiveOnes(nums []int) int {
	res:= 0

	max:=0
	current:=0

	for i := 0 ; i < len (nums) ; i++ {
		
		
		if nums[i]==1{
			current++
			if current>max{
				max=current
			}
		}else if nums[i]==0{
			current=0
		}
	} 
	res=max

	return res
}
