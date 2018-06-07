package main

import (
	"fmt"
	"encoding/json"
)

// 结构体变量必须大写开头
type mapText struct {
	Lesson string `json:"lesson"`
	Blogtype string `json:"type"`
}

type blog struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Text mapText`json:"text"`
}


func main()  {

	c:=make(map[string]interface{})
	c["name"]="angus"
	c["title"]="go web program"
	c["text"]=map[string]interface{}{
		"lesson":"Json",
		"type":"makeup",
	}

	data,err:=json.MarshalIndent(c,"","	")
	if(err != nil){
		fmt.Println("Json makeip error ", err.Error())
		return
	}
	fmt.Println(string(data))

	var stblog blog
	//var test map[string]interface{}
	err=json.Unmarshal([]byte(data),&stblog)
	if(err != nil){
		fmt.Println("parse json error ",err.Error())
		return
	}
	fmt.Println(stblog)
}
