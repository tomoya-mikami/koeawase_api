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
