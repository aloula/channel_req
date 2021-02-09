### Parallel HTTP requests using Go channels

Basic routine to make parallel HTTP requests using Go channels

Builds for Linux, RPI and Windows can be found at **builds**

#### Usage

> `$ ./channel_req <url> <parallel requests> <print response>`

#### Examples:

+ Make 1 request printing the response:

> `$ ./channel_req https://www.google.com 1 1`

+ Make 100 parallel requests without print the response:

> `$ ./channel_req https://www.google.com 100 0`