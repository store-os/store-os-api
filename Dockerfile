ARG ARCH=

ARG TARGETOS
ARG TARGETARCH

RUN echo ${TARGETOS} 
RUN echo ${TARGETARCH}

FROM ${ARCH}golang:1.15-alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH}

WORKDIR /build
COPY go.mod .
COPY go.sum . 
RUN go mod download
COPY . .
RUN go build -o main main.go
RUN ls -ltr
WORKDIR /dist 

RUN cp /build/main .
EXPOSE 8080 

CMD ["/dist/main"]
