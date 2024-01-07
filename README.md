# Open Telemetry compatible SQL driver wrapper for Golang

This package is an implementation of a SQL driver wrapper with tracing capabilities, compatible with the Open Telemetry (OTEL) API.
It is a fork of https://github.com/inkbe/opentracing-sql but with changes to use the otel packages instead of opentracing.

## Installation

To use this package, first install it via go get:

```
go get github.com/jonas-jonas/otelsql/v2
```

## Usage

Register a new database driver by passing an instance created by calling `NewTracingDriver``:

```go

// ...
import (
    "github.com/jonas-jonas/otelsql/v2"
)
// ...

var driver *sql.Driver
var tracer otel.Tracer // e.g. otel.Tracer("sql-tracing")
// init driver, tracer.
// ...
sql.Register("otel-sql", otelsql.NewTracingDriver(driver, tracer))
db, err := sql.Open("otel-sql", ...)
// use db handle as usual.
```

By default, a runtime-based naming function will be used, which will set the span name according to the name of the
function being called (e.g. `conn.QueryContext`).

It's also possible to specify your own naming function:

```
otsql.NewTracingDriver(driver, tracer, otsql.SpanNameFunction(customNameFunction))
```

Example custom naming function:

```go
func spanNamingFunc(ctx context.Context) string {
	pc, _, _, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	return f.Name()
}
```

Note that only calls to context-aware DB functions will be traced (e.g. db.QueryContext).

## Comparison with existing packages

There is an existing package https://github.com/ExpansiveWorlds/instrumentedsql which uses the same approach by wrapping
an existing driver with a tracer, however the current implementation provides the following features:

- Pass custom naming function to name spans according to your needs.
- Option to enable/disable logging of SQL queries.

The following features from instrumentedsql package are not supported:

- Passing a custom logger.
- Support of cloud.google.com/go/trace.
- Don't log exact query args.
- Creating spans for LastInsertId, RowsAffected.

## Documentation

[GoDoc documentation](https://godoc.org/github.com/jonas-jonas/otelsql)

## References

[OTEL Go libraries](https://github.com/open-telemetry/opentelemetry-go)
