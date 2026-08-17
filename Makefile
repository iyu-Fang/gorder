.PHONY: gen
gen: genproto genopenapi

.PHONY: genproto
genproto:
	@bash ./scripts/genproto.sh

.PHONY: genopenapi
genopenapi:
	@bash ./scripts/genopenapi.sh