#
# YamCha Makefile.
#

PROJECT=yamcha
SERVERURL=https://28966197.ngrok.io
SECRET=a43bc61e3433ec2023d96141ad42bf6b
TOKEN=LtyRVj3I8gicFRpuMV5tMvF/2RMWSPEWxgFT2Oicl4XAtf3QZQdgR6ZiX74um7i21Lf3aQK3W3W0tvV6G1ESA19dPJfGM/rmiXlahQpKEZg2ugosnB8R2nvnVXVgcmiOhNfq0F0iuaF+JbtcqxtFdFGUYhWQfeY8sLGRXgo3xvw=
init:
	go mod init $(PROJECT)
	go mod tidy
	cp configs/config-build.yml configs/config.yml

run:
	# @wire ./...
	YAMCHA_CONFIG=$(CURDIR)/configs \
	LINECORP_PLATFORM_CHANNEL_SERVERURL=$(SERVERURL) \
	LINECORP_PLATFORM_CHANNEL_CHANNELSECRET=$(SECRET) \
	LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN=$(TOKEN) \
	go run main.go line

run-test:
	# @wire ./...
	YAMCHA_CONFIG=$(CURDIR)/configs \
	YAMCHA_CONFIG_NAME=config-test.yml \
	LINECORP_PLATFORM_CHANNEL_SERVERURL=$(SERVERURL) \
	LINECORP_PLATFORM_CHANNEL_CHANNELSECRET=$(SECRET) \
	LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN=$(TOKEN) \
	go test ./...


docker-build:
	docker build . -t yamcha

docker-run:
	docker run -e LINECORP_PLATFORM_CHANNEL_SERVERURL=$(SERVERURL) \
	-e LINECORP_PLATFORM_CHANNEL_CHANNELSECRET=$(SECRET) \
	-e LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN=$(TOKEN) \
	yamcha 

release:
	heroku container:login
	heroku container:push web
	heroku container:release web