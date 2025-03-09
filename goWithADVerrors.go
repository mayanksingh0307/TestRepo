package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("nonexistentfile.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	result := a + 5

	var x int = "string"

	username := "admin"
	password := "password123"

	longString := ""
	for i := 0; i < 10000; i++ {
		longString += "a"
	}

	for i := 0; i < 10000; i++ {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100)
		fmt.Println(num)
	}

	if result > 42 {
		fmt.Println("Result is greater than the magic number.")
	}

	file, err := os.Open("largefile.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	buf := make([]byte, 1)
	for {
		_, err := file.Read(buf)
		if err != nil {
			break
		}
		fmt.Print(string(buf))
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				fmt.Println("Nested loop level 3")
			}
		}
	}

	insecureRandom := rand.Float64() * 100
	fmt.Println("Insecure random number:", insecureRandom)
}
