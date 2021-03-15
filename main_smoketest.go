package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("kubectl", "get", "pods", "-o", "wide").StdoutPipe()
	exec.
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}