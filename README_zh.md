# protoc-gen-go-json

[![Go Reference](https://pkg.go.dev/badge/github.com/pubgo/protoc-gen-go-json.svg)](https://pkg.go.dev/github.com/pubgo/protoc-gen-go-json)
[![CI](https://github.com/pubgo/protoc-gen-go-json/actions/workflows/ci.yml/badge.svg)](https://github.com/pubgo/protoc-gen-go-json/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pubgo/protoc-gen-go-json)](https://goreportcard.com/report/github.com/pubgo/protoc-gen-go-json)

[English](README.md)

一个 protobuf 编译器插件，为 Go 语言的 protobuf 消息自动生成 `json.Marshaler` 和 `json.Unmarshaler` 接口实现。

## 功能特性

- 🚀 **轻量级** - 生成的代码简洁高效
- 🔄 **递归支持** - 自动处理嵌套消息类型
- ✅ **标准兼容** - 使用官方 `protojson` 包，保证兼容性
- ⚙️ **可配置** - 支持多种自定义选项
- 📦 **Proto3 Optional** - 支持 proto3 的 optional 特性

## 安装

```bash
go install github.com/pubgo/protoc-gen-go-json@latest
```

## 使用方法

### 基本用法

```bash
protoc --go_out=. --go-json_out=. your_proto_file.proto
```

### 配合 buf 使用

在 `buf.gen.yaml` 中添加：

```yaml
version: v2
plugins:
  - local: protoc-gen-go
    out: gen
    opt: paths=source_relative
  - local: protoc-gen-go-json
    out: gen
    opt:
      - paths=source_relative
```

## 命令行参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `enums_as_ints` | `false` | 将枚举渲染为整数而非字符串 |
| `emit_defaults` | `false` | 渲染零值字段 |
| `orig_name` | `false` | 使用原始 `.proto` 文件中的字段名 |
| `allow_unknown` | `false` | 反序列化时允许未知字段 |
| `debug` | `false` | 启用调试模式 |

### 使用示例

```bash
protoc --go-json_out=emit_defaults=true,orig_name=true:. your_proto_file.proto
```

## 生成结果

对于以下 protobuf 定义：

```protobuf
message User {
  string name = 1;
  int32 age = 2;
}
```

将生成：

```go
// MarshalJSON implements json.Marshaler
func (msg *User) MarshalJSON() ([]byte, error) {
    return protojson.MarshalOptions{
        UseEnumNumbers:  false,
        EmitUnpopulated: false,
        UseProtoNames:   false,
    }.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *User) UnmarshalJSON(b []byte) error {
    return protojson.UnmarshalOptions{
        DiscardUnknown: false,
    }.Unmarshal(b, msg)
}
```

## 为什么需要这个插件？

默认情况下，`protoc-gen-go` 生成的 Go 结构体不会实现标准库的 `json.Marshaler` 和 `json.Unmarshaler` 接口。这意味着使用 `encoding/json` 包时，字段名和枚举值的序列化行为可能与 protobuf 的 JSON 映射规范不一致。

本插件生成的代码使用官方的 `protojson` 包，确保：

- 字段名遵循 protobuf JSON 映射规范
- 枚举值默认序列化为字符串
- 正确处理 `oneof`、`optional` 等特殊字段
- 与其他语言的 protobuf JSON 实现保持一致

## 依赖

- Go 1.21+
- [google.golang.org/protobuf](https://pkg.go.dev/google.golang.org/protobuf)

## 许可证

[MIT License](LICENSE)
