CREATE TABLE IF NOT EXISTS package (
  id INTEGER PRIMARY KEY,
  name TEXT,
  localpath TEXT,
  vendor TEXT,
  status TEXT,
  created_at DATETIME,
  updated_at DATETIME,
  remote_url TEXT
)