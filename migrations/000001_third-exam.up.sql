CREATE TABLE Users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    username VARCHAR(32),
    bio VARCHAR(300),
    website VARCHAR(64),
    email VARCHAR(320),
    password VARCHAR(32),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE Posts (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    userID UUID,
    content TEXT,
    title VARCHAR(256),
    likes INTEGER,
    dislikes INTEGER,
    views INTEGER,
    categories VARCHAR(128),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE Comments (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    owner_id UUID,
    post_id UUID,
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

INSERT INTO Users (first_name, last_name, username, bio, website, email, password) VALUES
('John', 'Doe', 'johndoe', 'Software Engineer', 'www.johndoe.com', 'john@example.com', 'password123'),
('Jane', 'Smith', 'janesmith', 'Data Scientist', 'www.janesmith.com', 'jane@example.com', 'password456'),
('Michael', 'Johnson', 'michaelj', 'Web Developer', 'www.michaelj.com', 'michael@example.com', 'password789'),
('Emily', 'Brown', 'emilyb', 'UX Designer', 'www.emilyb.com', 'emily@example.com', 'passwordabc'),
('David', 'Wilson', 'davidw', 'Product Manager', 'www.davidw.com', 'david@example.com', 'passworddef'),
('Sarah', 'Anderson', 'saraa', 'Marketing Specialist', 'www.saraa.com', 'sarah@example.com', 'passwordghi'),
('Daniel', 'Martinez', 'danm', 'Graphic Designer', 'www.danm.com', 'daniel@example.com', 'passwordjkl'),
('Jessica', 'Taylor', 'jesst', 'Content Writer', 'www.jesst.com', 'jessica@example.com', 'passwordmno'),
('Christopher', 'Harris', 'chrish', 'Software Developer', 'www.chrish.com', 'chris@example.com', 'passwordpqr'),
('Amanda', 'Clark', 'amandac', 'Project Manager', 'www.amandac.com', 'amanda@example.com', 'passwordstu');

INSERT INTO Posts (userID, content, title, likes, dislikes, views, categories) VALUES
('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', 'Post 1', 20, 5, 100, 'Technology'),
('6512bf3e-4b37-47e3-80c9-0bb70693d81f', 'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'Post 2', 15, 3, 80, 'Science'),
('d17bc8db-0aee-4af1-b72f-2a746cc7880e', 'Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 'Post 3', 30, 8, 120, 'Business'),
('5e2d4b36-21b9-4c28-8325-4684a6a9d249', 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.', 'Post 4', 25, 6, 90, 'Education'),
('da42e522-4ff7-4c24-bd07-5ee9a02bc6e4', 'Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'Post 5', 18, 4, 110, 'Health'),
('23e68d5d-1e57-4ebf-a25e-500d3d03fd92', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', 'Post 6', 22, 7, 95, 'Technology'),
('38b7a68d-94eb-4806-967d-53d788e4e8bb', 'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'Post 7', 27, 9, 130, 'Science'),
('689472b4-76f2-42d2-90a1-6c47191c7d74', 'Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 'Post 8', 19, 3, 105, 'Business'),
('b966c5d7-cc4d-4b52-8cc7-75e5563e3151', 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.', 'Post 9', 24, 5, 85, 'Education'),
('b766f90e-40cc-4ac7-9364-e38f169b1b22', 'Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'Post 10', 16, 6, 115, 'Health');

INSERT INTO Comments (owner_id, post_id, text) VALUES
('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Great post!'),
('6512bf3e-4b37-47e3-80c9-0bb70693d81f', '6512bf3e-4b37-47e3-80c9-0bb70693d81f', 'Interesting topic!'),
('d17bc8db-0aee-4af1-b72f-2a746cc7880e', 'd17bc8db-0aee-4af1-b72f-2a746cc7880e', 'Well explained.'),
('5e2d4b36-21b9-4c28-8325-4684a6a9d249', '5e2d4b36-21b9-4c28-8325-4684a6a9d249', 'I have a question.'),
('da42e522-4ff7-4c24-bd07-5ee9a02bc6e4', 'da42e522-4ff7-4c24-bd07-5ee9a02bc6e4', 'Looking forward to more content!'),
('23e68d5d-1e57-4ebf-a25e-500d3d03fd92', '23e68d5d-1e57-4ebf-a25e-500d3d03fd92', 'This helped me a lot.'),
('38b7a68d-94eb-4806-967d-53d788e4e8bb', '38b7a68d-94eb-4806-967d-53d788e4e8bb', 'I disagree with some points.'),
('689472b4-76f2-42d2-90a1-6c47191c7d74', '689472b4-76f2-42d2-90a1-6c47191c7d74', 'Could you elaborate more?'),
('b966c5d7-cc4d-4b52-8cc7-75e5563e3151', 'b966c5d7-cc4d-4b52-8cc7-75e5563e3151', 'Waiting for your next post.'),
('b766f90e-40cc-4ac7-9364-e38f169b1b22', 'b766f90e-40cc-4ac7-9364-e38f169b1b22', 'Impressive content!');


