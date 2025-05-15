import cv2, base64, requests, time
import os
import uuid, json


BASE_DIR = os.path.dirname(os.path.abspath(__file__))
NUM_FRAMES = 3
FRAME_DELAY = 0.2  # 200ms giữa mỗi ảnh
CASCADE_PATH = os.path.join(BASE_DIR, "haarcascade_frontalface_default.xml")

# files = {
#     "images": open("image/avatar.png", "rb") 
# }

# data = {
#     "shop_id": "c4596303-de42-424b-afcb-ea5be63ab060"
# }
shop_id = "c4596303-de42-424b-afcb-ea5be63ab060"

def send_image ():
    face_cascade = cv2.CascadeClassifier(CASCADE_PATH)
    cap = cv2.VideoCapture(0)
    frames = []
    # cv2.imshow("Camera", cap.read())

    print(f"Đang lấy ảnh từ camera...")

    for i in range(NUM_FRAMES):
        ret, frame = cap.read()
        if not ret:
            print(f"Không lấy được ảnh {i+1}")
            continue
        
        gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
        faces = face_cascade.detectMultiScale(gray, scaleFactor=1.1, minNeighbors=5)
        
        for (x, y, w, h) in faces:
            # Vẽ khung lên màn hình
            cv2.rectangle(frame, (x, y), (x + w, y + h), (0, 255, 0), 2)

            # Cắt ảnh mặt
            face_img = frame[y:y + h, x:x + w]
            _, buffer = cv2.imencode('.jpg', face_img)
            img_b64 = base64.b64encode(buffer).decode()
            frames.append(img_b64)
            
        cv2.imshow("Check_face", frame)
        if cv2.waitKey(1) == ord('q'):
            break
        time.sleep(FRAME_DELAY)

    cap.release()
    cv2.destroyAllWindows()

    print("Gửi ảnh lên Gateway...")

    payload = {
        "images": frames,
        "shop_id": shop_id
    }

    res = requests.post("http://localhost:5000/check_face", json=payload)
    print("Phản hồi từ Gateway:", res.json())
    return res
