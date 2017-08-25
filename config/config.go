package config

import "github.com/kelseyhightower/envconfig"

type base struct {
	IncomingURL string `envconfig:"incomingURL"`
	Channel     string `envconfig:"channel"`
}

type user struct {
	Username  string `envconfig:"username" default:"slack-bot"`
	IconEmoji string `envconfig:"icon_emoji" default:":smile:"`
	IconURL   string `envconfig:"icon_url"`
}

var (
	bConf = &base{}
	uConf = &user{}
)

func init() {
	err := envconfig.Process("", bConf)
	if err != nil {
		panic(err)
	}
	err = envconfig.Process("", uConf)
	if err != nil {
		panic(err)
	}
}

func IncomingURL() string {
	return bConf.IncomingURL
}

func Channel() string {
	return bConf.Channel
}

func Username() string {
	return uConf.Username
}

func IconEmoji() string {
	return uConf.IconEmoji
}

func IconURL() string {
	return uConf.IconURL
}
