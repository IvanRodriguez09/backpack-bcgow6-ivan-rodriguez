-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema empresa_internet
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema empresa_internet
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `empresa_internet` DEFAULT CHARACTER SET utf8 ;
USE `empresa_internet` ;

-- -----------------------------------------------------
-- Table `empresa_internet`.`InternetPlans`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `empresa_internet`.`InternetPlans` (
  `id` INT NOT NULL,
  `megas_velocity` INT NOT NULL,
  `price` FLOAT NOT NULL,
  `discount` INT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `empresa_internet`.`Clients`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `empresa_internet`.`Clients` (
  `dni` INT NOT NULL,
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `birthday` DATE NOT NULL,
  `providence` VARCHAR(45) NOT NULL,
  `city` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`dni`),
  UNIQUE INDEX `idClients_UNIQUE` (`dni` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `empresa_internet`.`Clients_has_InternetPlans`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `empresa_internet`.`Clients_has_InternetPlans` (
  `Clients_dni` INT NOT NULL,
  `InternetPlans_id` INT NOT NULL,
  PRIMARY KEY (`Clients_dni`, `InternetPlans_id`),
  INDEX `fk_Clients_has_InternetPlans_InternetPlans1_idx` (`InternetPlans_id` ASC) VISIBLE,
  INDEX `fk_Clients_has_InternetPlans_Clients_idx` (`Clients_dni` ASC) VISIBLE,
  CONSTRAINT `fk_Clients_has_InternetPlans_Clients`
    FOREIGN KEY (`Clients_dni`)
    REFERENCES `empresa_internet`.`Clients` (`dni`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Clients_has_InternetPlans_InternetPlans1`
    FOREIGN KEY (`InternetPlans_id`)
    REFERENCES `empresa_internet`.`InternetPlans` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
