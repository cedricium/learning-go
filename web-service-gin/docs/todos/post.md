# Create New Todo

Create a new Todo to be added to the in-memory todos.

**URL**: `/todos`

**Method**: `POST`

**Data constraints**

Provide label of Todo to be created.

```json
{
  "label": "[unicode 64 chars max]",
  "priority": 0 | 1 | 2 "integer denoting (LOW | MEDIUM | HIGH)"
}
```

**Data example** All fields must be sent.

```json
{
  "label": "build API"
}
```

**Copyable example**

```sh
$ curl http://localhost:8080/todos \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"label": "build api", "priority": 2}'
```

## Success Response

**Condition**: If everything is OK and an Todo didn't already exist.

**Code**: `201 CREATED`

**Content example**

```json
{
  "id": "5NtXRzvVg",
  "label": "build api",
  "completed": false,
  "priority": 2
}
```

## Error Responses

**Condition**: If fields are missing.

**Code**: `400 BAD REQUEST`

**Content example**

```json

```
