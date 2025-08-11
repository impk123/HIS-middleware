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
2. Basic understanding of the golang syntax.
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

## ER Diagram
See [docs/er-diagram.png](docs/er-diagram.png)

## Project Stucture
See [docs/project-structure.md](docs/project-structure.md)




