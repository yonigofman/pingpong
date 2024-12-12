# PingPong

PingPong is a lightweight Go package that provides a simple data structure to enable back-and-forth (ping-pong) traversal over a slice of elements. The traversal alternates direction when the end or start of the slice is reached.

## Features

- **Ping-pong traversal**: Automatically switches traversal direction at the start and end of the slice.
- **Lightweight**: Minimal dependencies and easy-to-use interface.
- **Well-tested**: Includes unit tests to ensure stability.
- **Customizable**: Works with any slice of strings.

## Installation

To use PingPong in your Go project, install it using `go get`:

```bash
go get github.com/yonigofman/pingpong
```

Then import it into your project:

```go
import "github.com/yonigofman/pingpong"
```

## Usage

Here is an example of how to use the PingPong package:

```go
package main

import (
    "fmt"
    "github.com/yonigofman/pingpong"
)

func main() {
    ping, err := pingpong.NewPingPong([]string{"a", "b", "c"})
    
    if err != nil {
		fmt.Println(err)
		return
	}

    fmt.Println(ping.Next()) // Output: a
    fmt.Println(ping.Next()) // Output: b
    fmt.Println(ping.Next()) // Output: c
    fmt.Println(ping.Next()) // Output: b
    fmt.Println(ping.Next()) // Output: a
    fmt.Println(ping.Next()) // Output: b
}
```

## API Reference

### `NewPingPong(items []string) *PingPong`
Creates a new PingPong instance.

- **Parameters**:
  - `items []string`: A slice of strings to traverse.
- **Returns**:
  - `*PingPong`: A pointer to the PingPong instance.

### `func (p *PingPong) Next() string`
Returns the next element in the ping-pong traversal.

- **Returns**:
  - `string`: The current element in the traversal.

## Development

### Prerequisites

- Go 1.19 or later

### Building and Testing

To build the project:

```bash
go build
```

To run tests:

```bash
go test ./...
```

### Linting

Use `golangci-lint` for linting:

```bash
golangci-lint run ./...
```

## Continuous Integration

The project includes a GitHub Actions workflow for CI/CD:

- **Tests**: Runs `go test` on multiple Go versions.
- **Linting**: Ensures the code adheres to Go best practices.
- **Releases**: Automatically creates a release on `main` branch updates.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Thanks to the Go community for inspiration and support.

