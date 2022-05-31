package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Coloring the output
	cC := "\033[36m"

	fmt.Println("------------------------------------------------------")
	fmt.Println(string(cC), "	RSA Private Key Calculator")
	fmt.Println("------------------------------------------------------")

	// Taking user input for the value of n
	fmt.Print("Enter the value of n: ")
	rdr := bufio.NewReader(os.Stdin)
	n, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input => ", err)
		return
	}
	n = strings.TrimSuffix(n, "\n")

	// Taking user input for the value of e
	fmt.Print("Enter the value of e: ")
	rdr2 := bufio.NewReader(os.Stdin)
	e, err := rdr2.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input => ", err)
		return
	}
	e = strings.TrimSuffix(e, "\n")
	fmt.Println("=======================================================")
	fmt.Println("Value of n = ", n)
	fmt.Println("=======================================================")
	fmt.Println("Value of e = ", e)
	fmt.Println("=======================================================")

	// Conversion to integer
	n_val, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}

	e_val, err := strconv.Atoi(e)
	if err != nil {
		log.Fatal(err)
	}

	// Creating slice
	s := make([]int, 0, 4)

	// Calculating primes
	for i := 2; i < n_val+1; i++ {
		for n_val%i == 0 {
			s = append(s, i)
			n_val = n_val / i
		}
	}

	// Print Primes
	fmt.Println("Primes are = ", s)
	fmt.Println("=======================================================")
	fmt.Println("The value of prime p = ", s[0])
	fmt.Println("The value of prime q = ", s[1])
	fmt.Println("=======================================================")

	// Computing phi of n
	phi_n := (s[0] - 1) * (s[1] - 1)
	fmt.Println("The value of prime ϕ(n) = ", phi_n)
	fmt.Println("=======================================================")

	// Calculating Private Key (d)
	for j := 0; j < e_val+1; j++ {
		if (1+(j*phi_n))%e_val == 0 {
			fmt.Println("D is calculated by substituting i with ", j)
			fmt.Println("=======================================================")
			d := (1 + (j * phi_n)) / e_val
			fmt.Println("The value of private key d = ", d)
			fmt.Println("=======================================================")
		}
	}

}
