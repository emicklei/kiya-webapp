build:
	gcloud builds submit --tag gcr.io/kiya-webapp/kiya-webapp --project kiya-webapp

deploy:
	gcloud run deploy kiya-webapp --image gcr.io/kiya-webapp/kiya-webapp --platform managed --quiet --region=europe-west4 --project kiya-webapp
