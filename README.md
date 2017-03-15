# Statuses

HTTP status code utility for go.

```go
import (
  "VanishingDante/statuses"
)
```

### status.Status

If the param is a valid HTTP code, then the status message will be returned. Otherwise, an empty string will be returned.

```go
status.Status(302) // => "Found"
status.Status(9999) // => ""
```

func RedirectCode(code int) bool

Returns `true` if a status code is a valid redirect status.

```go
status.RedirectCode[200] // => false
status.RedirectCode[301] // => true
```

### func EmptyCode(code int) bool

Returns `true` if a status code expects an empty body.

```go
status.EmptyCode[200] // => false
status.EmptyCode[204] // => true
status.EmptyCode[304] // => true
```

### func RetryCode(code int) bool

Returns `true` if you should retry the rest.

```go
status.RetryCode[501] // => false
status.RetryCode[503] // => true
```
