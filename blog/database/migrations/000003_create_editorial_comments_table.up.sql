CREATE TABLE editorial_comments (
    id VARCHAR(26) NOT NULL UNIQUE PRIMARY KEY,
    post_id VARCHAR(26) NOT NULL,
    editor_id VARCHAR(26) NOT NULL,
    status ENUM('revision_requested', 'rejected'),
    comment TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,


    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (editor_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;