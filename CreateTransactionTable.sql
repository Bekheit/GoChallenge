create table transactions (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	currency STRING,
	amount DECIMAL,
	createdat TIMESTAMPTZ DEFAULT now()
);