from lib2to3.pytree import Base
from typing import List
from pydantic import BaseModel

class User(BaseModel):
    id: str
    name: str
    full_name: str
    aliases: List[str]

class CreateUser(BaseModel):
    name: str
    full_name: str
    aliases: List[str]
