package main

import (
	"fmt"
	"os"
)

func main() {
	env_flag := os.Getenv("FLAG")
	if env_flag != "" {
		fmt.Printf("%s\n", env_flag)
	} else {
		init_regex(flag_format)
		flag := getUniqueFlag()
		if flag != "" {
			fmt.Printf("%s\n", flag)
		} else {
			fmt.Printf("No FLAG for you\n")
		}
	}
}
