📂 โครงสร้างโปรเจค

|── .github
|   └── go.yml    
├── api
│   ├── handlers       # ตัวจัดการ API
│   ├── middleware     # Middleware
│   └── routes.go      # กำหนด routing
├── config
│   └── config.go      # การตั้งค่าระบบ
├── db
│   ├── migrations     # Database migrations
│   └── models         # Database models
├── docker
│   └── nginx          # Nginx configuration
├── docs               # Api spec / ER-diagram docs / Project-Structer
├── pkg                
|   └── his            # Utility functions
├── tests
│   ├── api_test.go    # Unit tests
│   └── test_curl.go   
├── .env.example
├── docker-compose.yml
├── Dockerfile
└── go.mod
└── go.sum
└── README.md

