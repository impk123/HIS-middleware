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

# ติดตั้ง Go dependencies
    go download
## Setup
1. Clone the repository
2. Copy `.env.example` to `.env` and configure
3. Run Docker (bash)
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



