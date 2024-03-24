package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
)

/* TO-DO's

* Message Sending: DONE
	Need the bot to send a message when someone joins a chat channel
	Update on this, totally unnecessary. See the readme.

* Bot needs to join a voice channel whenever someone joins one: DONE
	Will still need to send the message as well - No it won't.

* Small issue with above, the bot never actually leaves the channel.
	Need to find a way to check how many members are in the channel at any given time.
	When it's <= 1, the bot can leave/turn off?

* Need to connect to OpenAI API:
	Not happening

* Maybe deploy this to an EC2 instance?
	Not 100% on this yet. Ideally, this won't run from my computer. I do plan on using it.
	Just not sure if that's the best way to do it, but it would be a fun AWS project.

*** I'm going to need to rethink a lot of this. Encoding and Decoding audio is a lot more difficult than I thought it
	it would be. I thought I was 90% done, I'm probably 30% done. Sigh.

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
		fmt.Print("User has joined voice channel")
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
	} else {
		fmt.Println("User not in voice channel")
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
