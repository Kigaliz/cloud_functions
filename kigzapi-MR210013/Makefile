.PHONY: gomodgen deploy delete

gomodgen:
	GO111MODULE=on go mod init

deploy:
	gcloud functions deploy fetchData --entry-point fetchData --runtime go111 --trigger-http

delete:
	gcloud functions delete fetchData --entry-point fetchData --runtime go111 --trigger-http
