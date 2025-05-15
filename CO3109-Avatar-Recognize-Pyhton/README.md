# Avatar Recognition Service

## Giới thiệu
Avatar Recognition Service là một phần của dự án CO3109-Multidisciplinary_Project, cung cấp dịch vụ nhận diện khuôn mặt thông qua API. Service này sử dụng các công nghệ AI tiên tiến để nhận diện và phân tích khuôn mặt từ hình ảnh được gửi đến.

## Tính năng chính
- API endpoint cho nhận diện khuôn mặt
- Sử dụng DeepFace và MTCNN cho việc nhận diện khuôn mặt
- Tích hợp với Pinecone cho vector database
- Hỗ trợ GPU acceleration
- RESTful API với Flask
- Docker support cho deployment

## Yêu cầu hệ thống
- Python 3.x
- NVIDIA GPU (khuyến nghị)
- Docker và Docker Compose
- CUDA support (cho GPU acceleration)

## Cài đặt

### Sử dụng Python Virtual Environment
1. Tạo môi trường ảo:
```bash
python -m venv myenv
source myenv/bin/activate  # Linux/Mac
myenv\Scripts\activate     # Windows
```

2. Cài đặt dependencies:
```bash
pip install -r requirements.txt
```

### Sử dụng Docker
1. Build và chạy container:
```bash
docker-compose up --build
```

## Cấu trúc thư mục
```
CO3109-Avatar-Recognize-Pyhton/
├── src/                    # Mã nguồn chính
│   ├── controllers/        # Xử lý request
│   ├── services/          # Logic xử lý AI
│   └── app.py             # Entry point
├── test_img/              # Hình ảnh test
├── test_zone/             # Scripts test
├── requirements.txt       # Python dependencies
├── Dockerfile            # Docker configuration
├── docker-compose.yml    # Docker Compose configuration
└── README.md             # Documentation
```

## API Endpoints

### Base URL
```
http://localhost:8000/ai
```

### Endpoints
- POST `/recognize`: Nhận diện khuôn mặt từ hình ảnh
  - Request: Multipart form data với hình ảnh
  - Response: Kết quả nhận diện khuôn mặt

## Công nghệ sử dụng
- **Flask**: Web framework
- **DeepFace**: Framework nhận diện khuôn mặt
- **MTCNN**: Phát hiện khuôn mặt
- **Pinecone**: Vector database cho face embeddings
- **TensorFlow**: Deep learning framework
- **OpenCV**: Xử lý hình ảnh
- **Docker**: Containerization

## Phát triển
1. Clone repository
2. Cài đặt dependencies
3. Chạy server development:
```bash
flask run --debug --port 8000
```

## Deployment
Service có thể được triển khai bằng Docker:
```bash
docker-compose up -d
```

Service sẽ chạy trên port 8000 và sẵn sàng nhận requests.

## Testing
- Thư mục `test_img/` chứa các hình ảnh mẫu để test
- Thư mục `test_zone/` chứa các scripts test khác nhau
- Collection Postman được cung cấp trong file `DADN-AIBackend.postman_collection.json`

## Lưu ý
- Đảm bảo GPU và CUDA được cài đặt đúng cách nếu muốn sử dụng GPU acceleration
- Kiểm tra cấu hình môi trường trong file `.env` trước khi chạy
- Service yêu cầu đủ RAM và GPU memory cho việc xử lý model AI
