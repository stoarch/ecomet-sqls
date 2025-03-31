# ecomet-sqls: Ecomet DB Query Language Server

>[ecomet db](https://github.com/vzroman/ecomet)
>[ecomet query language](https://github.com/stoarch/ecomet-sqls/blob/master/ecomet-query-language-en.md)

An implementation of the Language Server Protocol Ecomet Query Language.

Fork of sqls for this specific db

## Note

This project is currently under development and there is no stable release. Therefore, destructive interface changes and configuration changes are expected.

## Features

ecomet-sqls aims to provide advanced intelligence for you to edit ecomet queries in your own editor.

### Support RDBMS
- Ecomet([Driver](https://github.com/vzroman/ecomet))
- SQLite3([go-sqlite3](https://github.com/mattn/go-sqlite3))

### Language Server Features

#### Auto Completion

- DML(Data Manipulation Language)
    - [x] Get
    - [x] INSERT
    - [x] Set
    - [x] DELETE


#### CodeAction

- [x] Execute Query
- [x] Switch Connection(Selected Database Connection)
- [x] Switch Database

#### Hover

#### Signature Help


#### Document Formatting

## Installation

```shell
go install github.com/stoarch/ecomet-sqls@latest
```

## Editor Plugins

- [sqls.vim](https://github.com/sqls-server/sqls.vim)
- [vscode-sqls](https://github.com/lighttiger2505/vscode-sqls)
- [sqls.nvim](https://github.com/nanotee/sqls.nvim)
- [Emacs LSP mode](https://emacs-lsp.github.io/lsp-mode/page/lsp-sqls/)

## DB Configuration

The connection to the RDBMS is essential to take advantage of the functionality provided by `sqls`.
You need to set the connection to the RDBMS.

### Configuration Methods

There are the following methods for RDBMS connection settings, and they are prioritized in order from the top.
Whichever method you choose, the settings you make will remain the same.

1. Configuration file specified by the `-config` flag
1. `workspace/configuration` set to LSP client
1. Configuration file located in the following location
    - `$XDG_CONFIG_HOME`/sqls/config.yml ("`$HOME`/.config" is used instead of `$XDG_CONFIG_HOME` if it's not set)

### Configuration file sample

```yaml
# Set to true to use lowercase keywords instead of uppercase.
lowercaseKeywords: false
connections:
  - alias: dsn_ecomet
    driver: eql 
    dataSourceName: wss://127.0.0.1:13306/websocket
```

### Workspace configuration Sample

- setting example with vim-lsp.

```vim
if executable('sqls')
    augroup LspSqls
        autocmd!
        autocmd User lsp_setup call lsp#register_server({
        \   'name': 'sqls',
        \   'cmd': {server_info->['sqls']},
        \   'whitelist': ['sql'],
        \   'workspace_config': {
        \     'sqls': {
        \       'connections': [
        \         {
        \           'driver': 'ecomet',
        \           'dataSourceName': 'wss://127.0.0.1:13306/websocket',
        \         }
        \       ],
        \     },
        \   },
        \ })
    augroup END
endif
```

### Configuration Parameters

The first setting in `connections` is the default connection.

| Key         | Description          |
| ----------- | -------------------- |
| connections | Database connections |

### connections

`dataSourceName` takes precedence over the value set in `proto`, `user`, `passwd`, `host`, `port`, `dbName`, `params`.

| Key            | Description                                 |
| -------------- | ------------------------------------------- |
| alias          | Connection alias name. Optional.            |
| driver         | `ecomet`, `sqlite3` Required. |
| dataSourceName | Data source name.                           |
| proto          | `tcp`, `udp`, `unix`, `websocket`.                       |
| user           | User name                                   |
| passwd         | Password                                    |
| host           | Host                                        |
| port           | Port                                        |
| path           | unix socket path                            |
| dbName         | Database name                               |
| params         | Option params. Optional.                    |
| sshConfig      | ssh config. Optional.                       |

#### sshConfig

| Key        | Description                 |
| ---------- | --------------------------- |
| host       | ssh host. Required.         |
| port       | ssh port. Required.         |
| user       | ssh user. Optional.         |
| privateKey | private key path. Required. |
| passPhrase | passPhrase. Optional.       |

#### DSN (Data Source Name)

See also.

- <https://pkg.go.dev/github.com/jackc/pgx/v4>
- <https://github.com/mattn/go-sqlite3#connection-string>


## Inspired

I created ecomet-sqls inspired by the sqls.
