package main

import "fmt"


func BinarySerach()int{
	arr := [7]int{20,30,40,45,89,92,101}
	tar := 10;
	start := 0;
	end := len(arr) - 1;

	for start <= end{
		mid := (start + end) / 2;
		if tar < arr[mid]{
			end = mid - 1;
		}else if  tar > arr[mid]{
			start = mid + 1;
		}else{
			return mid
		}
	}
	return -1;
}



func main(){
	ans := BinarySerach()
	if ans == -1{
		fmt.Println("Element is not present in the given array")
	}else{
		fmt.Println("Element present in the array :",ans)
	}
}