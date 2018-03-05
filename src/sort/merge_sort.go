package main

import "fmt"

func merge(input1 []int,input2 []int) []int{
	var length int = len(input2)+len(input1)
	var res []int = make([]int,length)
	i:=0
	j:=0
	z:=0
	for ;i<len(input1)&&j<len(input2);{
		if(input1[i]>input2[j]){
			res[z]=input2[j]
			j++
		}else{
			res[z]=input1[i]
			i++
		}
		z++
	}
	for ;i<len(input1);{
		res[z]=input1[i]
		i++
		z++
	}
	for ;j<len(input2); {
		res[z] = input2[j]
		j++
		z++
	}
	return res
}


func merge_sort(inputdata []int,begin int,end int) []int{
	if(end-begin==0){
		var res []int=make([]int,1)
		res[0]=inputdata[begin]
		return res
	}
	mid:=(begin+end)/2
	input1:=merge_sort(inputdata,begin,mid)
	input2:=merge_sort(inputdata,mid+1,end)
	inputdata=merge(input1,input2)
	//fmt.Println(inputdata)
	return inputdata
}

func main()  {
	var data []int=make([]int,9)
	for i:=0;i<9;i++{
		data[i]=9-i
	}
	data[2]=5
	data[5]=9
	fmt.Println(data)
	data1:=merge_sort(data,0,len(data)-1)
	fmt.Println(data1)
}
