package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}

	query := args[0]

	/*	english := []string{"good", "great", "perfect"}
		turkish := []string{"iyi", "harika", "mükemmel"}

		for i, w := range english {
			if query == w {
				fmt.Printf("%q means %q\n", w, turkish[i])
				return
			}
		}

		fmt.Printf("%q not found \n", query)*/

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükummel",
		"awesome": "mükummel",
	}

	delete(dict, "awesome")

	turkish := make(map[string]string)

	for k, v := range dict {
		turkish[v] = k
	}
	/*
		fmt.Printf("english: %q\nturkish: %q\n", dict, turkish)
		dict["up"] = "yukari"
		dict["down"] = "asagi"
		dict["mistake"] = ""

		fmt.Printf("%#v\n", dict)

		//for k, v := range dict {
		//	fmt.Printf("%q means %#v\n", k, v)
		//}
	*/

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)

	fmt.Printf("# of keys: %#v\n", len(dict))

	//fmt.Printf("Zero Value: %#v\n", dict)
	//fmt.Printf("# of keys: %#v\n", len(dict))
}
