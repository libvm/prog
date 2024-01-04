package main

import "fmt"

func tests() {
	result1, err1 := encodeOrDecode("text.txt.out", "hello.txt", "decode")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := encodeOrDecode("text.txt.out", "", "decode")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(result2)
	}
	result3, err3 := encodeOrDecode("text.txt", "hello.txt", "encode")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(result3)
	}
	result4, err4 := encodeOrDecode("text.txt", "", "encode")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(result4)
	}
}
