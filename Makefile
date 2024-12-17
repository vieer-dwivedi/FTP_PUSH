.DEFAULT_RUN := deploy
init:
	go mod init automations
	go mod tidy
	go get ./
build: clean
	@echo Building application
	make init
	GOOS=linux GOARCH=amd64 go build -o ftp_uploader main.go
	make clean

clean:
	rm -rf main.zip
	rm -rf go.*