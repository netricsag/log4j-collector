# Update a specific report

Update a specific report with the new values provided. You can update everything or just a specific fields in the report. But the vulnerable Files will be flushed and set with the new values.

**URL** : `/api/v1/reports/:id`

**URL Parameters** : `id=[integer]` where `id` is the ID of the report.

**Method** : `POST`

**Auth required** : NO

**Permissions required** : None

**Data constraints** :

```json
{
    "servername": "[unicode 64 chars max]",
    "vulnerableFiles": 
    [
        {
            "fileName": "[unicode 64 chars max]"
        }       
    ]
}
```

**Data example** :

```json
{
    "servername": "srv-test-01",
    "vulnerableFiles": 
    [
        {
            "fileName": "test.txt"
        },
        {
            "filename": "text2.txt"
        }
        
    ]
}
```

## Success Response

**Condition** : If everything is OK and the Report already exists and gets updated.

**Code** : `200 OK`

**Content example**

```json
{
    "ID": 1,
    "CreatedAt": "2021-12-15T18:14:40.70316+01:00",
    "UpdatedAt": "2021-12-15T18:14:40.70316+01:00",
    "DeletedAt": null,
    "serverName": "srv-test-01",
    "vulnerableFiles": [
        {
            "ID": 1,
            "CreatedAt": "2021-12-15T18:14:40.704016+01:00",
            "UpdatedAt": "2021-12-15T18:14:40.704016+01:00",
            "DeletedAt": null,
            "ReportID": 1,
            "fileName": "test.txt"
        },
        {
            "ID": 2,
            "CreatedAt": "2021-12-15T18:14:40.704016+01:00",
            "UpdatedAt": "2021-12-15T18:14:40.704016+01:00",
            "DeletedAt": null,
            "ReportID": 1,
            "fileName": "text2.txt"
        }
    ]
}
```

### Or

**Condition** : If you send only the servername without the vulnerableFiles, the report will be deleted.

**Code** : `200 OK`

**Content example**

```json
{
    "message": "Report deleted"
}
```

## Error Responses

**Condition** : If ID is not found.

**Code** : `500 Internal Server Error`

**Content** : 
```json
{
    "message": "Report not found"
}
```

### Or

**Condition** : If JSON parsing fails.

**Code** : `500 Internal Server Error`

**Content** :

```json
{
    "message": "Error parsing request body"
}
```