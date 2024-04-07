
# DCS API

DCS API é uma API bem simples em Go projetada para fornecer funcionalidades relacionadas ao [DesClientSider (DCS)](https://github.com/Paique/DesClientSider), uma ferramenta para gerenciamento de mods client-side em servidores de Minecraft.

## Funcionalidades

- **Rota /keywords**: Retorna as palavras-chave atualmente configuradas no DCS para identificar mods incompatíveis.
- **Rota /contra**: Retorna as palavras-chave configuradas para identificar mods contra os quais o DCS não deve agir.

## Endpoints

### GET /keywords

Retorna as palavras-chave atualmente configuradas no DCS.

#### Exemplo de Requisição:

```
GET /keywords
```

#### Exemplo de Resposta:

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

Retorna as palavras-chave configuradas para identificar mods contra os quais o DCS não deve agir.

#### Exemplo de Requisição:

```
GET /contra
```

#### Exemplo de Resposta:

```json
[
    {
        "id": "0",
        "keyword": "ftb-essentials"
    }
]
```

## Requisitos

- Go 1.22
- MySql

## Instalação e Uso

1. Clone o repositório da DCS API.
2. Certifique-se de ter Go instalado em seu sistema.
3. Faça build do binário para o seu OS utilizando:

```bash
go build
```
4. Realize a configuração do Mysql, crie o seu usuário, e defina uma senha, também crie a database dcs, e as tabelas Keywords, e ContraKeywords, as duas contendo:
```mysql
USE dcs;
CREATE TABLE ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);
CREATE TABLE Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);
```
Não se esqueça de definir as keys desejadas na database!

5. Defina as variáveis de ambiente:
```
 DB_USERNAME, DB_PASSWORD
```

Opcionalmente você pode definir as variáveis:
```
DCS_LISTEN_PORT, DB_NAME, DB_PORT
```

As variáveis opcionais possuem valores padrões pré-definidos que serão substituídos se as variáveis forem setadas:
DCS_LISTEN_PORT = 8080
DB_NAME=dcs
DB_PORT=3306

Isso iniciará o servidor da API na porta desejada (ou na padrão 8080 se não definir).
É recomendado utilizar Docker.

## Contribuição

Contribuições são bem-vindas! Se você encontrar bugs ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request para o repositório oficial da DCS API.

## Licença

Este projeto é licenciado sob a [MIT License](https://opensource.org/licenses/MIT).

---

Este projeto faz parte do DCS [repositório oficial do DCS client](https://github.com/Paique/DesClientSider).
