package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Portnumber: ")
	portStr, _ := reader.ReadString('\n')
	portStr = strings.TrimSpace(portStr)

	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 48 {
		fmt.Println("Invalid port: 1-48")
		os.Exit(1)
	}

	ip := fmt.Sprintf("172.16.76.%d", 48+port)
	fmt.Printf("Portnumber %d -> %s\n", port, ip)
}