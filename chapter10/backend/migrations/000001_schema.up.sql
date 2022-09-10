CREATE SCHEMA IF NOT EXISTS gowebapp;

CREATE TABLE gowebapp.users (
    User_ID        BIGSERIAL PRIMARY KEY,
    User_Name      TEXT NOT NULL,
    Password_Hash  TEXT NOT NULL,
    Name           TEXT NOT NULL,
    Config         JSONB DEFAULT '{}'::JSONB NOT NULL,
    Created_At     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL,
    UNIQUE (User_Name)
);
-- SQLc converts snake_case to CamelCase

-- Simple image blob demo
CREATE TABLE gowebapp.images (
    Image_ID        BIGSERIAL PRIMARY KEY,
    User_ID         BIGINT NOT NULL,
    Content_Type    TEXT NOT NULL DEFAULT 'image/png',
    Image_Data      BYTEA NOT NULL
);


CREATE TABLE gowebapp.workouts (
    Workout_ID  BIGSERIAL PRIMARY KEY,
    User_ID     BIGINT NOT NULL,
    Start_Date  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE gowebapp.exercises (
    User_ID         BIGINT NOT NULL,
    Exercise_Name   TEXT NOT NULL,
    PRIMARY KEY(User_ID, Exercise_Name)
);

CREATE TABLE gowebapp.sets (
    Set_ID          BIGSERIAL PRIMARY KEY,
    Workout_ID      BIGINT NOT NULL,
    Exercise_Name   TEXT NOT NULL,
    Weight          INT NOT NULL DEFAULT 0, -- this can go up in decimal amounts so we can just divide/multiply to keep things simple
    Set1            BIGINT NOT NULL DEFAULT 0,
    Set2            BIGINT NOT NULL DEFAULT 0,
    Set3            BIGINT NOT NULL DEFAULT 0
);

