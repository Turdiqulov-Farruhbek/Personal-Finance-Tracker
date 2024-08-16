

DBURL="postgres://postgres:root@localhost:5432/delivery_product?sslmode=disable"
exp:
	export DBURL="postgres://postgres:root@localhost:5432/delivery_product?sslmode=disable"

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh $(CURRENT_DIR)
	
swag-gen:
	~/go/bin/swag init -g ./api/router.go -o docs force 1	