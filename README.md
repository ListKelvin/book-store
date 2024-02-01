## BOOK-STORE

This is the Book-Store BE

## Note

- Circular dependency issue ( Golang programs must be acyclic. In Golang cyclic imports are not allowed (That is its import graph must not contain any loops) )
- The `mapstructure` tag will be used to map the environment variables to the fields in the struct. Once that is done, the function will return the populated `Config` struct.
- The DB.AutoMigrate() function will automatically create the “users” table in the database if it doesn’t already exist.
- The GetGoogleOauthToken function will use the authorization code, client ID, and secret to retrieve an access token from the Google OAuth2 token endpoint
- GetGoogleUser function will use the access token to fetch the user’s account information.
- ...

issues:

- create authController

## Tech Stack
