package main

import (
	"fmt"
	"sort"
)

func main(){
	//sortKey()
	sortValue()
}

func sortKey(){
	m := map[string]int{
		"b":1,
		"a":2,
		"d":3,
		"c":4,
	}
	var s = make([]string,len(m))
	var i = 0
	for key := range m{
		s[i] = key
		i++
	}
	sort.Strings(s)
	fmt.Println(s)
	var newm = make(map[string]int,len(m))
	for _,k := range s{
		newm[k] = m[k]
	}
	fmt.Println(newm)
}

type Pair struct{
	val int
	key string
}
type Pairs []Pair

func (p Pairs) Len() int{
	return len(p)
}
func (p Pairs) Less(i,j int) bool{
	return p[i].val > p[j].val
}

func (p Pairs)Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

func sortValue(){
	m := map[string]int{
		"b":1,
		"a":2,
		"d":3,
		"c":4,
	}
	//转化为结构体排序

	var pairsSlice Pairs
	for k,v := range m{
		pairsSlice = append(pairsSlice,Pair{val:v,key:k})
	}
	sort.Sort(pairsSlice)
	fmt.Println(pairsSlice)
	var newm = make(map[string]int,len(m))
	for _,val := range pairsSlice{
		newm[val.key] = val.val
		fmt.Println(val.val,val.key)
	}
	//由于map是无序的 所以这里打印出来仍然是无序的
	fmt.Println(newm)
}
