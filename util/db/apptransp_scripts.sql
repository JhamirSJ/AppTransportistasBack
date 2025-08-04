USE sanjorge_transportistas_db;

SELECT * FROM banco
SELECT * FROM deposito
SELECT * FROM guia
SELECT * FROM producto
SELECT * FROM prueba_entrega

SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE banco;
TRUNCATE TABLE guia;
TRUNCATE TABLE producto;
TRUNCATE TABLE deposito;
TRUNCATE TABLE prueba_entrega;
SET FOREIGN_KEY_CHECKS = 1;

