BEGIN;

DELETE FROM objectType
WHERE typeId IN ('role', 'permission', 'tenant', 'user', 'pricing-tier', 'feature', 'float', 'find', 'emerald-id', 'fungible-token');

COMMIT;
