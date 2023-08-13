CREATE TABLE `user` (
  `userID` int NOT NULL AUTO_INCREMENT,
  `Uname` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `admin` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`userID`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `books` (
  `bookID` int NOT NULL AUTO_INCREMENT,
  `BookName` varchar(255) NOT NULL,
  `quantity` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`bookID`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `request` (
  `reqID` int NOT NULL AUTO_INCREMENT,
  `UID` int DEFAULT NULL,
  `BookID` int DEFAULT NULL,
  `Status` int DEFAULT NULL,
  PRIMARY KEY (`reqID`),
  KEY `UID` (`UID`),
  KEY `BookID` (`BookID`),
  CONSTRAINT `request_ibfk_1` FOREIGN KEY (`UID`) REFERENCES `user` (`userID`),
  CONSTRAINT `request_ibfk_2` FOREIGN KEY (`BookID`) REFERENCES `books` (`bookID`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
