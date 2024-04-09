bs:
	go build -o build/server cmd/server/main.go

rs:
	build/server -c config/server/example.yml

bc:
	go build -o build/client cmd/client/main.go

rc:
	build/client -c config/client/example.yml

gen:
	cd proto; \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto