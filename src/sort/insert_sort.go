package main

import (
	"fmt"
)
func change( first *int,second *int){
	temp:=*first
	*first=*second
	*second=temp
}
func change1(first int,second int){
	temp:=first
	first=second
	second=temp
}
func insert_sort(input []int) []int{
	fmt.Println(len(input))
	length:=len(input)
	var pos int=0
	for i:=0;i<length;i++{
		pos=i
		for j:=i;j<length;j++ {
			if input[pos]>input[j]{
				pos=j
			}
		}
		change(&input[i],&input[pos])
		//change1(input[i],input[pos])
	}
	fmt.Println(input)
	return input
}

func main()  {
	var data []int=make([]int,9)
	for i:=0;i<9;i++{
		data[i]=9-i
	}
	data[2]=5
	fmt.Println(data)
	insert_sort(data)
}
