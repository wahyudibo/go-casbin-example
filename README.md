# go-casbin-example

Simple proof-of-concept that implements casbin to governs route authorization. Built using go-casbin, gin, and mysql (gorm).

# Run it on your local

1. make sure you've got empty database and the database connection is ready.
2. copy `.envrc.example` and rename it to `.envrc`. Change the values as you see fit.
3. run `source .envrc`. The environment variables should be applied to your local session.
4. run the app using `go run cmd/go-casbin-example/main.go`. Check if the necessary tables are created in your database.
5. execute `permissions__users.sql` in `internal/database/mysql/seeder` to fill up the data in `permissions_users` table.
6. to check if the authorization is working, change the value of `userID` in `internal/middlewares/authz`. Only user `123` that has permissions to access protected endpoint.
7. check the endpoint by running `curl http://localhost:8080/authorized-only`. If success, the response will be
    ```
    {"message":"only authorized user is allowed here"}
    ```

