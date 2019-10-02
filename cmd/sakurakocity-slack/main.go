package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
)

func main() {
	//client := slack.New(os.Getenv("xoxb-109620016308-587541767702-3LVCDbfSxvqf5rxxSallsSie"))
	client := slack.New(
		"xoxb-109620016308-587541767702-3LVCDbfSxvqf5rxxSallsSie",
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot:", log.Lshortfile|log.LstdFlags)),
	)

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			//rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))
		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello world!", ev.Channel))
		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)
		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)
		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())
		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		default:
			//fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
