COMMAND=go build
TARGET=SimpleFileTransfer

ifeq ($(OS),Windows_NT)
	GOFILES = sft/windows.go
else
	GOFILES = sft/unix.go
endif

GOFILES += sft/math.go sft/gui.go sft/main.go

all: $(TARGET)

$(TARGET): $(GOFILES)
	go get fyne.io/fyne/v2@latest
	go install fyne.io/fyne/v2/cmd/fyne@latest
	go mod tidy
	$(COMMAND) -o $@ $^

test:
	./$(TARGET)

clean:
	rm -rf $(TARGET)