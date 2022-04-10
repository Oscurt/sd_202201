FROM  golang:1.18.0-alpine3.15

WORKDIR /app

COPY . .

RUN go mod tidy \
&& go build -o init main.go

CMD [ "./init" ]