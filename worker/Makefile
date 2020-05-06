PROJECT_ID:=$(shell gcloud config get-value project 2>/dev/null)
IMAGE_NAME:=us.gcr.io/$(PROJECT_ID)/worker
IMAGE_TAG:=latest
REGION:=us-central1

get-project-id:
	@echo $(PROJECT_ID)

helper-snap:
	sudo ln -s /snap/google-cloud-sdk/current/bin/docker-credential-gcloud /usr/local/bin

configure:
	gcloud auth configure-docker

build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)

run:
	docker run -it --rm -p 8080:8080 \
		-e PROJECT_ID=$(PROJECT_ID) \
		-e REMO_ACCESS_TOKEN=$(REMO_ACCESS_TOKEN) \
		-e GOOGLE_APPLICATION_CREDENTIALS=/service-account.json \
		-v $(PWD)/service-account.json:/service-account.json \
		$(IMAGE_NAME):$(IMAGE_TAG)

deploy:
	gcloud run deploy worker-service \
		--image $(IMAGE_NAME):$(IMAGE_TAG) \
		--region $(REGION) \
		--platform managed

test:
	go test -v ./...