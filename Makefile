.PHONY: build
build:
	docker build -t koeawase_api .

.PHONY: run
run:
	docker run -p 8080:8080 -e FIRESTORE_EMULATOR_HOST=host.docker.internal:8812 -v $(PWD):/workdir --name api --rm koeawase_api

.PHONY: stop
stop:
	docker stop api

.PHONY: firestore-start
firestore-start:
	gcloud beta emulators firestore start --host-port=localhost:8812

.PHONY: wire
wire:
	$(GOPATH)/bin/wire

.PHONY: bash
bash:
	docker exec -it api bash

.PHONY: set-project
set-project:
	gcloud config set project koeawase

.PHONY: deploy
deploy:
	gcloud app deploy
