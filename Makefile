# some common tasks and shortcuts
OUTPUTDIR=bin
RELEASE=-ldflags "-s -w"

.PHONY:app test

app:
	go build -tags app $(RELEASE) -o $(OUTPUTDIR)/brainFuck
	strip -s $(OUTPUTDIR)/brainFuck

test:
	go test -count=1 ./...

clean:
	rm bin/brainFuck
