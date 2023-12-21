package main

import (
	"math/rand"
	"time"
)

var (
	// generated using ChatGPT
	names = []string{
		"Alex",
		"Mia",
		"Juan",
		"Aisha",
		"Mohammad",
		"Isabella",
		"Ahmed",
		"Mei-Ling",
		"Leonardo",
		"Amara",
		"Rajesh",
		"Fatima",
		"Mateo",
		"Priya",
		"Carlos",
		"Lila",
		"Felix",
		"Gabriela",
		"Arjun",
		"Anika",
		"Giovanni",
		"Leila",
		"Manuel",
		"Isla",
		"Ali",
		"Lina",
		"Hugo",
		"Freya",
		"Javier",
		"Aylin",
		"Diego",
		"Emilia",
		"Ibrahim",
		"Yuki",
		"Aiden",
		"Elina",
		"Zhihao",
		"Anaya",
		"Mustafa",
		"Sienna",
		"Lily",
		"Amelie",
		"Maya",
		"Eva",
		"Oliver",
		"Samuel",
		"Liam",
		"Daniel",
		"Elijah",
		"Anna",
	}

	// generated using ChatGPT
	items = []string{
		"football",
		"box",
		"watermelon",
		"teddy bear",
		"basketball",
		"book",
		"gourmet chocolates",
		"holiday-themed socks",
		"personalized ornament",
		"miniature christmas tree",
		"holiday scented candles",
		"christmas-themed mug",
		"handmade soap set",
		"puzzle",
		"coffee sampler",
		"cozy knit scarf",
		"mini bottle of champagne",
		"essential oil diffuser",
		"festive cookie cutters",
		"mini photo album",
		"handwritten holiday card",
		"custom-made keychain",
		"a small plant",
		"pocket-sized board games",
		"holiday-themed puzzle",
		"popcorn seasoning kit",
		"mini holiday wreath",
		"wine sampler",
		"mini art supplies kit",
		"snow globe",
		"mini gingerbread house kit",
		"pocket-sized sketchbook",
		"festive face mask",
		"pocket-sized umbrella",
		"mini cheese and charcuterie board",
		"festive cocktail mixers",
		"mini holiday music cd",
		"handmade jewelry",
		"mini fairy lights",
		"miniature snowman kit",
		"funny holiday socks",
		"mini magnetic dartboard",
		"scratch-off world map",
		"mini photo frame",
		"reusable shopping bag",
		"mini hot sauce sampler",
		"pocket-sized tool kit",
		"mini bonsai tree kit",
		"holiday-themed coasters",
		"custom engraved keyring",
	}
)

func generateCustomerWithRandomWait() Customer {
	// to simulate a random customer arrival while the post office is open
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return randomCustomer()
}

func randomCustomer() Customer {
	return Customer{
		Name: randomName(),
		Item: randomItem(),
	}
}

func randomName() string {
	return names[rand.Intn(len(names))]
}

func randomItem() string {
	return items[rand.Intn(len(items))]
}
