class BaseInterface(object):
    def __init__(self, name):
        self.data = {}
        self.name = name

class ObjectDetectionInterface(BaseInterface):
    # Image -> Coordinates
    pass 

class ImageSegmentationInterface(BaseInterface):
    # Image -> Pixel Map
    pass

class ImagePreprocessingInterface(BaseInterface):
    # Image -> Image
    pass
