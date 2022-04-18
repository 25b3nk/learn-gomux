# learn-gomux
Learning Rest-API using gorialla mux and postgreSQL from [Semaphore CI gorilla mux tutorial](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)


## Getting started

1. Setup environment variables
    ```bash
    export APP_DB_USERNAME=postgres
    export APP_DB_PASSWORD=bhaskar
    export APP_DB_NAME=postgres
    ```

1. Run postgres docker
    ```bash
    docker run --rm --name "psql" -e POSTGRES_PASSWORD=bhaskar -it -p 5432:5432  postgres
    ```

1. Run tests
    ```bash
    go test -v
    ```
