## Wait a minute!

- Make sure you have go installed
- Make sure to have docker properly installed and run
- Make sure you have golang-migrate installed

<br/>

## How to run

On the server directory, run:

```
make postgresinit
```

If nothing goes wrong, proceed to create database, run:

```
make createdb
```

To make sure database has been created, run:

```
# Go inside container
make postgres

# Check list of database
\l
```

If you can see go-realtimechat database, you're ready for the next step, migrating tables. Please make sure that you have already installed and configured golang-migrate CLI

```
migrate -path migrations -database "postgresql://root:root@localhost:1234/go-realtimechat?sslmode=disable" -verbose up
```

If you want to check if table is already created, run:

```
# Go inside the container
make postgres

# Go inside the go-realtimechat database
\c go-realtimechat

# Check all tables
\dt

# To describe users table
\d users

# To check users table data
SELECT * FROM users;

# To quit
\q
```

Now, it's time to run the app

```
go run main.go
```
