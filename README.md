# HIS-middleware
# Hospital System


A middleware system for hospital patient information management.

## Features
- Patient information retrieval from HIS
- Staff management and authentication
- Patient search functionality

## Tech Stack
- Go 1.20
- Gin Framework
- PostgreSQL
- Docker
- Nginx

## Setup
1. Clone the repository
2. Run go mod tidy
3. Copy `.env.example` to `.env` and configure
4. Run Docker (bash)
    สร้างและรัน containers 
    -   Run `docker-compose up --build`
    ตรวจสอบการทำงาน
    -   Run `docker-compose ps`
    หยุดการทำงาน
    -   Run `docker-compose down`
    ลบ volumes (หากต้องการเริ่มใหม่)
    -   Run `docker-compose down -v`

## API Documentation
See [docs/api-spec.md](docs/api-spec.md)

## Project Stucture
- See [docs/project-structure.md](docs/project-structure.md)
- See [docs/project-structure.png](docs/Project-structure.png)

## ER Diagram
See [docs/ERdiagram.png](docs/ERdiagram.png)



