import os
from urllib.request import urlopen

import numpy as np
import PIL.Image
import requests
from tqdm import tqdm


def load_image_file(file, mode='RGB'):
    im = PIL.Image.open(file)
    if mode:
        im = im.convert(mode)
    return np.array(im)


class Downloader(object):
    def __init__(self):
        pass

    def download(self, url, target):
        # check if target folder exists
        if not os.path.isdir(target):
            os.makedirs(target)
        filename = url.split('/')[-1]
        file_size = int(urlopen(url).info().get('Content-Length', -1))
        chunk_size = 1024
        dest = os.path.join(target, filename)
        if os.path.exists(dest):
            first_byte = os.path.getsize(dest)
        else:
            first_byte = 0
        if first_byte >= file_size:
            return file_size
        header = {"Range": "bytes=%s-%s" % (first_byte, file_size)}
        pbar = tqdm(
            total=file_size,
            initial=first_byte,
            unit='B',
            unit_scale=True,
            desc=filename)
        req = requests.get(url, headers=header, stream=True)
        with open(dest, 'ab') as f:
            for chunk in req.iter_content(chunk_size=chunk_size):
                if chunk:
                    f.write(chunk)
                    pbar.update(1024)
        pbar.close()
        return file_size
