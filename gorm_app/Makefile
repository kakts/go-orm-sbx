.PHONY: run
run:
	go run main.go

TOOLS=\
	github.com/swaggo/swag/cmd/swag@latest

.PHONY: install-tools
install-tools:
	@for tool in $(TOOLS) ; do \
		go install -v $$tool; \
	done

.PHONY: doc
doc: # swagger用にコードからdocsにswagger情報を書き出す
	swag init --parseDependency true