CREATE TABLE `taskinfo` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `subject` VARCHAR(64) NOT NULL,
    `subtitle` TEXT NOT NULL,
    `filetype` VARCHAR(64) NOT NULL,
    `start` DATE NOT NULL,
    `end` DATE NOT NULL
);

CREATE TABLE `userinfo` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `email` TEXT NOT NULL UNIQUE,
    `studentid` TEXT NOT NULL UNIQUE
);

CREATE TABLE `uploadlog` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `time` DATE NOT NULL,
    `submitto` INTEGER NOT NULL,
    `email` TEXT NOT NULL,
    `filemd5` TEXT NOT NULL
);
