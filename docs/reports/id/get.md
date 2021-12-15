# Show a report

Show a single report by ID.

**URL** : `/api/v1/reports/:id`

**URL Parameters** : `id=[integer]` where `id` is the ID of the report.

**Method** : `GET`

**Auth required** : NO

**Permissions required** : None

**Data**: `{}`

## Success Responses

**Condition** : User can see one or more reports

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