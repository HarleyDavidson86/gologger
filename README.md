# gologger
This is a simple wrapper around the log package of go  
to provide different loglevels.

gologger provides the following log levels:

- Panic
- Error
- Warn
- Info
- Debug
- Trace

Customize Timestamp with func "SetTimestampFormat(format string)"

Set log level with func "SetLogLevel(level Level)"

Logs are printed out with log.Printf.  
By default, the format of the log messages is:

```2006/01/02 15:04:05 [<Loglevel>] - <Logmessage>```

If you set the flags of default log package to 0 (```log.SetFlag(0)```) then gologger  
is using the default timestamp:

```2006-01-02 15:04:05.000 [DEBUG  ] - Logmessage```

So you can customize the prefixes over the functionality of the standard [log package of go](https://pkg.go.dev/log) or by set a [correct timeformat](https://go.dev/src/time/format.go over the function ```SetTimestampFormat```.

This gives you also the possibility to connect a packages like [lumberjack](https://github.com/natefinch/lumberjack) to gain the functionality of a rolling log for example.

## Changlog

v1.0.2 -> v1.0.3: Inserted function ```SetTimestampFormat``` again with the possibility to use prefixes of the log package or define an own timestamp format.

v1.0.1 -> v1.0.2: Using default go log package, so the default timestamp changed and could not be set over the funtion ```SetTimestampFormat``` any more.