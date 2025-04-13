# Go Weather API Assessment

This is a Go API that fetches data from the National Weather Service.

### Set up and usage information

Assuming that Docker is already installed and the daemon is running, start the Docker container with:

- `docker build -t api .`
- `docker run -d -p 8080:8080 api`
- Or all together: `docker build -t api . && docker run -d -p 8080:8080 api`

This will run the API. Requests can then be made to it.

Here is a request example:

```
curl \
-X "GET" \
"http://localhost:8080/weather-today?latitude=30.2667&longitude=-97.7500"
```

Visit http://localhost:8080/swagger/index.html for more detailed API documentation.

To see the Docker container output:

- Find the container ID with `docker ps`
- `docker logs <container_id>`
- Or all together: `docker logs $(docker ps -n 1 -a -q)`

Stop the container with `docker stop $(docker ps -n 1 -a -q)`

### National Weather Service API documentation

- https://www.weather.gov/documentation/services-web-api
- https://api.weather.gov/openapi.json
