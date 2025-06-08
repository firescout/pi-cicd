Repo Manager is a lightweight Go web service for managing repositories and triggering CI/CD pipelines via HTTP webhooks. It is designed to run on Raspberry Pi OS and can be integrated into automation workflows.

## Features

- Exposes HTTP endpoints for repository management.
- Handles push events via `/repo/push?repo=RepoName`.
- Provides a shutdown endpoint.
- Uses a simple JSON mapping for repositories.
- Logs all HTTP requests.
- Easily extensible and suitable for CI/CD pipelines.

## Project Structure

```
.gitignore
go.mod
go.sum
main.go
test.http
common/
    repomap.json
handler/
    handler.go
restserver/
    api_default.go
    api.go
    error.go
    helpers.go
    impl.go
    logger.go
    routers.go
service/
    service.go
```

- **main.go**: Entry point, starts the service.
- **service/**: Service lifecycle management.
- **handler/**: Business logic for HTTP endpoints.
- **restserver/**: HTTP server, routing, and helpers.
- **repomap.json**: Repository mapping configuration.

## Endpoints

- `GET /repo/push?repo=RepoName`  
  Triggers a push event for the specified repository.

- `GET /shutdown`  
  Initiates a shutdown of the service.

## Usage
1. **Setup Json**
```json
{
    "system": "windows",
    "clone_path": "C:\\path\\to\\clone",
    "repos": [
        {
            "name": "REPONAME",
            "url": "https://github.com/example/REPONAME",
            "after_script": [
                {
                    "command": "rmdir",
                    "args": [
                        "/var/www/REPONAME"
                    ]
                }
            ],
            "path": "/var/www/REPONAME"
        }
    ]
}
```

2. **Build and Run Locally**

   ```sh
   go build -o repo-manager
   ./repo-manager
   ```

3. **Test Endpoints**

   Use test.http or curl:

   ```sh
   curl "http://localhost:9090/repo/push?repo=RepoName"
   ```

## GitHub Actions CI/CD

This app can be built and tested automatically using GitHub Actions. Below is an example workflow:

````yaml
name: Go CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Build
      run: go build -v ./...

    - name: Run Tests
      run: go test -v ./...
````

- On every push or pull request to `main`, this workflow will:
  - Check out the code.
  - Set up Go.
  - Build the project.
  - Run tests (if any are present).

## Configuration

- Edit repomap.json to add or modify repositories.

## Logging

All HTTP requests are logged with method, URI, handler name, and duration.

## License

MIT

---

For more details, see the code in main.go, service.go, and handler.go.

Similar code found with 2 license types
