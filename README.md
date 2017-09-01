# HTTP-Something-Wrong

Supported tags and respective `Dockerfile` links:  
・latest ([prod/Dockerfile](https://github.com/pottava/http-something-wrong/blob/master/prod/Dockerfile))

## Usage

### 1. Set environment variables

Environment Variables     | Description                                       |
------------------------- | ------------------------------------------------- |
ACCESS_LOG                | Send access logs to /dev/stdout. (default: true) | 
CONTENT_ENCODING          | Compress response data if the request allows. (default: true) |

### 2. Run the application

`$ docker run -d -p 9000:80 pottava/http-sw`

* with docker-compose.yml:  

```
errors:
  image: pottava/http-sw
  ports:
    - 9000:80
  environment:
    - ACCESS_LOG=false
  container_name: errors
```
