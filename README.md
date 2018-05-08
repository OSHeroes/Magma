# MagmaServer
``Import Packages are broken at the moment , Download and place the whole folder in your GOPATH
  like ¨yourGoWorkspace¨/src/github.com/Synaxis/Magma``

MagmaServer handles REST requests from the game client and server.
It is another framework used for login/store/inventoryList/XML retrieving and more
It can be used by both BFHeroes and BFPlay4Free


## Configuration

Below there is table with the config variables used by the `Magma`.
Please check ./config.go for more Info

| Name               | Default value   |
|--------------------|-----------------|
| `LOG_LEVEL`        | `debug`         |
| `HTTP_BIND`        | `127.0.0.1:80`  |//you can use 8080
| `HTTPS_BIND`       | `127.0.0.1:443` |
| `CERT_PATH`        | ./fixtures      |
| `PRIVATE_KEY_PATH` | ./fixtures      |

### Dependencies
Golang dependencies are resolved with [glide](https://github.com/Masterminds/glide).

### How to Compile
```bash
glide init
glide install
go build -a -ldflags= -v -o ./magma.exe ./cmd/heroes-api
```
#### Windows

```=>go build (remember change your IP inside ./config.go)```

## Credits

This was created by Synaxis and "neqnil" based on BF2BC backend,
credits:the1domo ,freaky123. Aluigi, mDawg
