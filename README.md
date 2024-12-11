# myChat API

myChat の REST API

# Endpoints

## Basic Functions(機能)

- `POST /threads`: Create a thread. (Need authorized)
- `GET /threads?offset=&limit=`: Read thread list.
- `GET /threads/{uuid}`: Read a thread detail.

- `GET /users/{username}`: Read a user data.
- `POST /users`: Create user data.

- WebSocket
- `/ws`: Send post and Recieve posts.

## Authentication(認証): Incomplete

- `GET /signup`
- `GET /login`
