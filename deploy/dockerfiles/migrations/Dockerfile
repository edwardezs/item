FROM dangdancheg/goose:0.0.2
COPY migrations /migrations
CMD goose -dir /migrations postgres "postgresql://${PG_USER}:${PG_PASS}@postgres:${PG_PORT}/${PG_DB}?sslmode=disable" up
