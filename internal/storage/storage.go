package storage

import (
	"github.com/bwmarrin/discordgo"
	"github.com/reasje/go_discord_thread_saver/internal/protocols"
	"time"
)

var (
	// channels
	ID          string
	Name        string
	Description string
	Type        string

	Channels []protocols.Channel

	// This app is designed to work with a single ticker
	// This ticker handles the message sending action
	Ticker *time.Ticker

	// Config
	Token     string //To store value of Token from config.json .
	BotPrefix string // To store value of BotPrefix from config.json.

	Config protocols.Config //To store value extracted from config.json.

	// Bot
	BotId string
	GoBot *discordgo.Session
)
