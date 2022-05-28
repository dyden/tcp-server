FROM golang:latest

EXPOSE 9000

WORKDIR /app

ENV HOST="localhost"
ENV PORT="9000"
ENV CONNECTION_TYPE="tcp"
ENV MAX_CONNECTIONS=1000        
ENV MAX_MESSAGES_HANDLERS=100         
ENV MAX_OCCURRENCES_HANDLERS=10          
ENV MAX_STATS_HANDLERS=10          


COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go


ENTRYPOINT [ "./app" ]

