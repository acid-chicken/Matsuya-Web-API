#! /usr/bin/make

REPO:=github.com/makotia/Matsuya-Web-API

gen: clean generate

build:
	@go build -o api

bootstrap:
	@goagen bootstrap -d $(REPO)/design

main:
	@goagen main -d $(REPO)/design

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d $(REPO)/design
	@goagen swagger -d $(REPO)/design
	@goagen client  -d $(REPO)/design
