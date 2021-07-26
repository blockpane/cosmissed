# 1st stage, build app
FROM golang:latest as builder
RUN apt-get update && apt-get -y upgrade
COPY . /build/app
WORKDIR /build/app

RUN go get ./... && go build -ldflags "-s -w" -o cosmissedd cmd/cosmissedd/main.go

# 2nd stage, create a user to copy, and install libraries needed if connecting to upstream TLS server
FROM debian:10 AS ssl
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get -y upgrade && apt-get install -y ca-certificates && \
    addgroup --gid 1317 --system cosmissed && adduser -uid 1317 --ingroup cosmissed --system --home /var/lib/cosmissed cosmissed

# 3rd and final stage, copy the minimum parts into a scratch container, is a smaller and more secure build.
FROM scratch
COPY --from=ssl /etc/ca-certificates /etc/ca-certificates
COPY --from=ssl /etc/ssl /etc/ssl
COPY --from=ssl /usr/share/ca-certificates /usr/share/ca-certificates
COPY --from=ssl /usr/lib /usr/lib
COPY --from=ssl /lib /lib
COPY --from=ssl /lib64 /lib64

COPY --from=ssl /etc/passwd /etc/passwd
COPY --from=ssl /etc/group /etc/group
COPY --from=ssl --chown=cosmissed:cosmissed /var/lib/cosmissed /var/lib/cosmissed

COPY --from=builder /build/app/cosmissedd /cosmissedd

EXPOSE 8080
USER cosmissed
WORKDIR /var/lib/cosmissed

ENTRYPOINT ["/cosmissedd"]
