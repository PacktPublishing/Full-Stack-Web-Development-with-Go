CREATE SCHEMA IF NOT EXISTS gowebapp;

-- ************************************** gowebapp.users

CREATE TABLE gowebapp.users
(
    User_ID        bigserial NOT NULL,
    User_Name      text NOT NULL,
    Pass_Word_Hash text NOT NULL,
    Name           text NOT NULL,
    Config         jsonb NOT NULL DEFAULT '{}'::JSONB,
    Created_At     timestamp NOT NULL DEFAULT NOW(),
    Is_Enabled     boolean NOT NULL DEFAULT TRUE,
    CONSTRAINT PK_users PRIMARY KEY ( User_ID )
);


-- ************************************** gowebapp.exercises

CREATE TABLE gowebapp.exercises
(
    Exercise_ID   bigserial NOT NULL,
    Exercise_Name text NOT NULL,
    CONSTRAINT PK_exercises PRIMARY KEY ( Exercise_ID )
);


-- ************************************** gowebapp.images

CREATE TABLE gowebapp.images
(
    Image_ID     bigserial NOT NULL,
    User_ID      bigserial NOT NULL,
    Content_Type text NOT NULL DEFAULT 'image/png',
    Image_Data   bytea NOT NULL,
    CONSTRAINT PK_images PRIMARY KEY ( Image_ID, User_ID ),
    CONSTRAINT FK_65 FOREIGN KEY ( User_ID ) REFERENCES gowebapp.users ( User_ID )
);

CREATE INDEX FK_67 ON gowebapp.images
    (
     User_ID
        );


-- ************************************** gowebapp.sets

CREATE TABLE gowebapp.sets
(
    Set_ID      bigserial NOT NULL,
    Exercise_ID bigserial NOT NULL,
    Weight      int NOT NULL DEFAULT 0,
    CONSTRAINT PK_sets PRIMARY KEY ( Set_ID, Exercise_ID ),
    CONSTRAINT FK_106 FOREIGN KEY ( Exercise_ID ) REFERENCES gowebapp.exercises ( Exercise_ID )
);

CREATE INDEX FK_108 ON gowebapp.sets
    (
     Exercise_ID
        );

-- ************************************** gowebapp.workouts

CREATE TABLE gowebapp.workouts
(
    Workout_ID  bigserial NOT NULL,
    Set_ID    bigserial NOT NULL,
    User_ID   bigserial NOT NULL,
    Exercise_ID bigserial NOT NULL,
    Start_Date  timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT PK_workouts PRIMARY KEY ( Workout_ID, Set_ID, User_ID, Exercise_ID ),
    CONSTRAINT FK_71 FOREIGN KEY ( Set_ID, Exercise_ID ) REFERENCES gowebapp.sets ( Set_ID, Exercise_ID ),
    CONSTRAINT FK_74 FOREIGN KEY ( User_ID ) REFERENCES gowebapp.users ( User_ID )
);

CREATE INDEX FK_73 ON gowebapp.workouts
    (
     Set_ID,
     Exercise_ID
        );

CREATE INDEX FK_76 ON gowebapp.workouts
    (
     User_ID
        );