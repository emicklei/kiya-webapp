deploy:
	cd go-app && gcloud app deploy

inuse:
	gcloud compute addresses list --global

clean:
	gcloud app versions list | grep -v SERVING | awk '{print $2}' | tail -n +1 | xargs -I {} gcloud app versions delete {}