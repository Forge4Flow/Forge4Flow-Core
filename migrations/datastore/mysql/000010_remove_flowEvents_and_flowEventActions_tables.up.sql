BEGIN;

# NOTE: This migration is not fully reversible and results in dataloss.
# Ensure you have backed up your database and followed the v0.0.4 migration guide.
# It will drop all existing flowEvent and flowEventActions.

DROP TABLE IF EXISTS flowEvent;
DROP TABLE IF EXISTS flowEventActions;

COMMIT;