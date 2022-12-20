-- name: ListUsers :many
SELECT *
FROM gowebapp.users
ORDER BY user_name;

-- name: ListImages :many
SELECT *
FROM gowebapp.images
ORDER BY image_id;


-- name: GetUser :one
SELECT *
FROM gowebapp.users
WHERE user_id = $1;

-- name: GetUserByName :one
SELECT *
FROM gowebapp.users
WHERE user_name = $1;

-- name: GetUserImage :one
SELECT u.name, u.user_id, i.image_data
FROM gowebapp.users u,
     gowebapp.images i
WHERE u.user_id = i.user_id
  AND u.user_id = $1;

-- name: DeleteUsers :exec
DELETE
FROM gowebapp.users
WHERE user_id = $1;

-- name: DeleteUserImage :exec
DELETE
FROM gowebapp.images i
WHERE i.user_id = $1;

-- name: DeleteUserWorkouts :exec
DELETE
FROM gowebapp.workouts w
WHERE w.user_id = $1;


-- name: CreateUserImage :one
INSERT INTO gowebapp.images (User_ID, Content_Type, Image_Data)
values ($1,
        $2,
        $3) RETURNING *;

-- name: UpsertUserImage :one
INSERT INTO gowebapp.images (Image_Data)
VALUES ($1) ON CONFLICT (Image_ID) DO
UPDATE
    SET Image_Data = EXCLUDED.Image_Data
    RETURNING Image_ID;

-- name: CreateUsers :one
INSERT INTO gowebapp.users (User_Name, Password_Hash, name)
VALUES ($1,
        $2,
        $3) RETURNING *;
