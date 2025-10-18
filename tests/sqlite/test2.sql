
-- SQLite Full Feature Demonstration Script
-- Demonstrates all major SQLite statement types with English comments

-- 1. Database initialization and settings
PRAGMA foreign_keys = ON;  -- Enable foreign key constraints
PRAGMA journal_mode = WAL; -- Use Write-Ahead Logging for better concurrency

-- 2. Database attachment/detachment
ATTACH DATABASE ':memory:' AS demo_db; -- Create in-memory database
DETACH DATABASE demo_db;              -- Detach the database

-- 3. Transaction control statements
BEGIN TRANSACTION; -- Start transaction
SAVEPOINT sp1;     -- Create savepoint
RELEASE sp1;       -- Release savepoint
COMMIT;            -- Commit transaction

BEGIN TRANSACTION; -- Start another transaction
ROLLBACK;          -- Rollback transaction

-- 4. Table creation with constraints
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,      -- Auto-incrementing primary key
    username TEXT UNIQUE NOT NULL,    -- Unique username constraint
    email TEXT CHECK(email LIKE '%@%'), -- Email format validation
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP -- Automatic timestamp
);

CREATE TABLE posts (
    post_id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id), -- Foreign key reference
    title TEXT,
    content TEXT,
    views INTEGER DEFAULT 0 -- Default view count
);

-- 5. Index operations
CREATE INDEX idx_username ON users(username); -- Speed up username searches
CREATE INDEX idx_post_views ON posts(views DESC); -- Optimized for popular posts
ANALYZE;        -- Update statistics for query planner
REINDEX idx_username; -- Rebuild the username index

-- 6. View creation
CREATE VIEW popular_posts AS
SELECT p.post_id, u.username, p.title, p.views
FROM posts p JOIN users u ON p.user_id = u.user_id
WHERE p.views > 100
ORDER BY p.views DESC; -- Shows only posts with >100 views

-- 7. Trigger implementation
CREATE TRIGGER update_post_count
AFTER INSERT ON posts
FOR EACH ROW
BEGIN
    UPDATE users 
    SET post_count = COALESCE(post_count, 0) + 1
    WHERE user_id = NEW.user_id; -- Increment user's post count
END;

-- 8. Full-text search virtual table
CREATE VIRTUAL TABLE posts_search USING fts5(
    title, 
    content, 
    content='posts',
    content_rowid='post_id',
    tokenize='porter unicode61' -- Advanced tokenization
);

-- 9. Data manipulation (CRUD operations)
INSERT INTO users (username, email) VALUES 
('alice', 'alice@example.com'),
('bob', 'bob@example.com'); -- Sample user data

INSERT INTO posts (user_id, title, content) VALUES
(1, 'First Post', 'Hello world!'),
(2, 'SQLite Guide', 'Complete SQLite tutorial'); -- Sample posts

-- 10. Query examples with analysis
EXPLAIN QUERY PLAN
SELECT u.username, COUNT(p.post_id) as post_count
FROM users u LEFT JOIN posts p ON u.user_id = p.user_id
GROUP BY u.user_id
HAVING post_count > 0; -- Analyze join performance

SELECT * FROM users WHERE username LIKE 'a%'; -- Simple query
UPDATE users SET email = 'alice.new@example.com' WHERE user_id = 1; -- Data update
DELETE FROM posts WHERE post_id = 2; -- Data deletion

-- 11. Limited update/delete operations
UPDATE posts SET views = views + 1 WHERE post_id = 1 LIMIT 1; -- Safe update
DELETE FROM posts WHERE views < 5 LIMIT 10; -- Batch deletion with limit

-- 12. Schema modifications
ALTER TABLE users ADD COLUMN bio TEXT DEFAULT ''; -- Add new column
ALTER TABLE posts RENAME COLUMN content TO body;  -- Rename column

-- 13. Maintenance operations
VACUUM;        -- Rebuild database file
PRAGMA optimize; -- Optimize database automatically

-- 14. Object cleanup
DROP TRIGGER update_post_count; -- Remove trigger
DROP VIEW popular_posts;        -- Remove view
DROP INDEX idx_username;        -- Remove index
DROP TABLE posts;               -- Remove table
DROP TABLE users;               -- Remove table
