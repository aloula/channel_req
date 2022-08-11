### Parallel HTTP requests using Go channels

Basic routine to make parallel HTTP requests using Go channels

Builds for Linux, RPI and Windows can be found at **builds**

#### Usage:

`$ ./channel_req <url> <parallel requests> <print response>`

#### Examples:

Make 1 request printing the response:

`$ ./channel_req https://www.google.com 1 1`
```
Response: <!doctype html><html itemscope="" itemtype="http://schema.org/WebPage"...
# 1 -> Response Time: 261ms | Status Code: 200
```

Make 10 parallel requests without print the response:

`$ ./channel_req https://www.google.com 10 0`
```
# 1 -> Response Time: 253ms | Status Code: 200
# 2 -> Response Time: 256ms | Status Code: 200
# 3 -> Response Time: 259ms | Status Code: 200
...
```

