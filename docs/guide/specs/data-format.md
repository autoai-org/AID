# Data Formats

When you are using CVPM for training, CVPM requires you to specify the file path of the training, validation and test dataset. In this chapter, we will introduct how the dataset should  be organized.

## Overall

You dataset is a zip file (Do not use .rar or other format since it's not free format). You can also use .7z file.

Inside the file, your dataset should looks like:

```
- train
    - img_0001.jpg (filenames can be changed if needed)
    - img_0002.png 
- test
    - img_0001.jpg (filenames can be changed if needed)
    - img_0002.png
- label_map.txt (if needed)
- annotations.json
```

```annotations.json``` is required for cvpm to understand your dataset. After the uncompression of your zip file, cvpm will check if this file exists, if it does not exist, cvpm will mark it as warning.

```annotations.json``` is a list which looks like:

``` json
[
    anno_obj_1,
    anno_obj_2
    ...
]
```

Since there are many different types of computer vision tasks, we hereby define several common data formats for the task.

## Classification

``` json
{
    "folder": "train_or_test_or_val",
    "filename": "img_0001.jpg",
    "size": { // not required
        "width": 600,
        "height": 400,
        "depth":3
    },
    "class": {
        "label": "class_label"
    }
}
```

## Detection

```json
{
    "folder": "train_or_test_or_val",
    "filename": "img_0001.jpg",
    "size": { // not required
        "width": 600,
        "height": 400,
        "depth":3
    },
    "boundbox": [{
        "label": "class_label",
        "xmin": "xmin",
        "ymin": "ymin",
        "xmax": "xmax",
        "ymax": "ymax"
    }]
}
```

## Segmentation

``` json
{
    "folder": "train_or_test_or_val",
    "filename": "img_0001.jpg",
    "size": { // not required
        "width": 600,
        "height": 400,
        "depth":3
    },
    "segmentation": {
        "label": "path_to_label_map_file"
    }
}
```

## Image-to-Image

``` json
{
    "folder": "train_or_test_or_val",
    "filename": "img_0001.jpg",
    "size": { // not required
        "width": 600,
        "height": 400,
        "depth":3
    },
    "translation": {
        "target": "path-to-target-image-file"
    }
}
```