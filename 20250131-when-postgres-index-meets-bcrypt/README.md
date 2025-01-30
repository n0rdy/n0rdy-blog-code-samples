A code to the blog post [When Postgres index meets Bcrypt](https://n0rdy.foo/posts/20250131/when-postgres-index-meets-bcrypt/)

## How to run

- Run PostgreSQL database via Docker compose:

```bash
docker compose up -d
```
- Run the code from the `main.go` file to start the app. This will trigger the database schema creation and data insertion, so, please, be patient, as it will take some time.
- Run the code from the `thirdpartyapi/main.go` file to start the third-party API.
- Have fun! =)