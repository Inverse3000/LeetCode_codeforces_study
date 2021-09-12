package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	b := a[2:5]
	fmt.Println(len(b))
	b = append(b, 2)
	b = append(b, 102)
	b = append(b, 203)
	b = append(b, 6832)
	b[0] = 6
	fmt.Println(b)
	fmt.Println(a)
	/*c := []int{1,2}
	fmt.Println(cap(c))
	c = append(c,4)
	fmt.Println(cap(c))*/
	fmt.Println("Copying Slice")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s3)

	//Delete elements
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
	fmt.Println(cap(s2))
	//Pop front
	s2 = s2[1:]
	fmt.Println(s2)
	//Pop back
	s2 = s2[:len(s2)-1]
	fmt.Println(s2)
	mp := map[string]int{}
	mp["lol"] = 1
	fmt.Println(mp["lol"])
	if _, ok := mp["haha"]; ok {
		fmt.Println("nice")
	}

}
