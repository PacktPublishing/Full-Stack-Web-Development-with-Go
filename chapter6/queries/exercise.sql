-- name: CreateUserExercise :one
INSERT INTO gowebapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    $1,
    $2
) ON CONFLICT (Exercise_Name) DO NOTHING RETURNING (
    User_ID, Exercise_Name
);

-- name: ListUserExercises :many
SELECT Exercise_Name
FROM gowebapp.exercises
WHERE User_ID = $1;

-- name: DeleteUserExercise :exec
DELETE FROM gowebapp.exercises
WHERE User_ID = $1 AND Exercise_Name = $2;


-- name: CreateUserDefaultExercise :exec
INSERT INTO gowebapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    1,
    'Bench Press'
),(
    1,
    'Barbell row'
);

-- name: CreateUserWorkout :one
INSERT INTO gowebapp.workouts (
    User_ID,
    Start_Date
) VALUES (
    $1,
    NOW()
) RETURNING *;

-- name: CreateDefaultSetForExercise :one
INSERT INTO gowebapp.sets (
    Workout_ID,
    Exercise_Name,
    Weight
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: CreateSetForExercise :one
INSERT INTO gowebapp.sets (
    Workout_ID,
    Exercise_Name, 
    Weight,
    Set1,
    Set2,
    Set3
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: UpdateSet :one
UPDATE gowebapp.sets SET
    Weight = $2,
    Set1 = $3,
    Set2 = $4,
    Set3 = $5
WHERE Set_ID = $1 RETURNING *;