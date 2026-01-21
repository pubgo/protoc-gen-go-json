# Example Proto Files

This directory contains example proto files demonstrating the usage of `protoc-gen-go-json`.

## Files

- [user.proto](proto/example/user.proto) - User-related message definitions
- [order.proto](proto/example/order.proto) - Order-related message definitions with well-known types

## Generate Code

Using protobuild:

```bash
protobuild gen
```

Or using protoc directly:

```bash
protoc \
  --go_out=./gen \
  --go_opt=paths=source_relative \
  --go-json_out=./gen \
  --go-json_opt=paths=source_relative \
  -I./proto \
  ./proto/example/*.proto
```

## Generated Files

After generation, you will find:
- `gen/example/user.pb.go` - Go structs for user messages
- `gen/example/user.json.pb.go` - JSON marshaler/unmarshaler implementations
- `gen/example/order.pb.go` - Go structs for order messages
- `gen/example/order.json.pb.go` - JSON marshaler/unmarshaler implementations

## Usage Example

```go
package main

import (
    "encoding/json"
    "fmt"
    
    "github.com/pubgo/protoc-gen-go-json/example/gen/example"
)

func main() {
    user := &example.User{
        Id:    "123",
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   30,
        Role:  example.Role_ROLE_USER,
    }
    
    // Marshal to JSON using the generated MarshalJSON method
    data, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))
    // Output: {"id":"123","name":"John Doe","email":"john@example.com","age":30,"role":"ROLE_USER"}
    
    // Unmarshal from JSON using the generated UnmarshalJSON method
    var user2 example.User
    err = json.Unmarshal(data, &user2)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", user2)
}
```
