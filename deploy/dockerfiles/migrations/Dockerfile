FROM dangdancheg/goose:0.0.2
ARG BUILD_DATE_TIME
ARG BUILD_GIT_HASH
ENV BUILD_DATE_TIME=${BUILD_DATE_TIME}
ENV BUILD_GIT_HASH=${BUILD_GIT_HASH}
COPY migrations /migrations
CMD goose -dir /migrations postgres "postgresql://${PG_USER}:${PG_PASS}@postgres:${PG_PORT}/${PG_DB}?sslmode=disable" up
