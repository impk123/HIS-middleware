# API Specification Document

This document provides the detailed specification for the API endpoints, including authentication, staff management, and patient management. It also includes examples of API usage.

## 1. Authentication

### 1.1 Staff Login

- **Endpoint:** `POST /staff/login`
- **Description:** This endpoint allows staff members to authenticate themselves and receive a JWT token.
- **Request Body:** 
  ```json
  {
    "username": "string",
    "password": "string",
    "hospital": "string"
  }
  ```
- **Responses:**
  - **200 Success:**
    ```json
    {
      "token": "string",
      "staff": {
        "id": "integer",
        "username": "string",
        "hospital": "string"
      }
    }
    ```
  - **400 Bad Request**
  - **401 Unauthorized**
  - **500 Internal Server Error**

## 2. Staff Management

### 2.1 Create Staff

- **Endpoint:** `POST /staff/create`
- **Authentication:** Required (Admin only)
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string",
    "hospital": "string"
  }
  ```
- **Responses:**
  - **201 Created**
  - **400 Bad Request**
  - **401 Conflict (username already exists)**

## 3. Patient Management

### 3.1 Search Patients

- **Endpoint:** `GET /patients/search`
- **Authentication:** Required
- **Query Parameters:**
  - `first_name_th` (optional)
  - `last_name_th` (optional)
  - `first_name_en` (optional)
  - `last_name` (optional)
  - `date_of_birth` (optional)
  - `gender` (optional)
  - `patient_hn` (optional)
  - `phone_number` (optional)
  - `email` (optional)
- **Responses:**
  - **200 Success:**
    ```json
    [
      {
        "first_name_th": "string",
        "last_name_th": "string",
        "first_name_en": "string",
        "last_name": "string",
        "date_of_birth": "date (YYYY-MM-DD)",
        "gender": "string",
        "patient_hn": "string",
        "phone_number": "string",
        "email": "string"
      }
    ]
    ```
  - **401 Unauthorized**

## 7. Example API Usage

### 7.1 Successful Login Flow

1. ** StaffCreate (success) สร้าง Staff
 ```bash
   curl -X POST http://localhost:8080/staff/create \
    -H "Content-Type: application/json" \
    -d "{
         "username": "admin",
         "password": "admin123",
         "hospital": "12123"
     }"
   ```
2. **Login to get token:**
   ```bash
   curl -X POST http://localhost:8080/staff/login \
     -H "Content-Type: application/json" \
     -d '{"username":"admin","password":"admin123","hospital":"12123"}'
   ```

3. **Use token to access protected endpoint:**
   ```bash
   curl -X GET http://localhost:8080/patients/search \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
     -H "Content-Type: application/json"
   ```



Feel free to reach out if you have any further questions or need additional information regarding the API specification!
