import os
from dotenv import load_dotenv

load_dotenv()

from deepface import DeepFace

detector_backend = 'ssd'


metrics = ["cosine", "euclidean", "euclidean_l2"]

result = DeepFace.verify(
  img1_path = "./../test_img/norm.jpg", img2_path = "./../test_img/up.jpg", distance_metric = metrics[1], detector_backend = detector_backend
)

# dfs = DeepFace.find(
#   img_path = "img1.jpg", db_path = "C:/my_db", distance_metric = metrics[2]
# )

print("Is verified: ", result["verified"])
print("Distance: ", result["distance"])
