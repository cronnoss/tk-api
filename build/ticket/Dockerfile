# Собираем в гошке
FROM golang:1.22 as build

ENV BIN_FILE /opt/ticket/ticket-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# We build a static Go binary (without dependencies on the C API),
# otherwise it will not work in the Alpine image.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/ticket/*

# The output is a subtle image.
FROM alpine:3.20

LABEL ORGANIZATION="Ticket project"
LABEL SERVICE="ticket"
LABEL MAINTAINERS="ticket@ticket.com"

ENV BIN_FILE "/opt/ticket/ticket-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/ticket/config.toml
COPY ./build/ticket/config.toml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}
