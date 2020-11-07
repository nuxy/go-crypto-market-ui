VERSION = 0.0.1
PACKAGE = github.com/nuxy/go-crypto-market-ui/cmd/crypto-market-ui
LDFLAGS = "-X main.Version=$(VERSION)"

run:
	go run $(GOFLAGS) -ldflags $(LDFLAGS) ./cmd/crypto-market-ui/main.go

build:
	go build -x $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/crypto-market-ui $(PACKAGE)

install:
	go install -x $(GOFLAGS) -ldflags $(LDFLAGS) $(PACKAGE) 

buildall:
	GOOS=darwin  GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/crypto-market-ui-$(VERSION)-osx-64         $(PACKAGE)
	GOOS=freebsd GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/crypto-market-ui-$(VERSION)-freebsd-64     $(PACKAGE)
	GOOS=linux   GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/crypto-market-ui-$(VERSION)-linux-64       $(PACKAGE)
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/crypto-market-ui-$(VERSION)-windows-64.exe $(PACKAGE)
	GOOS=windows GOARCH=386   go build $(GOFLAGS) -ldflags $(LDFLAGS) -o ./bin/crypto-market-ui-$(VERSION)-windows-32.exe $(PACKAGE)
