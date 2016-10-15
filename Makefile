.PHONY: service
service:
	cd docker;docker-compose up -d

.PHONY: log_server
log_server:
	cd docker;docker-compose logs -f goserver

clean_containers:
	docker rm $$(docker stop $$(docker ps -q -a))
