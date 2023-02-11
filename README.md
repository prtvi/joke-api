# Joke API using Golang and MongoDB

A simple API built using the Echo framework in Golang.

## USING THE API

- [Getting started](#getting-started) introduces you to the operations offered by the API.
- [API calls](#api-calls) gives you examples of those operations

## Types of jokes

`type`: identifier of group/type of joke

| #   | Type        |
| --- | ----------- |
| 1   | General     |
| 2   | Knock-Knock |
| 3   | Programming |

# Getting Started

| Endpoint                                  |                                        What it does                                        |
| ----------------------------------------- | :----------------------------------------------------------------------------------------: |
| `GET` [`/random`](#randomjokes)           |        Returns a joke object that contains a `setup`, `punchline`, `type` and `id`         |
| `GET` [`/random/:n`](#randomjokescount)   |                        Returns an array with up to 44 joke objects.                        |
| `GET` [`/joke/:id`](#jokesid)             |                         Returns a joke object with a specific id.                          |
| `POST` [`/new`](#jokescreate)             |      Provided a `setup`, `punchline` and `type` it will insert the joke into the db.       |
| `DELETE` [`/remove/:id`](#jokeremoveid)   |                              Remove a joke with the given id                               |
| `UPDATE` [`/update/:id`](#randomtypetype) | Provided a new `setup` or `punchline` or `type` it will update the joke with the given id. |

# API calls

This API supports a data response in JSON format.

### /random

```json
{
  "id": 43,
  "type": "general",
  "punchline": "To prove that he was framed!",
  "setup": "Why did the burglar hang his mugshot on the wall?"
}
```

### /random/:count

`count`: number of jokes to be fetched

`/random/2`

```json
[
  {
    "id": 8,
    "type": "general",
    "punchline": "To get to the other slide.",
    "setup": "Why did the kid cross the playground?"
  },
  {
    "id": 40,
    "type": "general",
    "punchline": "Lots of training",
    "setup": "How do locomotives know where they're going?"
  }
]
```

### /joke/:id

`id`: unique identifier for a joke

`/joke/1`

```json
[
  {
    "id": 1,
    "type": "general",
    "setup": "What did the fish say when it hit the wall?",
    "punchline": "Dam."
  }
]
```

### /new

`Post request body:`

```json
{
  "setup": "this is a test",
  "punchline": "this is a funny test",
  "type": "test"
}
```

`Response`

```json
{
  "responseMsg": "Success",
  "statusCode": 200,
  "joke": {
    "id": 134,
    "setup": "this is a test",
    "punchline": "this is a funny test",
    "type": "test"
  }
}
```

### /remove/:id

`/remove/3`

```json
{
  "statusCode": 200,
  "message": "Delete operation successful!"
}
```

### /update/:id

`/update/3`

`Request body`

```json
{
  "type": "general",
  "punchline": "Guilty",
  "setup": "Do I enjoy making courthouse puns?"
}
```

`Response`

```json
{
  "responseMsg": "Success",
  "statusCode": 200,
  "joke": {
    "id": 3,
    "type": "general",
    "punchline": "Guilty",
    "setup": "Do I enjoy making courthouse puns?"
  }
}
```
