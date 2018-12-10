import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="cvpm",
    version="0.0.1",
    author="Xiaozhe Yao",
    author_email="xiaozhe.yaoi@gmail.com",
    description="Computer Vision Package Manager",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/unarxiv/cvpm",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Intended Audience :: Developers",
        "Intended Audience :: Education",
        "Intended Audience :: Information Technology",
        "Intended Audience :: Science/Research",
        "Topic :: Scientific/Engineering",
        "Topic :: Scientific/Engineering :: Image Recognition",
        "Topic :: Software Development",
    ],
    install_requires=[
        "flask",
        "tqdm",
        "toml",
        "requests",
        "pillow",
        "numpy",
        "gevent",
        "flask-cors",
    ],
)