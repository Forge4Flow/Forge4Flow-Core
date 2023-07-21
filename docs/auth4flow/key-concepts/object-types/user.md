# User

In many cases, applications require access control rules to be defined per user. To make this easier, Auth4Flow comes with a built-in `user` object type. This object type has one `parent` relation which makes it easy to associate users to a parent object such as a tenant, a parent user or an account (in a B2B context). The full representation of the user object type is:

```json
{
  "type": "user",
  "relations": {
    "parent": {
      "inheritIf": "parent",
      "ofType": "user",
      "withRelation": "parent"
    }
  }
}
```

Using this object type, we can create warrants for individual users like:

```
[user:7] is a [parent] of [user:84]
[tenant:A] is a [parent] of [user:1]
```

### Auto Registration vs Manual Registration

Auth4Flow provides two methods to register users, the first is Auto Registration which will automatically create an account for users on their initial login. The second is Manual Registration which allows you to manually register users using the Forge4Flow dashboard or the Users API regardless if Auto Registration is enabled or not.

See the Forge4Flow [configuration guide](../../../getting-started/configuration.md) to learn how to setup Auto Registration

### Data Integrity

By default, Auth4Flow operates with 'strict data integrity' with respect to users. This means that you can only create warrants for users that also exist in Auth4Flow. So before we can create user-specific warrants, we need to 'register' our application users. This can be done using Auto/Manual Registration, the Forge4Flow dashboard, or the Users API.

