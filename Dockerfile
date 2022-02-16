# Server
FROM golang:1.17-alpine as server

WORKDIR /app
COPY . .

RUN go build -o bin ./cmd/server.go

# FRONT
FROM node:16-alpine as front

WORKDIR /app

RUN npm -g i pnpm

COPY web/package.json web/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY web .
RUN pnpm run build

# RUnner
FROM alpine

WORKDIR /app
COPY --from=server /app/bin ./bin
COPY --from=front /app/build ./web/build

EXPOSE 8000

CMD ["./bin"]
