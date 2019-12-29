CREATE TABLE IF NOT EXISTS package (
  id INTEGER PRIMARY KEY,
  name TEXT,
  localpath TEXT,
  vendor TEXT,
  status TEXT,
  created_at datetime,
  updated_at datetime,
  remote_url TEXT
)