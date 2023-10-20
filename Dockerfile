FROM golang:1.21

WORKDIR /
COPY . .
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/swaggo/http-swagger
RUN go get -d github.com/swaggo/swag
RUN go get -d github.com/swaggo/files
RUN go get -d golang.org/x/tools
RUN go get -d github.com/KyleBanks/depth
RUN go get -d github.com/go-openapi/spec
RUN go get -d github.com/go-openapi/jsonpointer
RUN go get -d github.com/go-openapi/jsonreference
RUN go get -d github.com/go-openapi/swag
RUN go get -d golang.org/x/net
RUN go get -d gopkg.in/yaml.v3
RUN go get -d github.com/mailru/easyjson
RUN go get -d golang.org/x/sys
RUN go get -d github.com/josharian/intern
CMD ["go","run","main.go"]