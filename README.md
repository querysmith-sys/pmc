<!-- Folder Structure -->
pmc/
├── main.go              # entry point, just boots the app
├── cmd/
│   ├── root.go          # root cobra command setup
│   ├── get.go           # GET command
│   ├── post.go          # POST command
│   ├── put.go           # PUT command
│   ├── patch.go         # PATCH command
│   └── delete.go        # DELETE command
├── internal/
│   ├── request/
│   │   └── request.go   # builds and fires http requests
│   ├── response/
│   │   └── response.go   # formats and displays responses
│   ├── auth/
│   │   └── auth.go      # handles bearer, basic, cookie auth
│   ├── history/
│   │   └── history.go   # saves and retrieves request history
│   └── cookies/
│       └── jar.go       # cookie storage and handling
└── go.mod