# Show all reports 

Show all Accounts the active User can access and with what permission level.
Includes their own Account if they have one.

**URL** : `/api/v1/reports`

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data** : `{}`

## Success Responses

**Condition** : User can see the report of the provided ID

**Code** : `200 OK`

**Content** :

```json
[
    {
        "ID": 1,
        "CreatedAt": "2021-12-15T14:19:04.698346+01:00",
        "UpdatedAt": "2021-12-15T15:46:31.255535+01:00",
        "DeletedAt": null,
        "serverName": "srv-test-01",
        "vulnerableFiles": [
            {
                "ID": 1,
                "CreatedAt": "2021-12-15T15:46:31.25599+01:00",
                "UpdatedAt": "2021-12-15T15:46:31.25599+01:00",
                "DeletedAt": null,
                "ReportID": 1,
                "fileName": "test1.txt"
            },
            {
                "ID": 2,
                "CreatedAt": "2021-12-15T15:46:31.25599+01:00",
                "UpdatedAt": "2021-12-15T15:46:31.25599+01:00",
                "DeletedAt": null,
                "ReportID": 1,
                "fileName": "text2.txt"
            }
        ]
    }
]
```

## Error Responses

**Condition** : If rport does not exist with `ID` of provided `id` parameter.

**Code** : `500 Internal Server Error`

**Content** : 

```json
{
    "message": "Report not found"
}
```