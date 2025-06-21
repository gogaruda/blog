CREATE TABLE post_series_item (
    series_id VARCHAR(26) NOT NULL,
    post_id VARCHAR(26) NOT NULL,
    order_index INT DEFAULT 0,
    PRIMARY KEY (series_id, post_id),
    FOREIGN KEY (series_id) REFERENCES post_series(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;