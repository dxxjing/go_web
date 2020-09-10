package main

import (
	"fmt"
	"sort"
)

//使用sort包 自定义类型排序 该类型必须实现 sort.Interface接口 才可使用其包内的方法
type mySlice []int

func (m mySlice) Len() int{
	return len(m)
}

func (m mySlice) Less(i,j int) bool{
	return m[i] < m[j]
}

func (m mySlice) Swap(i,j int){
	m[i],m[j] = m[j],m[i]
}

func main(){
	//intSliceSort()
	mysort()
}

//自定义类型排序
func mysort(){
	var s = []int{1,4,9,5,3}
	sort.Sort(mySlice(s))
	fmt.Println(s)
	sort.Sort(sort.Reverse(mySlice(s)))
	fmt.Println(s)
}


func intSliceSort(){
	var s = []int{1,4,9,5,3}
	//默认升序
	sort.IntSlice(s).Sort()
	fmt.Println(s)
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}