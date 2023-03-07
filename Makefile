.DEFAULT_GOAL = run
production : VERSION = $(shell date +'%Y-%m-%d %H:%M:%S')

prep:
	go get github.com/joho/godotenv
	go get github.com/gin-gonic/gin
	go get gorm.io/gorm
	go get gorm.io/driver/postgres

production:
	go build -ldflags \
		"-X 'main.VERSION=Production build on $(VERSION)'" \
		"-X 'main.DEFAULTCONFIGFILE=/etc/galt/galt.yaml'"

release: production

buildgrpc:
	protoc  --go_out=paths=source_relative,plugins=grpc:. server/galt.proto

build:
	go build .

clean:
	rm -f galt

run: clean buildgrpc build
	go run .

br: clean build
	go run .
