.PHONY: test
test:
	@go test -v ./coinspaid

.PHONY: integration
integration:
	@go test -v -tags=integration ./test/integration