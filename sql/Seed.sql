
-- Seeding Data: Users
INSERT INTO users (id, email, username, gender, country, degree, major)
VALUES ('b9ba95ec-3041-708f-44b0-bfad168dc0ca', 'john.doe@gmail.com', 'John Doe', 'Man', 'US', 'Bachelor', 'Science'),
       ('c9eac5bc-d071-70f5-9ece-1ace39ec4cf2', 'jane.smith@gmail.com', 'Jane Smith', 'Women', 'Germany', 'Master',
        'Art'),
       ('599ac52c-a0c1-70ca-3505-5cdf9f2c54539', 'amir.malaysia@gmail.com', 'Amir Khan', 'Man', 'Malaysia', 'Doctoral',
        'Social'),
       ('09fad50c-b061-7089-68b0-b900aca5551e', 'lisa.lee@gmail.com', 'Lisa Lee', 'Women', 'Australia', 'Diploma',
        'Art'),
       ('094ae51c-7091-70da-8ed5-3c8665a3968b', 'emma.clark@gmail.com', 'Emma Clark', 'Women', 'US', 'Bachelor',
        'Social');

-- Seeding Data: Posts
INSERT INTO posts (user_id, title, content, views)
VALUES ('b9ba95ec-3041-708f-44b0-bfad168dc0ca', 'My First Post', 'This is the content of my first post.', 10),
       ('c9eac5bc-d071-70f5-9ece-1ace39ec4cf2', 'Art and Science', 'Exploring the connection between art and science.',
        25),
       ('599ac52c-a0c1-70ca-3505-5cdf9f2c54539', 'Social Studies in Malaysia',
        'A deep dive into social studies curriculum in Malaysia.', 30),
       ('09fad50c-b061-7089-68b0-b900aca5551e', 'Creative Design', 'Tips for creative design projects.', 5),
       ('094ae51c-7091-70da-8ed5-3c8665a3968b', 'Breaking Barriers',
        'Innovative ideas to break stereotypes in education.', 50);

-- Seeding Data: Comments
INSERT INTO comments (post_id, user_id, content)
VALUES (1, 'c9eac5bc-d071-70f5-9ece-1ace39ec4cf2', 'Great first post! Keep it up.'),
       (1, '599ac52c-a0c1-70ca-3505-5cdf9f2c54539', 'Nice content. Looking forward to more.'),
       (2, '09fad50c-b061-7089-68b0-b900aca5551e', 'Amazing insights into art and science.'),
       (3, 'b9ba95ec-3041-708f-44b0-bfad168dc0ca', 'This is very informative. Thanks for sharing.'),
       (4, '094ae51c-7091-70da-8ed5-3c8665a3968b', 'Excellent advice for creative projects!');

-- Seeding Data: Likes
INSERT INTO likes (post_id, user_id)
VALUES (1, 'c9eac5bc-d071-70f5-9ece-1ace39ec4cf2'),
       (1, '599ac52c-a0c1-70ca-3505-5cdf9f2c54539'),
       (2, 'b9ba95ec-3041-708f-44b0-bfad168dc0ca'),
       (3, '09fad50c-b061-7089-68b0-b900aca5551e'),
       (3, '094ae51c-7091-70da-8ed5-3c8665a3968b'),
       (4, 'c9eac5bc-d071-70f5-9ece-1ace39ec4cf2');