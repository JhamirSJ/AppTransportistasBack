INSERT INTO guia (numero, fecha, codigo_cliente, nombre_cliente, nro_comprobante, importe_x_cobrar)
VALUES 
('G-001', '2025-06-25', 'CL001', 'Bodega María', 'F001-0001', 145.90),
('G-002', '2025-06-26', 'CL002', 'Tienda Los Andes', 'F001-0002', 210.50),
('G-003', '2025-06-27', 'CL003', 'Minimarket Sol', 'F001-0003', 89.90),
('G-004', '2025-06-27', 'CL004', 'Bodega La Grande', 'F001-0004', 320.75),
('G-005', '2025-06-27', 'CL005', 'Distribuidora Lima Norte', 'F001-0005', 156.40);

INSERT INTO producto (id_guia, nombre, cantidad) VALUES
(1, 'Galleta Rellenitas 150g', 5),
(1, 'Galleta Soda 250g', 3),
(1, 'Galleta Animalitos 180g', 2);

INSERT INTO producto (id_guia, nombre, cantidad) VALUES
(2, 'Galleta Naranja 150g', 4),
(2, 'Galleta Integral 120g', 10),
(2, 'Galleta Tradicional 200g', 6);

INSERT INTO producto (id_guia, nombre, cantidad) VALUES
(3, 'Galleta Limón 120g', 2),
(3, 'Galleta Soda 250g', 3);

INSERT INTO producto (id_guia, nombre, cantidad) VALUES
(4, 'Panetón Clásico 900g', 10),
(4, 'Galleta Jengibre 150g', 30),
(4, 'Galleta Especial 200g', 5);

INSERT INTO producto (id_guia, nombre, cantidad) VALUES
(5, 'Galleta Vainilla 150g', 8),
(5, 'Galleta Animalitos 180g', 4),
(5, 'Galleta Naranja 150g', 6);

INSERT INTO banco (nombre) VALUES
('BCP'),
('BBVA'),
('Interbank'),
('Scotiabank'),
('Banco de la Nación')
ON DUPLICATE KEY UPDATE nombre = nombre;
