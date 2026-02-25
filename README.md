<!-- Folder Structure -->
```pmc/
├── main.go              # Entry point, boots the CLI app
├── cmd/
│   ├── root.go          # Root Cobra command setup
│   ├── get.go           # GET command
│   ├── post.go          # POST command
│   ├── put.go           # PUT command
│   ├── patch.go         # PATCH command
│   └── delete.go        # DELETE command
├── internal/
│   ├── request/
│   │   └── request.go   # Builds and executes HTTP requests
│   ├── response/
│   │   └── response.go  # Formats and prints HTTP responses
│   ├── auth/
│   │   └── auth.go      # Handles Bearer, Basic & Cookie auth
│   ├── history/
│   │   └── history.go   # Stores and retrieves request history
│   └── cookies/
│       └── jar.go       # Cookie storage and management
└── go.mod
```
