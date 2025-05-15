
import os    
import cv2
import numpy as np
import requests
from io import BytesIO

    

def face_check(image, shop_id: str, test_mode=False) -> str:
    try:
        url = f"http://localhost:8085/api/v1/user/check-in/{shop_id}"

        if test_mode:
            # Lấy ảnh test từ file
            CURRENT_DIR = os.path.dirname(os.path.abspath(__file__))
            ROOT_DIR = os.path.abspath(os.path.join(CURRENT_DIR, "..", ".."))
            image_path = os.path.join(ROOT_DIR, "image", "avatar.png")

            with open(image_path, "rb") as f:
                files = {
                    "image_file": ("avatar.png", f, "image/png")
                }
                response = requests.post(url, files=files)
        else:
            # Encode ảnh từ numpy array (OpenCV)
            success, encoded_img = cv2.imencode('.jpg', image)
            if not success:
                print(" Không thể encode ảnh")
                return "failure"

            img_bytes = BytesIO(encoded_img.tobytes())
            img_bytes.name = "image.jpg"

            files = {
                "image_file": (img_bytes.name, img_bytes, "image/jpeg")
            }
            response = requests.post(url, files=files)

        if response.status_code == 200:
            result = response.json()
            return result.get("status", "unknown")
        else:
            print(" API lỗi mã:", response.status_code)
            return "failure"

    except Exception as e:
        print(f" Lỗi trong face_check: {e}")
        return "failure"
