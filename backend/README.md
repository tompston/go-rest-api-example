# backend

#### Init project structure

```bash
├── main.go             # holds the server
├── app                 # holds the default config for server + all routes
│── public              # optionally serve some of the generated files for easier accesibility for frontend people            
├── settings            # holds util functions for loading .env vars
│   └── database        # holds a function that creates db connection with .env vars
└── lib
    ├── response        # predefined response function and messages
    └── validate        # holds functions for validating payload structs
    └── utils           # holds placeholder function used inside the controllers
```

#### Update on change

```bash
# https://github.com/cespare/reflex
reflex -r '\.go' -s -- sh -c "go run main.go"
```

#### Docker commands

```bash
# NOTE. Dockerfile can be improved
# build the image
docker build -t backend .

# run and publish with the name of backend
docker run --publish 4444:4444 --name backend backend

# stop
docker stop backend

# remove 
docker image rm backend
```

#### Renaming generated sql files

```bash
# run from the root of the project
for i in ./db/sql/*.sql.gen.txt;  do mv "$i" "${i/.sql.gen.txt}.sql"; done
```