package main

import "fmt"

func main() {
	v1 := 1
	var v2 int64 = 2
	var v3 string = "hello"
	var v4 float32 = 3.2
	MyPrintf(v1, v2, v3, v4)
}

func MyPrintf(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Println(v, "is int.")
		case int64:
			fmt.Println(v, "is int64.")
		case string:
			fmt.Println(v, "is string.")
		case float32:
			fmt.Println(v, "is float32.")
		}
	}
}
