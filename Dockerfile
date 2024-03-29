########################################
FROM golang:1.22.1 as base

RUN apt update && \
    apt install bc # make git # golangci-lint

########################################
#FROM base as builder

########################################
FROM base as runner

#COPY --from=builder /usr/local/go /usr/local/go
#COPY --from=builder /go /go

#ENV PATH /go/bin:/usr/local/go/bin:$PATH
#ENV GOPATH /go

WORKDIR /app