# Get Specific Todo By ID

Fetches a specific todo given the ID if it exists.

**URL**: `/todos/:id`

**URL Parameters**: `id=[string]` where `id` is the ID of the Todo on the
server.

**Method**: `GET`

**Data**: `{}`

## Success Responses

**Condition**: If Todo exists.

**Code**: `200 OK`

**Content example**:

```json
{
  "id": "rWM8Rkv4Rm",
  "label": "learn Go",
  "completed": true,
  "priority": 2
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
