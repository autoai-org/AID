from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {
        "message":"it works!"
    }

@app.get("/extensions")
async def get_all_extensions():
    return []

@app.get("/events")
async def get_all_events():
    return []

@app.post("/events")
async def post_events():
    return False

@app.post("/extensions")
async def post_extensions():
    return False