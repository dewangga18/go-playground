package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	mainSort()
// }

func mainSort() {
	
	// sort int slice
	ages := []int{10, 20, 30, 5, 15}
	sort.Ints(ages)
	fmt.Println(ages)

	// reverse int slice
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	fmt.Println(ages)
	
	// sort string slice
	names := []string{"John", "Doe", "Jane", "Bob"}
	sort.Strings(names)
	fmt.Println(names)
	
	// sort float64 slice
	floats := []float64{1.0, 2.0, 3.0, 5.0, 1.5}
	sort.Float64s(floats)
	fmt.Println(floats)
	
	users := []User{
		{"John", "20"},
		{"Doe", "25"},
		{"Jane", "22"},
		{"Bob", "28"},
	}

	fmt.Println(users)

	// sort using sort.Sort (implement sort.Interface)
	sort.Sort(UserSlice(users))

	fmt.Println(users)

	// sort slice of struct using sort.Slice (no need to implement sort.Interface)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println(users)
}

// make contract of sort interface
type User struct {
	Name string
	Age  string
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}