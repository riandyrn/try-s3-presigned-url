build:
	docker build -t try-s3-presigned-url .
run:
	make build
	docker run -it --rm \
		-e AWS_REGION=${AWS_REGION} \
		-e AWS_ACCESS_KEY=${AWS_ACCESS_KEY} \
		-e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
		--env-file env.list \
		-p 8765:8765 \
		try-s3-presigned-url