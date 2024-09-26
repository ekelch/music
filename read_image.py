import cv2
import requests
import matplotlib.pyplot as plt
from deepface import DeepFace

# read image
img = cv2.imread("pics\happytest.jpg")

# display image
plt.imshow(img[:,:,:: -1])
plt.show()
