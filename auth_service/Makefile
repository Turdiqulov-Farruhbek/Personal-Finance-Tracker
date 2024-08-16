

DBURL="postgres://postgres:root@localhost:5432/financer_auth?sslmode=disable"
exp:
	export DBURL="postgres://postgres:root@localhost:5432/financer_auth?sslmode=disable"

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




# doc-run:
# 	docker run --network="host" --name libCont lib_service

# doc-con:
# 	docker run --network="host" --publish  9090:3030 recruitment-service

# doc-network:
# 	docker network create library

# doc-run-allow:
# 	docker run --network="library" -p 9090:2121 library_service

# doc-continer-all-del:
# 	docker rm $(docker ps -aq)
# 	docker rmi $(docker ps images)

# doc-network-all:
# 	docker network ls


# doc-query:
# 	docker ps
# 	docker image rmi auth_service
# 	docker container ls
# 	docker container prune
# 	docker container rm libCont
# 	docker build -t lib_service .
# 	docker network ls
# 	docker images
# 	docker exec -it post-cont psql -U postgres -d library_db
# 	docker system prune -af
#  	docker volume prune -af