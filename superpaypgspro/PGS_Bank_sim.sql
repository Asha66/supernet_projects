CREATE DATABASE  IF NOT EXISTS `PGS_Bank_Sim` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `PGS_Bank_Sim`;
-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: localhost    Database: PGS_Bank_Sim
-- ------------------------------------------------------
-- Server version	8.0.35-0ubuntu0.20.04.1

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
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account_holder_name` varchar(255) NOT NULL,
  `account_number` varchar(255) NOT NULL,
  `balance` double NOT NULL,
  `bankname` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL,
  `customerid` varchar(255) NOT NULL,
  `account_category` enum('PGS','MERCHANT','CUSTOMER') NOT NULL,
  `card_number` varchar(45) DEFAULT NULL,
  `password` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES (1,'Supernet','1123456789',10000,'icici','india','19876543211','PGS',NULL,'Test@123'),(2,'Merchant','1223456789',0,'icici','india','19876543221','MERCHANT',NULL,'Test@123'),(3,'Customer','1234567890',4899,'icici','india','29876543211','CUSTOMER','4312 1234 1234 1234','Test@123'),(4,'Customer','1234567891',6920,'icici','US','29876543212','CUSTOMER','3412 123456 78901','Test@123'),(5,'Customer','1234567892',19989.9,'icici','Ivory Coast','29876543213','CUSTOMER','5412 1234 1234 1234','Test@123');
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `account_txn`
--

DROP TABLE IF EXISTS `account_txn`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_txn` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `credit_acc_id` varchar(255) NOT NULL,
  `debit_acc_id` double NOT NULL,
  `amount` varchar(255) NOT NULL,
  `pgs_txn_id` varchar(255) NOT NULL,
  `description` varchar(45) NOT NULL,
  `created_at` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account_txn`
--

LOCK TABLES `account_txn` WRITE;
/*!40000 ALTER TABLE `account_txn` DISABLE KEYS */;
INSERT INTO `account_txn` VALUES (7,'1',4,'101','01887170297615397463','TXN_SUCCESS','2023-12-19 14:26:16'),(8,'1',5,'101','01847170297624485815','TXN_SUCCESS','2023-12-19 14:27:45'),(9,'1',3,'101','06059170297673782782','TXN_SUCCESS','2023-12-19 14:36:40');
/*!40000 ALTER TABLE `account_txn` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction` (
  `id` int NOT NULL AUTO_INCREMENT,
  `amount` double NOT NULL,
  `bank_transaction_id` varchar(255) NOT NULL,
  `pgs_txn_id` varchar(255) NOT NULL,
  `status` enum('TXN_SUCCESS','TXN_FAILED','SETTLEMENT_SUCCESS','SETTLEMENT_FAILED','REFUND_SUCCESS','REFUND_FAILED') NOT NULL,
  `transaction_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `description` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction`
--

LOCK TABLES `transaction` WRITE;
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
INSERT INTO `transaction` VALUES (9,101,'MHZcYLOT','03300170297518015965','TXN_FAILED','2023-12-19 08:39:57','Insufficient Balance'),(10,101,'j40FtzTm','04694170297567981886','TXN_FAILED','2023-12-19 08:50:19','Insufficient Balance'),(11,101,'kF18pALz','05081170297585184988','TXN_FAILED','2023-12-19 08:51:19','Insufficient Balance'),(12,101,'cgQA9XkF','05081170297604570937','TXN_FAILED','2023-12-19 08:54:30','Insufficient Balance'),(13,101,'pZdAqmVt','01887170297615397463','TXN_SUCCESS','2023-12-19 08:56:16','SUCCESS'),(14,101,'ikcjTncF','01847170297624485815','TXN_SUCCESS','2023-12-19 08:57:45','SUCCESS'),(15,101,'zeZ8IH2t','06059170297673782782','TXN_SUCCESS','2023-12-19 09:06:40','SUCCESS');
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'PGS_Bank_Sim'
--

--
-- Dumping routines for database 'PGS_Bank_Sim'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-19 15:08:57
