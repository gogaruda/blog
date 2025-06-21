CREATE TABLE comments (
    id VARCHAR(26) NOT NULL UNIQUE PRIMARY KEY,
    post_id VARCHAR(26) NOT NULL,
    parent_id VARCHAR(26),
    author_name VARCHAR(100),
    author_email VARCHAR(150),
    content TEXT NOT NULL,
    status ENUM('approved', 'pending', 'spam') DEFAULT 'pending',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (parent_id) REFERENCES comments(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;