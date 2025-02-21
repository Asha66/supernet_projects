-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: PGS_Recon
-- ------------------------------------------------------
-- Server version	8.0.40-0ubuntu0.22.04.1

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
-- Table structure for table `Bank_Transactions`
--

DROP TABLE IF EXISTS `Bank_Transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Bank_Transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `bank_txn_date` varchar(45) DEFAULT NULL,
  `cnps_txn_date` varchar(45) DEFAULT NULL,
  `cnps_txn_num` varchar(45) DEFAULT NULL,
  `bank_txn_num` varchar(45) DEFAULT NULL,
  `cnps_company_num` varchar(45) DEFAULT NULL,
  `declaration_type` varchar(45) DEFAULT NULL,
  `payment_method` varchar(45) DEFAULT NULL,
  `txn_type` varchar(45) DEFAULT NULL,
  `amount` varchar(45) DEFAULT NULL,
  `currency` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `description` varchar(45) DEFAULT NULL,
  `reserved1` varchar(45) DEFAULT NULL,
  `reserved2` varchar(45) DEFAULT NULL,
  `recon_file_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Bank_Transactions_1_idx` (`recon_file_id`),
  CONSTRAINT `fk_Bank_Transactions_1` FOREIGN KEY (`recon_file_id`) REFERENCES `Recon_Files` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Bank_Transactions`
--

LOCK TABLES `Bank_Transactions` WRITE;
/*!40000 ALTER TABLE `Bank_Transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `Bank_Transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Bank_Transactions_Pro`
--

DROP TABLE IF EXISTS `Bank_Transactions_Pro`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Bank_Transactions_Pro` (
  `id` int NOT NULL AUTO_INCREMENT,
  `recon_status` varchar(45) DEFAULT NULL,
  `recon_description` varchar(45) DEFAULT NULL,
  `recon_date_time` timestamp NULL DEFAULT NULL,
  `recon_file_id` int DEFAULT NULL,
  `bank_id` int NOT NULL,
  `txn_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_Bank_Transactions_Pro_1_idx` (`recon_file_id`),
  KEY `fk_Bank_Transactions_Pro_2_idx` (`txn_id`),
  KEY `fk_Bank_Transactions_Pro_3_idx` (`bank_id`),
  CONSTRAINT `fk_Bank_Transactions_Pro_1` FOREIGN KEY (`recon_file_id`) REFERENCES `Recon_Files` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_Bank_Transactions_Pro_2` FOREIGN KEY (`txn_id`) REFERENCES `Transactions` (`id`),
  CONSTRAINT `fk_Bank_Transactions_Pro_3` FOREIGN KEY (`bank_id`) REFERENCES `Bank_Transactions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Bank_Transactions_Pro`
--

LOCK TABLES `Bank_Transactions_Pro` WRITE;
/*!40000 ALTER TABLE `Bank_Transactions_Pro` DISABLE KEYS */;
/*!40000 ALTER TABLE `Bank_Transactions_Pro` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Recon_Files`
--

DROP TABLE IF EXISTS `Recon_Files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Recon_Files` (
  `id` int NOT NULL AUTO_INCREMENT,
  `file_name` varchar(45) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` varchar(45) DEFAULT NULL,
  `created_by` varchar(45) DEFAULT NULL,
  `updated_by` varchar(45) DEFAULT NULL,
  `status` int NOT NULL COMMENT '1=>File Read Success , 2=> invalid file,3=>UnKnownError',
  `remarks` varchar(45) DEFAULT NULL,
  `file_records` int DEFAULT NULL,
  `file_size` int DEFAULT NULL COMMENT 'File Size in Bytes',
  `response_dump` longtext,
  `bank_name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Recon_Files`
--

LOCK TABLES `Recon_Files` WRITE;
/*!40000 ALTER TABLE `Recon_Files` DISABLE KEYS */;
/*!40000 ALTER TABLE `Recon_Files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Transactions`
--

DROP TABLE IF EXISTS `Transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `transaction_time` timestamp NULL DEFAULT NULL,
  `transaction_deatils` text,
  `cnps_txn_number` varchar(45) DEFAULT NULL,
  `pg_txn_number` varchar(45) DEFAULT NULL,
  `bank_txn_number` varchar(45) DEFAULT NULL,
  `entity_id` varchar(45) DEFAULT NULL,
  `entity_name` varchar(155) DEFAULT NULL,
  `operator` varchar(45) DEFAULT NULL,
  `channel` varchar(45) DEFAULT NULL,
  `amount` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Transactions`
--

LOCK TABLES `Transactions` WRITE;
/*!40000 ALTER TABLE `Transactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `Transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'PGS_Recon'
--

--
-- Dumping routines for database 'PGS_Recon'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-15 16:12:42
