# myChat API

myChat の REST API

# Endpoints

## Functions(機能)

- `POST /threads`: Create a thread. (Need authorized)
- `GET /threads?offset=&limit=`: Read thread list.
- `GET /threads/{uuid}`: Read a thread detail.

- WebSocket
- `/ws`: Send post and Recieve posts.

## Authentication(認証)

- `GET /signup`
- `GET /login`
