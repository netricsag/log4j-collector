# Create or Update a report

If ServerName doesn't exist, create it will be created. If ServerName exists, it will be updated with the new values.
All the old values will be flushed.

**URL** : `/api/v1/reports`

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

**Condition** : If everything is OK or the Report already exists and gets updated.

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

**Condition** : If you send only the servername without the vulnerableFiles, the linked vulnerable files will be deleted, but the report with the servername will be kept.

**Code** : `200 OK`

**Content example**

```json
{
    "message": "Report deleted"
}
```

## Error Responses

**Condition** : If ServerName is not provided.

**Code** : `500 Internal Server Error`

**Content** : 
```json
{
    "message": "Server name cannot be empty"
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