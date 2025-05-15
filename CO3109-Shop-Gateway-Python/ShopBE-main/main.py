from gateway import *
from Tools.send_image.send_image import *
import cv2
import base64
import time
from threading import Thread

def main():
    gateway_thread = Thread(target=start_gateway, daemon=True)
    gateway_thread.start()
    time.sleep(2)
    print("Gõ 'start' để bắt đầu | 'exit' để thoát.")
    while True:
        cmd = input("Nhập lệnh: ").strip().lower()
        if cmd == "start":
            result = send_image()
            print("Kết quả từ Gateway:", result)
        elif cmd == "exit":
            print("Thoát chương trình.")
            break
        else:
            print("Lệnh không hợp lệ.")

if __name__ == "__main__":
    main()