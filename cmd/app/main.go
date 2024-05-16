package main

import (
	"fmt"
	"wb_test_L0/internal/config"
	"log/slog"
)


const (
	configPath = "config/config.yaml"
)

func main(){
	cfg := config.MustLoad(configPath)
	fmt.Println(cfg)
}

func setupLogger(env string){

}