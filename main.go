package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	log.Println("Por favor elige 1 o 2:")

	for scanner.Scan() {

		if election := scanner.Text(); election == "1" || election == "2" {

			log.Println("Test passed")

			break

		} else {

			log.Println("Please choose a valid option")

		}

	}

	log.Println("eneded")

}

var scanner = bufio.NewScanner(os.Stdin)
