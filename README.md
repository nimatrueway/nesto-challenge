# Readcommend

Readcommend is a book recommendation web app for the true book aficionados and disavowed human-size bookworms. It allows
to search for book recommendations with best ratings, based on different search criteria.

# Development environment

## Docker Desktop

Make sure you have the latest version of Docker Desktop installed, with sufficient memory allocated to it, otherwise you
might run into errors such as:

```
app_1         | Killed
app_1         | npm ERR! code ELIFECYCLE
app_1         | npm ERR! errno 137.
```

If that happens, first try running the command again, but if it doesn't help, try increasing the amount of memory
allocated to Docker in Preferences > Resources.

## Colima

If you are using Colima as a substitute for Docker Desktop, you might run into an issue
with [finding socket address](https://github.com/testcontainers/testcontainers-go/issues/2264#issuecomment-2131233614)
while running testcontainers-go on macOS. That can be worked around by running the following command:

```
sudo ln -sf $HOME/.colima/default/docker.sock /var/run/docker.sock
```

## Starting front-end app, back-end app and database

In this repo's root dir, run this command to start the front-end app (on port 8080), back-end app (on port 5001), and
PostgreSQL database (on port 5432):

```bash
$ docker-compose up --build
```

(later you can press Ctrl+C to stop this docker composition when you no longer need it)

Wait for everything to build and start properly.

## Connecting to database

During development, you can connect to and experiment with the PostgreSQL database by running this command:

```bash
$ ./psql.sh
```

Then, on the psql CLI, test as follows:

```psql
readcommend=# \dt
```

If everything went well, you should get this result:

```psql
    List of relations
 Schema |  Name  | Type  |  Owner
--------+--------+-------+----------
 public | author | table | postgres
 public | book   | table | postgres
 public | era    | table | postgres
 public | genre  | table | postgres
 public | size   | table | postgres
(5 rows)
```

To exit the PostgreSQL session, type `\q` and press `ENTER`.

## Accessing front-end app

Point your browser to http://localhost:8080

Be patient, the first time it might take up to 1 or 2 minutes for parcel to build and serve the front-end app.

You should see the front-end app appear, with all components displaying error messages because the back-end service does
not exist yet.

# Developing back-end service

change current directory to 'service' directory by running:

```
cd service
```

## Database schema migration

Skip this step if you already have the database schema set up. Database schema migration is implemented
using [goose](https://github.com/pressly/goose). Migration scripts are located in the `cmd/migrate/scripts` directory.
To run migration scripts, run the following command:

```
make run-db-migrations
```

## Running the service

The service runtime configuration is defined in the `config.yaml` file, through which you can configure the database
connection, server port, logging and other settings. The service can be run using the following:

```
make run
```

## Development tools and dependencies

To install development tools and dependencies, run the following:

```
make deps
```

## Running tests

```
make test
```

# Future works

- [ ] add more validation for non-existent authors or genres
- [ ] add open-telemetry and context logging
- [ ] add better error handling
- [ ] explain design decisions