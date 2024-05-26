from fastapi import APIRouter, HTTPException

from app.simpleapi.models import Item
from app.simpleapi.dto import AddItem, ReadItem

router = APIRouter()


@router.post("/items", response_model=ReadItem)
def add_item(item: AddItem):
    item_obj = Item.objects.create(**item.dict())
    return ReadItem(**item_obj.__dict__)


@router.get("/items/{item_id}", response_model=ReadItem)
def read_item(item_id: int):
    try:
        item = Item.objects.get(pk=item_id)
    except Item.DoesNotExist:
        raise HTTPException(status_code=404, detail="Item with such id does not exist")
    return ReadItem(**item.__dict__)
