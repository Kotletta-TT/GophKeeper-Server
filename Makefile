bs:
	go build -o build/server cmd/server/main.go

rs:
	build/server -c config/server/example.yml

bc:
	go build -o build/client cmd/client/main.go

rc:
	build/client -c config/client/example.yml