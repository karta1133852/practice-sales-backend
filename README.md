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