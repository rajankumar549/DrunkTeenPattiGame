
run:
	@go build -race && ./DrunkTeenPattiGame

test:
	@echo " >> running tests"
	@go test -v -race ./...

hooks:
	@echo "Copying [pre-commit] hook..."
	@cp -f ./pre-commit .git/hooks/pre-commit
	@echo "Making [pre-commit] hook executable..."
	@chmod +x .git/hooks/pre-commit
	@echo "Done."

config-test:
	@go build -race && ./DrunkTeenPattiGam -t