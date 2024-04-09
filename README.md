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

- Docker

## Installation and Usage

1. Clone the DCS API repository.
2. Make sure you have Docker, and git installed on your system.
3. Build the binary and execute using:

```bash
git clone https://github.com/Paique/DesClientSider-API.git
cd DesClientSider-API
docker compose up
```


The optional variables have predefined default values that will be replaced if the variables are set in compose.yaml:
```yaml
#These are the default values
environment:
  DCS_LISTEN_PORT: 8080
  DB_NAME: dcs
  DB_PORT: 3306
```

This will start the API server on the desired port (or on the default 8080 if not defined).
## Contribution

Contributions are welcome! If you find bugs or have suggestions for improvements, feel free to open an issue or send a pull request to the official DCS API repository.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

---

This project is part of the DCS [official DCS client repository](https://github.com/Paique/DesClientSider).
