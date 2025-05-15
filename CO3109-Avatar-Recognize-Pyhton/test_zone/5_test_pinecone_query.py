# PINECONE DB SETUP
from pinecone import Pinecone

from dotenv import load_dotenv
import os

load_dotenv()
PINECONE_KEY = os.getenv("PINECONE_KEY")
INDEX_HOST = os.getenv("INDEX_HOST")

# Initialize a Pinecone client with your API key
pc = Pinecone(api_key=PINECONE_KEY)
index = pc.Index(host=INDEX_HOST)

# DEEPFACE SETUP
from deepface import DeepFace

detector_backend = 'ssd'

embedding_objs = DeepFace.represent(img_path = "./test_zone./test_img/image.jpg", detector_backend=detector_backend)
embedding_vector = embedding_objs[0]['embedding']


response = index.query(
    vector=embedding_vector,
    top_k=5,
    include_values=True
)

print(response)