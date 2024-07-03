package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func uploadToS3(sess *session.Session, bucket, key, content string) error {
	svc := s3.New(sess)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader([]byte(content)),
	})
	return err
}

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

		s3Session, err := session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		})
		if err != nil {
			fmt.Println("S3 error", err)
			return
		}
		bucket := os.Getenv("S3_BUCKET")
		key := fmt.Sprintf("text_files/%s-2.txt", user.Username)
		content := user.Username

		err = uploadToS3(s3Session, bucket, key, content)
		if err != nil {
			fmt.Println("Error uploading to S3\n", err)
			return
		}

		fmt.Println("Message uploaded to S3")
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

func playMessage(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	// Return Something
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
