package main

import "fmt"

func main() {
	authorsMap := make(map[string]anime)

	addAuthors(authorsMap)

	// Trying to update a value
	authorsMap["Masashi Kishimoto"] = anime{
		name:    "Naruto Shipudden",
		genre:   "Shonen",
		ratings: 5,
	}

	printAuthors(authorsMap)

	// a struct can be used as key
	animeMap := make(map[anime]string)
	animeMap[authorsMap["Masashi Kishimoto"]] = KISHIMOTO_SENSAI
	animeMap[authorsMap["Eichiro Oda"]] = ODA_SENSAI
	animeMap[authorsMap["Tite Kobo"]] = KOBO_SENSAI
	animeMap[authorsMap["Haruichi Furudate"]] = FURUDATE_SESNSAI

	printAnimes(animeMap)
}

type anime struct {
	name    string
	genre   string
	ratings int
}

func addAuthors(authors map[string]anime) {
	authors["Masashi Kishimoto"] = anime{
		name:    "Naruto",
		genre:   "Shonen",
		ratings: 5,
	}

	authors["Eichiro Oda"] = anime{
		name:    "One Piece",
		genre:   "Shonen",
		ratings: 5,
	}

	authors["Tite Kobo"] = anime{
		name:    "Bleach",
		genre:   "Shonen",
		ratings: 5,
	}

	authors["Haruichi Furudate"] = anime{
		name:    "Haikyuu",
		genre:   "Sports",
		ratings: 5,
	}
}

func printAnimes(animes map[anime]string) {
	for anime, author := range animes {
		fmt.Println("<-<->->")
		fmt.Println("Anime", anime.name, "is written by", author, "and it has a rating of", anime.ratings)
	}
}

func printAuthors(authors map[string]anime) {
	for name, anime := range authors {
		fmt.Println("<-<->->")
		fmt.Println("Author", name, "has written the anime", anime.name, "with a rating of", anime.ratings)
	}
}

const KISHIMOTO_SENSAI = "Kisihimoto Sensai"
const ODA_SENSAI = "Oda Sensai"
const KOBO_SENSAI = "Kobo Sensai"
const FURUDATE_SESNSAI = "Furudate Sensai"
