package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("============================================")
	fmt.Println("             Main test GitHUB               ")
	fmt.Println("============================================")

	for {
		time.Sleep(3 * time.Second)
		answers := []string{
			"глазницы",
			"кармы",
			"души",
			"ноги",
			"руки",
			"почки",
			"печени",
			"пятки",
			"причёски",
			"любви",
			"кармы",
		}

		fmt.Println("Размер моей ", answers[rand.Intn(len(answers))], " равен ", rand.Intn(200), " см")
	}

}
