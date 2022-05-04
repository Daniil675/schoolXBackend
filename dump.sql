-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by phantoms at 04-05-2022 11:33.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.


-- BEGIN TABLE City
DROP TABLE IF EXISTS City;
CREATE TABLE `City` (
  `id` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Inserting 5 rows into City
-- Insert batch #1
INSERT INTO City (id, Name) VALUES
(1, 'Shakhty'),
(2, 'Los Angeles'),
(3, 'Moscow'),
(4, 'Sochi'),
(5, 'Tumen');

-- END TABLE City

-- BEGIN TABLE Users
DROP TABLE IF EXISTS Users;
CREATE TABLE `Users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `Phone` varchar(15) DEFAULT NULL,
  `City_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Inserting 10 rows into Users
-- Insert batch #1
INSERT INTO Users (id, Name, Email, Phone, City_id) VALUES
(1, 'Daniil', 'daniil.popenko@yandex.ru', '89287750060', 1),
(2, 'Vlad', 'Vlad@yandex.ru', '1', 2),
(3, 'Nikita', 'Nikita@yandex.ru', '2', 3),
(4, 'Mark', 'Mark@yandex.ru', '3', 4),
(5, 'Nastya', 'Nastya@yandex.ru', '4', 5),
(6, 'Valya', 'Valya@yandex.ru', '5', 1),
(7, 'Dima', 'Dima@yandex.ru', '6', 2),
(8, 'Maxim', 'Maxim@yandex.ru', '7', 3),
(9, 'Andrey', 'Andrey@yandex.ru', '8', 4),
(10, 'Masha', 'Masha@yandex.ru', '9', 5);

-- END TABLE Users

