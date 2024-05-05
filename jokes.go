package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomJoke() string {
	jokes := [20]string{
		"Why shouldn't you marry a calendar? It's days are numbered",
		"What do you call people who sleep in their socks? Tiny.",
		"Why was the broom late for school? It over-swept.",
		"What did the comforter say after falling off the bed? Oh, sheet!",
		"How much do you pay deer for a day's work? A hundred bucks.",
		"Why don't trees watch scary movies? They get petrified.",
		"Are all math puns bad? No, just sum.", "What does a house wear? Address.",
		"I got rid of my vacuum cleaner. It was just gathering dust.",
		"What did the Dalmatian say after dinner? \"Woof, that hit the spot!\"",
		"What kind of bug tells time? A clock roach.",
		"What did one beer say to the other? It's ale good.",
		"What do you call coffee with a six sense? Déjà brew.",
		"What's a llama's favorite movie? \"Alpaca-lypse Now.\"",
		"Why shouldn't you make a dinosaur mad? Because you'll get Jurass-kicked.",
		"What do math books wear under their covers? Alge-bras.",
		"How do movie stars stay cool? They have many fans.",
		"It was an emotional wedding. Even the cake was in tiers.",
		"Where do rabbits go for breakfast? IHOP.",
		"Why did the belt get arrested? It held up a pair of pants.",
	}

	joke = jokes[rand.Intn(len(jokes))]
	
	return joke

}
