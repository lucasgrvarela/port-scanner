package main

import "fmt"

func main() {
	result := ScanPort("tcp", "localhost", 1024)
	fmt.Printf("Port result: %v\n", result)

	results := InitialScan("localhost")
	fmt.Println(results)
}
