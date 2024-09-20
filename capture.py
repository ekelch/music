import cv2

#open default cam
cam = cv2.VideoCapture(0)

# check to see if camera works
if not cam.isOpened():
    print("Cannot open camera")
    exit()

# Set camera dimensions
cam.set(3, 640)
cam.set(4, 480)

# Initialize face capture variables
count = 0

# Path to the Haar cascade file for face detection
face_cascade_Path = "haarcascade_frontalface_default.xml"

# Create a face cascade classifier
faceCascade = cv2.CascadeClassifier(face_cascade_Path)

while True:
    ret, frame = cam.read()

    # convert frame to gray scale (not sure why yet) https://github.com/medsriha/real-time-face-recognition/blob/master/face_taker.py
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    # Detect faces in the frame
    faces = faceCascade.detectMultiScale(gray, scaleFactor=1.3, minNeighbors=5)
    # Process each detected face
    for (x, y, w, h) in faces:
        # Draw a rectangle around the detected face
        cv2.rectangle(frame, (x, y), (x+w, y+h), (255, 0, 0), 2)

        # Increment the count for naming the saved images
        count += 1

        # display captured frame
        cv2.imshow('Camera', frame)

    # to break loop press q
    if cv2.waitKey(1) == ord('q'):
        break

# release capture and writer objects (free)
cam.release()
cv2.destroyAllWindows()
