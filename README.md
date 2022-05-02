# User test data
## To login, use these two accounts below:
```JSON
{
  "username": "user01",
  "password": "4b8f353889d9a05d17946e26d014efe99407cba8bd9d0102d4aab10ce6229043"
  // original password: "password01"
}
```
```JSON
{
  "username": "user02",
  "password": "08f0d4cb02352f2f7fd251fbbe1c9aa5fd176bb0c7f1bd35e4f71a8dcb820852"
  // original password: "password02"
}
```

# API Routes
## Authentication
- Login: ```POST /api/auth/login```
  - Request body
    ```JSON
    {
      "username": "usr",
      "password": "pwd"
    }
    ```
  - Response
    ```JSON
    {
      "uid": 1,
      "token": "jwtToken"
    }
    ```
## Users
- Get user's data: ```GET /api/users/:uid```
  - Response
    ```JSON
    {
      "uid": 1,
      "username": "user01",
      "coin": 655,
      "point": 470,
      "vipType": "Normal",
      "accumulatedSpent": 345
    }
    ```
- Create new user: ```POST /api/users```
  - Request body
    ```JSON
    {
      "username": "usr",
      "password": "pwd"
    }
    ```
  - Response
    ```JSON
    {
      "uid": 1
    }
    ```
- Get user's orders list: ```GET /api/users/:uid/orders```
  - Response
    ```JSON
    [
      {
        "orderId": 1,
        "time": "2022-04-30T16:44:38Z"
      },
      {
        "orderId": 3,
        "time": "2022-05-01T09:57:35Z"
      },
    ]
    ```
- Create new order: ```POST /api/users/:uid/orders```
  - Request body
    ```JSON
    {
      "originalPrice": 150,
      "payedCoin": 115,
      "payedPoint": 10,
      "exchange": 200,  // point's exchange rate (%)
      "discount": 10,   // discount off (%)
      "products": [
        {
          "productNo": 2,
          "quantity": 5
        },
        {
          "productNo": 3,
          "quantity": 1
        }
      ]
    }
    ```
  - Response
    ```JSON
    {
      "uid": 1,
      "coin": 655,
      "point": 470,
      "accumulatedSpent": 345,
      "orderId": 4
    }
    ```
  
# Build
## Local
- Run
  ```bash
  make build run
  ```

## Docker Compose
- Build
  ```bash
  make docker-compose
  ```

- Run
  ```bash
  make run-docker-compose
  ```

## Docker
- Build
  ```bash
  # make clean-docker if needed
  make [clean-docker] build-docker
  ```

- Run
  ```bash
  make run-docker
  ```