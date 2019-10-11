package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

//LinkCustomer : A chatbot DB to store account link information.
type LinkCustomer struct {
	//Data from CustData from provider.
	Name   string
	Age    int
	Desc   string
	Nounce string
	//For chatbot linked data.
	LinkUserID string
}

var linkedCustomers []LinkCustomer

var jsonData = []byte(`{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "http://linecorp.com/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "飲茶 ver.0.0.1 beta",
        "weight": "bold",
        "size": "xl"
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "新增訂單",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "查詢訂單",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "刪除訂單",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "WEBSITE",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "spacer",
        "size": "sm"
      }
    ],
    "flex": 0
  }
}`)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch {
				case strings.EqualFold(message.Text, "快速回復按鈕"):
					_ = quickReplyButton(bot, event)
					return
				case strings.EqualFold(message.Text, "t"):
					date := "20191010"
					res, err := bot.GetNumberMulticastMessages(date).Do()
					fmt.Println(res, err)
					return
				case strings.HasPrefix(message.Text, "飲茶"):
					_ = responseTemplate(bot, event)
					return
				}
			}

		} else if event.Type == linebot.EventTypeAccountLink {
			log.Println("event.Type == linebot.EventTypeAccountLink")
		}
	}
}

func responseTemplate(bot *linebot.Client, event *linebot.Event) error {
	if container, err := linebot.UnmarshalFlexMessageJSON(jsonData); err != nil {
		log.Println("err:", err)
		return err
	} else {
		bot.ReplyMessage(
			event.ReplyToken,
			linebot.NewFlexMessage("alt text", container),
		).Do();
	}
	// container := &linebot.BubbleContainer{
    //     Type: linebot.FlexContainerTypeBubble,
    //     Body: &linebot.BoxComponent{
    //         Type:   linebot.FlexComponentTypeBox,
    //         Layout: linebot.FlexBoxLayoutTypeHorizontal,
    //         Contents: []linebot.FlexComponent{
    //             &linebot.TextComponent{
    //                 Type: linebot.FlexComponentTypeText,
    //                 Text: "Hello,",
    //             },
    //             &linebot.TextComponent{
    //                 Type: linebot.FlexComponentTypeText,
    //                 Text: "World!",
    //             },
    //         },
    //     },
    // }
    // if _, err := bot.ReplyMessage(
    //     event.ReplyToken,
    //     linebot.NewFlexMessage("alt text", container),
	// ).Do(); err != nil {
	// 	log.Println("err:", err)
	// 	return err
	// }
	// message:= linebot.NewFlexMessage("alt text", container)
	
    return nil
}




func quickReplyButton(bot *linebot.Client, event *linebot.Event) error {
	_, err := bot.ReplyMessage(
		event.ReplyToken,
		linebot.NewTextMessage("Welcome").
			WithQuickReplies(linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton(
					"",
					linebot.NewMessageAction("account link", "link")),
				linebot.NewQuickReplyButton(
					"",
					linebot.NewMessageAction("list user", "list")),
			)),
	).Do()
	if err != nil {
		log.Println("err:", err)
		return err
	}

	return nil
}
