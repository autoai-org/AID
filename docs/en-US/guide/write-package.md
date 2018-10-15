# Write a Package

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

``` python
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

### Bundle
    
::: tip 注意 请确保你的 Node.js 版本 >= 8。 :::



### cvpm.toml

### Pre-trained models

## Publish to Model Hub

## Publish to GitHub

When publishing to GitHub, mention "CVPM Available" at the top of your README file, only if you want you model to be searchable by the Model Hub.

Your README file should be look like

```
## PROJECT NAME

[CVPM-Available](https://hub.autoai.org)
```

You can look at [Face Utility](https://github.com/cvmodel/Face_Utility) for reference that we provided as an official repository.

## Best Practice