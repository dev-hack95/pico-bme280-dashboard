CREATE TABLE sensor_readings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    humidity TEXT,
    pressure TEXT,
    temperature TEXT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	active BOOLEAN DEFAULT TRUE
)

CREATE TABLE users (
id INTEGER PRIMARY KEY AUTOINCREMENT,
user_name TEXT,
password TEXT,
active BOOLEAN DEFAULT TRUE)