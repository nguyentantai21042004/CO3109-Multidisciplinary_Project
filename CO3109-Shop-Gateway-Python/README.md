# Shop Gateway - Face Recognition System

## Giới thiệu
Shop Gateway là một thành phần trong hệ thống nhận diện khuôn mặt của dự án CO3109-Multidisciplinary_Project. Module này đóng vai trò như một gateway để xử lý và chuyển tiếp dữ liệu hình ảnh từ camera để nhận diện khuôn mặt khách hàng trong cửa hàng.

## Tính năng chính
- Nhận và xử lý hình ảnh từ camera
- Thực hiện nhận diện khuôn mặt
- Gửi kết quả nhận diện thông qua MQTT
- API endpoint cho việc kiểm tra khuôn mặt
- Lưu trữ hình ảnh đầu vào để kiểm tra

## Yêu cầu hệ thống
- Python 3.x
- Các thư viện Python được liệt kê trong `requirements.txt`:
  - Flask: Web framework
  - OpenCV: Xử lý hình ảnh
  - NumPy: Xử lý dữ liệu số
  - Paho-MQTT: Giao tiếp MQTT
  - Python-dotenv: Quản lý biến môi trường
  - Requests: Gửi HTTP requests

## Cài đặt
1. Clone repository
2. Cài đặt các thư viện cần thiết:
```bash
pip install -r requirements.txt
```

## Cấu trúc thư mục
```
CO3109-Shop-Gateway-Python/
├── Tools/              # Thư viện và công cụ hỗ trợ
├── image/              # Thư mục lưu trữ hình ảnh
├── main.py             # File chính để chạy ứng dụng
├── gateway.py          # Xử lý API và logic gateway
├── requirements.txt    # Danh sách thư viện cần thiết
└── README.md          # Tài liệu hướng dẫn
```

## Cách sử dụng
1. Chạy ứng dụng:
```bash
python main.py
```
2. Gateway sẽ chạy tại `http://localhost:5000/`
3. Sử dụng lệnh trong terminal:
   - Gõ 'start' để bắt đầu chụp và xử lý ảnh
   - Gõ 'exit' để thoát chương trình

## API Endpoints
### POST /check_face
- Nhận dữ liệu hình ảnh và ID cửa hàng
- Request body (JSON):
  ```json
  {
    "images": ["base64_encoded_image_1", "base64_encoded_image_2", ...],
    "shop_id": "shop_id_here"
  }
  ```
- Response:
  ```json
  {
    "result": "person_id_or_unknown"
  }
  ```
