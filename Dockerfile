FROM golang:latest

#RUN apt-get  make curl gcc 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD go test ./.. -v 
