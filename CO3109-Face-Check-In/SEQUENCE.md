# Sequence Diagrams - ESP32 Face Check-In Device

## 1. Device Initialization Flow

```mermaid
sequenceDiagram
    participant ESP as ESP32 Device
    participant WiFi as WiFi Network
    participant NTP as NTP Server
    participant MQTT as MQTT Broker
    participant DHT as DHT Sensor
    participant TFT as TFT Display

    ESP->>TFT: Initialize Display (240x320)
    ESP->>TFT: Set Rotation & SPI Speed
    ESP->>WiFi: Begin Connection
    Note over ESP,WiFi: SSID & Password
    WiFi-->>ESP: Connected
    ESP->>NTP: Configure Time (GMT+7)
    NTP-->>ESP: Time Sync
    ESP->>MQTT: Connect (AIO credentials)
    MQTT-->>ESP: Connected
    ESP->>MQTT: Subscribe to Feed Topic
    ESP->>DHT: Begin Sensor
    
    par Task Creation
        ESP->>ESP: Create MQTT Task (Core 1)
        ESP->>ESP: Create Sensor Task (Core 1)
        ESP->>ESP: Create Display Task (Core 1)
    end
```

## 2. Main Operation Flow

```mermaid
sequenceDiagram
    participant MQTT as MQTT Task
    participant Sensor as Sensor Task
    participant Display as Display Task
    participant Screen as TFT Screen
    participant LED as RGB LED
    participant DHT as DHT11 Sensor
    participant Broker as MQTT Broker

    loop Every 10ms
        MQTT->>Broker: Check Connection
        alt Disconnected
            MQTT->>Broker: Reconnect
        end
        MQTT->>Broker: MQTT Loop
    end

    loop Every 2 minutes
        Sensor->>DHT: Read Temperature
        alt Valid Reading
            Sensor->>Sensor: Update currentTemp
        end
    end

    loop Every 1 minute
        Display->>Screen: Update Normal Display
        Display->>Screen: Draw Time
        Display->>Screen: Draw Date
        Display->>Screen: Draw Temperature
        Display->>Screen: Draw WiFi Status
    end

    Broker-->>MQTT: Check-in Result
    MQTT->>Screen: Display Result Screen
    MQTT->>LED: Set LED Color
    Note over MQTT: Green for Success<br/>Red for Failure
    MQTT->>MQTT: Beep Sound
    MQTT->>Display: Return to Normal after 2s
```

## 3. Display Update Flow

```mermaid
sequenceDiagram
    participant D as Display Task
    participant TFT as TFT Screen
    participant Time as Time Service
    participant WiFi as WiFi Service
    participant Temp as Temperature Service

    D->>TFT: Clear Display
    D->>TFT: Draw Background Image
    
    par Time and Date
        D->>Time: Get Local Time
        Time-->>D: Current Time
        D->>TFT: Draw Time (HH:MM)
        D->>TFT: Draw Date (DD/MM/YYYY)
    end

    par Temperature
        D->>Temp: Get Current Temperature
        alt Temperature Available
            D->>TFT: Draw Temperature (°C)
        end
    end

    par WiFi Status
        D->>WiFi: Get RSSI
        WiFi-->>D: Signal Strength
        D->>TFT: Draw WiFi Bars
    end
```

## System Components

### Hardware Components
- ESP32 Development Board
- TFT ST7789 Display (240x320)
- DHT11 Temperature Sensor
- RGB LED
- Buzzer
- GPIO Connections:
  - TFT_CS: GPIO5
  - TFT_DC: GPIO16
  - TFT_RST: GPIO17
  - TFT_BL: GPIO4
  - TFT_SCK: GPIO18
  - TFT_MOSI: GPIO23
  - LED_R: GPIO19
  - LED_G: GPIO21
  - LED_B: GPIO22
  - BUZZER: GPIO27
  - DHT: GPIO14

### Software Tasks
1. **MQTT Task**
   - Priority: 1
   - Core: 1
   - Stack: 4096 bytes
   - Function: Handle MQTT communication

2. **Sensor Task**
   - Priority: 1
   - Core: 1
   - Stack: 2048 bytes
   - Function: Read temperature every 2 minutes

3. **Display Task**
   - Priority: 1
   - Core: 1
   - Stack: 4096 bytes
   - Function: Update display every minute

### Network Configuration
- WiFi Connection
- NTP Time Sync (GMT+7)
- MQTT Connection to Adafruit IO
  - Port: 1883
  - Subscribe to feed topic for check-in results

### Display States
1. **Normal Display**
   - Time (top-left)
   - Date (top-right)
   - Temperature (bottom-right)
   - WiFi signal strength (bottom-left)
   - Background image

2. **Result Display**
   - Success/Failure image
   - LED indication (Green/Red)
   - Buzzer feedback
   - Auto-return to normal after 2 seconds 