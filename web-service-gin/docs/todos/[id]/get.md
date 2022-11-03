# Get Specific Todo By ID

Fetches a specific todo given the ID if it exists.

**URL**: `/todos/:id`

**URL Parameters**: `id=[string]` where `id` is the ID of the Todo on the
server.

**Method**: `GET`

**Data**: `{}`

**`curl` copypasta**:

```sh
$ curl http://localhost:8080/todos/-testing-
```

## Success Responses

**Condition**: If Todo exists.

**Code**: `200 OK`

**Content example**:

```json
{
  "id": "-testing-",
  "label": "clean room",
  "completed": false,
  "priority": 0
}
```

## Error Responses

**Condition**: If Todo does not exist with provided `id` parameter.

**Code**: `404 NOT FOUND`

**Content**:

```json
{
  "message": "todo not found"
}
```
