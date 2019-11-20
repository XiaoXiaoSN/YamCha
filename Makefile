#
# YamCha Makefile.
#

PROJECT=yamcha
SERVERURL=https://28966197.ngrok.io
SECRET=a43bc61e3433ec2023d96141ad42bf6b
TOKEN=9RCh7QaXx25rPTgyeDuqF1jjSzexpcSQubLkbxjfjYRM7gTeQkekJsY2QEexXqNX1Lf3aQK3W3W0tvV6G1ESA19dPJfGM/rmiXlahQpKEZgeRg6FXOZ0obiPFUdWSIrAuskBZAIuyT1O9Zu9IRBaAQdB04t89/1O/w1cDnyilFU=
PORT=18180

init:
	go mod init $(PROJECT)
	go mod tidy
	cp configs/config.dev.yml configs/config.yml

run:
	LINECORP_PLATFORM_CHANNEL_SERVERURL=$(SERVERURL) \
	LINECORP_PLATFORM_CHANNEL_CHANNELSECRET=$(SECRET) \
	LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN=$(TOKEN) \
	PORT=$(PORT) \
	go run cmd/line/*.go

release:
	heroku container:login
	heroku container:push web
	heroku container:release web