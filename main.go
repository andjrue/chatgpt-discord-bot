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

* Bot needs to join a voice channel whenever someone joins one: DONE
	Will still need to send the message as well

* Small issue with above, the bot never actually leaves the channel.
	Need to find a way to check how many members are in the channel at any given time.
	When it's <= 1, the bot can leave/turn off?

* Need to connect to OpenAI API:
	The idea is that when someone joins a voice channel, the bot follows and will text-to-speech
	the "Someone Joined" message

* Maybe deploy this to an EC2 instance?
	Not 100% on this yet. Ideally, this won't run from my computer. I do plan on using it.
	Just not sure if that's the best way to do it, but it would be a fun AWS project.

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

	/*
	                             __xxxxxxxxxxxxxxxx___.
	                        _gxXXXXXXXXXXXXXXXXXXXXXXXX!x_
	                   __x!XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!x_
	                ,gXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx_
	              ,gXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!_
	            _!XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!.
	          gXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXs
	        ,!XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!.
	       g!XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	      iXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	     ,XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx
	     !XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx
	   ,XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx
	   !XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXi
	  dXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	  !XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	   XXXXXXXXXXXXXXXXXXXf~~~VXXXXXXXXXXXXXXXXXXXXXXXXXXvvvvvvvvXXXXXXXXXXXXXX!
	   !XXXXXXXXXXXXXXXf`       'XXXXXXXXXXXXXXXXXXXXXf`          '~XXXXXXXXXXP
	    vXXXXXXXXXXXX!            !XXXXXXXXXXXXXXXXXX!              !XXXXXXXXX
	     XXXXXXXXXXv`              'VXXXXXXXXXXXXXXX                !XXXXXXXX!
	     !XXXXXXXXX.                 YXXXXXXXXXXXXX!                XXXXXXXXX
	      XXXXXXXXX!                 ,XXXXXXXXXXXXXX                VXXXXXXX!
	      'XXXXXXXX!                ,!XXXX ~~XXXXXXX               iXXXXXX~
	       'XXXXXXXX               ,XXXXXX   XXXXXXXX!             xXXXXXX!
	        !XXXXXXX!xxxxxxs______xXXXXXXX   'YXXXXXX!          ,xXXXXXXXX
	         YXXXXXXXXXXXXXXXXXXXXXXXXXXX`    VXXXXXXX!s. __gxx!XXXXXXXXXP
	          XXXXXXXXXXXXXXXXXXXXXXXXXX!      'XXXXXXXXXXXXXXXXXXXXXXXXX!
	          XXXXXXXXXXXXXXXXXXXXXXXXXP        'YXXXXXXXXXXXXXXXXXXXXXXX!
	          XXXXXXXXXXXXXXXXXXXXXXXX!     i    !XXXXXXXXXXXXXXXXXXXXXXXX
	          XXXXXXXXXXXXXXXXXXXXXXXX!     XX   !XXXXXXXXXXXXXXXXXXXXXXXX
	          XXXXXXXXXXXXXXXXXXXXXXXXx_   iXX_,_dXXXXXXXXXXXXXXXXXXXXXXXX
	          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXP
	          XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	           ~vXvvvvXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXf
	                    'VXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXvvvvvv~
	                      'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX~
	                  _    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXv`
	                 -XX!  !XXXXXXX~XXXXXXXXXXXXXXXXXXXXXX~   Xxi
	                  YXX  '~ XXXXX XXXXXXXXXXXXXXXXXXXX`     iXX`
	                  !XX!    !XXX` XXXXXXXXXXXXXXXXXXXX      !XX
	                  !XXX    '~Vf  YXXXXXXXXXXXXXP YXXX     !XXX
	                  !XXX  ,_      !XXP YXXXfXXXX!  XXX     XXXV
	                  !XXX !XX           'XXP 'YXX!       ,.!XXX!
	                  !XXXi!XP  XX.                  ,_  !XXXXXX!
	                  iXXXx X!  XX! !Xx.  ,.     xs.,XXi !XXXXXXf
	                   XXXXXXXXXXXXXXXXX! _!XXx  dXXXXXXX.iXXXXXX
	                   VXXXXXXXXXXXXXXXXXXXXXXXxxXXXXXXXXXXXXXXX!
	                   YXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXV
	                    'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX!
	                    'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXf
	                       VXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXf
	                         VXXXXXXXXXXXXXXXXXXXXXXXXXXXXv`
	                          ~vXXXXXXXXXXXXXXXXXXXXXXXf`
	                              ~vXXXXXXXXXXXXXXXXv~
	                                 '~VvXXXXXXXV~~
	                                       ~~
	*/

	select {} // DO NOT REMOVE THIS

	/*
	* After hours of googling, someone suggested this on a random stackoverflow post.
	* The bot was automatically shutting off before any voice updates could be detected.
	* Adding the select statement fixed the issue, messages are now sending
	* when someone joins a call.
	 */
}
