import depthai as dai
import cv2

BLOB_PATH = "./best_openvino_2022.1_6shave.blob"
CLASS_NAMES = ["coca_cola_500ml", "monster_energy_drink_mango_loco", "kyknos_ketchup_classic"]
CONFIDENCE = 0.5

pipeline = dai.Pipeline()

# Camera
cam = pipeline.create(dai.node.ColorCamera)
cam.setPreviewSize(640, 640)
cam.setInterleaved(False)
cam.setFps(30)

# YOLO detection network
nn = pipeline.create(dai.node.YoloDetectionNetwork)
nn.setBlobPath(BLOB_PATH)
nn.setConfidenceThreshold(CONFIDENCE)
nn.setNumClasses(len(CLASS_NAMES))
nn.setCoordinateSize(4)
nn.setIouThreshold(0.5)
nn.input.setBlocking(False)

# Outputs
cam_out = pipeline.create(dai.node.XLinkOut)
cam_out.setStreamName("rgb")
nn_out = pipeline.create(dai.node.XLinkOut)
nn_out.setStreamName("detections")

# Links
cam.preview.link(nn.input)
cam.preview.link(cam_out.input)
nn.out.link(nn_out.input)

with dai.Device(pipeline) as device:
    rgb_q = device.getOutputQueue("rgb", maxSize=4, blocking=False)
    det_q = device.getOutputQueue("detections", maxSize=4, blocking=False)

    print("Camera running — press Q to quit")

    while True:
        frame = rgb_q.get().getCvFrame()
        detections = det_q.get().detections

        for det in detections:
            h, w = frame.shape[:2]
            x1, y1 = int(det.xmin * w), int(det.ymin * h)
            x2, y2 = int(det.xmax * w), int(det.ymax * h)
            label = CLASS_NAMES[det.label]
            conf = det.confidence

            cv2.rectangle(frame, (x1, y1), (x2, y2), (0, 255, 0), 2)
            cv2.putText(frame, f"{label} {conf:.0%}",
                        (x1, y1 - 8), cv2.FONT_HERSHEY_SIMPLEX, 0.6, (0, 255, 0), 2)

        cv2.imshow("STEVE Cart — OAK-D Lite", frame)
        if cv2.waitKey(1) == ord("q"):
            break
