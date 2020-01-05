# Publish Packages

## Prepare Your Model Locally

### Project Structure

Your package is firstly a Python Package. Therefore, it will look like:

```
- example_package
    - example_package
        - solver.py
        - bundle.py
    - pretrained
        - pretrained.toml
    - tests // optional
    - examples // optional
    - cvpm.toml
    - requirements.txt
    - README.md
```

### Solvers

Solver is a ```class``` extended from ```cvpm.Solver``` class. In this class, you need to implement 2 functions: ```__init__``` and ```infer```. As their literal meanings, in the ```init``` function, parameters that your infer functions use should be initiate, and in the ```infer``` function, a image file path will be passed in the parameter, and the ```infer``` function is supposed to return your result in a dict or list.

For example, a simple solver frame will look like 

```python
from cvpm.solver import Solver

class SampleSolver(Solver):
    def __init__(self, toml_file=None):
        super().__init__(toml_file)
        # Do you Init Work here
        self.classifer = get_classifier()
        self.set_ready()
    def infer(self, image_file, config):
        # Use your own load_image method, or from cvpm.utility import load_image_file
        image = load_image(image_file)
        result = self.classifier(image)
        return result
```

### cvpm.toml

cvpm.toml is the entry file that tells the cvpm system which solver it should looking for. It uses [toml](https://github.com/toml-lang/toml) syntax and should looks like the following format:

``` toml
[package]
name="Your_Package_Name"

[[solvers]] 
name="Your_Solver_1_Name"
class="path/to_your_folder/Solver_Class_Name"

[[solvers]]
name="Your_Solver_2_Name"
class="path/to_your_folder/Solver_Class_Name"

[[solvers]]
name="Your_Solver_3_Name"
class="path/to_your_folder/Solver_Class_Name"
```

For example, the Face Utility that we provided as one of our official models have the [cvpm.toml](https://github.com/cvmodel/Face_Utility/blob/master/cvpm.toml) file as following:

``` toml
[package]
name="Face_utility"

[[solvers]]
name="Face_Detection"
class="face_utility/solver/FaceDetectionSolver"

[[solvers]]
name="Face_Landmark"
class="face_utility/solver/FaceLandmarkSolver"

[[solvers]]
name="Face_Encoding"
class="face_utility/solver/FaceEncodingSolver"
```

Make sure this file is located at the root path of your project.

### Pre-trained models

In most occasions, your algorithms (or neural networks etc.) should be uploaded along with some pre-trained models. We encourage you to upload it with our CLI (Not Done yet. [#72](https://github.com/unarxiv/CVPM/issues/72)), but it is also allowed if you upload it another downloadable services.

You need to define your pre-trained models in ```pretrained/pretrained.toml``` file as following using [toml syntax](https://github.com/toml-lang/toml):

``` toml
[[models]]
name = "model_file_name"
url = "https://www.example.com/your_model_file.zip" # Remote File

[[models]]
name = "model_file_name"
url = "pretrained/name_of_your_modelfile" # Local File
```

For example, the Face Utility that we provided as one of our official models have the [pretrained.toml](https://github.com/cvmodel/Face_Utility/blob/master/pretrained.toml) file as following:

``` toml
[[models]]
name = "dlib_face_recognition_resnet_model_v1"
url = "https://premium.file.cvtron.xyz/data/dlib_face_recognition_resnet_model_v1.dat"

[[models]]
name = "mmod_human_face_detector"
url = "https://premium.file.cvtron.xyz/data/mmod_human_face_detector.dat"

[[models]]
name = "shape_predictor_5_face_landmarks"
url = "https://premium.file.cvtron.xyz/data/shape_predictor_5_face_landmarks.dat"

[[models]]
name = "shape_predictor_68_face_landmarks"
url = "https://premium.file.cvtron.xyz/data/shape_predictor_68_face_landmarks.dat"
```

Make sure this file is located at the ```pretrained``` subfolder of your project.

### Pre/Post Install Scripts

It is an ongoing plan and have not been implemented yet. ([#49](https://github.com/unarxiv/CVPM/issues/49))

## Publish to Model Hub


## Publish to GitHub

When publishing to GitHub, mention "CVPM Available" at the top of your README file, only if you want you model to be searchable by the Model Hub.

Your README file should be look like the following:

```
## PROJECT NAME

[CVPM-Available](https://hub.autoai.org)
```

You can look at [Face Utility](https://github.com/cvmodel/Face_Utility) for reference as an official repository.

## Best Practice

TBD