1. Install GoLang from official site

2. Check your personal path on terminal
   `$(go env GOPATH)/bin`

3. Insert two line in bashrc

```
vi ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bi
```

4. Run the command to create mocks
   `mockgen -destination=application/mocks/application.go -source=application/product.go application/`

#### If you prefer don't install golang you can run it on docker

`docker-compose up -d`
`docker-compose ps`
`docker exec -it appproduct`
`docker-compose down`

curl -X POST http://localhost:9000/product -H 'Content-Type: application/json' -d '{"name":"fone ble - motorola", "price": 197.90}'

curl -X GET http://localhost:9000/product/{id}

curl -X PUT http://localhost:9000/product/81ecd763-97c9-4669-981d-ae46a660cd1f/enable

curl -X PUT \
 http://localhost:9000/product/81ecd763-97c9-4669-981d-ae46a660cd1f/disable \
-d '{"ID":"81ecd763-97c9-4669-981d-ae46a660cd1f","Name":"fone ble - motorola","Price":197.9,"Status":"enable"}'

#### Example:

- Creating one register
  `curl -X POST http://localhost:9000/product -H 'Content-Type: application/json' -d '{"name":"smart watch", "price": 1557.54}'`
  `{"ID":"70763b80-532d-431e-82b5-9ba1faee64ea","Name":"smart watch","Price":1557.54,"Status":"disabled"}`

- Getting the last one
  `curl -X GET http://localhost:9000/product/70763b80-532d-431e-82b5-9ba1faee64ea`
  `{"ID":"70763b80-532d-431e-82b5-9ba1faee64ea","Name":"smart watch","Price":1557.54,"Status":"disabled"}`

- Change status to enabled
  `curl -X PUT http://localhost:9000/product/70763b80-532d-431e-82b5-9ba1faee64ea/enable`
  `{"ID":"70763b80-532d-431e-82b5-9ba1faee64ea","Name":"smart watch","Price":1557.54,"Status":"enable"}`

- Checking the last update
  `curl -X GET http://localhost:9000/product/70763b80-532d-431e-82b5-9ba1faee64ea`
  `{"ID":"70763b80-532d-431e-82b5-9ba1faee64ea","Name":"smart watch","Price":1557.54,"Status":"enable"}`
