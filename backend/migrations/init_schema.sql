CREATE TABLE IF NOT EXISTS targets (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	url TEXT NOT NULL,
	interval INTEGER NOT NULL,
	created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS target_logs (
	id TEXT PRIMARY KEY,
	target_id TEXT,
	status TEXT,
	timestamp DATETIME,
	FOREIGN KEY (target_id) REFERENCES targets(id)
);
