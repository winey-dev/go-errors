# go-errors

`go-errors` is a Go package designed to simplify error handling and enhance error reporting in Go applications. It provides structured error types, stack trace support, and integration with common libraries like SQL drivers, HTTP, and gRPC.

## Features

- **Structured Errors**: Create errors with custom codes, messages, and additional context.
- **Error Formatting**: Flexible formatting options for error messages, including stack traces and file/line information.
- **Error Wrapping**: Wrap errors with additional context while preserving the original error.
- **Integration**:
  - SQL error handling for MySQL and PostgreSQL.
  - HTTP error formatting for JSON responses.
  - gRPC error conversion with detailed messages.

## Installation

To install the package, use `go get`:

```bash
go get github.com/winey-dev/go-errors
```

## Usage

### Creating Errors

You can create errors with custom messages or codes:

```go
import "github.com/winey-dev/go-errors"

err := errors.New("something went wrong")
errWithCode := errors.Errorc(errors.Internal, "internal server error")
```

### Wrapping Errors

Wrap existing errors with additional context:

```go
wrappedErr := errors.Wrap(err, "failed to process request")
wrappedErrWithCode := errors.Wrapc(err, errors.InvalidRequest, "invalid input provided")
```

### SQL Error Handling

`SQLError` is designed to parse SQL errors provided by solutions like MySQL or PostgreSQL and convert them into your custom error structure. This method is meaningful only when your application uses the custom error structure provided by this package. For example:

```go
import (
	"database/sql"
	"github.com/winey-dev/go-errors"
)

func QueryUser() error {
	err := sql.ErrNoRows // Simulating a SQL error
	return errors.SQLError(err)
}

func main() {
	err := QueryUser()
	if err != nil {
		log.Printf("SQL Error: %+v", err)
	}
}
```

**Output Example:**
```
[NotFound] sql: no rows in result set (/path/to/file.go:10)
```

### HTTP Error Formatting

`JSONHTTPErrorHandle` is used to convert your custom error structure into an HTTP-friendly JSON response. This method is meaningful only when your application uses the custom error structure provided by this package. For example:

```go
import (
	"net/http"
	"github.com/winey-dev/go-errors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := errors.Errorc(errors.NotFound, "user not found")
	errors.JSONHTTPErrorHandle(w, err)
}
```

**Output Example:**
```json
{
	"err_code": "NotFound",
	"message": "user not found",
	"detailed": []
}
```

### gRPC Error Conversion

`GRPCError` maps your custom error structure to gRPC-compatible error responses. This method is meaningful only when your application uses the custom error structure provided by this package. For example:

```go
import (
	"github.com/winey-dev/go-errors"
	"google.golang.org/grpc/status"
)

func GetUser() error {
	err := errors.Errorc(errors.NotFound, "user not found")
	return errors.GRPCError(err)
}

func main() {
	err := GetUser()
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		log.Printf("gRPC Error: %s", grpcStatus.Message())
	}
}
```

**Output Example:**
```
gRPC Error: user not found
```

### Error Formatting

To debug errors effectively, you can use flexible formatting options. The package supports the following formats:

1. **Default Error Message (`%v`)**:
   ```go
   log.Printf("%v", err)
   ```
   **Output Example:**
   ```
   [NotFound] user not found
   ```

2. **Error with File and Line Information (`%#v`)**:
   ```go
   log.Printf("%#v", err)
   ```
   **Output Example:**
   ```
   [NotFound] user not found (/path/to/file.go:10)
   ```

3. **Error with Stack Trace (`%+v`)**:
   ```go
   log.Printf("%+v", err)
   ```
   **Output Example:**
   ```
   [NotFound] user not found
   	-> [Internal] database connection failed
   ```

4. **Error with Stack Trace and File/Line Information (`%+#v`)**:
   ```go
   log.Printf("%+#v", err)
   ```
   **Output Example:**
   ```
   [NotFound] user not found (/path/to/file.go:10)
   	-> [Internal] database connection failed (/path/to/file.go:5)
   ```

### Explanation of Formatting Options

- **`+`**: Includes the stack trace of the error.
- **`#`**: Includes the file name and line number where the error occurred.
- **Combination (`%+#v`)**: Provides both the stack trace and file/line information for detailed debugging.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
