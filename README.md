# Try S3 Presigned URL

This is experiment to use S3 Presigned URL for uploading file. Built using Go & Docker.

To run the server:

```bash
make run
```

Make sure to set `AWS_REGION`, `AWS_ACCESS_KEY`, `AWS_SECRET_ACCESS_KEY` on env before running the server.

Also make sure the used IAM user at least has `s3:PutObject` permission.
