# run with the command "docker build ."
# builds the code and tests random inputs 10 seconds
FROM golang:1.18-alpine3.17
WORKDIR /app
COPY . ./
RUN go test ./... --fuzz=Fuzz --fuzztime=10s
