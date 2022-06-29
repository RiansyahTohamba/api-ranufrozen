FROM golang:latest

LABEL maintener="Riansyah Tohamba <mriansyah93@gmail.com>"


WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD [ "./ranufrozen" ]

