# heroes-api

Heroes API handles REST requests from the game client and server.

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
glide install
```

### Start

```bash
go build -o heroes-api cmd/heroes-api/main.go && ./heroes-api
```

It will load `.env` file by default.

#### Linux

As a std user (with a custom config file):

```bash
go build -o heroes-api cmd/heroes-api/main.go && ./heroes-api --config dev.env
```

As a root user:

```bash
go build -o heroes-api cmd/heroes-api/main.go && sudo -H ./heroes-api
```

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

This repository was based on the implementation, which can be found in `github.com/HeroesAwaken/GoFesl`.
