package linebot

import (
	"fmt"
	"log"
	"net/http"

	"yamcha/pkg/storage"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBot define basic line bot interface
type LineBot interface {
	CallbackHandle(http.ResponseWriter, *http.Request)
}

// YamchaLineBot app
type YamchaLineBot struct {
	storage storage.Storage
	bot     *linebot.Client
}

// NewYambotLineBot create a Yamcha line bot
func NewYambotLineBot(channelSecret, channelToken string, storage storage.Storage) (LineBot, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	log.Println("Channel Secret:", channelSecret)
	log.Println("Channel Token:", channelToken)

	return &YamchaLineBot{
		storage: storage,
		bot:     bot,
	}, nil
}

// CallbackHandle function for http server
func (app *YamchaLineBot) CallbackHandle(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		log.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if err := app.replyText(event.ReplyToken, "我...我看不懂貼圖 QQ"); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := app.replyText(event.ReplyToken, "Got followed event"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeUnfollow:
			log.Printf("Unfollowed this bot: %v", event)
		case linebot.EventTypeJoin:
			if err := app.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if err := app.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeBeacon:
			if err := app.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
				log.Print(err)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}
