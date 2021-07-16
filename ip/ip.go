package main

import "fmt"

type IpAddress [4]byte

func (ip IpAddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", int(ip[0]), int(ip[1]), int(ip[2]), int(ip[3]))
}

func main() {
	hosts := map[string]IpAddress{
		"baires": {10, 2, 3, 4},
		"google": {1, 2, 3, 4},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
