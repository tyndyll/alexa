package main

import (
	"log"

	"github.com/tyndyll/alexa"
)

func main() {
	certURL := "https://s3.amazonaws.com/echo.api/echo-api-cert.pem"
	if valid, err := alexa.ValidateCertificate(certURL, "RjwqmCJgXPJSUW+Zi8s1CzbfAoYO/TtaRl05zw+Hw5OO6jzMNGg/OKOAIhMo8wTVvoPSragaKUr0THii7ulrY50pwpSBmOR6fpYZDHemkQck4/DBCxIGrBVtsXra4geHkZ/1pwtlqLvgJzQIJ2FxCwWemkLwhOLW0ecphSFkXUR3PdDk+Dawt3GpJiLMpbh5B3Ge9ppx78RNcyzp8EAkXag8w9GeOxUGEpAddz/QgMH2NbTCsTR/E6agFNV0/vHm/+zQTVQDw9GIo38lw7VMZrFVFfQkSlYEqRMgGeV37iShMrhxCY5vSYvw0EKxRorHx+9djV8sC5E86hTf+X2xaQ=="); err != nil {
		log.Fatalln(err)
	} else {
		if valid {
			log.Println("Valid certificate")
		} else {
			log.Println("Invalid certificate")
		}
	}
}
