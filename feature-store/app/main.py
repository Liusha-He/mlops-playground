from django.core.wsgi import get_wsgi_application
from fastapi import FastAPI


def get_fastapi_application() -> FastAPI:
    from app.simpleapi.routers import router as simpleRouter
    
    api = FastAPI()
    api.include_router(simpleRouter, 
                       prefix="/api/v1",
                       tags=["simple-api"])
    return api


app = get_wsgi_application()
api = get_fastapi_application()
