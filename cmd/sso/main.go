package main

import (
	"fmt"
	"gRPCLearn/iternal/config"
	"log/slog"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New()
	}
}
