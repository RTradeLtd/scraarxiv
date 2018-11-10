#LENSVERSION=`git describe --tags`
LENSVERSION="testing"

lens:
	@make cli

# Build lens cli
.PHONY: cli
cli:
	@echo "===================  building Scraarxiv CLI  ==================="
	rm -f lens temporal-lens
	go build -ldflags "-X main.Version=$(LENSVERSION)" ./cmd/temporal-scraarxiv
	@echo "===================          done           ==================="

# Build our vendor
.PHONY: vendor
vendor:
	@echo "===================  rebuilding vendor  ==================="
	rm -rf vendor
	dep ensure -v
	@echo "===================          done           ==================="

# Set up test environment
.PHONY: testenv
WAIT=3
testenv:
	@echo "===================   preparing test env    ==================="
	docker-compose -f scraarxiv.yml up -d
	sleep $(WAIT)
	@echo "===================          done           ==================="
