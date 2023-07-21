# Role

Roles are one of three building blocks for implementing role based access control. They can be thought of as 'containers' or 'groups' of (typically) users. In most rbac implementations, the set of roles is finite and usually based on some organizational structure and/or role (ex. admin, owner, member etc.). The `role` object type in Warrant has four relations: `owner` (direct relation), `editor`, `viewer` and `member`. The full representation of the object type is:

```
{
  "type": "role",
  "relations": {
    "owner": {},
    "editor": {
      "inheritIf": "owner"
    },
    "viewer": {
      "inheritIf": "editor"
    },
    "member": {
      "inheritIf": "member",
      "ofType": "role",
      "withRelation": "member"
    }
  }
}
```

### Data Integrity <a href="#data-integrity" id="data-integrity"></a>

Similar to users and tenants, Auth4Flow operates with 'strict data integrity' with respect to roles. This means that you can only create warrants for roles that exist in Auth4Flow.

However, unlike users and tenants, the expectation is that roles are not already defined in your own system and so do not have to be 'synced' or transferred into Auth4Flow. Instead, you can rely on the RBAC APIs, specifically the Roles API, to manage roles within Auth4Flow.
