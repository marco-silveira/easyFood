/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=1 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

SET foreign_key_checks = 1;

-- Dumping database structure for easyfood
CREATE DATABASE IF NOT EXISTS `easyfood` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `easyfood`;

-- Dumping structure for table easyfood.categorias
CREATE TABLE IF NOT EXISTS `categorias` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `nome` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `categorias` (`id`, `nome`) VALUES
  (1, "Brasileira"),
  (2, "Massas"),
  (3, "Japonesa");

CREATE TABLE IF NOT EXISTS `dim_cidade` (
    `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
    `codigo` INT(7) unsigned NOT NULL DEFAULT 0,
    `nome` CHAR(255) NOT NULL DEFAULT '0',
    `uf` CHAR(2) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

LOAD DATA INFILE 'easyfood/cidades/cidades.csv'
    INTO TABLE `dim_cidade`
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    IGNORE 1 ROWS;

-- Dumping structure for table easyfood.restaurantes
CREATE TABLE IF NOT EXISTS `restaurantes` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `horario_abertura` TIME,
  `horario_fechamento` TIME,
  `id_cidade` INT(11) unsigned,
  `dias_funcionamento` TINYINT(3) unsigned NOT NULL COMMENT "Representacao binaria : 1 = segunda , 2 = terca, 4 = quarta etc",
  `nome` varchar(32) NOT NULL,
  `descricao` TEXT,
  `telefone` varchar(11) NOT NULL,
  `endereco` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id_cidade` (`id_cidade`),
  CONSTRAINT `restaurante_cidade_fk` FOREIGN KEY (`id_cidade`) REFERENCES `dim_cidade` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

INSERT INTO `restaurantes` (`horario_abertura` ,`horario_fechamento` ,`dias_funcionamento` ,`id_cidade` ,`nome` ,`descricao` ,`telefone` ,`endereco`) VALUES
	("15:00:00", "21:30:00", 3, 2, "Restaurante do zé", "Melhor comida feita pelo zé", "31985467513", "Rua das flores, numero 12, bairro Sagrada Familia"),
	("10:45:00", "16:00:00", 9, 3, "Maria das Massas", "Massas artesanais", "33985467513", "Rua das flores, numero 12, bairro Sagrada Familia");

-- Dumping data for table easyfood.categorias: ~0 rows (approximately)
/*!40000 ALTER TABLE `categorias` DISABLE KEYS */;
/*!40000 ALTER TABLE `categorias` ENABLE KEYS */;

CREATE TABLE IF NOT EXISTS `restaurante-categoria` (
  `id_restaurante` INT(11) unsigned NOT NULL,
  `id_categoria` INT(11) unsigned NOT NULL,
  PRIMARY KEY (`id_restaurante`, `id_categoria`),
  CONSTRAINT `restaurante_fk` FOREIGN KEY (`id_restaurante`) REFERENCES `restaurantes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `categoria_fk` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

-- Dumping structure for table easyfood.pratos
CREATE TABLE IF NOT EXISTS `pratos` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `id_restaurante` INT(11) unsigned NOT NULL DEFAULT 0,
  `id_categoria` INT(11) unsigned DEFAULT NULL,
  `nome` char(50) NOT NULL DEFAULT '0',
  `preco` decimal(10,2) unsigned NOT NULL DEFAULT 0.00,
  `tempo_de_preparo` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT 'MINUTOS',
  PRIMARY KEY (`id`),
  KEY `id_restaurante` (`id_restaurante`),
  KEY `id_categoria` (`id_categoria`),
  CONSTRAINT `pratos_categoria_fk` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `pratos_restaurante_fk` FOREIGN KEY (`id_restaurante`) REFERENCES `restaurantes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table easyfood.pratos: ~0 rows (approximately)
/*!40000 ALTER TABLE `pratos` DISABLE KEYS */;
INSERT INTO `pratos` (`id`, `id_restaurante`, `id_categoria`, `nome`, `preco`, `tempo_de_preparo`) VALUES
	(1, 1, NULL, 'tilápia', 45.00, 30),
	(2, 1, NULL, 'batata frita', 15.00, 15),
	(3, 1, NULL, 'batata frita com queijo', 18.00, 16),
	(4, 2, NULL, 'iscas de frango acebolada', 20.00, 20),
	(5, 2, NULL, 'filé parmegiana', 40.00, 40);
/*!40000 ALTER TABLE `pratos` ENABLE KEYS */;

-- Dumping data for table easyfood.restaurantes: ~0 rows (approximately)
/*!40000 ALTER TABLE `restaurantes` DISABLE KEYS */;
/*!40000 ALTER TABLE `restaurantes` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

CREATE TABLE IF NOT EXISTS `usuarios` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `primeiro_nome` varchar(32) NOT NULL,
  `ultimo_nome` varchar(32) NOT NULL,
  `telefone` varchar(11) NOT NULL,
  `email` varchar(50) NOT NULL UNIQUE,
  `senha_hash` BINARY(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

INSERT INTO `usuarios` (`primeiro_nome`, `ultimo_nome`, `telefone`, `email`, `senha_hash`) VALUES
  ('Joao', 'Pedro', '31984464729', 'joaopedro@gmail.com', SHA1('joaopedro2010')),
  ('Maria', 'Lima', '35987432164', 'marialima@hotmail.com', SHA1('ml15122015')),
  ('Carlos', 'Antunes', '37984455792', 'carlos_antunes12@outlook.com', SHA1('carlitos1212'));

-- Dumping structure for table easyfood.reviews
CREATE TABLE IF NOT EXISTS `reviews` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `id_usuario` INT(11) unsigned NOT NULL DEFAULT 0,
  `id_restaurante` INT(11) unsigned NOT NULL DEFAULT 0,
  `nota` TINYINT(1) unsigned NOT NULL DEFAULT 0,
  `comentario` TEXT,
  PRIMARY KEY (`id`),
  KEY `id_usuario` (`id_usuario`),
  KEY `id_restaurante` (`id_restaurante`),
  CONSTRAINT `review_id_usuario_fk` FOREIGN KEY (`id_usuario`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `review_id_restaurante_fk` FOREIGN KEY (`id_restaurante`) REFERENCES `restaurantes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

INSERT INTO `reviews` (`id_usuario`, `id_restaurante`, `nota`, `comentario`) VALUES
	(1, 1, 5, "Melhor restaurante da região."),
	(3, 2, 3, "Bons pratos, péssimo atendimento.");

INSERT INTO `restaurante-categoria` (`id_restaurante`, `id_categoria`) VALUES
	(1, 1),
	(2, 2),
  (1, 2),
  (2, 1);

CREATE TABLE IF NOT EXISTS `dim_cidade` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `codigo` INT(7) unsigned NOT NULL DEFAULT 0,
  `nome` CHAR(255) NOT NULL DEFAULT '0',
  `uf` CHAR(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

LOAD DATA INFILE 'easyfood/cidades/cidades.csv'
	INTO TABLE `dim_cidade`
	FIELDS TERMINATED BY ','
	LINES TERMINATED BY '\n'
	IGNORE 1 ROWS;
