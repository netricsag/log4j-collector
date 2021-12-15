# Delete a report

Delete a specific report by ID.

**URL** : `/api/v1/reports/:id`

**URL Parameters** : `id=[integer]` where `id` is the ID of the report.

**Method** : `DELETE`

**Auth required** : NO

**Permissions required** : None

**Data** : `{}`

## Success Response

**Condition** : If the report exists.

**Code** : `200 OK`

**Content** :

```json
{
    "message": "Report deleted"
}
```

## Error Responses

**Condition** : If there was no report available to delete.

**Code** : `500 Internal Server Error`

**Content** : 

```json
{
    "message": "Report not found"
}
```

## Notes

* Will remove all VulnerableFiles associated with the report.