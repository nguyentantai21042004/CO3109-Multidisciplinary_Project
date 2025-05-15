# Sequence Diagrams - Shop Gateway System

## 1. Manual Image Capture Flow

```mermaid
sequenceDiagram
    participant U as User
    participant M as Main Program
    participant G as Gateway Server
    participant T as Tools/send_image
    participant AI as AI Service
    participant MQTT as MQTT Broker

    U->>M: Start Command
    M->>T: send_image()
    T->>T: Capture image from camera
    T->>G: POST /check_face
    Note over T,G: Image + Shop ID
    G->>AI: face_check(image, shop_id)
    AI-->>G: Recognition Result
    G->>MQTT: Publish Result
    G-->>T: Return Result
    T-->>M: Display Result
    M-->>U: Show Result
```

## 2. API Face Check Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant G as Gateway Server
    participant FC as Face Check
    participant IU as Image Utils
    participant MQTT as MQTT Client

    C->>+G: POST /check_face
    Note over C,G: {images: [...], shop_id: "..."}
    
    loop For each image
        G->>IU: decode_base64_image()
        IU-->>G: NumPy Array
        G->>IU: save_image()
        Note over G,IU: Save for debugging
        G->>FC: face_check(image, shop_id)
        FC-->>G: Recognition Result
    end

    G->>G: Process Results
    Note over G: Get most common result
    G->>MQTT: publish_result()
    G-->>-C: Return Result (200/404)
```

## 3. Gateway Startup Flow

```mermaid
sequenceDiagram
    participant M as Main Program
    participant G as Gateway Thread
    participant F as Flask Server
    participant MQTT as MQTT Client
    participant AI as AI Service

    M->>+G: Start Gateway Thread
    G->>F: Initialize Flask
    G->>F: Register Routes
    F->>F: Configure Host/Port
    G->>F: Start Server (0.0.0.0:5000)
    Note over G,F: Running in background

    M->>M: Wait 2 seconds
    M->>U: Display Command Menu
    
    loop Until exit
        U->>M: Input Command
        alt start command
            M->>T: Execute send_image()
        else exit command
            M->>M: Break loop
        end
    end
```

## 4. Face Check Processing Flow

```mermaid
sequenceDiagram
    participant FC as Face Check
    participant IU as Image Utils
    participant AI as AI Service
    participant MQTT as MQTT Client

    FC->>FC: Receive Image & Shop ID
    
    alt Debug Mode
        FC->>IU: Save Input Image
    end

    FC->>AI: Process Face Recognition
    Note over FC,AI: Using AI model
    
    AI-->>FC: Recognition Result
    
    FC->>MQTT: Publish Result
    FC-->>FC: Return Result
```

## System Components

### 1. Main Components
- **Gateway Server**: Flask-based HTTP server
- **Face Check Module**: Face recognition processing
- **Image Utils**: Image handling utilities
- **MQTT Client**: Communication with IoT devices
- **Tools**: Testing and utility scripts

### 2. Data Flow
1. **Image Input**
   - Camera capture
   - Base64 encoded images
   - Image file upload

2. **Processing**
   - Image decoding
   - Face recognition
   - Result aggregation

3. **Output**
   - MQTT publishing
   - HTTP response
   - Debug image saving

### 3. Communication Protocols
- **HTTP/REST**
  - POST /check_face
  - Multipart form data
  - JSON responses

- **MQTT**
  - Topic-based messaging
  - Result publishing
  - Device communication

### 4. Error Handling
- Image validation
- Face detection errors
- Connection issues
- Service errors

## API Specifications

### POST /check_face
- **Purpose**: Process face recognition request
- **Input**:
  ```json
  {
    "images": ["base64_image1", "base64_image2", ...],
    "shop_id": "shop_identifier"
  }
  ```
- **Output**:
  ```json
  {
    "result": "recognition_result"
  }
  ```
- **Status Codes**:
  - 200: Success
  - 400: Invalid input
  - 404: Face not found
  - 500: Server error

## Development Setup
1. **Environment**:
   - Python 3.x
   - Flask server
   - OpenCV
   - MQTT client

2. **Configuration**:
   - Server port: 5000
   - Debug mode
   - MQTT settings
   - Image storage paths

3. **Testing**:
   - Manual camera test
   - API endpoint testing
   - MQTT communication test 