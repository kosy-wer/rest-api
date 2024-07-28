# Rest-api - Simple Go REST Client 
![Version](https://img.shields.io/badge/version-1.0.0-%2330333a)



rest-api is my first project using Golang. It provides CRUD features for sending emails and several other functionalities.

What Rest-api fitur:
- CRUD (Create ,Remove ,Update ,Delete).
- Basic Authentication.
- Model data with validation.
- environment-based configuration.


## Compatibility note

This project is a simple implementation of CRUD operations for sending emails and authentication. It primarily focuses on displaying JSON responses. The current version is basic and lacks advanced features. Additionally, some features are not fully flexible and may lack robust support. The project has not undergone thorough testing for various conditions and edge cases.

Compatibility Note:
- The project primarily supports displaying JSON.
- It is not highly advanced or feature-rich.
- Some features are not fully flexible and well-supported.
- The project has not been thoroughly tested for various conditions.

# User Management API

This API provides functionalities for managing users, including creating, retrieving, updating, and deleting user data.

## Base URL

http://localhost:8080/api

## Endpoints
GET /users/{userEmail}

### Get User by Email

Retrieves a user by its email.

#### Parameters

| Name       | In     | Type   | Required | Description                |
|------------|--------|--------|----------|----------------------------|
| X-API-Key  | header | string | true     | API key for authorization  |
| userEmail  | path   | string | true     | Email of the user to retrieve |

#### Responses

| Code | Description                      | Example Value               |
|------|----------------------------------|-----------------------------|
| 200  | Successfully retrieved user.     | `"string"`                  |
| 404  | User not found.                  | `"string"`                  |
| 500  | Internal server error.           | `"string"`                  |

### Create a New User

Creates a new user in the system.

POST /users

#### Parameters

| Name       | In     | Type   | Required | Description                |
|------------|--------|--------|----------|----------------------------|
| X-API-Key  | header | string | true     | API key for authorization  |
| body       | body   | object | true     | The user object to create. |

#### UserCreateRequest Object

| Field | Type   | Required | Description                |
|-------|--------|----------|----------------------------|
| email | string | true     | Email of the user          |
| name  | string | true     | Updated name of the user   |

#### Responses

| Code | Description                      | Example Value               |
|------|----------------------------------|-----------------------------|
| 200  | Successfully created user.       | `"user details"`            |
| 400  | Bad request.                     | `"error message"`           |
| 500  | Internal server error.           | `"error message"`           |

### Update User

Updates an existing user.

PUT /users/{userEmail}


#### Parameters

| Name       | In     | Type   | Required | Description                |
|------------|--------|--------|----------|----------------------------|
| X-API-Key  | header | string | true     | API key for authorization  |
| userEmail  | path   | string | true     | Email of the user to update |
| body       | body   | object | true     | The user object to update. |

#### Responses

| Code | Description                      | Example Value               |
|------|----------------------------------|-----------------------------|
| 200  | Successfully updated user.       | `"user details"`            |
| 400  | Bad request.                     | `"error message"`           |
| 404  | User not found.                  | `"error message"`           |
| 500  | Internal server error.           | `"error message"`           |

### Delete User

Deletes an existing user.

DELETE /users/{userEmail}


#### Parameters

| Name       | In     | Type   | Required | Description                |
|------------|--------|--------|----------|----------------------------|
| X-API-Key  | header | string | true     | API key for authorization  |
| userEmail  | path   | string | true     | Email of the user to delete |

#### Responses

| Code | Description                      | Example Value               |
|------|----------------------------------|-----------------------------|
| 200  | Successfully deleted user.       | `"user details"`            |
| 404  | User not found.                  | `"error message"`           |
| 500  | Internal server error.           | `"error message"`           |

# Tables with Custom Background

Tables aren't part of the core Markdown spec, but they are part of GFM and Markdown Here supports them. They are an easy way of adding tables to your email -- a task that would otherwise require copy-pasting from another application.

Colons can be used to align columns.

<style>
  table {
    width: 100%;
    border-collapse: collapse;
  }
  th, td {
    border: 1px solid white;
    padding: 8px;
  }
  th {
    background-color: black;
    color: white;
  }
  td {
    background-color: black;
    color: white;
  }
</style>

<table>
  <tr>
    <th>Tables</th>
    <th>Are</th>
    <th>Cool</th>
  </tr>
  <tr>
    <td>col 3 is</td>
    <td>right-aligned</td>
    <td>$1600</td>
  </tr>
  <tr>
    <td>col 2 is</td>
    <td>centered</td>
    <td>$12</td>
  </tr>
  <tr>
    <td>zebra stripes</td>
    <td>are neat</td>
    <td>$1</td>
  </tr>
</table>

There must be at least 3 dashes separating each header cell. The outer pipes (|) are optional, and you don't need to make the raw Markdown line up prettily. You can also use inline Markdown.

<table>
  <tr>
    <th>Markdown</th>
    <th>Less</th>
    <th>Pretty</th>
  </tr>
  <tr>
    <td>*Still*</td>
    <td>`renders`</td>
    <td>**nicely**</td>
  </tr>
  <tr>
    <td>1</td>
    <td>2</td>
    <td>3</td>
  </tr>
</table>

