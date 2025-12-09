ALTER TABLE users DROP CONSTRAINT users_tenant_email_unique;
ALTER TABLE users ADD CONSTRAINT users_email_key UNIQUE (email);