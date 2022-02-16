CREATE SCHEMA IF NOT EXISTS gowebapp;

CREATE TABLE gowebapp.users (
User_ID        BIGSERIAL PRIMARY KEY,
User_Name      text NOT NULL,
Pass_Word_Hash text NOT NULL,
Name           text NOT NULL,
Config         JSONB DEFAULT '{}'::JSONB NOT NULL,
Created_At     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
Is_Enabled     BOOLEAN DEFAULT TRUE NOT NULL
);
-- SQLc converts snake_case to CamelCase

-- Simple image blob demo
CREATE TABLE gowebapp.images (
 Image_ID BIGSERIAL PRIMARY KEY,
 User_ID BIGINT NOT NULL,
 Content_Type TEXT NOT NULL DEFAULT 'image/png',
 Image_Data BYTEA NOT NULL
);

CREATE TABLE gowebapp.exercises (
Exercise_ID   BIGSERIAL PRIMARY KEY,
Exercise_Name text NOT NULL
);


CREATE TABLE gowebapp.workouts (
Workout_ID  BIGSERIAL PRIMARY KEY,
User_ID     BIGINT NOT NULL,
Set_ID      BIGINT NOT NULL,
Start_Date  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE gowebapp.sets (
Set_ID      BIGSERIAL PRIMARY KEY,
Exercise_ID BIGINT NOT NULL,
Weight      INT NOT NULL DEFAULT 0 -- this can go up in decimal amounts so we can just divide/multiply to stay easy
);