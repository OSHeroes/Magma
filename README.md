# MagmaServer
``Import Packages are broken at the moment , Download and place the whole folder in your GOPATH
  like ¨yourGoWorkspace¨/src/github.com/Synaxis/bfheroesMagma``

MagmaServer handles REST requests from the game client and server.
It is another framework used for login/store and more
It can be use by both BFHeroes , and BFPlay4Free ( and more ?)

there is also "PLASMA" framework which we don't know how to make it (not docummented)
at the moment (? ) 

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
Golang dependencies are resolved thanks to [glide](https://github.com/Masterminds/glide).

```bash
glide init
glide install
```

### How to Compile

#### Windows

```=>run build.bat (remember change your IP inside ./config.go)```

## Credits

This was created by Synaxis and "neqnil" based on BF2BC backend, from the1domo ,freaky123(they recorded live traffic) .And using the implemenation and research from Aluigi. Also credits to MakaHost and mDawg
