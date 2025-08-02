# API của app Công việc
API của app công việc cung cấp các api phục vụ quản lý cho các công việc của người dùng

## Công nghệ sử dụng
- Ngôn ngữ: Golang
- 
## Cấu trúc dự án
api/
├───cmd                     # Khởi chạy dự án
└───internal
    ├───entity              # các entity ánh xạ các bảng (collection) trong Database
    ├───handler             # Xử lý các HTTP Request
    ├───repository          # Tương tác với database hoặc nguồn dữ liệu
    └───service             # Xử lý nghiệp vụ (business logic)