package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	disKey = os.GetEnv("DISCORD_KEY")

	discord, err := discord.go("Bot " + disKey)
	if err != nil {
		fmt.Println("Error creating session")
	}
}
