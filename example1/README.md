Example1
===

Upload files in order of 1 MB, 10 MB, 100 MB, 200 MB to S3.

Usage:

```console
$ go run example1.go -b test-bucket -r us-east-1 -t 60
upload ../testdata/1M...
request 1DE404D8F72EFE24 took 1.046535818s to complete
successfully uploaded file to sdk-go-test/../testdata/1M
upload ../testdata/10M...
request 49B33B89D8D51BC9 took 8.532994279s to complete
successfully uploaded file to sdk-go-test/../testdata/10M
upload ../testdata/100M...
request  took 1m0.001056708s to complete
upload canceled due to timeout, RequestCanceled: request context canceled
caused by: context deadline exceeded
upload ../testdata/200M...
request  took 1m0.003835726s to complete
upload canceled due to timeout, RequestCanceled: request context canceled 
caused by: context deadline exceeded
```