CREATE DATABASE IF NOT EXISTS go_stock_api;
USE go_stock_api;

-- Users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Stocks table (reference for tickers)
CREATE TABLE stocks (
    ticker VARCHAR(10) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Transactions table (purchase/sale history)
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    ticker VARCHAR(10) NOT NULL,
    type ENUM('buy', 'sell') NOT NULL,
    quantity DECIMAL(10,4) NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10,2) NOT NULL,
    date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ticker) REFERENCES stocks(ticker)
);

-- Optional: Historical prices
CREATE TABLE historical_prices (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ticker VARCHAR(10) NOT NULL,
    date DATE NOT NULL,
    close_price DECIMAL(10,2) NOT NULL,

    FOREIGN KEY (ticker) REFERENCES stocks(ticker),
    UNIQUE(ticker, date)
);
