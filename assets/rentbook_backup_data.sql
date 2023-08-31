-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: localhost    Database: rentbook
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `books` (
  `book_id` varchar(255) NOT NULL,
  `book_name` varchar(255) DEFAULT NULL,
  `book_publisher` varchar(255) DEFAULT NULL,
  `book_author` varchar(255) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`book_id`),
  KEY `fk_userbook` (`user_id`),
  CONSTRAINT `fk_userbook` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES ('2eeb296f-313e-4254-ac9e-b71b4642cb50','Golang Programming part1','Gramedia','nikola',0,'2023-08-30 23:21:23','2023-08-30 23:21:23','0e2b13a3-70d7-4f78-beeb-2e223b2dfdea'),('40258561-3534-4e9d-aa02-160d25c0a91c','C# Programming part1','Erlangga','nikola',0,'2023-08-30 23:23:26','2023-08-30 23:23:26','7dd5f83f-cbb5-4df6-9ac7-25a9217afe10'),('4c988bc5-b61d-428f-9257-16dead18d493','Ruby Programming part1','Bukunesia','nikola',0,'2023-08-30 23:22:21','2023-08-30 23:22:21','0e2b13a3-70d7-4f78-beeb-2e223b2dfdea'),('6cc2128d-1988-4a9c-bd65-103362aa299d','C++ Programming part1','Erlangga','einstein',0,'2023-08-30 23:20:00','2023-08-30 23:20:00','6ea79d28-7e70-4ad9-a49b-956c37e5ce26'),('7466d035-6248-4d2a-8229-5e8e8d7b49df','NodeJs Programming part1','Gramedia','rudolf',0,'2023-08-30 23:24:40','2023-08-30 23:24:40','7dd5f83f-cbb5-4df6-9ac7-25a9217afe10'),('792ce93a-67ee-4b4d-a8ca-f0a7c92fb286','Python Programming part2','Bukunesia','einstein',0,'2023-08-30 23:22:01','2023-08-30 23:22:01','0e2b13a3-70d7-4f78-beeb-2e223b2dfdea'),('d51ac5ec-33d5-4a08-8d7e-d25f647b790d','C Programming part1','Bukunesia','rudolf',0,'2023-08-30 23:24:23','2023-08-30 23:24:23','7dd5f83f-cbb5-4df6-9ac7-25a9217afe10'),('f729ca57-29ea-4735-bdc6-f40a4e89c8bf','Java Programming part1','Erlangga','einstein',0,'2023-08-30 23:23:53','2023-08-30 23:23:53','7dd5f83f-cbb5-4df6-9ac7-25a9217afe10'),('fe006e50-79ce-4aa2-ae94-46bf180592e7','Python Programming part1','Gramedia','rudolf',0,'2023-08-30 23:20:30','2023-08-30 23:20:30','6ea79d28-7e70-4ad9-a49b-956c37e5ce26');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rents`
--

DROP TABLE IF EXISTS `rents`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rents` (
  `rent_id` varchar(255) NOT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `book_id` varchar(255) DEFAULT NULL,
  `rent_start_date` timestamp NULL DEFAULT NULL,
  `rent_end_date` timestamp NULL DEFAULT NULL,
  `rent_status` varchar(255) DEFAULT NULL,
  `rent_qty` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`rent_id`),
  KEY `fk_userrent` (`user_id`),
  KEY `fk_bookrent` (`book_id`),
  CONSTRAINT `fk_bookrent` FOREIGN KEY (`book_id`) REFERENCES `books` (`book_id`),
  CONSTRAINT `fk_userrent` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rents`
--

LOCK TABLES `rents` WRITE;
/*!40000 ALTER TABLE `rents` DISABLE KEYS */;
INSERT INTO `rents` VALUES ('096523d2-5dd1-4e8e-b090-6ace511b7002','0e2b13a3-70d7-4f78-beeb-2e223b2dfdea','7466d035-6248-4d2a-8229-5e8e8d7b49df','2023-08-30 22:30:10','2023-09-01 16:01:10','Active',3,'2023-08-30 23:28:28','2023-08-30 16:28:53'),('ca5d0b47-adf2-44d3-afba-5296a2b1e9a3','7dd5f83f-cbb5-4df6-9ac7-25a9217afe10','2eeb296f-313e-4254-ac9e-b71b4642cb50','2023-08-30 23:30:10','2023-09-01 15:01:10','Active',2,'2023-08-30 23:27:00','2023-08-30 16:40:25'),('cc4cd481-7663-4dbe-8797-580913387b91','6ea79d28-7e70-4ad9-a49b-956c37e5ce26','792ce93a-67ee-4b4d-a8ca-f0a7c92fb286','2023-08-28 22:30:10','2023-08-30 16:01:10','Expired',1,'2023-08-30 23:30:44','2023-08-30 16:34:39');
/*!40000 ALTER TABLE `rents` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` varchar(255) NOT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `user_email` varchar(255) DEFAULT NULL,
  `user_password` varchar(255) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('0e2b13a3-70d7-4f78-beeb-2e223b2dfdea','fikri','fikri@mailsac.com','$2a$10$de.M1nHcaBXU2Xxyxm6cPO/qH6YD66VZRcsO97naajMvXG6dAHY5e',0,'2023-08-30 23:09:48','2023-08-30 16:14:37'),('6ea79d28-7e70-4ad9-a49b-956c37e5ce26','ahmad','ahmad@mailsac.com','$2a$10$de.M1nHcaBXU2Xxyxm6cPO/qH6YD66VZRcsO97naajMvXG6dAHY5e',0,'2023-08-30 23:09:23','2023-08-30 23:09:23'),('7dd5f83f-cbb5-4df6-9ac7-25a9217afe10','yusnar','yusnar@mailsac.com','$2a$10$NMcuyONdrPWEjUaS0Fr4oO2qFiLOEZ2WZWXvudhz2o5QNursT.AYa',0,'2023-08-30 23:08:54','2023-08-30 16:11:46');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-31  6:43:35
