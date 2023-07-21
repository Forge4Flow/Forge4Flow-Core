# Object Types

Object types are the basic building block of any authorization scheme in Warrant. Represented as JSON, each object type defines a type of resource (e.g. stores, items, etc.) in an application along with the relationships that can exist between that and other resources in the application (e.g. users). Object types are an incredibly flexible way to define authorization models and even allow you to express complex hierarchical and inherited relationships.

Object types can be created directly in the Forge4Flow dashboard or via the Object Types API.

In this overview, we'll explain the key attributes of object types by creating an authorization model for a simple eCommerce application with three object types: user, store, and item.

### Type

The first attribute of an object type is its `type`. Each object type must have a unique string as its type. In our eCommerce app, we'll have the following object types:

```json
{
    "type": "user"
},
{
    "type": "store"
},
{
    "type": "item"
}
```

### Relations <a href="#relations" id="relations"></a>

By defining the object types above, we've started building an authorization model for our application that will allow us to create fine-grained access control rules for stores, items, and users, helping us answer questions like:

```
Does [user:1] have the ability to [edit] [item:x]?
is [user:1] the [owner] of [store:3]?
```

In order to create access rules using our object types, we first need to add relations to them. The relations of an object type define the relationships available on an object of that type. For example, if we want to specify that `[user:A] is an [owner] of [store:S]`, we must add an `owner` relation to the `store` object type.

There are two types of relations that can be defined on object types:

* [Direct Relations](./#direct-relations)
* [Inherited Relations](./#inherited-relations)

We'll start by adding some direct relations to our object types.

#### Direct Relations <a href="#direct-relations" id="direct-relations"></a>

```json
{
    "type": "user",
    "relations": {
        "manager": {}
    }
},
{
    "type": "store",
    "relations": {
        "owner": {},
        "editor": {},
        "viewer": {}
    }
},
{
    "type": "item",
    "relations": {
        "owner": {},
        "editor": {},
        "viewer": {},
        "parent": {}
    }
}
```

With these object types, we can now create authorization rules that specify exactly which users are `owners`, `editors`, and `viewers` of each store and item. We can also assign stores as `parents` of items, and users as `managers` of other users.

#### Inherited Relations <a href="#inherited-relations" id="inherited-relations"></a>

Using only direct relations to build your authorization model can be powerful, but explicitly creating warrants for each and every relationship in an application can become tedious or infeasible in larger, more complex use cases. That's why relations can define conditions under which they will be inherited (e.g. `a user is an editor of a store if they're an owner of that store`). There are two ways to specify how relations can be inherited:

* Inherited Hierarchical Relations
* Inherited Object Relations

#### **Inherited Hierarchical Relations**

In practice, it's common for relations to have overlap (e.g. an `owner` has the same privileges as an `editor` + other privileges). For example, in many applications a user with write privileges often inherits read privileges too. In our example application, an `owner`is also both an `editor` and a `viewer`, and an `editor` is also a `viewer`. Instead of having to explicitly assign each of the `owner`, `editor`, and `viewer` relations to a user who is an `owner`, object types allow you to specify a relation hierarchy (e.g. the `editor`relation is inherited if the user is an `owner`) using the `inheritIf` property. Let's add `inheritIf` rules to our `store` and `item`object types specifying that:

* `owners` are also `editors`
* `editors` are also `viewers`

```json
{
    "type": "user",
    "relations": {
        "manager": {}
    }
},
{
    "type": "store",
    "relations": {
        "owner": {},
        "editor": {
            "inheritIf": "owner"
        },
        "viewer": {
            "inheritIf": "editor"
        }
    }
},
{
    "type": "item",
    "relations": {
        "owner": {},
        "editor": {
            "inheritIf": "owner"
        },
        "viewer": {
            "inheritIf": "editor"
        },
        "parent": {}
    }
}
```

With our `inheritIf` rules in place, we can simply grant a user the `editor` relation and they will implicitly inherit the `viewer`relation. `inheritIf` rules also work recursively on other inherited relations, so assigning a user the `owner` relation will implicitly grant that user _both_ the `editor` and `viewer` relations. This is because `owner` will inherit `editor` and `editor` will in turn inherit `viewer`. This will simplify our access checks and cut down on the number of warrants we need to create for each user.

#### **Inherited Object Relations**

In many applications, resources have their own hierarchy (e.g. a document belongs to a folder) and the access rules for these resources follow that hierarchy (e.g. the owner of a folder is the owner of any document in that folder). By combining the `inheritIf`, `ofType`, and `withRelation` properties, you can specify that a relation will be inherited when a user has a specified relation (`inheritIf`) on another type of object (`ofType`) with a hierarchical relation (`withRelation`) on the current object. For example, a user is an `editor` of a document if they are an `editor` of a `folder` that is the document's `parent`. In our example app, let's define three inherited object relations:

1. A user is an `owner` of an item if that user is an `owner` of a `store` that is the item's `parent`.
2. A user is an `editor` of an item if that user is an `editor` of a `store` that is the item's `parent`.
3. A user is an `editor` of an item if that user is the `manager` of the `user` that is the item's `owner`.

**NOTE:** Some of the relations below will be [composing multiple inheritance rules together using logical operators.](./#composing-inherited-relations-using-logical-operators) We'll cover this in detail later.

```json
{
    "type": "user",
    "relations": {
        "manager": {}
    }
},
{
    "type": "store",
    "relations": {
        "owner": {},
        "editor": {
            "inheritIf": "owner"
        },
        "viewer": {
            "inheritIf": "editor"
        }
    }
},
{
    "type": "item",
    "relations": {
        "owner": {
            "inheritIf": "owner",
            "ofType": "store",
            "withRelation": "parent"
        },
        "editor": {
            "inheritIf": "anyOf",
            "rules": [
                {
                    "inheritIf": "owner"
                },
                {
                    "inheritIf": "editor",
                    "ofType": "store",
                    "withRelation": "parent"
                },
                {
                    "inheritIf": "manager",
                    "ofType": "user",
                    "withRelation": "owner"
                }
            ]
        },
        "viewer": {
            "inheritIf": "editor"
        },
        "parent": {}
    }
}
```

The `inheritIf`, `ofType`, and `withRelation` properties make it easy to define inheritance rules for complex relationships between objects so we don't have to create a large number of explicit warrants. Without them, we'd need to create a warrant for every item ↔ store ↔ user relationship in our application. This could easily be thousands, if not hundreds of thousands of rules.

### Composing Inherited Relations Using Logical Operators <a href="#composing-inherited-relations-using-logical-operators" id="composing-inherited-relations-using-logical-operators"></a>

With both direct and inherited relations in our toolkit, we can create authorization models for a majority of use cases, but there are still some scenarios in practice that require a combination of inheritance rules (e.g. a user is an `editor` of an item if they are an `owner` of that item **OR** they are the `manager` of another user who is an `editor` of that item). To support designing authorization models that cover such scenarios, relations can compose multiple inheritance rules using logical operations to form more complex conditions.

The three supported logical operations are `anyOf`, `allOf`, and `noneOf`.

#### `anyOf` <a href="#anyof" id="anyof"></a>

The `anyOf` operation allows you to specify a set of rules that are considered fulfilled if _at least one of_ the rules in the set is satisfied. In other words, it works like the logical _OR_ operation. The following object type specifies an `editor-or-viewer`relation that is inherited if the user is an `editor` **OR** if the user is a `viewer`:

```json
{
  "type": "item",
  "relations": {
    "editor": {},
    "viewer": {},
    "editor-or-viewer": {
      "inheritIf": "anyOf",
      "rules": [
        {
          "inheritIf": "editor"
        },
        {
          "inheritIf": "viewer"
        }
      ]
    }
  }
}
```

#### `allOf` <a href="#allof" id="allof"></a>

The `allOf` rule type allows you to specify a set of rules that are considered fulfilled if _all of_ the rules in the set are satisfied. In other words, it works like the logical _AND_ operation. The following object type specifies an `editor-and-viewer` relation that is implicitly granted if the user is an `editor` **AND** the user is a `viewer`:

```json
{
  "type": "item",
  "relations": {
    "editor": {},
    "viewer": {},
    "editor-and-viewer": {
      "inheritIf": "allOf",
      "rules": [
        {
          "inheritIf": "editor"
        },
        {
          "inheritIf": "viewer"
        }
      ]
    }
  }
}
```

#### `noneOf` <a href="#noneof" id="noneof"></a>

The `noneOf` rule type allows you to specify a set of rules that are considered fulfilled if _none of_ the rules in the set are satisfied. In other words, it works like the logical _NOR_ operation. The following object type specifies a `not-editor-and-not-viewer` relation that is implicitly granted if the user is _not_ an `editor` **AND** the user is _not_ a `viewer`:

```json
{
  "type": "item",
  "relations": {
    "editor": {},
    "viewer": {},
    "not-editor-and-not-viewer": {
      "inheritIf": "noneOf",
      "rules": [
        {
          "inheritIf": "editor"
        },
        {
          "inheritIf": "viewer"
        }
      ]
    }
  }
}
```

### Built-in Object Types <a href="#built-in-object-types" id="built-in-object-types"></a>

By default, Auth4Flow provides a core set of built-in object types that make it easy to implement well-defined, common authorization schemes like role based access control and multi-tenancy. These built-in types are:

* [User](user.md)
* [Tenant](tenant.md)
* [Role](role.md)
* [Permission](permission.md)
* [Pricing Tier](pricing-tier.md)
* [Feature](feature.md)

We'll cover each of these built-in object types in detail in the next sections.

### Summary <a href="#summary" id="summary"></a>

* Object types are composed of `type` and `relations` attributes.
* Relations can be granted explicitly or inherited via a combination of the `inheritIf`, `ofType`, and `withRelation` properties.
* Relation inheritance rules can be composed using logical operators to build more complex access control models.
* Auth4Flow provides several built-in object types like user, tenant, role, permission, feature, and pricing-tier that make it easy to setup common authorization use-cases.

Next, we'll dive into each of the built-in object types that Auth4Flow provides.
