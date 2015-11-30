# API List

```
POST   : /auth/register
```
```
POST   : /auth/login
```
```
HEAD   : /auth/logout
```
```
POST   : /apps
```
```
GET    : /apps
```
```
PATCH  : /apps/:appID
```
```
DELETE : /apps/:appID
```
```
POST   : /apps/:appID/token
```
```
GET    : /apps/:appID/token/:token
```
```
POST   : /apps/:appID/releases
```
```
GET    : /apps/:appID/releases
```
```
PATCH  : /apps/:appID/releases/:releaseID
```
```
POST   : /apps/:appID/releases/:releaseID/bundles
```
```
GET    : /apps/:appID/releases/:releaseID/bundles
```
```
PUT    : /apps/:appID/releases/:releaseID/lock
```
```
POST   : /apps/:appID/releases/:releaseID/bundles/download
```

# Releasifier APIs

## Authentication

### - Create a new user

#### Request

`POST : /auth/register`

```
{
  "fullname": "Mr. Robot",
  "email": "mr.robot@savetheworld.com",
  "password": "saveUS"
}
```

#### Response

**Success**

200

```
{
  "id": 1,
  "jwt": "token...."
}
```

**Failure**

400

409

```
{
  "error": "duplicate email"
}
```

### - Login

#### Request

`POST : /auth/login`

```
{
  "email": "mr.robot@savetheworld.com",
  "password": "saveUS"
}
```

#### Response

**Success**

200

```
{
  "id": 1,
  "jwt": "1273651735172351725317531735173"
}
```

**Failure**

400

```
{
  "error": "request body has some issues"
}
```

401

```
{
  "error": "unauthorized access"
}
```

### - Logout

#### Request

`HEAD : /auth/logout`

#### Response

**Success**

200

**Failure**

## Apps

### - Create a new app

#### Request

`POST : /apps`

```
{
  "name": "Save the World",
  "public_key": "you can't catch me",
  "private_key": "hahaha"
}
```

#### Response

**Success**

201

```
{
  "id": 1,
  "name": "Save the World",
  "public_key": "you can't catch me",
  "created_at": "2015"
}
```

**Failure**

### - Get all apps

#### Request

`GET : /apps`

#### Response

**Success**

200

```
[
  {
    "id": 1,
    "name": "Save the World",
    "public_key": "you can't catch me",
    "created_at": "2015"
  },
  ...
]
```

```
[]
```

**Failure**

401

```
{
  "error": "unauthorized access"
}
```

### - Update an existing app

#### Request

`PATCH : /apps/:appID`

```
{
  "name": "Save the World",
  "public_key": "you can't catch me",
  "private_key": "key"
}
```

#### Response

**Success**

200

**Failure**

400

```
{
  "error": "request body has some issues"
}
```

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

### - Delete an existing app

#### Request

`DELETE : /apps/:appID`

#### Response

**Success**

200

**Failure**

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

### - Generate access token for an existing user

#### Request

`POST : /apps/:appID/token`

#### Response

**Success**

201

```
{
  "token": "12121212"
}
```

**Failure**

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

### - Assign permission to an existing user based on token

#### Request

`GET : /apps/:appID/token/:token`

#### Response

**Success**

200

**Failure**

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

## Apps / Releases

### - Create a new release

#### Request

`POST : /apps/:appID/releases`

```
{
  "version": "1.0.0",
  "platform": "ios",
  "note": "this version includes"
}
```

#### Response

**Success**

201

```
{
  "id": 1,
  "version": "1.0.0",
  "platform": "ios",
  "note": "this version includes",
  "created_at": "2015"
}
```

**Failure**

400

```
{
  "error": "request body has some issues"
}
```

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

409

```
{
  "error": "duplicate version"
}
```

```
{
  "error": "duplicate name"
}
```

### - Get all releases for a specific app

#### Request

`GET : /apps/:appID/releases`

#### Response

**Success**

200

```
[
  {
    "id": 1,
    "version": "1.0.0",
    "platform": "ios",
    "note": "this version includes",
    "created_at": "2015"
  },
  ...
]
```

```
[]
```

**Failure**

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

### - Update an existing release

#### Request

`PATCH : /apps/:appID/releases/:releaseID`

```
{
  "version": "1.0.0",
  "platform": "ios",
  "note": "this version includes"
}
```

#### Response

**Success**

200

**Failure**

400

```
{
  "error": "request body has some issues"
}
```

401

```
{
  "error": "unauthorized access"
}
```

403

```
{
  "error": "release is locked"
}
```

404

```
{
  "error": "app not found"
}
```

```
{
  "error": "release not found"
}
```

### - Upload one or many files as part of bundles

#### Request

`POST : /apps/:appID/releases/:releaseID/bundles`

#### Response

**Success**

201

```
[
  {
    "id": 1,
    "name": "main.jsbundle",
    "type": "javascript",
    "hash": "12763tuyg2bhjwquid7",
    "created_at": "2015"
  },
  ...
]
```

**Failure**

400

```
{
  "error": "files are too big"
}
```

401

```
{
  "error": "unauthorized access"
}
```

403

```
{
  "error": "release is locked"
}
```

404

```
{
  "error": "app not found"
}
```

```
{
  "error": "release not found"
}
```

### - Get all bundles for specific app and release

#### Request

`GET : /apps/:appID/releases/:releaseID/bundles`

#### Response

**Success**

200

```
[
  {
    "id": 1,
    "name": "main.jsbundle",
    "type": "javascript",
    "hash": "12763tuyg2bhjwquid7",
    "created_at": "2015"
  },
  ...
]
```

```
[]
```

**Failure**

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

```
{
  "error": "release not found"
}
```

### - Lock release

#### Request

`PUT : /apps/:appID/releases/:releaseID/lock`

#### Response

**Success**

200

**Failure**

400

```
{
  "error": "release is already locked"
}
```

401

```
{
  "error": "unauthorized access"
}
```

404

```
{
  "error": "app not found"
}
```

```
{
  "error": "release not found"
}
```

### - Request for a zip file contains all bundles

#### Request

`POST : /apps/:appID/releases/:releaseID/bundles/download`

#### Response

**Success**

200

```
bytes: bundle.zip
```

**Failure**

404

```
{
  "error": "app not found"
}
```

```
{
  "error": "release not found"
}
```

```
{
  "error": "release is not ready"
}
```
