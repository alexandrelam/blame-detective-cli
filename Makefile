BINARY = macos-blamed
GO_FILE = cmd/blamed.go

all: build copy

build:
	go build -o $(BINARY) $(GO_FILE)
	chmod +x $(BINARY)

copy:
	cp $(BINARY) ~/.local/bin/blamed

clean:
	rm -f $(BINARY)
