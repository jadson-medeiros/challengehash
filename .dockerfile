FROM postgres
ENV POSTGRES_PASSWORD 123
ENV POSTGRES_DB challengehash
COPY world.sql /docker-entrypoint-initdb.d/