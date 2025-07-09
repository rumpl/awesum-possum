.PHONY: build install clean

build:
	go build -o awesum-possum

install: build
	mv awesum-possum /usr/local/bin/

clean:
	rm -f awesum-possum