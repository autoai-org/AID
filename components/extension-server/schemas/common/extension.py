from enum import Enum
from typing import Optional
from sqlmodel import Field, SQLModel

class ExtensionType(str, Enum):
    DOCKER = 'DOCKER'
    WEBHOOK = 'WEBHOOK'
    BINARY = 'BINARY'

class Extension(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    # condition is a small piece of code that can be eval by the interpreter
    # i.e., ext-server will run if(eval(condition)): ...
    # in the condition, the extension has access to the event object, i.e., they can access as e.event_type=='build_image'
    condition: str
    extension_type: ExtensionType
    # ext-server will create new process to execute the entrypoint. For different types, different format of entrypoint will be needed.
    # 1. for Docker image, it should be like: docker exec xxx
    # 2. for webhook, it should be a url address.
    # 3. for a local binary, it should be a shell command
    entrypoint: str

