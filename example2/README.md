Example2
===

Upload files of 1 MB, 10 MB, 100 MB, 200 MB to S3 in parallel.

Usage:

```console
$ go run example2.go -b test-bucket -r us-east-1 -t 60
upload ../testdata/1M...
upload ../testdata/10M...
upload ../testdata/100M...
upload ../testdata/200M...
request 52A00C1826952483 took 1.116944884s to complete
successfully uploaded file to sdk-go-test/../testdata/1M
request 7CDC734EFE0F760D took 15.871513305s to complete
successfully uploaded file to sdk-go-test/../testdata/10M
request  took 1m0.003787591s to complete
request  took 1m0.003876892s to complete
upload canceled due to timeout, RequestCanceled: request context canceled
caused by: context deadline exceeded
```