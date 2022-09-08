# gologger
Simple logging library with loglevel functionalities.

gologger provides the following log levels:

- Panic
- Error
- Warn
- Info
- Debug
- Trace

Customize Timestamp with func "SetTimestampFormat(format string)"

Set log level with func "SetLogLevel(level Level)"

Logs are printed out with fmt.Printf.  
On default, the format of the log messages is:

```<Timestamp> [<Loglevel>] - <Logmessage>```

For example:

```2006-01-02 15:04:05.000 [DEBUG  ] - Logmessage```
