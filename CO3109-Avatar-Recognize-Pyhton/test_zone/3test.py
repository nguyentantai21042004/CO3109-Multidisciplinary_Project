# FIND EUCLIDEAN DISTANCE BETWEEN TWO FACE REPRESENTATIONS

from deepface.commons import distance as dst

distance = dst.findEuclideanDistance(img1_representation, img2_representation)
threshold = dst.findThreshold(model_name, 'euclidean')

if distance <= threshold:
	identified = True
else:
	identified = False