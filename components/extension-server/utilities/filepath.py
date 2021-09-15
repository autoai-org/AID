def get_home_dir():
    """
    Returns the home directory of the aid project.
    """
    import os
    homedir = os.path.expanduser("~")
    return os.path.join(homedir, ".autoai", "aid")
