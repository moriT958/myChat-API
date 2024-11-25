# myChat API

myChat の REST API

# Endpoints

## Functions(機能)

- `POST /threads`: Create a thread. (Need authorized)
- `GET /threads?offset=&limit=`: Read thread list.
- `GET /threads/{uuid}`: Read a thread detail.

- `POST /posts`: Create a post. (Need authorized)
- `GET /posts/{threadUuid}?offset=&limit=`: Read post list replied to the thread.

## Authentication(認証)

- `GET /signup`
- `GET /login`
