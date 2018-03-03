package main

import (
	"fmt"
	"math/rand"
)

func position(input []int,begin int,end int) int {
	x:=input[end]
	i :=begin-1
	for j:=begin;j<end ;j++ {
		if (input[j]<x){
			i++
			temp:=input[i]
			input[i]=input[j]
			input[j]=temp
		}
	}
	input[end]=input[i+1]
	input[i+1]=x
	return i+1
}

func sort(input []int,begin int,end int) []int{
	if(begin==end||begin>end){
		return input
	}
	q:=position(input,begin,end)
	sort(input,begin,q-1)
	sort(input,q+1,end)
	return input
}
func main(){
	fmt.Println("test")
	data:=make([]int,8)
	for i := 0; i<len(data);i++{
		data[i]=rand.Int()
		data[i]=9-i;
	}
	data[1]=5
	fmt.Printf("before sort %v\n",data)
	sort(data,0,len(data)-1)
	fmt.Printf("%v",data)
}