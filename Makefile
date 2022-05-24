# Makefile

.PHONY: build-smartcontracts-base
build-smartcontracts-base:
	docker build -t vegaprotocol/smartcontracts-base:local ./smartcontracts-base

.PHONY: build-ganache
build-ganache:
	docker build -t vegaprotocol/ganache:local ./ganache

	