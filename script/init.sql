CREATE DATABASE IF NOT EXISTS shiika;
USE shiika;

CREATE TABLE IF NOT EXISTS user (name text NOT NULL, password text DEFAULT NULL)DEFAULT CHARSET=utf8mb4;

CREATE TABLE  IF NOT EXISTS kaminoku (
  `id` varchar(36)  NOT NULL,
  `first` text NOT NULL,
  `second` text NOT NULL,
  `third` text NOT NULL,
  `userid` text DEFAULT NULL
)DEFAULT CHARSET=utf8mb4;

CREATE TABLE  IF NOT EXISTS simonoku (
  `id` varchar(36) NOT NULL,
  `fourth` text NOT NULL,
  `fifth` text NOT NULL,
  `kaminokuid` varchar(36) DEFAULT NULL,
  `userid` text DEFAULT NULL
)DEFAULT CHARSET=utf8mb4;

/*
ALTER TABLE `kaminoku`
  ADD CONSTRAINT `userid` FOREIGN KEY (`userid`) REFERENCES `user` (`id`) ON UPDATE CASCADE;

ALTER TABLE `simonoku`
  ADD CONSTRAINT `userid` FOREIGN KEY (`userid`) REFERENCES `user` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `kaminokuid` FOREIGN KEY (`kaminokuid`) REFERENCES `kaminoku` (`id`) ON UPDATE CASCADE;
  */