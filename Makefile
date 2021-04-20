SHELL := /bin/bash

PLIST=info.plist
ICON=icon.png
EXEC_BIN=slack-alfred-workflow
DIST_FILE=slack.alfredworkflow
GO_SRCS=$(shell find -f . \( -name \*.go \))

all: $(DIST_FILE)

$(EXEC_BIN): $(GO_SRCS)
	go build -o $(EXEC_BIN)

$(DIST_FILE): $(EXEC_BIN) $(CREDITS) $(PLIST)
	zip -r $(DIST_FILE) $(PLIST) $(ICON) $(EXEC_BIN)
