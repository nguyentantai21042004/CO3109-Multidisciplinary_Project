import os
import base64
import cv2
import numpy as np
import datetime

def decode_base64_image(img_b64):
    try:
        image_data = base64.b64decode(img_b64)
        image_np = np.frombuffer(image_data, np.uint8)
        img = cv2.imdecode(image_np, cv2.IMREAD_COLOR)
        return img
    except:
        return None

def save_image(img, shop_id, prefix="img"):
    folder = f"dataset/{shop_id}"
    os.makedirs(folder, exist_ok=True)
    timestamp = datetime.datetime.now().strftime("%Y%m%d_%H%M%S_%f")
    filename = os.path.join(folder, f"{prefix}_{timestamp}.jpg")
    cv2.imwrite(filename, img)
    print(f" Lưu ảnh vào: {filename}")
