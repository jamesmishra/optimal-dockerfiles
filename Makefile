help:
	@echo "Makefile commands:"
	@echo "\tmake build:\t\tBuild the containers in this repo."
	@echo "\tmake up:\t\tRun the 'hello' container service."

build:
	# Build and start the caching services.
	docker-compose up --build --detach apt-cacher-ng devpi-server
	# Wait 2 seconds for the caching services to come up.
	sleep 2
	# Build our hello container.
	-docker-compose build hello
	# Shut down the caches.
	docker-compose down --remove-orphans

up:
	-docker-compose up hello

.PHONY: help build up
