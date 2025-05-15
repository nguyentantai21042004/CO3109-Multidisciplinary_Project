from dotenv import load_dotenv
import os
import paho.mqtt.publish as publish
import ssl

# Load biến môi trường từ file .env
load_dotenv()

# Lấy biến
AIO_USERNAME = os.getenv("AIO_USERNAME")
AIO_KEY = os.getenv("AIO_KEY")
AIO_HOST = os.getenv("AIO_HOST")
AIO_PORT = int(os.getenv("AIO_PORT"))

# Gửi MQTT message

def publish_result(result: str):
    # Reset all
    # client.publish("nhatminh1710/feeds/1", "0")
    # client.publish("nhatminh1710/feeds/Failure", "0")
    # client.publish("nhatminh1710/feeds/Unknow", "0")
    # wait_for_mqtt_connection()  
    
    if result == "success":
        # result_code = client.publish("nhatminh1710/feeds/chamcong123", "1")
        publish.single(
            f"{AIO_USERNAME}/feeds/chamcong123",
            payload="1",
            hostname=AIO_HOST,
            port=AIO_PORT,
            auth={'username': AIO_USERNAME, 'password': AIO_KEY}
        )

    elif result == "failure":
        publish.single(
            f"{AIO_USERNAME}/feeds/chamcong123",
            payload="-1",
            hostname=AIO_HOST,
            port=AIO_PORT,
            auth={'username': AIO_USERNAME, 'password': AIO_KEY}
        )
    else:
        publish.single(
            f"{AIO_USERNAME}/feeds/chamcong123",
            payload="0",
            hostname=AIO_HOST,
            port=AIO_PORT,
            auth={'username': AIO_USERNAME, 'password': AIO_KEY}
        )
