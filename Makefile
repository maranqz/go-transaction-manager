DIR=$(PWD)

GO_TEST=cd ./sh && bash ./go.test.sh
GO_TEST_COVERAGE=cd ./sh && bash ./go.test.coverage.sh

GO_TEST_WITH_REAL_DB=--tags=with_real_db

test:
	$(GO_TEST)

test.with_real_db:
	$(GO_TEST) $(GO_TEST_WITH_REAL_DB)

test.coverage:
	$(GO_TEST_COVERAGE)

test.coverage.with_real_db:
	$(GO_TEST_COVERAGE) $(GO_TEST_WITH_REAL_DB)

fmt:
	go fmt ./...

lint:
	golangci-lint run -v --timeout=2m

generate:
	go generate ./...


go.mod.tidy:
	cd sh && sh ./go.mod.tidy.sh

go.mod.vendor:
	cd sh && sh ./go.mod.vendor.sh
