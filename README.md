# MagmaServer

MagmaServer handles REST requests from the game client and server.
It is another framework used for login/store and more
It can be use by both BFHeroes , and BFPlay4Free ( and more ?)

there is also "PLASMA" framework which we don't know how to make it (not docummented)
at the moment (or maybe just a weird name for Server ? ) 

## Configuration

Below there is table with all enviroment variables which are used by the `heroes-api`.
You can take a glance at `config/config.go` file if you need more information about specific variable.

Instead of editing `~/bashrc` you may create a `.env` file and insert there configuration of your choice.
Optionally, you can provide a custom path to a `.env` file by providing `--config` flag (i.e. `./backend --config dev.env`).

| Name               | Default value   |
|--------------------|-----------------|
| `LOG_LEVEL`        | `DEBUG`         |
| `HTTP_BIND`        | `127.0.0.1:80`  |
| `HTTPS_BIND`       | `127.0.0.1:443` |
| `CERT_PATH`        |                 |
| `PRIVATE_KEY_PATH` |                 |

### Dependencies

Currently golang dependencies are resolved thanks to [glide](https://github.com/Masterminds/glide).

Packages stored in `vendor` directory SHOULD NOT be pushed to the repository.

```bash
glide init
glide install
```

### Start

```bash
go build -o heroes-api cmd/heroes-api/main.go && ./heroes-api
```

It will load `.env` file by default.

#### Windows

Unfortunately, Windows is not really great for running any console-based applications, but if you use `powershell` you might also appreciate following command:

```powershell
go build -o heroes-api.exe cmd/heroes-api/main.go ; if ($?) { .\heroes-api.exe }
```

Or following if you are using custom `.env` file:

```powershell
go build -o heroes-api.exe cmd/heroes-api/main.go ; if ($?) { .\heroes-api.exe --config .dev.env }
```

Note: PowerShell has one big advantage over other terminal in Windows - text coloring of logs.

## Credits

This was created by mea and a friend of mine called "neqnil" based on BF2MASE backend, from the1domo ,freaky123(they recorded live traffic) . And using the implemenation and research from Aluigi
