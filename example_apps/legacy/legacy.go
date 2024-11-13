package main

import (
	"log"
	"math/rand/v2"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	for {
		time.Sleep(time.Duration(rand.Int64N(5 * time.Second.Nanoseconds())))

		time.Local = time.FixedZone("", gofakeit.IntRange(-20, 20)*30*60)

		logger.Printf("%s | %s | method: %s | user: %s | ip: %s | message: %s\n",
			time.Now().Format(time.Layout),
			gofakeit.LogLevel("syslog"),
			gofakeit.HTTPMethod(),
			gofakeit.Username(),
			gofakeit.IPv4Address(),
			gofakeit.Comment())
	}
}
