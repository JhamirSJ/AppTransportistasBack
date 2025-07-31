CREATE DATABASE IF NOT EXISTS sanjorge_despacho_db;
USE sanjorge_despacho_db;

-- Tabla banco
CREATE TABLE banco (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE
);

-- Tabla guía
CREATE TABLE guia (
    id INT AUTO_INCREMENT PRIMARY KEY,
    numero VARCHAR(50) NOT NULL,
    fecha DATE NOT NULL,
    codigo_cliente VARCHAR(50),
    nombre_cliente VARCHAR(100),
    nro_comprobante VARCHAR(50),
    importe_x_cobrar DECIMAL(10,2) NOT NULL,
    monto_cobrado DECIMAL(10,2) DEFAULT 0,
    entregada BOOLEAN DEFAULT FALSE
);

-- Tabla producto (relación con guía)
CREATE TABLE producto (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_guia INT NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    cantidad INT NOT NULL,
    FOREIGN KEY (id_guia) REFERENCES guia(id) ON DELETE CASCADE
);

-- Tabla depósito (relación con banco)
CREATE TABLE deposito (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nro_operacion VARCHAR(100) NOT NULL,
    fecha DATE NOT NULL,
    monto DECIMAL(10,2) NOT NULL,
    id_banco INT NOT NULL,
    comprobante_path VARCHAR(255),
    sincronizado BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (id_banco) REFERENCES banco(id)
);

CREATE DATABASE IF NOT EXISTS sanjorge_transportistas_db;
USE sanjorge_transportistas_db;

-- Tabla banco
CREATE TABLE banco (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE
);

-- Tabla guía
CREATE TABLE guia (
    id INT AUTO_INCREMENT PRIMARY KEY,
    numero VARCHAR(50) NOT NULL,
    fecha DATE NOT NULL,
    codigo_cliente VARCHAR(50),
    nombre_cliente VARCHAR(100),
    nro_comprobante VARCHAR(50),
    importe_x_cobrar DECIMAL(10,2) NOT NULL,
    monto_cobrado DECIMAL(10,2) DEFAULT 0,
    entregada BOOLEAN DEFAULT FALSE
);

-- Tabla producto (relación con guía)
CREATE TABLE producto (
    id INT AUTO_INCREMENT PRIMARY KEY,
    id_guia INT NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    cantidad INT NOT NULL,
    FOREIGN KEY (id_guia) REFERENCES guia(id) ON DELETE CASCADE
);

-- Tabla depósito (relación con banco)
CREATE TABLE deposito (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nro_operacion VARCHAR(100) NOT NULL,
    fecha DATE NOT NULL,
    monto DECIMAL(10,2) NOT NULL,
    id_banco INT NOT NULL,
    comprobante_path VARCHAR(255),
    sincronizado BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (id_banco) REFERENCES banco(id)
);

CREATE TABLE prueba_entrega (
    id INT AUTO_INCREMENT PRIMARY KEY,
    guia_id INT NOT NULL,
    firma BLOB,
    imagen_path VARCHAR(255),
    fecha_registro DATE NOT NULL,
    sincronizado BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (guia_id) REFERENCES guia(id)
);