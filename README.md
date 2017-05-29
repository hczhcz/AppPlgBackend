# Public APIs

## Basic Interface

Base URL: “...”

#### Request

Uses HTTPS POST & JSON.

| name       | type          |
| ---        | ---           |
| session_id | string / null |

session_id should be generated by secure pseudorandom number generator.

#### Response

Uses JSON.

| name   | type          |
| ---    | ---           |
| action | string / null |
| data   | {...}         |

If action == null, the data field is normal response.

Else, the corresponding action handler (typically an error handler) should be triggered, and the data field is the handler's required information.

## Data Structures

### User

| name        | type                        |
| ---         | ---                         |
| nickname    | string                      |
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

#### Response

(none)

Verification code is sent via email or SMS.

### /user_login

#### Request

| name  | type  |
| ---   | ---   |
| login | Login |

TODO: captcha?

#### Response

(none)

Verification code is sent via email or SMS.

### /user_verify

#### Request

| name | type   |
| ---  | ---    |
| code | string |

#### Response

| name    | type   |
| ---     | ---    |
| user_id | uint64 |

Called after /user_new or /user_login.

### /user_get

#### Request

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

## Action Handlers

### session_init

| name       | type   |
| ---        | ---    |
| session_id | string |

Triggered when the current session is null or invalid.

The client should reassign the session_id.

# Implementation Specifications

## Dependencies

* PostgreSQL (github.com/go-pg/pg)

## Database Schemes

### Common keys

| name        | type   |
| ---         | ---    |
| create_time | uint64 |
| deleted     | bool   |
