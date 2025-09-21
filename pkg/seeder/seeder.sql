BEGIN;

INSERT INTO public."role"
(id, "name")
VALUES(nextval('role_id_seq'::regclass), 'admin');

INSERT INTO public.role_right
(id, role_id, "section", route, r_create, r_read, r_update, r_delete)
VALUES(nextval('role_right_id_seq'::regclass), 1, 'be', '/users/user', 1, 1, 1, 1);

INSERT INTO public."user"
(id, role_id, "name", "password", email, last_access)
VALUES(nextval('user_id_seq'::regclass), 1, 'Admin', '$2a$10$3se4UX19yN9mA5JvkCT7/.rpnCemqIx2yJkhbQdYfXO2MsZW5Isuq', 'admin@admin.com', 0);

COMMIT;
