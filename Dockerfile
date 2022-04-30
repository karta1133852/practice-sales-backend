# Base image
FROM golang

# Create app directory
RUN     mkdir -p /app
WORKDIR /app

# Bundle app source
COPY . .

# Install app dependencies
RUN go mod tidy
# Compile
RUN GOARCH=amd64 GOOS=linux go build -o bin/app main.go

# Listening on port 3000
EXPOSE 3000

ENTRYPOINT ["./bin/app"]