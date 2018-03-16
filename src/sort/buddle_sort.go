package main

import "fmt"

func swap(first *int, second *int){
	temp:=*first
	*first=*second
	*second=temp
}

func buddle_sort(inputdata []int) []int{
	fmt.Println(len(inputdata))
	for i:=0;i<len(inputdata) ;i++  {
		for j:=0;j<len(inputdata)-1 ;j++  {
			if(inputdata[j]>inputdata[j+1]){
				swap(&inputdata[j],&inputdata[j+1])
			}
		}
	}
	return inputdata
}

func main(){
	var inputdat []int=make([]int,10)
	for i:=0;i<10;i++{
		inputdat[i]=i
	}
	inputdat[5]=10
	inputdat[2]=8
	fmt.Println(inputdat)
	buddle_sort(inputdat)
	fmt.Println(inputdat)
}
