--- Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
---
--- This software is released under the MIT License.
--- https://opensource.org/licenses/MIT
/* Create required tables */
CREATE TABLE IF NOT EXISTS package (
	id TEXT PRIMARY KEY,
	name TEXT,
	localpath TEXT,
	vendor TEXT,
	status TEXT,
	created_at DATETIME,
	updated_at DATETIME,
	remote_url TEXT
);

CREATE TABLE IF NOT EXISTS event (
	id TEXT PRIMARY KEY,
	title TEXT,
	data TEXT,
	source TEXT,
	status TEXT,
	created_at DATETIME,
	updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS log (
	id TEXT PRIMARY KEY,
	title TEXT,
	filepath TEXT,
	source TEXT,
	created_at DATETIME,
	updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS solver (
	id TEXT PRIMARY KEY,
	package TEXT,
	vendor TEXT,
	name TEXT,
	solverpath TEXT,
	created_at DATETIME,
	updated_at DATETIME,
	status TEXT
);

CREATE TABLE IF NOT EXISTS hook (
	id TEXT PRIMARY KEY,
	name TEXT,
	local_url TEXT,
	remote_url TEXT,
	created_at DATETIME,
	updated_at DATETIME,
	status TEXT
);

CREATE TABLE IF NOT EXISTS apikey (
	id TEXT PRIMARY KEY,
	aidkey TEXT,
	created_at DATETIME,
	updated_at DATETIME,
	status TEXT
);