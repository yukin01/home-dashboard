.PHONY: auth enable
enable:
	gcloud services enable cloudresourcemanager.googleapis.com
	gcloud services enable run.googleapis.com
	gcloud services enable iam.googleapis.com
	gcloud services enable cloudscheduler.googleapis.com
	gcloud services enable iap.googleapis.com

auth:
	gcloud auth application-default login
