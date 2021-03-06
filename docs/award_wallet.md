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

### Parameters

**platform** : _[required] Platform._ Type: string ; Default value: ""

**access** : _[required] Access._ Type: integer ; Default value: 0

**state** : _[required] State._ Type: string ; Default value: ""

**granularSharing** : _[required] Granular Sharing._ Type: bool ; Default value: false

### Examples

`GET /api/v1/award/connection?platform=mobile&access=1&state=xyz&granularSharing=true`

### Example Body Response

```json
{
    "statusCode" : 0,
    "response" : {
        "url" : ""
    },
}
```

---

## Get data account
`GET /api/v1/award/account`

### Parameters

**userID** : _[required] User ID._ Type: integer ; Default value: 0

### Examples

`GET /api/v1/award/account?userID=57342`
