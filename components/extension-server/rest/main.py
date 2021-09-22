from runner.pipeline import Pipeline
from storage.database import DB
from schemas.requests import CreateEventRequest, CreateExtensionRequest
from utilities.filepath import get_home_dir
from fastapi import FastAPI
from schemas.common.event import Event
from schemas.common.extension import Extension
from typing import List

app = FastAPI()
pipe = Pipeline()
@app.get("/")
async def root():
    return {
        "message":"it works!"
    }

@app.get("/events")
async def get_all_events():
    return []

@app.post("/events")
async def post_events(request: CreateEventRequest):
    event = Event(event_type=request.event_type, event_payload=str(request.event_payload))
    event.event_status = "Added"
    pipe.run_event_through_all_extensions(event)
    return False

@app.get("/extensions", response_model=List[Extension])
async def get_all_extensions():
    return DB().fetch_extensions()

@app.post("/extensions", response_model=Extension)
async def create_extension(request: CreateExtensionRequest):
    ext = Extension(
        **request.__dict__
    )
    ext.condition = str(request.condition)
    database = DB()
    ext = database.save(ext)
    return ext


