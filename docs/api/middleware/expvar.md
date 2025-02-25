---
id: expvar
---

# ExpVar

Expvar middleware for [Fiber](https://github.com/gofiber/fiber) that serves via its HTTP server runtime exposed variants in the JSON format. The package is typically only imported for the side effect of registering its HTTP handlers. The handled path is `/debug/vars`.

## Signatures

```go
func New() fiber.Handler
```

## Examples

Import the middleware package that is part of the Fiber web framework

```go
import (
  "github.com/hakantaskin/fiber"
  expvarmw "github.com/hakantaskin/fiber/middleware/expvar"
)
```

After you initiate your Fiber app, you can use the following possibilities:
```go
var count = expvar.NewInt("count")

app.Use(expvarmw.New())
app.Get("/", func(c *fiber.Ctx) error {
    count.Add(1)

    return c.SendString(fmt.Sprintf("hello expvar count %d", count.Value()))
})
```

Visit path `/debug/vars` to see all vars and use query `r=key` to filter exposed variables.

```bash
curl 127.0.0.1:3000
hello expvar count 1

curl 127.0.0.1:3000/debug/vars
{
    "cmdline": ["xxx"],
    "count": 1,
    "expvarHandlerCalls": 33,
    "expvarRegexpErrors": 0,
    "memstats": {...}
}

curl 127.0.0.1:3000/debug/vars?r=c
{
    "cmdline": ["xxx"],
    "count": 1
}
```

## Config

| Property | Type                    | Description                                                         | Default |
|:---------|:------------------------|:--------------------------------------------------------------------|:--------|
| Next     | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. | `nil`   |

## Default Config

```go
var ConfigDefault = Config{
    Next: nil,
}
```
