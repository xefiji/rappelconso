build: 
	go build -o deployment/dist/rappelconso cmd/server.go

build-linux: 
	GOOS=linux go build -o deployment/dist/rappelconso cmd/server.go

run: build
	deployment/dist/rappelconso

reset: build
	deployment/dist/rappelconso -reset

migrate-up:
	tern migrate -c migrations/tern.conf -m migrations

dist: build-linux
	@rm -rf deployment/dist/migrations
	@cp -R migrations/ deployment/dist/migrations 2>/dev/null || :
	@docker build -t gcr.io/rappelconso/rappelconso deployment

deploy:
	docker push gcr.io/rappelconso/rappelconso
	gcloud run deploy rappelconso --image gcr.io/rappelconso/rappelconso \
	--platform managed \
	--allow-unauthenticated \
	--region europe-west1
