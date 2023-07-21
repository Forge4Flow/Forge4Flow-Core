# Permission

Permissions are the final building block for implementing role based access control. Permissions represent specific permissions or actions that can be taken within a system (ex. 'create a report', 'edit a report') and can be assigned directly to roles and/or users. Like roles, systems typically have a finite set of permissions that may change infrequently. The full representation of the `permission` object type is:

```
{
  "type": "permission",
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
          "ofType": "permission",
          "withRelation": "member"
        },
        {
          "inheritIf": "member",
          "ofType": "role",
          "withRelation": "member"
        }
      ]
    }
  }
}
```

### Data Integrity <a href="#data-integrity" id="data-integrity"></a>

Similar to users and tenants, Auth4Flow operates with 'strict data integrity' with respect to permissions. This means that you can only create warrants for permissions that exist in Auth4Flow.

However, unlike users and tenants, the expectation is that permissions are not already defined in your own system and so do not have to be 'synced' or transferred into Warrant. Instead, you can rely on the RBAC APIs, specifically the Permissions API, to manage permissions within Auth4Flow.
