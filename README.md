# minio-logger
Logger webhook for MIN.IO deployment in Go-Lang

Add .env

    HOST=0.0.0.0:5656    
    LOG_FILE=minio.log

Sample cURL:

    curl --location --request POST 'localhost:5656/log' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "version": "1",
      "deploymentid": "bc0e4d1e-bacc-42eb-91ad-2d7f3eacfa8d",
      "time": "2019-08-12T21:34:37.187817748Z",
      "api": {
        "name": "PutObject",
        "bucket": "testbucket",
        "object": "hosts",
        "status": "OK",
        "statusCode": 200,
        "timeToFirstByte": "0s",
        "timeToResponse": "2.143308ms"
      },
      "remotehost": "127.0.0.1",
      "requestID": "15BA4A72C0C70AFC",
      "userAgent": "MinIO (linux; amd64) minio-go/v6.0.32 mc/2019-08-12T18:27:13Z",
      "requestHeader": {
        "Authorization": "AWS4-HMAC-SHA256 Credential=minio/20190812/us-east-1/s3/aws4_request,SignedHeaders=host;x-amz-content-sha256;x-amz-date;x-amz-decoded-content-length,Signature=d3f02a6aeddeb29b06e1773b6a8422112890981269f2463a26f307b60423177c",
        "Content-Length": "686",
        "Content-Type": "application/octet-stream",
        "User-Agent": "MinIO (linux; amd64) minio-go/v6.0.32 mc/2019-08-12T18:27:13Z",
        "X-Amz-Content-Sha256": "STREAMING-AWS4-HMAC-SHA256-PAYLOAD",
        "X-Amz-Date": "20190812T213437Z",
        "X-Amz-Decoded-Content-Length": "512"
      },
      "responseHeader": {
        "Accept-Ranges": "bytes",
        "Content-Length": "0",
        "Content-Security-Policy": "block-all-mixed-content",
        "ETag": "a414c889dc276457bd7175f974332cb0-1",
        "Server": "MinIO/DEVELOPMENT.2019-08-12T21-28-07Z",
        "Vary": "Origin",
        "X-Amz-Request-Id": "15BA4A72C0C70AFC",
        "X-Xss-Protection": "1; mode=block"
      }
    }'