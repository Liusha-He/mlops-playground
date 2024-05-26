from pydantic import BaseModel
from datetime import datetime


class AddItem(BaseModel):
    title: str
    description: str


class ReadItem(BaseModel):
    id: int
    title: str
    description: str
    created_at: datetime
