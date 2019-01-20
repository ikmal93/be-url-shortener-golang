package config

import "math/rand"

func RandomCharacter(n int) string {

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Redirect struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

type Response struct {
	Code    int    `json:"code"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
