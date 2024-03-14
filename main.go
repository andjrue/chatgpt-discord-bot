package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("env error", err)
		return
	}

	disKey := os.Getenv("DISCORD_KEY")
	sess, err := discordgo.New("Bot " + disKey)
	if err != nil {
		fmt.Println("Error creating session", err)
		return
	}
	err = sess.Open()
	if err != nil {
		fmt.Println("Oopsie", err)
		return
	}
	fmt.Println("Bot is online")
	sess.Close()

}
