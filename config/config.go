package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBTestName string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "HIS_hosp"),
		DBTestName: getEnv("DB_TEST_NAME", "test_his_hosp"),
	}
}

// // LoadTestConfig โหลดการตั้งค่าสำหรับการทดสอบ
// func LoadTestConfig() *Config {
// 	return &Config{
// 		DB: DBConfig{
// 			Host:     getEnv("TEST_DB_HOST", "localhost"),
// 			Port:     getEnv("TEST_DB_PORT", "5433"), // ใช้พอร์ตต่างจากของปกติ
// 			User:     getEnv("TEST_DB_USER", "test_user"),
// 			Password: getEnv("TEST_DB_PASSWORD", "test_password"),
// 			Name:     getEnv("TEST_DB_NAME", "test_his_hosp"),
// 		},
// 		AppEnv:  "test",
// 		AppPort: getEnv("TEST_APP_PORT", "8081"), // ใช้พอร์ตต่างจากของปกติ
// 	}
// }

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
