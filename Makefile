run:
	export LIVEPEER_COM_API_KEY=$(LIVEPEER_COM_API_KEY)
	export LIVEPEER_PRICING_TOOL=$(LIVEPEER_PRICING_TOOL)
	export POWERGATE_ADDR=$(POWERGATE_ADDR)
	export IPFS_REV_PROXY_ADDR=$(IPFS_REV_PROXY_ADDR)
	export TRUSTED_MINERS=$(TRUSTED_MINERS)
	export POLL_INTERVAL=$(POLL_INTERVAL)
	export PORT=$(PORT)
	export PINATA_API_KEY=$(PINATA_API_KEY)
	export PINATA_SECRET_KEY=$(PINATA_SECRET_KEY)
	export IPFS_GATEWAY=$(IPFS_GATEWAY)
	export POW_TOKEN=$(POW_TOKEN)
	export DEMUX_URL=$(DEMUX_URL)
	export MONGO_URI=$(MONGO_URI)
	go run cmd/main.go

build:
	go build cmd/main.go

test:
	go test -v ./dataservice
