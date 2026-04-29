# Katana: Split files to chunks and merge them

## Why?

Currenly there are some limitations sending large files on some IM platforms.  
With this, one can make a single file to 512KiB parts, and send them.  
The other party in chat can easily merge them.

## Build

```sh
go build -o bin/katana main.go
# or
make build
```

## How to

### `split`

The commands splits a file. Assume you have `input.txt`:

```sh
./katana split -f input.txt
```

Generates files like this:

```sh
input.txt.0
input.txt.1
input.txt.2
input.txt.3
```

### `merge`

To merge, you only pass the original filename, which is root name of files:

```sh
./katana merge -f input.txt
```

It finds the partions in order and merges them all. Output is: `katana_input.txt`


### `check`

This is just for test. Compares size of the original file (if available) and merge output.  

```sh
./katana check input.txt katana_input.txt
# output: true
```

## TODO

### Features
- [ ] Adding custom chunk size
- [ ] Adding unit tests
- [ ] E2E tests on different file formats

### Futures!
- [ ] Adding RSA E2E encryption for key exchange
- [ ] Adding AES encryption for chunks
- [ ] Adding Key-Gen for RSA and AES