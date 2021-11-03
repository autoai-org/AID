# Copyright (c) 2021 Xiaozhe Yao
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
import os
import numpy as np
from PIL import Image
import json

def load_image(image_path, resize_shape=None):
    if not os.path.exists(image_path):
        raise ValueError("No such file: "+image_path)
    img = Image.open(image_path)
    if resize_shape is not None:
        img = img.resize(resize_shape)
    try:
        data = np.asarray(img, dtype='uint8')
    except SystemError:
        data = np.asarray(img.getdata(), dtype='uint8')
    return data

def load_image_dataset(root_path, resize_shape):
    # read annotation file first
    # the format is at https://aid.autoai.org/#/guide/specs/data-format
    anno_file = os.path.join(root_path, "annotations.json")
    with open(anno_file, 'r') as file:
        anno_data = json.load(file)
    train_data = []
    test_data = []
    val_data = []
    for each in anno_data:
        img_path = os.path.join(root_path, each['folder'], each['filename'])
        data = load_image(img_path, resize_shape)
        if (each['folder']=='train'):
            train_data.append(data)
        elif(each['folder']=='test'):
            test_data.append(data)
        else:
            val_data.append(data)
    return np.array(train_data), np.array(test_data), np.array(val_data)