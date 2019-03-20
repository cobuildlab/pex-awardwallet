# End Point Award

## Get account data
`GET /api/v1/award/account`

### Examples

`GET /api/v1/award/account`

### Example Error Body Response

```json
{
    "statusCode" : 0,
    "response" : {
        "errAward" : {
            "error" : ""
        }
    }
}
```

---

## Get connection link
`GET /api/v1/award/connection`

### Examples

`GET /api/v1/award/connection`

### Parameters

**platform** : _[required] Platform._ Type: string ; Default value: ""

**access** : _[required] Access._ Type: integer ; Default value: 0

### Example Body Response

```json
{
    "statusCode" : 0,
    "response" : {
        "url" : ""
    },
}
```