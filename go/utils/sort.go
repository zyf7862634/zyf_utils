package utils

import (
	"sort"
	"fmt"
)

//结构体排序
type Person struct {
	Name string
	Age  int
}
type PersonSlice []*Person  //也可以不用指针

func (a PersonSlice) Len() int {
	return len(a)
}

func (a PersonSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PersonSlice) Less(i, j int) bool { return a[j].Age < a[i].Age } //按照年龄逆序


func test() {

	var bb []*Person
	people1 := &Person{"zhang san", 12}
	bb = append(bb, people1)
	people2 := &Person{"li si", 30}
	bb = append(bb, people2)
	people3 := &Person{"wang wu", 52}
	bb = append(bb, people3)
	people4 := &Person{"zhao liu", 26}
	bb = append(bb, people4)

	//for i  := range bb {
	//	fmt.Println(bb[i].Name)
	//}
	sort.Sort(PersonSlice(bb)) // 按照 Age 的逆序排序
	for i  := range bb {
		fmt.Println(bb[i].Name)
	}
	sort.Sort(sort.Reverse(PersonSlice(bb))) // 按照 Age 的升序排序
	for i  := range bb {
		fmt.Println(bb[i].Name)
	}
}