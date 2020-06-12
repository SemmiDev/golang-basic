CREATE TABLE `mahasiswa` ( 
    `id` Int( 8 ) AUTO_INCREMENT NOT NULL,
    `nim` BigInt( 14 ) NOT NULL,
    `name` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `semester` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    CONSTRAINT `unique_id` UNIQUE( `id` ) )
CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB
AUTO_INCREMENT = 3;