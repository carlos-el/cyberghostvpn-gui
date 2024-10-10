run: # For developing with sudo activated
	sudo --preserve-env `which go` run .
clean:
	rm -rf bin/

build:
	mkdir -p bin
	go build -o bin/
build-prod:
	mkdir -p bin
	go build -ldflags "-s -w" -o bin/ -tags prod
build-linux:
	mkdir -p bin
	fyne package -os linux -icon assets/logo.png --release --tags prod
	mv cyberghostvpn-gui.tar.xz bin/
