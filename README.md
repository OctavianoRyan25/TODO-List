# Simple Rest API

## Features

-   Menggunakan validasi tiap-tiap field
-   Autentitakasi menggunakan JsonWebToken(Token berganti secara otomatis setelh 1 jam tergenerate)
-   Authorization

## API Reference

#### Get all items

```http
  GET /note
```

#### Get item

```http
  GET /note/${id}
```

| Parameter | Type  | Description                       |
| :-------- | :---- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

#### Store item

```http
  POST /note
```

#### Update item

```http
  PUT /note/${id}
```

| Parameter | Type  | Description                       |
| :-------- | :---- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

#### Delete item

```http
  DELETE /note/${id}
```

| Parameter | Type  | Description                       |
| :-------- | :---- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch |

## Tech

Beberapa library yang digunakan dalam project ini

-   [GinGonic]
-   [PostgresDriver]
-   [Swagger]
-   [Gorm]

[//]: # "These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax"
[swagger]: https://github.com/swaggo/swag
[postgresdriver]: https://github.com/lib/pq
[gingonic]: https://github.com/gin-gonic/gin
[gorm]: https://gorm.io/
