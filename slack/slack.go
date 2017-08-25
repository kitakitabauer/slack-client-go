package slack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kitakitabauer/slack-client-go/config"
	"github.com/pkg/errors"
)

type slack interface {
	Send(msg Msg) (string, error)
}

type Slack struct{}

type Msg struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconUrl   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func (s *Slack) Send(msg Msg) (string, error) {
	if msg.Text == "" || msg.Channel == "" {
		error := errors.Errorf("SlackMsg: %#v", msg)
		return "", errors.Wrap(error, "Illegal argument.")
	}

	params, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}

	resp, err := http.PostForm(
		config.IncomingURL(),
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
