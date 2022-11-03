# Show Accessible Accounts

Fetches all todos currently held in-memory.

**URL**: `/todos`

**Method**: `GET`

**Data constraints**: `{}`

## Success Responses

**Condition**: All Todos, regardless of status, are returned.

**Code**: `200 OK`

```json
[
  {
    "id": "-kYOjgD4R",
    "label": "clean room",
    "completed": false,
    "priority": 0
  },
  {
    "id": "fzYdjRDVR",
    "label": "exercise",
    "completed": false,
    "priority": 1
  },
  {
    "id": "BzYdjgv4Rz",
    "label": "learn Go",
    "completed": true,
    "priority": 2
  }
]
```
