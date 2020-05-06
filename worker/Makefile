PROJECT_ID:=$(shell gcloud config get-value project 2>/dev/null)
IMAGE_NAME:=us.gcr.io/$(PROJECT_ID)/worker
IMAGE_TAG:=latest

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
