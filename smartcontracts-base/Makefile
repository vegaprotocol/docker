# Makefile

#defauls
VERSION=0.8.13

.PHONY: default
default: abi

.PHONY: abi
abi: ABIDIR=contracts-solc-${VERSION}/build/doc/data/files/abi
abi: CONTRACTS=contracts-solc-${VERSION}/contracts
abi: SOLC_CONTAINER=ethereum/solc:${VERSION}
abi:
	@mkdir -p $(ABIDIR)
	@find $(ABIDIR) -iname '*.abi' -o -iname '*_abi.json' -print0 | xargs -0r rm
	docker run --rm \
		--user "$$(id -u):$$(id -g)" \
		-v "$$PWD:$$PWD" \
		-w "$$PWD" \
		$(SOLC_CONTAINER) \
			-o $(ABIDIR)/ --abi ${CONTRACTS}/*.sol
	@find $(ABIDIR) -name '*.abi' \
		| while read -r abifile ; \
	do \
		jq --sort-keys . <"$$abifile" >"$${abifile%.abi}_ABI.json" && rm "$$abifile" ; \
	done

