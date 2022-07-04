import grpc
from proto.user.internal_pb2_grpc import UserInternalServiceStub
from proto.user.internal_pb2 import GetUserRequest, CreateUserRequest
from models.users import User, CreateUser

def build_stub() -> UserInternalServiceStub:
    channel = grpc.insecure_channel('localhost:5050')
    stub = UserInternalServiceStub(channel)

    return stub

class UserBaseHandler():
    def __init__(self):
        self.stub = build_stub()

class GetUserHandler(UserBaseHandler):
    def __init__(self):
        super().__init__()

    async def get_user(self, id: str) -> User:
        request_payload = GetUserRequest(user_id = id)
        response = self.stub.GetUser(request_payload)

        aliases = []

        for alias in response.payload.user.aliases:
            aliases.append(alias)

        return User(
            id=response.payload.user.id,
            name=response.payload.user.name,
            full_name=response.payload.user.full_name,
            aliases=aliases,
        )

class CreateUserHandler(UserBaseHandler):
    def __init__(self):
        super().__init__()

    async def create_user(self, req: CreateUser) -> User:
        request_payload = CreateUserRequest(
            name = req.name,
            full_name = req.full_name,
            aliases = req.aliases,
        )

        response = self.stub.CreateUser(request_payload)

        aliases = []

        for alias in response.payload.user.aliases:
            aliases.append(alias)

        return User(
            id=response.payload.user.id,
            name=response.payload.user.name,
            full_name=response.payload.user.full_name,
            aliases=aliases,
        )
