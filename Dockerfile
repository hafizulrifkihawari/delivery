FROM golang as build

ENV GIN_MODE=release
ENV PORT=9004

WORKDIR /api

COPY . .

COPY go.mod go.sum ./
COPY .env ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -v .

FROM scratch

COPY --from=build /api/delivery /api/delivery

#uncomment below if .env needed
COPY --from=build /api/.env /api/.env

EXPOSE 9000
ENTRYPOINT ["/api/delivery"]

