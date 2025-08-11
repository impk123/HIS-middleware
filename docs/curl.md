package curl_test

//สร้างคำสั่ง curl สำหรับทดสอบ API

//  StaffCreate (success) สร้าง Staff

// curl -X POST http://localhost:8080/staff/create \
// -H "Content-Type: application/json" \
// -d "{
//     "username": "admin",
//     "password": "admin123",
//     "hospital": "12123"
// }"

//# StaffCreate (fail - missing fields) สร้าง Staff แบบผิด
// curl -X POST http://localhost:8080/staff/create \
// -H "Content-Type: application/json" \
// -d "{
//     "username": "admin",
//     "hospital": "12123"
// }"

//# Staff Login (Success) ล็อกอินเพื่อรับ Token
// curl -X POST http://localhost:8080/staff/login \
// -H "Content-Type: application/json" \
// -d "{
//     "username": "admin",
//     "password": "admin123",
//     "hospital": "12123"
// }"

//# Staff Login (Wrong Password)
// curl -X POST http://localhost:8080/staff/login \
// -H "Content-Type: application/json" \
// -d "{
//     "username": "admin",
//     "password": "wrongpass",
//     "hospital": "12123"
// }"

//# ใช้ Token ที่ได้เพื่อเข้าถึง API อื่น
// ค้นหาผู้ป่วย
// curl -X GET http://localhost:8080/patient/search \
//   -H "Authorization: Bearer <your_token_here>"
//   -d "{
//       "first_name": "สมศรี"
//   }"

// ทดสอบการเชื่อมต่อกับ Hospital A API (Optional)
// implement HIS integration
// curl -X GET "http://localhost/patient/search?national_id=1234567890123" \
//   -H "Authorization: Bearer <YOUR_TOKEN>"
