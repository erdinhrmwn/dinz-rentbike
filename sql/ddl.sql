-- Users
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	phone VARCHAR(20) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(20) NOT NULL DEFAULT 'customer',
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Vehicles
CREATE TABLE IF NOT EXISTS vehicles (
	id SERIAL PRIMARY KEY,
	type VARCHAR(20) NOT NULL,
	brand VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	category VARCHAR(100) NOT NULL,
	description TEXT,
	image_url TEXT,
	price_per_hour DECIMAL(12, 2) NOT NULL,
	status VARCHAR(20) NOT NULL DEFAULT 'available',
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Rentals
CREATE TABLE IF NOT EXISTS rentals (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users (id),
	vehicle_id INT NOT NULL REFERENCES vehicles (id),
	start_time TIMESTAMPTZ NOT NULL,
	end_time TIMESTAMPTZ NOT NULL,
	total_hours INT NOT NULL,
	total_price DECIMAL(12, 2) NOT NULL,
	status VARCHAR(20) NOT NULL DEFAULT 'pending',
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Payments
CREATE TABLE IF NOT EXISTS payments (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users (id),
	rental_id INT NOT NULL REFERENCES rentals (id),
	amount DECIMAL(12, 2) NOT NULL,
	status VARCHAR(20) NOT NULL DEFAULT 'pending',
	payment_method VARCHAR(50),
	xendit_invoice_id VARCHAR(255) UNIQUE,
	xendit_payment_url TEXT,
	paid_at TIMESTAMPTZ,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Reviews
CREATE TABLE IF NOT EXISTS reviews (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users (id),
	vehicle_id INT NOT NULL REFERENCES vehicles (id),
	rental_id INT NOT NULL REFERENCES rentals (id),
	rating SMALLINT NOT NULL,
	comment TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
