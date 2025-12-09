ALTER TABLE users DROP CONSTRAINT users_email_key;
ALTER TABLE users ADD CONSTRAINT users_tenant_email_unique UNIQUE (tenant_key, email);
