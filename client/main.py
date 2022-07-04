from fastapi import FastAPI
from handlers.users import GetUserHandler, CreateUserHandler
from models.users import User, CreateUser

app = FastAPI()

get_user_handler = GetUserHandler()
create_user_handler = CreateUserHandler()


@app.get("/users/{user_id}", response_model=User)
async def get_user(user_id: str):
    return await get_user_handler.get_user(user_id)

@app.post("/users",  response_model=User)
async def create_user(body: CreateUser):
    return await create_user_handler.create_user(body)