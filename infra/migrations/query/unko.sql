-- name: GetAllUnko :many
SELECT
       *
FROM
    unko u
ORDER BY
    u.id DESC
;
