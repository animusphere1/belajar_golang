package main

import "fmt";

func main(){
	var list = []string{"nama","saya"};

	for i := 0; i < len(list); i++ {
		fmt.Print(list[i])
	}
}