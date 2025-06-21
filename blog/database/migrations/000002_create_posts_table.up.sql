CREATE TABLE posts (
    id VARCHAR(26) NOT NULL UNIQUE PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    excerpt TEXT,
    content LONGTEXT NOT NULL,
    featured_image VARCHAR(255),
    status ENUM('draft', 'pending_review', 'revision_requested', 'rejected'),
    author_id VARCHAR(26) NOT NULL,
    editor_id VARCHAR(26) NOT NULL,
    published_at DATETIME,
    seo_title VARCHAR(60),
    seo_description VARCHAR(160),
    canonical_url VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY(author_id) REFERENCES users(id),
    FOREIGN KEY(editor_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;