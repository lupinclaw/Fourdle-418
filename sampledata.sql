-- Insert data into the User table
INSERT INTO User (Email, PasswordSaltAndHash) VALUES
    ('user1@example.com', 'salt1$hashed_password1'),
    ('user2@example.com', 'salt2$hashed_password2'),
    ('user3@example.com', 'salt3$hashed_password3');

-- Insert data into the SubscriptionType table
INSERT INTO SubscriptionType (Type, Fee) VALUES
    ('Basic', 10),
    ('Premium', 20),
    ('Pro', 30);

-- Insert data into the Subscription table
INSERT INTO Subscription (UserID, SubscriptionTypeID, Active, LastInvoiceDate, SignUpDate) VALUES
    (1, 1, 1, '2023-01-01', '2022-12-01'),
    (2, 2, 1, '2023-01-15', '2022-11-01'),
    (3, 3, 0, '2022-12-20', '2022-10-01');

-- Insert data into the Game table
INSERT INTO Game (WordID, WinOrLoss) VALUES
    (1, 1),
    (2, 0),
    (3, 1),
    (4, 0);

-- Insert data into the UserGame table
INSERT INTO UserGame (GameID, UserID, Date) VALUES
    (1, 1, '2023-01-02'),
    (2, 2, '2023-01-16'),
    (3, 3, '2022-12-21'),
    (4, 1, '2023-01-05');

-- Insert data into the UserMove table
INSERT INTO UserMove (UserID, GameID, Guess) VALUES
    (1, 1, 'abcd'),
    (2, 2, 'wxyz'),
    (3, 3, 'mnop'),
    (1, 4, 'qrst');
