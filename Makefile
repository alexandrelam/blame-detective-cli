# Variables
BINARY = blamed
LINUX_BINARY = linux-blamed
GO_FILE = cmd/blamed.go
INSTALL_PATH = ~/.local/bin

# Determine the operating system
UNAME_S := $(shell uname -s)

# Commands
ifeq ($(UNAME_S),Darwin)
	BINARY := macos-$(BINARY)
else ifeq ($(UNAME_S),Linux)
	BINARY := $(LINUX_BINARY)
endif

# Targets
all: build copy

build:
	go build -o $(BINARY) $(GO_FILE)
	chmod +x $(BINARY)

copy:
	cp $(BINARY) $(INSTALL_PATH)/blamed

clean:
	rm -f $(BINARY)
