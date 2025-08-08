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