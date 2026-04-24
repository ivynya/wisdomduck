package main

import "math/rand"

func generateWisdom() string {
	return generateSubject() + " " + generateVerb() + " " + generateNoun()
}

func generateSubject() string {
	subject := []string{
		"Duck",
		"He",
		"They",
		"It",
		"She",
		"That",
		"Wisdom",
		"You",
		"Knowledge",
		"Excellence",
		"One",
		"Friend",
		"Guardian",
		"Teacher",
		"Goose", // Sponsored by AidanJSmith
	}

	return subject[rand.Intn(len(subject))]
}

func generateVerb() string {
	verb := []string{
		"perceives",
		"understands",
		"values",
		"exemplifies",
		"is",
		"has",
		"delivers",
		"provides",
		"consumes",
		"abolishes",
		"accelerates",
		"achieves",
		"acts with",
		"pictures",
		"confronts",
		"locates",
		"teaches",
		"serves",
		"accesses",
		"incurs",
		"introduces",
		"shows",
		"extracts",
		"develops",
		"theorizes",
		"quacks",
		"meows",
		"barks",
		"enforces",
	}

	return verb[rand.Intn(len(verb))]
}

func generateNoun() string {
	noun := []string{
		"trans rights",
		"friendship",
		"kindness",
		"compassion",
		"love",
		"care",
		"entertainment",
		"judgment",
		"jurisdiction",
		"beauty",
		"play",
		"art",
		"understanding",
		"knowledge",
		"theory",
		"power",
		"development",
		"strategy",
		"professionalism",
		"speed",
		"tribulations",
		"potato",
		"communication",
		"warning",
		"income",
		"attention",
		"assistance",
		"perception",
		"imagination",
		"wrath",
		"cats",
		"catgirl",
		"catboy", // Sponsored by AidanJSmith
	}

	return noun[rand.Intn(len(noun))]
}
