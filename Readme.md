# Library Management System (MVS)

## It is a basic Library Management System build on MVC architecture

## Running Instructions

1. `git clone` repository
2. in terminal go to the directory of cloned repository
3. configure a `config.yaml` file according to `sampleConfig.yaml`.
4. run `migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST)/$DB_NAME" -verbose up
` (Replacing all variables according to .yaml file)
5. If it shows error then first run `migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp($DB_HOST)/$DB_NAME" force <VERSION>
` then again run previous command.
6. run `make all` to setup virtual hosting.
7. run `make goSetup` to get all dependencies install
8. run `go script` to start the server
