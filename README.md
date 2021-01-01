# HTTP-Something-Wrong

[![GitHub Actions](https://github.com/pottava/http-something-wrong/workflows/Publish%20artifacts/badge.svg?branch=master)](https://github.com/pottava/http-something-wrong/actions)

[gcr.io/pottava/errs](https://gcr.io/pottava/errs/)

Supported tags and respective `Dockerfile` links:  
・v1.1 ([Dockerfile](https://github.com/pottava/http-something-wrong/blob/master/Dockerfile))

## Usage

### 1. Set environment variables

Environment Variables     | Description                                       |
------------------------- | ------------------------------------------------- |
ACCESS_LOG                | Send access logs to /dev/stdout. (default: true) | 
CONTENT_ENCODING          | Compress response data if the request allows. (default: true) |

### 2. Run the application

```bash
$ docker run -d --rm -p 8080:8080 gcr.io/pottava/errs:v1.1
$ open http://localhost:8080
```
