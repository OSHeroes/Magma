# MagmaServer
``Import Packages are't working , Download and place the repository at your GOPATH
  like ¨yourGoWorkspace¨/src/github.com/Synaxis/Magma``

MagmaServer handles requests from the game client and server.
It's used for login/Store/inventoryList/ as parsed XML
It can be used by both BFHeroes and BFPlay4Free
### How to Build
(remember change your IP inside ./config.go)
```
glide init

glide install

go build 
```
## Configuration

Here's a table with the env config variables
Please check ./config.go for more Info

| Name               | Default value   |
|--------------------|-----------------|
| `LOG_LEVEL`        | `debug`         |
| `HTTP_BIND`        | `127.0.0.1:80`  |//you can use 8080
| `HTTPS_BIND`       | `127.0.0.1:443` |
| `CERT_PATH`        | ./fixtures      |
| `PRIVATE_KEY_PATH` | ./fixtures      |

### Dependencies
All dependencies are resolved with [glide](https://github.com/Masterminds/glide)

## Credits
This was created by Synaxis and "neqnil" based on BF2BC backend,
credits: mDawg,Makahost
