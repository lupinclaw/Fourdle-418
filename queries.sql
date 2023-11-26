/* Get total games played */
SELECT COUNT(GameID)
FROM UserGame
WHERE UserID = ?;

/* Get total games won */
SELECT COUNT(GameID)
FROM UserGame U, GAME G
WHERE G.GameID = U.GameID AND
U.UserID = ? AND
G.WinOrLoss = 1;

/* Get 10 most recent words */
SELECT W.Letters
FROM UserGame U, Game G, Word W
WHERE U.GameID = G.GameID AND G.WordID = W.WordID AND
U.UserID = ?
ORDER BY U.Date DESC LIMIT 10;

/* Get all seen words */
SELECT W.Letters
FROM UserGame U, Game G, Word W
WHERE U.GameID = G.GameID AND G.WordID = W.WordID
ORDER BY U.Date DESC;

/* Get number of games total for each of the last 30 days */
SELECT U.Date, COUNT(U.GameID)
FROM UserGame U, Game G
WHERE U.GameID = G.GameID AND
U.UserID = ?
GROUP BY U.Date
HAVING U.Date >= DATE('now', '-30 days');

/* Get number of games won for each of the last 30 days */
SELECT U.Date, COUNT(U.GameID)
FROM UserGame U, Game G
WHERE U.GameID = G.GameID AND
U.UserID = ? AND
G.WinOrLoss = 1
GROUP BY U.Date
HAVING U.Date >= DATE('now', '-30 days');