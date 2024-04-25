include ./scripts/rscli.mk

buildc:

runc:
	docker volume create \
		--driver local \
		--opt type=none \
		--opt device=$(shell pwd)/data \
		--opt o=bind \
		call_a_courier_be_sqlite

	docker run \
 		-v call_a_courier_be_sqlite:/app/data \
		-p 8080:8080 \
		 call-a-courier-be:local