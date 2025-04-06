# Voidserve

A dead simple local HTTP server for quickly serving a folder of files. Just run `voidserve` within the directory. It's meant for local development, files sharing, etc., not for production use and serving files publicaly.

## Installation

```bash
go install github.com/make0x20/voidserve@latest
```

## Usage

```bash
voidserve # Serve the current directory on port 8080

voidserve -p 8081 # Serve the current directory on port 8081

voidserve -a # Serve on all interfaces - when you want to access it from another device
```
