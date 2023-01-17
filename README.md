## Full-Stack app for creating transactions

https://user-images.githubusercontent.com/82293948/213023194-1b2dcc06-863b-4dfc-ad4b-f743eaaba685.mp4

- Backend
  - Golang, Chi Framework 
  - JWT for authentication
  - Cursor Pagination for Users and Transactions endpoints (using `uuids` and `created_at` values)
  - SQLC for getting the data, dbmate for managing migrations.
  - If you want to preview the backend endpoints, just copy `gomarvin.json` content in the [editor](https://gomarvin.pages.dev/) (Settings -> Import Tab)
- Frontend
  - Vue 3 + Vite + Tailwind.
  - Home view is guarded by authentication check. If the user has an invalid token, that route is not accessible.
- Other
  - Deno and Faker used for seeding the database (using the generated gomarvin client
  - Backend code baseline and frontend fetch functions generated with gomarvin.

### DISCLAIMER

A lot of parts are rough around the edges.

- JWT Auth flow in the frontend is lacking
  - `access_token` expiration is 15mins, no flow for re-authentication
- DB tables can be improved
- There are endpoints which don't execute any queries.
- Frontend is as minimal as possible
  - No loading/error states

### Setup

- Use the db dump to create the db schemas and rows.
- Edit `/backend/.env` if needed

```bash
# run backend
cd backend
go mod tidy
go mod download
go run main.go

# run frontend
cd frontend
npm i
npm run dev
```

#### Other

```bash
# seeder (using deno with faker)
deno run --allow-net ./seeder.ts

# dump database info for test_db
sudo pg_dump -U postgres test_db > ./dump.sql
```
