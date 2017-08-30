build: assets
	go build -o ./releases/pubpub main.go

assets:
	go-bindata -o pubpub/resources.go -pkg pubpub resources/

release:
	gox
