CREATE TABLE go_task.users (
	id INT auto_increment NOT NULL,
	username varchar(100) NOT NULL,
	email varchar(250) NOT NULL,
	first_name varchar(100) NULL,
	last_name varchar(100) NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NULL,
	updated_by INT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_unique UNIQUE KEY (username),
	CONSTRAINT users_unique_1 UNIQUE KEY (email)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
