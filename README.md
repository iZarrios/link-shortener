# link-shortener

## Setting up psql with docker

```bash
docker run --name auth-psql -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:14 -d <CONTAINER_NAME>
```
### Finding the local ip address of the container
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <CONTAINER_NAME>
```
