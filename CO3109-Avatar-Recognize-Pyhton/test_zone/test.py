import os
from dotenv import load_dotenv
import pickle
import json

load_dotenv()

from deepface import DeepFace

embedding_objs = DeepFace.represent(img_path = "image.jpg", detector_backend='ssd')

with open("embedding.json", "w") as f:
    f.write(json.dumps(embedding_objs, indent=4))

print(len(embedding_objs[0]['embedding']))

# pickle.dump(embedding_objs, open("embedding.pkl", "wb"))
