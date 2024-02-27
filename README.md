## Gin expense tracker api

This is a simple expense tracker api written in golang using gin web framework.

#### Features
- User authentication
- Expense item management
- Item category management


#### Setup

Clone this repository: 
```
git clone https://github.com/kipngeno-isaac/go-order-service.git
```

Configure environment variables:
```
cp .env.example .env
```

Run docker-compose: 
```
docker-compose up
```
Build app image:
```
docker build -f Dockerfile -t go-order-service:latest .
```

#### Running the Service
Start the service:
``` 
docker run -d -p 3000:8080 go-order-service:latest
```

Access the API using your preferred HTTP client.

#### API Endpoints
**Sign up**

Sample request
```
curl --location 'http://localhost:8080/sign-up' \
--header 'Content-Type: application/json' \
--data '{
    "FirstName": "John",
    "LastName": "Doer",
    "Email":"jdoe@mail.com",
    "Password":"secret"
}'
```
sample response
```
{
    "user": {
        "FirstName": "John",
        "LastName": "Doer",
        "Email": "jdoe@mail.com",
        "Password": "$2a$14$Cfs/.K9T2yPlEQ0wR6Sxm.DiofPDMWUyJs/rC6Jd/zMl1RSdx5aOy"
    }
}
```

**Login**

sample request
```
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "Email":"jdoe@mail.com",
    "Password":"secret"
}'
```
Sample response
```
{
    "user": {
        "FirstName": "John",
        "LastName": "Doer",
        "Email": "jdoe@mail.com",
        "Password": "$2a$14$Cfs/.K9T2yPlEQ0wR6Sxm.DiofPDMWUyJs/rC6Jd/zMl1RSdx5aOy"
    }
}
```