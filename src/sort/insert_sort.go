package main

import "fmt"

func  insert_sort(inputdata []int) []int  {
	length:=len(inputdata)
	var temp int
	for i:=0;i<length ;i++  {
		temp=inputdata[i]
		j:=i-1
		for ;j>=0 ;j--  {
			if(inputdata[j]>temp){
				inputdata[j+1]=inputdata[j]
			}else{
				break
			}
		}
		inputdata[j+1]=temp
	}
	fmt.Println(inputdata)
	return inputdata
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
