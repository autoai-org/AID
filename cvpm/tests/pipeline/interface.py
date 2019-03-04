import unittest

import numpy as np 
import PIL.Image

from cvpm.interface import ImageProcessingInterface
from cvpm.interface import ObjectDetectionInterface

def load_image_file(file, mode='RGB'):
    im = PIL.Image.open(file)
    if mode:
        im = im.convert(mode)
    return np.array(im)

class InterfaceTester(unittest.TestCase):
    def test_detect_interface(self):
        image_np = load_image_file('tests/assets/lenna.jpg')
        x_start, x_stop, y_start, y_stop = 12,13,14,15
        postDetectionInterface = ObjectDetectionInterface('post_detection')
        postDetectionInterface.addObjectCoordinates(x_start, x_stop, y_start, y_stop, 'face')
        print(postDetectionInterface.data)

if __name__ == "__main__":
    unittest.main()