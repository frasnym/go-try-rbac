## Seeder
```
-- SQLite
INSERT INTO permissions (id, name)
VALUES (1, 'GET /admin'), (2, 'GET /user'), (3, 'POST /user'), (4, 'GET /admin-or-user');

INSERT INTO roles (id, name)
VALUES (1, 'admin'), (2, 'user');

INSERT INTO role_permissions (role_id, permission_id)
VALUES (1, 1), (1, 3), (1, 4), (2,2), (2,4);
```

## cURL
```
curl --location 'localhost:8080/admin-or-user' \
--header 'X-User-Role: admin'
```