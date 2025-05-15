# Sequence Diagrams

## 1. Save Employee Face Image Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant AC as AI Controller
    participant AS as AI Service
    participant DF as DeepFace
    participant PC as Pinecone DB

    C->>+AC: POST /ai/save/{shopID}/{employeeID}
    Note over C,AC: Multipart form data with image
    AC->>+AS: save(shopID, employeeID, image)
    AS->>AS: Convert image to NumPy array
    AS->>+DF: represent(image, detector='ssd')
    DF-->>-AS: Return face embedding vector
    AS->>+PC: upsert(namespace=shopID, vectors=[{id: employeeID, values: embedding}])
    PC-->>-AS: Confirmation
    AS-->>-AC: Success/Error
    AC-->>-C: 200 OK / Error Response
```

## 2. Find Employee by Face Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant AC as AI Controller
    participant AS as AI Service
    participant DF as DeepFace
    participant PC as Pinecone DB

    C->>+AC: POST /ai/find/{shopID}
    Note over C,AC: Multipart form data with image
    AC->>+AS: find(shopID, image)
    AS->>AS: Convert image to NumPy array
    AS->>+DF: represent(image, detector='ssd')
    DF-->>-AS: Return face embedding vector
    AS->>+PC: query(namespace=shopID, vector=embedding, top_k=1)
    PC-->>-AS: Return closest match
    Note over AS: Save response to pickle file
    alt Match found
        AS-->>AC: {user_id: employeeID}, 200
    else No match
        AS-->>AC: {"error": "No match found"}, 404
    end
    AC-->>-C: JSON Response
```

## 3. System Architecture Flow

```mermaid
sequenceDiagram
    participant C as Client Application
    participant API as REST API Layer
    participant S as AI Service Layer
    participant ML as ML Processing Layer
    participant DB as Vector Database

    C->>+API: HTTP Request
    API->>+S: Process Request
    S->>+ML: Face Detection & Embedding
    ML-->>-S: Feature Vector
    S->>+DB: Store/Query Vector
    DB-->>-S: Database Response
    S-->>-API: Service Response
    API-->>-C: HTTP Response
```

## API Endpoints

### 1. Save Employee Face
- **Endpoint**: POST `/ai/save/{shopID}/{employeeID}`
- **Input**: 
  - Path Parameters:
    - `shopID`: Shop identifier
    - `employeeID`: Employee identifier
  - Body: Multipart form data
    - `image`: Image file
- **Process**:
  1. Validate image input
  2. Convert image to NumPy array
  3. Generate face embedding using DeepFace
  4. Store embedding in Pinecone DB with shop namespace
- **Response**:
  - Success: 200 OK
  - Error: 400 Bad Request / 500 Server Error

### 2. Find Employee
- **Endpoint**: POST `/ai/find/{shopID}`
- **Input**:
  - Path Parameters:
    - `shopID`: Shop identifier
  - Body: Multipart form data
    - `image`: Image file
- **Process**:
  1. Validate image input
  2. Convert image to NumPy array
  3. Generate face embedding using DeepFace
  4. Query Pinecone DB for closest match
  5. Save response for debugging
- **Response**:
  - Success: 200 OK with employee ID
  - Not Found: 404 Not Found
  - Error: 400 Bad Request / 500 Server Error

## Technical Details

### Face Processing
- Using DeepFace library with SSD detector backend
- Face embedding generation for feature extraction
- Vector dimension based on DeepFace model output

### Vector Database
- Using Pinecone for vector similarity search
- Namespace separation by shop ID
- Top-1 nearest neighbor search for identification
- Vector upsert for employee registration

### Error Handling
- Image validation
- Face detection verification
- Database operation error handling
- Detailed error logging for debugging 