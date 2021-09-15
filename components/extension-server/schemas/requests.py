from schemas.common.extension import ExtensionType
from typing import List
from pydantic import BaseModel

class CreateExtensionRequest(BaseModel):
    condition: List[str] = None
    extension_type: ExtensionType
    remote_url: str
    entrypoint: str
    vendor: str
    name: str

class CreateEventRequest(BaseModel):
    event_type: str
    event_payload: dict