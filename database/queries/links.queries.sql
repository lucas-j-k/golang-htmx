-- name: ListLinksForUser :many
SELECT
  l.id,
  l.url,
  l.link_type_id,
  l.published,
  lt.name AS link_type_name,
  lt.icon_class
FROM
  link l
  LEFT JOIN link_type lt ON l.link_type_id = lt.id
WHERE
  l.user_id = ?;

-- name: ListPublicLinksForUser :many
SELECT
  l.id,
  l.url,
  l.link_type_id,
  lt.name AS link_type_name,
  lt.icon_class
FROM
  link l
  LEFT JOIN link_type lt ON l.link_type_id = lt.id
WHERE
  l.user_id = ?
AND
  l.published = 1;

-- name: ListLinkTypes :many
SELECT
  id,
  name
FROM
  link_type;

-- name: FindLinkById :one
SELECT
  id,
  url,
  link_type_id,
  published
FROM
  link
WHERE
  id = ?
  AND user_id = ?;

-- name: InsertLink :execresult
INSERT INTO
  link (user_id, url, link_type_id, published)
VALUES
  (?, ?, ?, ?);

SELECT
  id
FROM
  link
WHERE
  id = LAST_INSERT_ID();

-- name: DeleteLink :exec
DELETE FROM
  link
WHERE
  id = ?;

-- name: UpdateLink :exec
UPDATE
  link
SET
  url = ?,
  link_type_id = ?,
  published = ?
WHERE
  id = ?;