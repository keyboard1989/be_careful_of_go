package main

import "fmt"
import "strings"

//range map和数组
//map 不是按顺序导出的
//数组是

func main() {
	d := []string{
		"hello1",
		"hello2",
		"hello3",
		"hello4",
		"hello5",
		"hello6",
		"hello7",
	}

	for _, key := range d {
		fmt.Println(key)
	}

	//该输出是有序的
	// hello1
	// hello2
	// hello3
	// hello4
	// hello5
	// hello6
	// hello7
	fmt.Println(strings.Repeat("=", 30))

	d1 := map[string]string{
		"hello1": "world1",
		"hello2": "world2",
		"hello3": "world3",
		"hello4": "world4",
		"hello5": "world5",
		"hello6": "world6",
		"hello7": "world7",
	}

	for key, _ := range d1 {
		fmt.Println(key)
	}
	//该输出是序的，每次都不一样
}
