FROM postgres:16-alpine

# Environment variable for container, default user is postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=app_db


# Init the data with scripts
COPY ./postgres-script/*.sql /docker-entrypoint-initdb.d/

#Expose the PostgreSQL default port
EXPOSE 5432
