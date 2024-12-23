FROM 1.23-bookworm AS build

WORKDIR /app

COPY . .

RUN go mod download

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/go-hexagonal-sample-exam ./cmd/http/main.go

# final stage
FROM alpine:latest AS final
LABEL maintainer="andy82115"

# set working directory
WORKDIR /app

# copy binary
COPY --from=build /app/bin/go-hexagonal-sample-exam ./

EXPOSE 8080

ENTRYPOINT [ "./go-hexagonal-sample-exam" ]
