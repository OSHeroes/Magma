# MagmaServer
``Import Packages are't working , Download and place the repository at your GOPATH
  like ¨yourGoWorkspace¨/src/github.com/Synaxis/Magma``

MagmaServer handles requests from the game client and server.
It's used for login/Store/inventoryList/ as parsed XML
It can be used by both BFHeroes and BFPlay4Free
### How to Build
(change your IP inside ./config.go)
```
glide init

glide install

go build 
```
## Configuration

Here are all the .env (file) config variables
check ./config.go for more Info
To create a .env file just type in bash (or gitbash)
touch .env

| Name               | Value           |
|--------------------|-----------------|
| `LOG_LEVEL`        | `debug`         |
| `HTTP_BIND`        | `127.0.0.1:80`  |//you can use 8080
| `HTTPS_BIND`       | `127.0.0.1:443` |
| `CERT_PATH`        | ./fixtures      |
| `PRIVATE_KEY_PATH` | ./fixtures      |
## Credits
This is created by Synaxis and "neqnil" based on BF2BC backend,
credits: mDawg,Makahost
