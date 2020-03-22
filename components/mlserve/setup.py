import setuptools

long_description = "For full description, please go to [Documentation](https://aid.autoai.org)"

setuptools.setup(
    name="mlpm",
    version="1.0.0.2",
    author="Xiaozhe Yao",
    author_email="xiaozhe.yaoi@gmail.com",
    description="Machine Learning Package Manager",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://aid.autoai.org",
    packages=['mlpm'],
    license="MIT",
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
        "sanic",
        "toml",
        "requests",
        "numpy",
        "werkzeug"
    ],
    project_urls={
        "Bug Tracker": "https://github.com/autoai-org/aid/issues",
        "Documentation": "https://aid.autoai.org",
        "Source Code": "https://github.com/autoai-org/aid",
    },
)