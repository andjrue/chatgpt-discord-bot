package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

/* TO-DO's

* I think I figured out a better way to do this. We could set up an S3 bucket to store the .mp3's, then include a username
  lookup to see if that user has an existing file. If they do, we play the file. If they don't, we make an API call and
  have a new one created.

* This is good for a few reasons, the main of which being cutting down on cost. OpenAI's API isn't initially expensive,
  but that could add up quickly over time. I also don't have any plans of making this public, so there would only be 6ish
  calls ever made.

*/

func userHasJoinedVoice(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("env error - userHasJoinedvoice", err)
		return
	}

	chanID := os.Getenv("CHANNEL_ID")
	// fmt.Println("Function call is working") DEBUG

	if v.BeforeUpdate == nil {
		if v.UserID == s.State.User.ID { // Bot was sending messages when they joined
			return
		}
		fmt.Println("User has joined voice channel")
		user, err := s.User(v.VoiceState.UserID)
		if err != nil {
			fmt.Println("Error getting user ID", err)
			return
		}

		message := fmt.Sprintf("%s has joined the channel", user.Username)

		_, err = s.ChannelMessageSend(chanID, message)
		if err != nil {
			fmt.Println("Error sending message", err)
			return
		}
	}
}

func botJoinChannel(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	if v.UserID == s.State.User.ID {
		return
	}

	if v.ChannelID != "" {
		_, err := s.ChannelVoiceJoin(v.GuildID, v.ChannelID, false, true)
		if err != nil {
			fmt.Println("Error joining voice channel", err)
		}
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("env error - Main", err)
		return
	}

	disKey := os.Getenv("DISCORD_KEY")
	sess, err := discordgo.New("Bot " + disKey)
	if err != nil {
		fmt.Println("Error creating session", err)
		return
	}

	sess.AddHandler(userHasJoinedVoice)
	sess.AddHandler(botJoinChannel)

	err = sess.Open()
	if err != nil {
		fmt.Println("Oopsie", err)
		return
	}

	fmt.Println("Bot is online")

	select {} // DO NOT REMOVE THIS

	/*
		The select statement prevents the bot from turning off.
		Thank you random person on stackoverflow.
	*/
}
