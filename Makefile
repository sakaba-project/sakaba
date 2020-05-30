# build target files
GOFILES := $(shell find . -name "*.go")
# peoject name
PROJECT_NAME = sakaba
# executable binary file name
EXECUTE_FILENAME = $(PROJECT_NAME)

# make run -> ./sakaba
.PHONY: run
run: $(EXECUTE_FILENAME)
	./$(EXECUTE_FILENAME)

# link with run command.
# do `make build` before `make run`
$(EXECUTE_FILENAME): build

# make build -> go build
# output `sakaba` by binary image
.PHONY: build
build: $(GOFILES)
	go build -o $(EXECUTE_FILENAME)

# make clean -> rm sakaba
.PHONY: clean
clean:
	rm $(EXECUTE_FILENAME)

# make test -> go test
.PHONY: test
test:
	go test -v -covermode=count

# make fmt -> go fmt
.PHONY: fmt
fmt:
	gofmt -s -w $(GOFILES)

