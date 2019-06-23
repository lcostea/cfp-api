FROM alpine


WORKDIR /app
COPY bin/cfp-api /app/
ENTRYPOINT ["/app/cfp-api"]
