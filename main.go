package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kitakitabauer/slack-client-go/config"
	"github.com/kitakitabauer/slack-client-go/slack"
	"github.com/pkg/errors"
)

// endPoint is the URL which you want to check API response
const endPoint = "http://example.com"

type status int

const (
	normal status = iota
	alert
)

var interval = 1 * time.Minute

var iconEmojiMap = map[status]string{
	normal: config.IconEmoji(),
	alert:  ":cry:",
}

var currentStatus status = normal

var msg = slack.Msg{
	Username: config.Username(),
	Channel:  config.Channel(),
}

func sendToSlack(text string, st status) error {
	if currentStatus == st {
		return nil
	}

	s := slack.Slack{}
	msg.Text = text + "\nendPoint: " + endPoint
	msg.IconEmoji = iconEmojiMap[st]

	_, err := s.Send(msg)
	if err != nil {
		err := errors.Wrap(err, "Failed to send to slack.")
		return err
	}

	currentStatus = st
	return nil
}

func exec() {
	res, err := http.Get(endPoint)
	if err != nil {
		err := sendToSlack("Request to example.com failed.", alert)
		if err != nil {
			fmt.Printf("error: %#v", err)
			return
		}
		return
	}
	defer res.Body.Close()

	err = sendToSlack("example.com is back to normal!", normal)
	if err != nil {
		fmt.Printf("error: %#v", err)
		return
	}

	return
}

func main() {
	t := time.NewTicker(interval)
	for {
		select {
		case <-t.C:
			exec()
		}
	}
}
