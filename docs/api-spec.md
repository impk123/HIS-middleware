# API Specifications

## Staff Management

### Create Staff
`POST /staff/create`

Request Body:
```json
{
    "username": "string",
    "password": "string",
    "hospital": "string"
}
Respond 200 Success:
```json
{
    "CreatedAt": "string(YYYY-MM-DD)",
    "UpdatedAt": "string(YYYY-MM-DD)",
    "DeletedAt": null,
    "ID": "integer",
    "Username": "string",
    "Password": "string",
    "Hospital": "string",
    "FirstNameTh": "string",
    "LastNameTh": "string"
}
