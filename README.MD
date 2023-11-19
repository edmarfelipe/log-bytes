# Log Bytes Counter

[![Go Report Card](https://goreportcard.com/badge/github.com/edmarfelipe/log-bytes)](https://goreportcard.com/report/github.com/edmarfelipe/log-bytes)

`log-bytes` counts the bytes from the stdout of a command, and prints the result to stdout, it is useful to measure the size of a service log.


## Install

```bash
go install github.com/edmarfelipe/log-bytes@latest
```

## Usage

```
<command> | log-bytes [flags]

flags:
  -template string
        Template to print [default, minimal] (default "default")
  -unit string
        Data unit to print  [B, KB, MB, GB] (default "kb")
```

**Using with default**

```bash
$ server | log-bytes
```

**Example with minimal template**

```bash
$ server | log-bytes -template minimal
```

**Example with minimal template and**

```bash
$ server | log-bytes -template minimal -unit mb
```


