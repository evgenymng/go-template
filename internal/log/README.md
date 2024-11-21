# Log

## Tips & Best Practices

The logs are uniform and structured. This is achieved by using an
abstraction called `LogObject`. This approach is handy in case you want to do
log aggregation, and generally easier on eyes.

Basically, you should always specify `Tag` for a log message:

```go
log.L().Tag(log.TagRequest)
```

For your convenience, you can store the prepared `LogObject`s in a variable,
I would recommend you naming it `l`:

```go
l := log.L().Tag(log.TagRequest)
if err != nil {
    log.Error("Alarm!", l.Error(err))
}
```

There are number of methods, through which you can attach different
information to the `LogObject`:
+ `Tag()`;
+ `Error()`;
+ `TraceId()`;
+ `Add()` — add an arbitrary key-value pair.

All this information will be inlined in logs.
