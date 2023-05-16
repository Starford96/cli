build:
	docker build -f Dockerfile.build -t starford96/buffalo:latest .

push:
	docker push starford96/buffalo:latest