.PHONY: test
test:
	go build -o .test tests/main/** && ./.test