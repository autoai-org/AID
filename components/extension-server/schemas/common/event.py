from enum import Enum
from typing import Dict, Optional, Union
from sqlmodel import Field, SQLModel, create_engine

# https://stackoverflow.com/questions/65209934/pydantic-enum-field-does-not-get-converted-to-string
class EventType(str, Enum):
    BUILD_IMAGE = 'build_image'
    CREATE_CONTAINER = 'create_container'

class Event(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    event_type: EventType
    event_payload: str
    event_status: Optional[int] = Field(default=None)

