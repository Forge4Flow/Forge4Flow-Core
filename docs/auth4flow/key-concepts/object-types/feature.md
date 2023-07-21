# Feature

Features represent the specific features within an application or system (ex. 'dashboard', 'report-builder') and can be used to implement feature flagging by user, tenant and/or pricing tier. The full representation of the `feature` object type is:

```
{
  "type": "feature",
  "relations": {
    "owner": {},
    "editor": {
      "inheritIf": "owner"
    },
    "viewer": {
      "inheritIf": "editor"
    },
    "member": {
      "inheritIf": "anyOf",
      "rules": [
        {
          "inheritIf": "member",
          "ofType": "feature",
          "withRelation": "member"
        },
        {
          "inheritIf": "member",
          "ofType": "pricing-tier",
          "withRelation": "member"
        }
      ]
    }
  }
}
```
