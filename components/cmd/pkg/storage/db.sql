--- Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
---
--- This software is released under the MIT License.
--- https://opensource.org/licenses/MIT

/* Create required tables */
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

CREATE TABLE IF NOT EXISTS event (
  id INTEGER PRIMARY KEY,
  title TEXT,
  data TEXT,
  from TEXT,
  status TEXT,
  created_at DATETIME,
  updated_at DATETIME
)

CREATE TABLE IF NOT EXISTS log (
  id INTEGER PRIMARY KEY,
  title TEXT,
  filepath TEXT,
  from TEXT,
  created_at DATETIME,
  updated_at DATETIME
)

/* Insert default values */
INSERT INTO log (
  title, filepath, from, created_at, updated_at
) values (
  'system', 
)