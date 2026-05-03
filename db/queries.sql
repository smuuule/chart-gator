-- name: GetSong :one
SELECT * FROM songs
WHERE id = $1 LIMIT 1;

-- name: CreateSong :one
INSERT INTO songs (
    title, artist, normalized_title, normalized_artist
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT (normalized_title, normalized_artist) DO NOTHING
RETURNING *;

-- name: ListSongs :many
SELECT * FROM songs
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetChartsBySongId :many
SELECT * FROM charts
WHERE song_id = $1
ORDER BY difficulty;

-- name: CreateChart :one
INSERT INTO charts (
    song_id, charter, difficulty, source_id, download_url
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;
