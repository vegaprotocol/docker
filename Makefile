# Makefile

.PHONY: build-smartcontracts-base
build-smartcontracts-base:
	docker build --platform=linux/arm64 -t vegaprotocol/smartcontracts-base:v1.5.0 ./smartcontracts-base

.PHONY: build-ganache
build-ganache:
	docker build --platform=linux/arm64 --progress=plain -t vegaprotocol/ganache:v1.5.0 ./ganache
	
.PHONY: build-timescaledb
build-timescaledb:
	docker build -t vegaprotocol/timescaledb:local ./timescaledb
