		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(30) DEFAULT 'user' CHECK (role IN ('admin', 'greetings_manager', 'user')),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS greetings (
			id SERIAL PRIMARY KEY,
			message TEXT NOT NULL,
			created_by INTEGER REFERENCES users(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);

		INSERT INTO greetings (message, created_by) 
		 SELECT 'Looks like a nice day', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Looks like a nice day');

		INSERT INTO greetings (message, created_by) 
		 SELECT 'Pretty nice morning, uh?', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Pretty nice morning, uh?');

		INSERT INTO greetings (message, created_by) 
		 SELECT 'Have a good day', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Have a good day');

		INSERT INTO greetings (message, created_by) 
		 SELECT 'Check back later if you like it', NULL 
		 WHERE NOT EXISTS (SELECT 1 FROM greetings WHERE message = 'Check back later if you like it');
