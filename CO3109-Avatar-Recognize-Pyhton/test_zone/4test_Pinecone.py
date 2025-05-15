# PINECONE DB SETUP
from pinecone import Pinecone

from dotenv import load_dotenv
import os

load_dotenv()
PINECONE_KEY = os.getenv("PINECONE_KEY")


# Initialize a Pinecone client with your API key
pc = Pinecone(api_key=PINECONE_KEY)
index = pc.Index(host="https://dadn-deepface-bfhi09b.svc.aped-4627-b74a.pinecone.io")

# DEEPFACE SETUP
from deepface import DeepFace

detector_backend = 'ssd'

embedding_objs = DeepFace.represent(img_path = "image.jpg", detector_backend=detector_backend)
embedding_vector = embedding_objs[0]['embedding']

# PINECONE UPLOAD
# Insert the embedding vector into the Pinecone index
vectors = [
    {
        "id": "1",
        "values": embedding_vector
    }
]

index.upsert(vectors)
