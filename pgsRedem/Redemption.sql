-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: Redemption
-- ------------------------------------------------------
-- Server version	8.0.39-0ubuntu0.22.04.1

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
-- Table structure for table `Error_Codes`
--

DROP TABLE IF EXISTS `Error_Codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Error_Codes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(45) DEFAULT NULL,
  `language` varchar(45) DEFAULT NULL,
  `message` varchar(100) DEFAULT NULL,
  `desc` varchar(45) DEFAULT NULL,
  `create_date` varchar(45) DEFAULT NULL,
  `last_update` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Error_Codes`
--

LOCK TABLES `Error_Codes` WRITE;
/*!40000 ALTER TABLE `Error_Codes` DISABLE KEYS */;
INSERT INTO `Error_Codes` VALUES (1,'0010','en','invalid Cnps_entity_id','test','',''),(2,'0010','fr','Identifiant Employeur incorrecte','','',''),(3,'0011','en','invalid Entity_name','','',''),(4,'0011','fr','Nom d\'Entreprise incorrecte','','',''),(5,'0012','en','invalid Enity_mobile','','',''),(6,'0012','fr','Nom d\'operateur Télécoms incorrecte','','',''),(7,'0013','en','invalid Amount','','',''),(8,'0013','fr','Montant incorrecte','','',''),(9,'0014','en','invalid Currency','','',''),(10,'0014','fr','Devise incorrecte','','',''),(11,'0015','en','invalid Cnps_transaction_id','','',''),(12,'0015','fr','Identifiant de transaction CNPS incorrecte','','',''),(13,'0016','en','invalid Partial_payment','','',''),(14,'0016','fr','Paiement partiel incorrecte','','',''),(15,'0017','en','invalid Transaction_start_date','','',''),(16,'0017','fr','Date de début de transaction incorrecte','','',''),(17,'0018','en','invalid Notification_url','','',''),(18,'0018','fr','URL de Notification incorrecte','','',''),(19,'0019','en','invalid Return_url','','',''),(20,'0019','fr','URL  Retour ne doit pas être vide','','',''),(21,'0020','en','invalid Cancel_url','','',''),(22,'0020','fr','URL d\'Annulation incorrecte','','',''),(23,'0021','en','invalid Sign','','',''),(24,'0021','fr','Signature incorrecte','','',''),(25,'0022','en','invalid Description','','',''),(26,'0022','fr','Description incorrecte','','',''),(27,'0023','en','Merchant Not Registered with PG','','',''),(28,'0023','fr','Le partenaire n\'est pas enregistré sur la PG','','',''),(29,'0024','en','Merchant is Inactive ','','',''),(30,'0024','fr','Le partenaire est inactif','','',''),(31,'0027','en','Batch sign not matched','','',''),(32,'0027','fr','Le signe de lot ne correspond pas','','',''),(33,'0028','en','Atleast one Transaction sign not matched','','',''),(34,'0028','fr ','Au moins un signe de transaction ne correspond pas','','',''),(35,'0029','en','invalid cnps txn id','','',''),(36,'0029','fr','Id txn cnps incorrecte','','',''),(37,'0030','en','invalid user id','','',''),(38,'0030','fr','id utilisateur incorrecte','','',''),(39,'0031','en','parameter not match','','',''),(40,'0031','fr','Le parametre ne concorde pas','','',''),(41,'0032','en','version cannot be empty','','',''),(42,'0032','fr','Le champ Version ne peut être vide','','',''),(43,'0033','en','invalid version','','',''),(44,'0033','fr','Version incorrecte','','',''),(45,'0034','en','General Error :','','',''),(46,'0034','fr','Erreur Système :','','',''),(47,'0035','en','invalid Email','','',''),(48,'0035','fr','Email incorrecte','','',''),(49,'0036','en','invalid cnps declaration number','','',''),(50,'0036','fr','Le numéro de déclaration CNPS est incorrecte','','',''),(51,'0037','en','invalid decleration period','','',''),(52,'0037','fr','La période de déclaration est incorrecte','','',''),(53,'0038','en','invalid nature code','','',''),(54,'0038','fr','Le code nature est incorrecte','','',''),(55,'0039','en','invalid nature name  ','','',''),(56,'0039','fr','le nom de la nature est incorrecte','','',''),(57,'0040','en','Invalid payment method','','',''),(58,'0040','fr','Organisme Payeur invalide','','',''),(59,'0041','en','Invalid payment mode','','',''),(60,'0041','fr','Mode de paiement invalide','','',''),(61,'0042','en','Invalid Transacion end date','','',''),(62,'0042','fr','Date de fin de Transaction incorrecte','','',''),(63,'0043','en','Invalid customer mobile number','','',''),(64,'0043','fr','Numero de mobile du client incorrect','','',''),(65,'0044','en','Channel is not active','','',''),(66,'0044','fr','Canal inactif','','',''),(67,'0045','en','Benefits cannot be empty','','',''),(68,'0045','fr','Les avantages ne peuvent pas être vides','','',''),(69,'0046','en','Invalid User Id','','',''),(70,'0046','fr','Identifiant invalide','','',''),(71,'0047','en','Benefit type cannot be empty','','',''),(72,'0047','fr','Le type d\'avantage ne peut pas être vide','','',''),(74,'0048','en','Invalid CNPS transaction id','','',''),(75,'0048','fr','Identifiant de transaction CNPS invalide','','',''),(76,'0049','en','Invalid cnps entity id','','',''),(77,'0049','fr','ID d\'entité cnps non valide','','',''),(78,'0050','en','Invalid entity name','','',''),(79,'0050','fr','Nom d\'entité non valide','','',''),(80,'0051','en','Invalid entity phone','','',''),(81,'0051','fr','Numéro de téléphone de l\'entité non valide','','',''),(82,'0052','en','Invalid Benefit Code','','',''),(83,'0052','fr','Code d\'avantage invalide','','',''),(84,'0053','en','Invalid benefit label','','',''),(85,'0053','fr','Libellé d\'avantage non valide','','',''),(86,'0054','en','Professional category cannot be empty','','',''),(87,'0054','fr','La catégorie professionnelle ne peut pas être vide','','',''),(88,'0055','en','Invalid benefit account number','','',''),(89,'0055','fr','Numéro de compte d\'avantages invalide','','',''),(90,'0056','en','Invalid benefit msisdn','','',''),(91,'0056','fr','Msisdn d\'avantage non valide','','',''),(92,'0057','en','Invalid Benefit Amount','','',''),(93,'0057','fr','Montant de la prestation invalide','','',''),(94,'0058','en','Total amount is not equal to sum of benefit amount','','',''),(95,'0058','fr','Le montant total n\'est pas égal à la somme des prestations','','',''),(96,'0059','en','Batch already exists','','',''),(97,'0059','fr','Le lot existe déjà','','',''),(98,'0060','en','Invalid update status sent','','',''),(99,'0060','fr','État de mise à jour non valide envoyé','','',''),(100,'0061','en','Invalid CNPS application transaction status','','',''),(101,'0061','fr','Statut de transaction d\'application CNPS invalide','','',''),(102,'0062','en','Invalid Transaction Status','','',''),(103,'0062','fr','Statut de transaction invalide','','',''),(104,'0063','en','Transaction is in completed state , status cant be updated','','',''),(105,'0063','fr','La transaction est à l\'état terminé, le statut ne peut pas être mis à jour','','',''),(106,'0064','en','Cnps Transaction Id cannot be repeated in same batch','','',''),(107,'0064','fr','L\'identifiant de transaction Cnps ne peut pas être répété dans le même lot','','',''),(108,'0065','en','Batch with cnps application transaction id does not exist','','',''),(109,'0065','fr','Le lot avec l\'ID de transaction d\'application cnps n\'existe pas','','',''),(110,'0066','en','Cnps Transaction id does not exist in given batch','','',''),(111,'0066','fr','L\'identifiant de transaction Cnps n\'existe pas dans le lot donné','','',''),(112,'0067','en','Cheque Number cannot be null for transaction to  make Ready Offline',NULL,NULL,NULL),(113,'0067','fr','Le numéro de chèque ne peut pas être nul pour que la transaction soit prête hors ligne',NULL,NULL,NULL),(114,'0068','en','Invalid  Mode or Method for given batch',NULL,NULL,NULL),(115,'0068','fr','Mode ou méthode non valide pour un lot donné',NULL,NULL,NULL),(116,'0069','en','All transaction status in request must be same for the batch',NULL,NULL,NULL),(117,'0069','fr','Tous les statuts de transaction dans la demande doivent être identiques pour le lot',NULL,NULL,NULL),(118,'0070','en','Invalid Transaction count for batch',NULL,NULL,NULL),(119,'0070','fr','Nombre de transactions non valide pour le lot',NULL,NULL,NULL),(120,'0071','en','Invalid Bank Name',NULL,NULL,NULL),(121,'0071','fr','Nom de banque invalide',NULL,NULL,NULL),(122,'0072','en','Invalid Bank Code',NULL,NULL,NULL),(123,'0072','en','Code bancaire invalide',NULL,NULL,NULL),(124,'0073','en','Enter Payment Method',NULL,NULL,NULL),(125,'0073','fr','Entrez le mode de paiement',NULL,NULL,NULL),(126,'0074','en','Enter Payment Mode',NULL,NULL,NULL),(127,'0074','fr','Entrer en mode de paiement',NULL,NULL,NULL);
/*!40000 ALTER TABLE `Error_Codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `banks`
--

DROP TABLE IF EXISTS `banks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `banks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `bank_name` varchar(45) DEFAULT NULL,
  `bank_code` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `created_by` int DEFAULT NULL,
  `created_at` varchar(45) DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  `updated_at` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_banks_1_idx` (`created_by`,`updated_by`),
  KEY `fk_banks_2_idx` (`updated_by`),
  CONSTRAINT `fk_banks_1` FOREIGN KEY (`created_by`) REFERENCES `sysuser` (`id`),
  CONSTRAINT `fk_banks_2` FOREIGN KEY (`updated_by`) REFERENCES `sysuser` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `banks`
--

LOCK TABLES `banks` WRITE;
/*!40000 ALTER TABLE `banks` DISABLE KEYS */;
INSERT INTO `banks` VALUES (1,'Ascoma','121233','ACTIVE','rti@gmail.com',1,'2023-07-19 17:09:09',NULL,NULL),(2,'ECOBANK','523641','ACTIVE','eco@gmail.com',13,'2024-05-02 16:26:56',NULL,NULL);
/*!40000 ALTER TABLE `banks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `batch_txn_details`
--

DROP TABLE IF EXISTS `batch_txn_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `batch_txn_details` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cnps_batch_txn_app_id` varchar(45) DEFAULT NULL,
  `sys_bacth_id` varchar(45) DEFAULT NULL,
  `create_date` varchar(45) DEFAULT NULL,
  `txn_start_date` varchar(45) DEFAULT NULL,
  `total_amount` varchar(45) DEFAULT NULL,
  `benfit_type` varchar(45) DEFAULT NULL,
  `mode` varchar(45) DEFAULT NULL,
  `method` varchar(45) DEFAULT NULL,
  `sign` varchar(150) DEFAULT NULL,
  `batch_txn_count` varchar(45) DEFAULT NULL,
  `benfits` longtext,
  `status` enum('REDEM_ORDER_RECEIVED','REDEM_ORDER_SUCCESS','REDEM_ORDER_REJECT','REDEM_ORDER_CANCELLED') DEFAULT NULL,
  `description` varchar(45) DEFAULT NULL,
  `currency` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `batch_txn_details`
--

LOCK TABLES `batch_txn_details` WRITE;
/*!40000 ALTER TABLE `batch_txn_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `batch_txn_details` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `clients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `mobile_number` varchar(45) DEFAULT NULL,
  `status` varchar(45) DEFAULT NULL,
  `user_id` varchar(45) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `created_by` int DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  `created_at` varchar(45) DEFAULT NULL,
  `updated_at` varchar(45) DEFAULT NULL,
  `pass_key` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_clients_2_idx` (`updated_by`),
  KEY `fk_clients_1` (`created_by`),
  CONSTRAINT `fk_clients_1` FOREIGN KEY (`created_by`) REFERENCES `sysuser` (`id`),
  CONSTRAINT `fk_clients_2` FOREIGN KEY (`updated_by`) REFERENCES `sysuser` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clients`
--

LOCK TABLES `clients` WRITE;
/*!40000 ALTER TABLE `clients` DISABLE KEYS */;
INSERT INTO `clients` VALUES (16,'aone','kaushikp@proyava.com','1211111111','ACTIVE',NULL,'Test@123',1,NULL,'2023-07-20 11:45:23',NULL,NULL),(17,'new','kp@gmail.com','1211111111','ACTIVE','UID05511','Irc1UAbj2j4SyK+adRg9SVeb4kqGbmnkguRntiOntxS/jorjs3gFsWv5AioFiwGn',1,1,'2023-07-20 13:15:10','2023-07-20 16:47:18',NULL),(18,'qwss','W@gmail.com','1212311111','ACTIVE','kaushikp','Irc1UAbj2j4SyK+adRg9SVeb4kqGbmnkguRntiOntxS/jorjs3gFsWv5AioFiwGn',1,NULL,'2023-07-20 16:46:52',NULL,NULL),(19,'VK','vk18@gmail.com','8521448699','ACTIVE','vk18','Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9Hixkmo/MqK2a78qoL2tO55MrYG',13,NULL,'2024-10-02 16:26:39',NULL,NULL);
/*!40000 ALTER TABLE `clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `events` (
  `id` int NOT NULL AUTO_INCREMENT,
  `version` varchar(45) DEFAULT NULL,
  `create_date` varchar(45) DEFAULT NULL,
  `user_id` varchar(45) DEFAULT NULL,
  `cnps_app_txn_id` varchar(45) DEFAULT NULL,
  `event` longtext,
  `request` longtext,
  `response` longtext,
  `status` varchar(45) DEFAULT NULL,
  `request_type` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=815 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(100) DEFAULT NULL,
  `privilege` longtext,
  `status` varchar(100) DEFAULT NULL,
  `created_at` varchar(100) DEFAULT NULL,
  `updated_at` varchar(100) DEFAULT NULL,
  `created_by` int DEFAULT NULL,
  `updated_by` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_roles_1_idx` (`created_by`),
  KEY `fk_roles_2_idx` (`updated_by`),
  CONSTRAINT `fk_roles_1` FOREIGN KEY (`created_by`) REFERENCES `sysuser` (`id`),
  CONSTRAINT `fk_roles_2` FOREIGN KEY (`updated_by`) REFERENCES `sysuser` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'SYSADMIN','{\"Menus\":[\"dashboard-menu\",\"sysusermanagement-menu\",\"systemmanagement-menu\",\"clientsmanagement-menu\",\"bankmanagement-menu\",\"reports-menu\"],\"Submenus\":[\"searchsysuser-active\",\"searchrole-active\",\"searchclientsmanagement-active\",\"Searchbank-active\",\"batchreport-active\",\"transactionreport-active\"]}','ACTIVE','2023-05-22 15:56:08.441822+05:30','2023-05-26 16:10:18',NULL,NULL),(2,'Adminrwee','{\"Menus\":[\"sysusermanagement-menu\",\"systemmanagement-menu\"],\"Submenus\":[\"searchsysuser-active\",\"searchrole-active\"]}','ACTIVE','2023-05-23 15:30:38',NULL,NULL,NULL),(3,'wwwSYSADMIN','{\"Menus\":[\"sysusermanagement-menu\"],\"Submenus\":[\"searchsysuser-active\"]}','ACTIVE','2023-05-23 15:31:21','2023-05-26 14:56:20',NULL,NULL),(4,'TestNew','{\"Menus\":[\"reports-menu\"],\"Submenus\":[\"transactionreport-active\"]}','ACTIVE','2023-05-23 16:19:38',NULL,NULL,NULL),(5,'rolefive','{\"Menus\":[\"systemmanagement-menu\"],\"Submenus\":[\"searchrole-active\"]}','ACTIVE','2023-05-23 16:26:12',NULL,NULL,NULL),(6,'guest','{\"Menus\":[\"sysusermanagement-menu\"],\"Submenus\":[\"searchsysuser-active\"]}','ACTIVE','2023-05-23 17:07:47',NULL,NULL,NULL),(7,'guestone','{\"Menus\":[\"sysusermanagement-menu\",\"systemmanagement-menu\",\"reports-menu\"],\"Submenus\":[\"searchsysuser-active\",\"searchrole-active\",\"batchreport-active\"]}','ACTIVE','2023-05-23 17:08:19','2023-05-26 16:31:17',NULL,NULL),(8,'UY','{\"Menus\":[\"sysusermanagement-menu\"],\"Submenus\":[\"searchsysuser-active\"]}','ACTIVE','2023-05-26 16:36:21',NULL,NULL,NULL),(9,'YYT','{\"Menus\":[\"sysusermanagement-menu\"],\"Submenus\":[\"searchsysuser-active\"]}','ACTIVE','2023-05-26 16:37:02',NULL,NULL,NULL),(10,'guestFFFSS','{\"Menus\":[\"sysusermanagement-menu\",\"systemmanagement-menu\",\"reports-menu\"],\"Submenus\":[\"searchsysuser-active\",\"searchrole-active\",\"batchreport-active\",\"transactionreport-active\"]}','ACTIVE','2023-05-26 16:39:27','2023-05-29 17:37:52',NULL,NULL),(11,'sysadminnew','{\"Menus\":[\"dashboard-menu\",\"sysusermanagement-menu\",\"systemmanagement-menu\",\"clientsmanagement-menu\",\"bankmanagement-menu\",\"reports-menu\"],\"Submenus\":[\"searchsysuser-active\",\"searchrole-active\",\"searchclientsmanagement-active\",\"Searchbank-active\",\"batchreport-active\",\"transactionreport-active\"]}','ACTIVE','2023-07-20 15:57:35',NULL,1,NULL);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sysuser`
--

DROP TABLE IF EXISTS `sysuser`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sysuser` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fullname` varchar(255) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `mobile` varchar(255) DEFAULT NULL,
  `address` varchar(256) DEFAULT NULL,
  `pass_count` varchar(255) DEFAULT NULL,
  `created_at` varchar(255) DEFAULT NULL,
  `updated_at` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `language` varchar(255) DEFAULT NULL,
  `password_set` varchar(255) DEFAULT NULL,
  `password_updated_date` varchar(255) DEFAULT NULL,
  `txn_id` varchar(255) DEFAULT NULL,
  `role_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sysuser_1_idx` (`role_id`),
  CONSTRAINT `fk_sysuser_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sysuser`
--

LOCK TABLES `sysuser` WRITE;
/*!40000 ALTER TABLE `sysuser` DISABLE KEYS */;
INSERT INTO `sysuser` VALUES (1,'demo','h9vueqgEoRwg6TKf3BFb+2rSQfgHMT2h+kZTqglI9Zd+GMTAquNihFuxQylYHKht','kaushikp@proyava.com','9872222210','AFRICA','0','2023-02-23 12:29:07','2023-06-08 15:52:08','ACTIVE',NULL,'1','english','1','2023-06-29 17:57:25',NULL,1),(13,'another','h9vueqgEoRwg6TKf3BFb+2rSQfgHMT2h+kZTqglI9Zd+GMTAquNihFuxQylYHKht','rti@gmail.com','2123113455','TT','0','2024-10-23 11:07:20','2024-10-24 13:24:21','ACTIVE','1','1','english','1','2024-10-23 11:07:20',NULL,1);
/*!40000 ALTER TABLE `sysuser` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction_deatils`
--

DROP TABLE IF EXISTS `transaction_deatils`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction_deatils` (
  `id` int NOT NULL AUTO_INCREMENT,
  `txn_number` varchar(45) DEFAULT NULL,
  `cnps_txn_id` varchar(45) DEFAULT NULL,
  `cnps_entity_id` varchar(45) DEFAULT NULL,
  `cnps_entity_name` varchar(45) DEFAULT NULL,
  `cnps_prestation_number` varchar(45) DEFAULT NULL,
  `prestation_period` varchar(45) DEFAULT NULL,
  `benefit_code` varchar(45) DEFAULT NULL,
  `benefit_label` varchar(45) DEFAULT NULL,
  `benefit_account_number` varchar(45) DEFAULT NULL,
  `benefit_msisdn` varchar(45) DEFAULT NULL,
  `benefit_amount` varchar(45) DEFAULT NULL,
  `bank_code` varchar(45) DEFAULT NULL,
  `cheque_number` varchar(45) DEFAULT NULL,
  `benefit_sign` varchar(150) DEFAULT NULL,
  `txn_status` enum('REDEM_APPROVED','REDEM_DECLINED','REDEM_PENDING','REDEM_CANCELLED','REDEM_READY','READM_READY_OFFLINE') DEFAULT NULL,
  `bank_status` varchar(45) DEFAULT NULL,
  `create_date` varchar(45) DEFAULT NULL,
  `benefit_details` longtext,
  `batch_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_batch_id_1_idx` (`batch_id`),
  CONSTRAINT `fk_batch_id_1` FOREIGN KEY (`batch_id`) REFERENCES `batch_txn_details` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction_deatils`
--

LOCK TABLES `transaction_deatils` WRITE;
/*!40000 ALTER TABLE `transaction_deatils` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaction_deatils` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `web_event`
--

DROP TABLE IF EXISTS `web_event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `web_event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `event` varchar(255) DEFAULT NULL,
  `created_on` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `web_event`
--

LOCK TABLES `web_event` WRITE;
/*!40000 ALTER TABLE `web_event` DISABLE KEYS */;
/*!40000 ALTER TABLE `web_event` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'Redemption'
--

--
-- Dumping routines for database 'Redemption'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-10-28 11:15:27
