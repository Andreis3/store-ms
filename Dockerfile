##################################
# STEP 1 build executable binary #
##################################
FROM golang:1.22.1-alpine as builder

WORKDIR /app
COPY . .

# Fetch dependencies.
RUN go mod download

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -o /go/bin/server cmd/server/main.go

##############################
# STEP 2 time zone           #
##############################
FROM alpine:latest as time-zone
RUN apk --no-cache add tzdata zip
WORKDIR /usr/share/zoneinfo
# -0 means no compression.  Needed because go's
# tz loader doesn't handle compressed data.
RUN zip -q -r -0 /zoneinfo.zip .

##############################
# STEP 3 build a small image #
##############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/server /go/bin/server

# Copy the zoneinfo.zip file from the time-zone stage.
ENV ZONEINFO /zoneinfo.zip
COPY --from=time-zone /zoneinfo.zip /

EXPOSE 8080

ENTRYPOINT ["/go/bin/server"]