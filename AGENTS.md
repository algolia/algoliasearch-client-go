# GO CLIENT - AI AGENT INSTRUCTIONS

## ⚠️ CRITICAL: CHECK YOUR REPOSITORY FIRST

Before making ANY changes, verify you're in the correct repository:

```bash
git remote -v
```

- ✅ **CORRECT**: `origin .../algolia/api-clients-automation.git` → You may proceed
- ❌ **WRONG**: `origin .../algolia/algoliasearch-client-go.git` → STOP! This is the PUBLIC repository

**If you're in `algoliasearch-client-go`**: Do NOT make changes here. All changes must go through `api-clients-automation`. PRs and commits made directly to the public repo will be discarded on next release.

## ⚠️ BEFORE ANY EDIT: Check If File Is Generated

Before editing ANY file, verify it's hand-written by checking `config/generation.config.mjs`:

```javascript
// In generation.config.mjs - patterns WITHOUT '!' are GENERATED (do not edit)
'clients/algoliasearch-client-go/algolia/**',                   // Generated
'!clients/algoliasearch-client-go/algolia/transport/**',        // Hand-written ✓
'!clients/algoliasearch-client-go/algolia/errs/**',             // Hand-written ✓
```

**Hand-written (safe to edit):**

- `algolia/transport/**` - HTTP transport, retry, requester
- `algolia/errs/**` - Error types
- `algolia/call/**` - Call type definitions
- `algolia/compression/**` - Gzip compression
- `algolia/debug/**` - Debug utilities
- `algolia/utils/**` - Utility functions

**Generated (DO NOT EDIT):**

- `algolia/search/**` - Search client and models
- `algolia/insights/**`, `algolia/recommend/**`, etc. - Other API clients
- Root config files

## Language Conventions

### Naming

- **Files**: `snake_case.go`
- **Packages**: `lowercase` single word
- **Variables/Functions**: `camelCase` (unexported), `PascalCase` (exported)
- **Types/Interfaces**: `PascalCase`
- **Constants**: `PascalCase` (exported), `camelCase` (unexported)

### Formatting

- `gofmt` / `goimports` standard formatting
- Run: `yarn cli format go clients/algoliasearch-client-go`

### Go Idioms

- Accept interfaces, return structs
- Errors as values, not exceptions
- Use `context.Context` for cancellation
- Prefer composition over inheritance

### Dependencies

- **HTTP**: Standard `net/http` with custom transport
- **JSON**: Standard `encoding/json`
- **Build**: Go modules
- **Min version**: Go 1.21

## Client Patterns

### Transport Architecture

```go
// Core transport in algolia/transport/
type Transport struct {
    requester      Requester
    retryStrategy  *RetryStrategy
    compression    compression.Compression
    connectTimeout time.Duration
}

func New(cfg Configuration) *Transport {
    // Configurable requester, timeouts, compression
}
```

### Requester Interface

```go
type Requester interface {
    Request(req *http.Request) (*http.Response, error)
}

// Default uses http.Client, but injectable for testing
```

### Retry Strategy

- Host states: `isUp`, `isTimedOut`
- Automatic host failover on failure
- Configurable read/write timeouts
- Retries on network errors, not on 4xx

### Error Handling

```go
// algolia/errs/
var (
    ErrUnreachableHosts = errors.New("all hosts are unreachable")
)

type APIError struct {
    Status  int
    Message string
}
```

## Common Gotchas

### Context Usage

```go
// Always pass context for cancellation
ctx := context.Background()
response, err := client.Search(ctx, params)

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
response, err := client.Search(ctx, params)
```

### Error Checking

```go
// ALWAYS check errors
response, err := client.Search(ctx, params)
if err != nil {
    // Handle error - don't ignore!
    var apiErr *errs.APIError
    if errors.As(err, &apiErr) {
        // Handle API error
    }
    return err
}
```

### Pointer vs Value

```go
// Use pointers for optional fields in structs
type SearchParams struct {
    Query          string  // Required
    HitsPerPage    *int    // Optional - use utils.ToPtr(10)
}

// Helper in algolia/utils/
import "github.com/algolia/algoliasearch-client-go/v4/algolia/utils"
params := SearchParams{
    HitsPerPage: utils.ToPtr(10),
}
```

### JSON Marshaling

```go
// Models use json tags
type Hit struct {
    ObjectID string `json:"objectID"`
}

// Custom marshaling in model files - don't modify generated code
```

### Module Path

```go
// Import path includes version
import "github.com/algolia/algoliasearch-client-go/v4/algolia/search"
```

## Build & Test Commands

```bash
# From repo root (api-clients-automation)
yarn cli build clients go                      # Build Go client
yarn cli cts generate go                       # Generate CTS tests
yarn cli cts run go                            # Run CTS tests
yarn cli playground go search                  # Interactive playground
yarn cli format go clients/algoliasearch-client-go

# From client directory
cd clients/algoliasearch-client-go
go build ./...                                 # Build all packages
go test ./...                                  # Run tests
golangci-lint run                              # Run linter
```
