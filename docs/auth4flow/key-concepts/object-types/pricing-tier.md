# Pricing Tier

Pricing tiers represent specific 'packages' of features within an application that can be assigned to specific users and tenants to gate access based on payment plan and entitlements. They are typically used in SaaS apps to implement and manage different pricing plans (ex. 'free', 'growth', 'enterprise'). The full representation of the `pricing-tier` object type is:

```
{
  "type": "pricing-tier",
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
      "ofType": "pricing-tier",
      "withRelation": "member"
    }
  }
}
```
