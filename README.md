# DCS API

DCS API is a simple Go API designed to provide functionalities related to [DesClientSider (DCS)](https://github.com/Paique/DesClientSider), a tool for managing client-side mods on Minecraft servers.

## Features

- **/keywords Endpoint**: Returns the keywords currently configured in DCS to identify incompatible mods.
- **/contra Endpoint**: Returns keywords configured to identify mods against which DCS should not act.

## Endpoints

### GET /keywords

Returns the keywords currently configured in DCS.

#### Request Example:

```
GET /keywords
```

#### Response Example:

```json
[
    {
        "id": "0",
        "keyword": "advancementplaques"
    },
    {
        "id": "1",
        "keyword": "advdebug"
    },
    {
        "id": "2",
        "keyword": "afkpeace"
    }
]
```

### GET /contra

Returns keywords configured to identify mods against which DCS should not act.

#### Request Example:

```
GET /contra
```

#### Response Example:

```json
[
    {
        "id": "0",
        "keyword": "ftb-essentials"
    }
]
```

## Requirements

- Go 1.22
- MySql

## Installation and Usage

1. Clone the DCS API repository.
2. Make sure you have Go installed on your system.
3. Build the binary for your OS using:

```bash
go build
```

4. Configure MySql, create your user, set a password, create the `dcs` database, and the `Keywords` and `ContraKeywords` tables, both containing:
```mysql
CREATE DATABASE dcs;
USE dcs;
CREATE TABLE ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);
CREATE TABLE Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);
```
Don't forget to define the desired keys in the database!

5. Set the environment variables:
```
DB_USERNAME, DB_PASSWORD
```

Optionally, you can set the variables:
```
DCS_LISTEN_PORT, DB_NAME, DB_PORT
```

The optional variables have predefined default values that will be replaced if the variables are set:
DCS_LISTEN_PORT = 8080
DB_NAME=dcs
DB_PORT=3306

This will start the API server on the desired port (or on the default 8080 if not defined).
It is recommended to use Docker.

## Contribution

Contributions are welcome! If you find bugs or have suggestions for improvements, feel free to open an issue or send a pull request to the official DCS API repository.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

---

This project is part of the DCS [official DCS client repository](https://github.com/Paique/DesClientSider).
