'''
@Author: Xiaozhe Yao
Interface is the definition of data structure between two solvers.
'''
class BaseInterface(object):
    def __init__(self, name):
        self.name = name
        self._data = {}
    @property
    def data(self):
        return self._data

class ObjectDetectionInterface(BaseInterface):
    # Image -> Coordinates
    # Data: {'x_start':Integer, 'x_stop':Integer, 'y_start':Integer, 'y_stop':Integer}
    def __init__(self,name):
        super().__init__(name)
        self.data['data'] = []
    def validate(self):
        pass
    def addObjectCoordinates(self, x_start, x_stop, y_start, y_stop, name):
        self.data['data'].append({
            "x_start": x_start,
            "x_stop": x_stop,
            "y_start": y_start,
            "y_stop": y_stop,
            "name": name
        })

class ImageSegmentationInterface(BaseInterface):
    # Image -> Pixel Map
    pass

class ImageProcessingInterface(BaseInterface):
    # Image -> Image
    pass
