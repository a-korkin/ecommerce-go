# eCommerce App

## Preparing
```console
docker compose up -d
make prepare
# seed mock data if needed
make seed_data
```
## Starting
```console
# start web interface
make run_web
# start kafka consumer
make run_consumer
# start grpc server
make run_grpc
```
## Testing
```console
make test
```
## Examples of requests
```http
### expect code 201 and info about new product
POST localhost:8080/products
Content-Type: application/json

{
    "title": "product#1",
    "category": "1c27e74d-e71c-4dd4-8832-e1b8b67b2d97",
    "price": 723.21
}

### expect list of products by pages
GET localhost:8080/products?page=1&limit=20

### expect code 200 and info about updated product
PUT localhost:8080/products/bb5bc9e0-b3b6-4e2f-bdbc-ba95a38c0bed 
Content-Type: application/json

{
    "title": "upd product",
    "category": "1c27e74d-e71c-4dd4-8832-e1b8b67b2d97",
    "price": 321.5
}

### expect category
GET localhost:8000/categories/1c27e74d-e71c-4dd4-8832-e1b8b67b2d97

### expect code 204
DELETE localhost:8080/users/e2c39673-bb0e-4b93-9f64-15e4fd4990ab 
```
