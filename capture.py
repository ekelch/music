import cv2

#open default cam
cam = cv2.VideoCapture(0)

# check to see if camera works
if not cam.isOpened():
    print("Cannot open camera")
    exit()

while True:
    ret, frame = cam.read()

    # display captured frame
    cv2.imshow('Camera', frame)

    # to break loop press q
    if cv2.waitKey(1) == ord('q'):
        break

# release capture and writer objects (free)
cam.release()
cv2.destroyAllWindows()
