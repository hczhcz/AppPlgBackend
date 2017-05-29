# Public APIs

## Basic Interface

Base URL: “...”

#### Request

Uses HTTPS POST & JSON.

| name       | type          |
| ---        | ---           |
| session_id | string / null |

#### Response

Uses JSON.

| name  | type   |
| ---   | ---    |
| error | uint64 |
| ...   | ...    |

If error == 0, “...” is normal response.

Else, “...” is error response corresponding to the error code.

## Error Codes

### 001 Invalid session_id

| name       | type   |
| ---        | ---    |
| session_id | string |

## Data Structures

### User

| name        | type                        |
| ---         | ---                         |
| gender      | “male” / “female” / “other” |
| description | string                      |

### Login

| name     | type          |
| ---      | ---           |
| email    | string / null |
| phone    | string / null |

Requires (email != null) xor (phone != null).

### Post

TODO

## Commands

### /user_new

#### Request

| name  | type  |
| ---   | ---   |
| login | Login |
| user  | User  |

TODO: captcha?

Response:

(none)

Note: Verify user via email or SMS.

### /user_login

#### Request

| name  | type  |
| ---   | ---   |
| login | Login |

TODO: captcha?

Response:

(none)

Note: Verify user via email or SMS.

### /user_verify

#### Request

| name | type   |
| ---  | ---    |
| code | string |

#### Response

| name    | type   |
| ---     | ---    |
| user_id | string |

### /user_get

Request:

| name    | type   |
| ---     | ---    |
| user_id | uint64 |

#### Response

| name | type        |
| ---  | ---         |
| user | User / null |

### /user_update

#### Request

| name | type |
| ---  | ---  |
| user | User |

#### Response

(none)

# Implementation Specifications

## Dependencies

* PostgreSQL (github.com/go-pg/pg)

## Database Schemes

### Common keys

| name        | type  |
| ---         | ---   |
| create_time | int64 |
| deleted     | bool  |
