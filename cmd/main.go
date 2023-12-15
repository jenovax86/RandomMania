package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomFactGenerator() {
	var facts = []string{"$30 of raw popcorn translates into $3,000 in sales at movie theaters",
		"'Barack Obama' born August 4 1961 was the 44th President of the United States.",
		"'Dina Thanthi' was founded by S. P. Adithanar a lawyer trained in Britain and practised in Singapore with its first edition from Madurai in 194",
		"'fan' is short for 'fanatic'",
		"'Katharine Hepburn' has won four OSCAR awards more than any actor in film industry.",
		"'Mumbai' is called as the 'Business Capital' of India.",
		"'Nitrous Oxide' is called as Laughing Gas.",
		"160 cars can drive side by side on the Monumental Axis in Brazil, the world's widest road.",
		"2 out of 5 people marry their first love",
		"2,500 newborn babies will be dropped in the next month.",
		"20% of China's plants are used in medicine",
		"20% of tuxedo rentals take place in May.",
		"20252 is Smokey the Bear's own zip code.",
		"21% of people don't make their bed in the morning",
		"160 cars can drive side by side on the Monumental Axis in Brazil, the world's widest road.",
		"17% of Intel employees are Indians.",
	}
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := newRand.Intn(len(facts))
	fmt.Println(facts[randomNumber])
}

func main() {
	randomFactGenerator()
}
