CREATE TABLE IF NOT EXISTS package (
  id INTEGER PRIMARY KEY,
  name TEXT,
  localpath TEXT,
  vendor TEXT,
  status VARCHAR(16),
  created_at datetime,
  updated_at datetime,
  remote_url TEXT
)