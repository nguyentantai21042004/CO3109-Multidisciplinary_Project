# CO3109 - Shop Gateway Python

Phần mềm này là một thành phần của dự án môn CO3109-Multidisciplinary Project, đảm nhiệm việc xử lý và so sánh khuôn mặt của nhân viên với cơ sở dữ liệu có sẵn.

## Tổng quan hệ thống

Hệ thống này hoạt động như một gateway để:
1. Nhận ảnh khuôn mặt từ các thiết bị
2. Xử lý và so sánh với dữ liệu khuôn mặt có sẵn của nhân viên
3. Gửi kết quả xác thực qua MQTT đến Adafruit

## Cấu trúc hệ thống

```
project_root/
├── main.py                 # File chính, chạy gateway + xử lý nhận diện
├── gateway.py              # Flask server nhận ảnh, gọi face_check
├── Tools/
│   ├── Gateway/
│   │   ├── face_check.py   # Module kiểm tra khuôn mặt với DB
│   │   ├── mqtt_client.py  # Kết nối và gửi MQTT đến Adafruit
│   │   └── image_utils.py  # Xử lý ảnh (giải mã, lưu trữ)
│   └── send_image/
│       ├── haarcascade_frontalface_default.xml  # Model phát hiện mặt
│       ├── mac_map.json    # Map MAC address → shopID
│       └── send_image.py   # Module gửi ảnh test từ webcam
├── dataset/               # Thư mục lưu ảnh tạm thời
├── .env                  # Cấu hình môi trường (MQTT, API keys)
└── requirements.txt      # Thư viện Python cần thiết

```

## Cài đặt và Cấu hình

### 1. Cài đặt thư viện

```bash
pip install -r requirements.txt
```

### 2. Cấu hình môi trường

Tạo file `.env` với các thông tin:

```
AIO_USERNAME=your_username
AIO_KEY=your_key
AIO_HOST=io.adafruit.com
AIO_PORT=1883
```

## Chạy hệ thống

1. Khởi động gateway:
```bash
python main.py
```

2. Gateway sẽ chạy tại `http://localhost:5000/` và sẵn sàng nhận requests

3. API Endpoints:
- POST `/check_face`: Nhận ảnh và shopID, trả về kết quả xác thực

## Luồng xử lý

1. Client gửi ảnh khuôn mặt và shopID đến gateway
2. Gateway xử lý ảnh và so sánh với database
3. Kết quả được gửi về client và đồng thời publish lên Adafruit qua MQTT
4. Các thiết bị khác có thể subscribe để nhận kết quả

## Testing

Sử dụng module `send_image.py` để test hệ thống:
1. Chạy main.py
2. Gõ 'start' để bắt đầu chụp và gửi ảnh test
3. Gõ 'exit' để thoát

## Đóng góp

Dự án này là một phần của môn CO3109-Multidisciplinary Project tại HCMUT.