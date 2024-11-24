# myChat API

myChat の REST API

# Endpoints

## Authentication(認証)

- `GET /signup`
- `GET /login`

## Functions(機能)

- `POST /threads`: Create a thread. (Need authorized)
- `GET /threads?offset=&limit=`: Read thread list.
- `GET /threads/{threadUuid}`: Read a thread detail.

- `POST /posts`: Create a post. (Need authorized)
- `GET /posts/{threadUuid}?offset=&limit=`: Read post list replied to the thread.
