# Warrants

The rules that specify relationships between objects in an application (e.g. `[store:A] is [parent] of [item:123]`) are called **warrants**. The Auth4Flow access check engine uses both an application's warrants and its object types to answer access check queries at runtime. Individual warrants define concrete relations between objects while object types define how some relations can be implied by others.

### Overview <a href="#overview" id="overview"></a>

Every warrant is composed of three required attributes and an optional _policy_ (more on this later):

* **Object** - This is the resource the warrant specifies a relationship on. The object is broken down into an `objectType` and an `objectId`. The `objectType` must refer to a valid object type defined for the application. The `objectId` is the identifier used by the application for the given resource.
* **Relation** - This is the relationship the warrant specifies between the object and subject. The relation must be one of the defined relations on the referenced `objectType`. It is often used to specify an action the warrant will grant the subject the ability to perform on the object (e.g. `editor`, `viewer`, etc.).
* **Subject** - This is the resource being granted the specified relationship on the object. The subject is broken down into an `objectType`, an `objectId`, and optionally a `relation` (to specify a group of subjects). Similar to the object, a subject's `objectType` must refer to a valid object type defined for the application, and its `objectId` is the identifier used by the application for the given resource. While the subject will often times be an individual user, note that the subject can be an arbitrary resource (e.g. a role, a tenant, a server, an access token, etc.) or even a group of subjects (e.g. in [group warrants](warrants.md#group-warrants)).
* **Policy** (optional) - A policy specifies an additional boolean expression to be evaluated _at the time of each access check_request. The provided expression can reference arbitrary variables which can then be optionally provided by access check requests as _context_. Given some context at access check time, a warrant's policy must evaluate to true in order for the warrant to be considered a match for the access check. If a warrant's object, relation, and subject attributes do not match an access check or the expression evaluates to false, the warrant is not matched against during resolution of the access check. The policy attribute can be used to implement attribute based access control (ABAC).

Here is an example warrant specifying that the subject `user:ABC` has the relation `editor` on object `item:123`:

```
{
  "objectType": "item",
  "objectId": "123",
  "relation": "editor",
  "subject": {
    "objectType": "user",
    "objectId": "ABC"
  }
}
```

### Direct Warrants <a href="#direct-warrants" id="direct-warrants"></a>

A direct warrant represents a relationship between two _specific_ objects. For example, we can define a warrant specifying that `[user:1] is a [member] of [role:admin]`:

```
{
  "objectType": "role",
  "objectId": "admin",
  "relation": "member",
  "subject": {
    "objectType": "user",
    "objectId": "1"
  }
}
```

### Group Warrants <a href="#group-warrants" id="group-warrants"></a>

In some scenarios, we need to specify a relationship between an object and a _group_ of subjects (e.g. `[member]s of [role:admin]` or `[manager]s of [tenant:acme]`). This is especially useful when implementing coarse grained access control schemes like role based access control (RBAC) or to grant access to a group of subjects without the overhead of creating a separate warrant for each subject. Group warrants are specified by including the optional `relation` attribute on the `subject` of the warrant. Providing the `relation` attribute on the `subject` specifies that the warrant applies to _all_ resources that match the given subject's `objectType`, `objectId`, and `relation`. For example, we can define a group warrant specifying that `a [member] of [role:admin] is an [editor] of [report:1]`:

```
{
  "objectType": "report",
  "objectId": "1",
  "relation": "editor",
  "subject": {
    "objectType": "role",
    "objectId": "admin",
    "relation": "member"
  }
}
```

### Wildcards <a href="#wildcards" id="wildcards"></a>

In other scenarios, we need to specify that a subject has a relation on **all** objects of a particular type. We can accomplish this without having to create a warrant for every object of that type by using a "wildcard" (`*`) `objectId`. Providing the wildcard character on the `objectId` attribute of a warrant specifies that the warrant applies to **all** objects of the given `objectType`. For example, we can define a warrant specifying that `[user:123] is an [editor] of [all reports]`:

```
{
  "objectType": "report",
  "objectId": "*",
  "relation": "editor",
  "subject": {
    "objectType": "user",
    "objectId": "123"
  }
}
```

### Policies & Context <a href="#policies--context" id="policies--context"></a>

Warrants can optionally include a policy. A **policy** is a [boolean expression](https://en.wikipedia.org/wiki/Boolean\_expression) specifying _additional_ conditions that must be satisfied in order for a warrant to match an access check. In other words, if a warrant is matched during an access check via its object, relation, and subject attributes, the warrant's policy (if it has one) will _also_ be evaluated in order to determine whether or not the warrant truly matches the access check. If the policy evaluates to `true` the warrant is considered a match. Otherwise, it is not considered a match.

and `[role:accountant] is a [member] of [permission:view-profits-and-losses]` _only when_ `customerId == "wayne-enterprises"`:

```
{
  "objectType": "permission",
  "objectId": "view-profits-and-losses",
  "relation": "member",
  "subject": {
    "objectType": "role",
    "objectId": "accountant"
  },
  "policy": "companyId == \"wayne-enterprises\""
}
```

Policies can reference variables whose values are dynamic. The value of each variable is provided by access check requests via a _context_ attribute (i.e. contextual information from your application such as a role, a tenant, a geographic location, etc). Before a policy is evaluated, the values provided via context are substituted into the expression so the expression can be properly evaluated. Policy expressions are statically type checked, so type mismatches between values will not evaluate to `true`. Policies with incomplete values (i.e. when the value for one or more variables in the expression is not provided in the context) or policies that encounter an error during evaluation will also not evaluate to `true`. Note that policies are compiled and statically checked for errors upon creation.

Policies have numerous uses but are most commonly used to implement forms of attribute based access control (ABAC). For example, we can create a warrant with a policy that only matches users coming from a specific ip address:

```
{
  "objectType": "database",
  "objectId": "prod",
  "relation": "admin",
  "subject": {
    "objectType": "user",
    "objectId": "ops-user"
  },
  "policy": "user.client_ip == \"192.168.1.1\""
}
```

When combined with role based access control (RBAC), policies allow an application to support different role/permission mappings per customer or tenant. For example, we can create a warrant specifying that `[role:accountant]` grants users `[permission:view-balance-sheet]` _only for_ `customerId == "wayne-enterprises`:

```
{
  "objectType": "permission",
  "objectId": "view-balance-sheet",
  "relation": "member",
  "subject": {
    "objectType": "role",
    "objectId": "accountant"
  },
  "policy": "companyId == \"wayne-enterprises\""
}
```

And another warrant specifying that `[role:accountant]` grants users `[permission:view-profits-and-losses]` only for `customerId == "daily-planet"`:

```
{
  "objectType": "permission",
  "objectId": "view-profits-and-losses",
  "relation": "member",
  "subject": {
    "objectType": "role",
    "objectId": "accountant"
  },
  "policy": "companyId == \"daily-planet\""
}
```

We can then make access checks, passing in different values for `companyId` via the `context` based on the company a user belongs to in our application:

```
{
  "warrants": [
    {
      "objectType": "permission",
      "objectId": "view-profits-and-losses",
      "relation": "member",
      "subject": {
        "objectType": "role",
        "objectId": "accountant"
      },
      "context": {
        "companyId": "wayne-enterprises"
      }
    }
  ]
}
```

This access check will return `false` because `[role:accountant]` only grants `[permission:view-profits-and-losses]`within the context of company `daily-planet`.

#### Built-in Methods <a href="#built-in-methods" id="built-in-methods"></a>

Warrant provides built-in methods that can be called from within policies to provide additional functionality:

**`expiresIn(duration)`**

The `expiresIn` method evaluates to `true` if the current time (the time the policy is being evluated) comes chronologically before the time the warrant was created + a specified `duration`. The `duration` must be specified as a string in a [format accepted by Go's time.ParseDuration method](https://pkg.go.dev/time#ParseDuration).

```
// warrant will be valid for 24 hours
expiresIn("24h")

// warrant will be valid for 150 milliseconds
expiresIn("150ms")

// warrant will be valid for 60 seconds
expiresIn("60s")
```

### Creating and Managing Warrants <a href="#creating-and-managing-warrants" id="creating-and-managing-warrants"></a>

Warrants can be created directly in the Forge4Flow dashboard or programmatically via API. Check out the [API Reference](../../reference/api-reference/) for more details.
