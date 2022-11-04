# Mark Todo Completed

Update the `completed` status of a Todo based on the given `id` if it exists.

**URL** : `/todos/:id`

**Method** : `PUT`

**Data constraints**

```json
{
  "completed": "[boolean—`true` if completed, `false` otherwise]"
}
```

**Data example**:

```json
{
  "completed": true
}
```

**`curl` copypasta**:

```sh
curl http://localhost:8080/todos/-testing- \
--include \
--header "Content-Type: application/json" \
--request "PUT" \
--data '{"completed": true}'
```

## Success Responses

**Code** : `200 OK`

**Content example** : For the example above, when `completed` is updated and
PUT to `/todos/-testing-`...

```json
{
  "id": "-testing-",
  "label": "clean room",
  "completed": true,
  "priority": 0
}
```

## Error Response

**Condition** : Todo does not exist with given `id`

**Code** : `404 NOT FOUND`

**Content** :

```json
{
  "message": "todo not found"
}
```

### Or

**Condition** : Request body does not contain valid `completed` field

**Code** : `400 BAD REQUEST`

**Content** :

## Notes

### Data Ignored

Endpoint will ignore irrelevant and read-only data such as parameters that
don't exist, or `id` and `label` fields which are not editable.
