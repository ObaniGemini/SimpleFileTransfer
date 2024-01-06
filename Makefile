COMMAND=go build
TARGET=SimpleFileTransfer

GOFILES=sft/main.go

all: $(TARGET)

$(TARGET): $(GOFILES)
	go get fyne.io/fyne/v2@latest
	go mod tidy
	$(COMMAND) -o $@ $^

test:
	./$(TARGET)

clean:
	rm -rf $(TARGET)