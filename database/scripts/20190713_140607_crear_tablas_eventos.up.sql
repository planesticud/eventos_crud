-- object: evento | type: SCHEMA --
-- DROP SCHEMA IF EXISTS evento CASCADE;
CREATE SCHEMA evento;
-- ddl-end --
COMMENT ON SCHEMA evento IS 'Esquema del core';
-- ddl-end --

-- SET search_path TO pg_catalog,public,evento;
-- ddl-end --

-- object: evento.tipo_evento | type: TABLE --
-- DROP TABLE IF EXISTS evento.tipo_evento CASCADE;
CREATE TABLE evento.tipo_evento (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	dependencia_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_recurrencia_id integer,
	CONSTRAINT pk_tipo_evento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evento.tipo_evento IS 'Tabla que permite registrar los diferentes tipos de sesion que tienen las diferentes organizaciones';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.id IS 'Identificador de la tabla tipo sesion';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.nombre IS 'Nombre del tipo de sesion';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.descripcion IS 'Descripcion del tipo de sesion que se referencia';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.codigo_abreviacion IS 'Codigo de abreviación del tipo de sesión que se maneja';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.activo IS 'Booleano que define si el tipo de sesión se encuentra actualmente activo o no';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.dependencia_id IS 'se referencia al esquema de dependencia del sistema OKIOS';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.fecha_creacion IS 'Fecha de creacion de un tipo de evento';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_evento.fecha_modificacion IS 'Fecha de modifiación de un tipo de evento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_evento ON evento.tipo_evento  IS 'Contrait para pk de la tabla';
-- ddl-end --

-- object: evento.calendario_evento | type: TABLE --
-- DROP TABLE IF EXISTS evento.calendario_evento CASCADE;
CREATE TABLE evento.calendario_evento (
	id serial NOT NULL,
	descripcion varchar(250),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	periodo_id integer NOT NULL,
	activo boolean NOT NULL,
	evento_padre_id integer,
	tipo_evento_id integer,
	CONSTRAINT pk_calendario_evento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evento.calendario_evento IS 'Tabla que almacena cada uno de las sesiones y sus datos';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.id IS 'Identificador de la sesión';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.descripcion IS 'Descripción de la sesión qeu se registra';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.fecha_creacion IS 'Fecha de creacion de una sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.fecha_modificacion IS 'Fecha de modifiación del evento para la organización';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.fecha_inicio IS 'Fecha de inicio de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.fecha_fin IS 'Fecha de finalización de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.periodo_id IS 'Se almacena el periodo al cual pertenece el evento';
-- ddl-end --
COMMENT ON COLUMN evento.calendario_evento.activo IS 'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_calendario_evento ON evento.calendario_evento  IS 'Restriccion pk de la tabla calendario evento';
-- ddl-end --

-- object: evento.encargado_evento | type: TABLE --
-- DROP TABLE IF EXISTS evento.encargado_evento CASCADE;
CREATE TABLE evento.encargado_evento (
	id serial NOT NULL,
	encargado_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	rol_encargado_evento_id integer,
	calendario_evento_id integer,
	CONSTRAINT pk_encargado_evento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evento.encargado_evento IS 'Tabla de rompimiento entre los diferentes encargados y el evento';
-- ddl-end --
COMMENT ON COLUMN evento.encargado_evento.id IS 'Identificador de la tabla responsable_sesion';
-- ddl-end --
COMMENT ON COLUMN evento.encargado_evento.encargado_id IS 'Se hace referencia al esquema de persona';
-- ddl-end --
COMMENT ON COLUMN evento.encargado_evento.fecha_creacion IS 'Fecha de creacion de un participante del evento';
-- ddl-end --
COMMENT ON COLUMN evento.encargado_evento.fecha_modificacion IS 'Fecha de modifiación de un participante del evento';
-- ddl-end --
COMMENT ON COLUMN evento.encargado_evento.activo IS 'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_encargado_evento ON evento.encargado_evento  IS 'Identificador de la tabla responsable evento';
-- ddl-end --

-- object: evento.rol_encargado_evento | type: TABLE --
-- DROP TABLE IF EXISTS evento.rol_encargado_evento CASCADE;
CREATE TABLE evento.rol_encargado_evento (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_de_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_rol_encargado_evento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evento.rol_encargado_evento IS 'Permite almacenar los diferentes roles que puede tener un encargado en un evento';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.id IS 'Identificador de la tabla rol_participante_sesion';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.nombre IS 'Nombre del rol que puede tener un participante de una sesion';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.descripcion IS 'Descripcion que permite especificar el rol y las funciones en la sesion';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.codigo_abreviacion IS 'Codigo de abreviacion del nombre del rol que cumple un participante en la sesion';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.activo IS 'Booleano que permite identificar si el rol se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.fecha_creacion IS 'Fecha de creacion de un rol participante del evento';
-- ddl-end --
COMMENT ON COLUMN evento.rol_encargado_evento.fecha_modificacion IS 'Fecha de modifiación de un rol participante del evento';
-- ddl-end --

-- object: evento.tipo_recurrencia | type: TABLE --
-- DROP TABLE IF EXISTS evento.tipo_recurrencia CASCADE;
CREATE TABLE evento.tipo_recurrencia (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_recurrencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evento.tipo_recurrencia IS 'Tabla que almacena los tipos de recurrencia en tiempo que se tienen (por minuto, hora, diario...)';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.id IS 'Identificador de la tabla tipo_recurrencia';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.nombre IS 'Nombre del tipo de recurrencia que se registra en la tabla';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.descripcion IS 'Breve descripcion del tipo de recurrencia que se registra';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.codigo_abreviacion IS 'Codigo de abreviacion del tipo de recurrencia';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.activo IS 'Booleano que permite identificar si el tipo de recurrencia se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.fecha_creacion IS 'Fecha de creacion de un tipo de recurencia';
-- ddl-end --
COMMENT ON COLUMN evento.tipo_recurrencia.fecha_modificacion IS 'Fecha de modifiación de un tipo de recurrencia';
-- ddl-end --

-- object: fk_encargado_evento_rol_encargado_evento | type: CONSTRAINT --
-- ALTER TABLE evento.encargado_evento DROP CONSTRAINT IF EXISTS fk_encargado_evento_rol_encargado_evento CASCADE;
ALTER TABLE evento.encargado_evento ADD CONSTRAINT fk_encargado_evento_rol_encargado_evento FOREIGN KEY (rol_encargado_evento_id)
REFERENCES evento.rol_encargado_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_tipo_evento_tipo_recurrencia | type: CONSTRAINT --
-- ALTER TABLE evento.tipo_evento DROP CONSTRAINT IF EXISTS fk_tipo_evento_tipo_recurrencia CASCADE;
ALTER TABLE evento.tipo_evento ADD CONSTRAINT fk_tipo_evento_tipo_recurrencia FOREIGN KEY (tipo_recurrencia_id)
REFERENCES evento.tipo_recurrencia (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: tipo_evento_uq | type: CONSTRAINT --
-- ALTER TABLE evento.tipo_evento DROP CONSTRAINT IF EXISTS tipo_evento_uq CASCADE;
ALTER TABLE evento.tipo_evento ADD CONSTRAINT tipo_evento_uq UNIQUE (tipo_recurrencia_id);
-- ddl-end --

-- object: fk_calendario_evento_tipo_evento | type: CONSTRAINT --
-- ALTER TABLE evento.calendario_evento DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_evento CASCADE;
ALTER TABLE evento.calendario_evento ADD CONSTRAINT fk_calendario_evento_tipo_evento FOREIGN KEY (tipo_evento_id)
REFERENCES evento.tipo_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_encargado_evento_calendario_evento | type: CONSTRAINT --
-- ALTER TABLE evento.encargado_evento DROP CONSTRAINT IF EXISTS fk_encargado_evento_calendario_evento CASCADE;
ALTER TABLE evento.encargado_evento ADD CONSTRAINT fk_encargado_evento_calendario_evento FOREIGN KEY (calendario_evento_id)
REFERENCES evento.calendario_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: evento.tipo_publico | type: TABLE --
-- DROP TABLE IF EXISTS evento.tipo_publico CASCADE;
CREATE TABLE evento.tipo_publico (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	codigo_abreviacion varchar(250),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	calendario_evento_id integer,
	CONSTRAINT pk_tipo_publico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE evento.tipo_publico IS 'Tabla para almacenar el tipo de publico que puede participar en un evento ';
-- ddl-end --

-- object: fk_tipo_publico_calendario_evento | type: CONSTRAINT --
-- ALTER TABLE evento.tipo_publico DROP CONSTRAINT IF EXISTS fk_tipo_publico_calendario_evento CASCADE;
ALTER TABLE evento.tipo_publico ADD CONSTRAINT fk_tipo_publico_calendario_evento FOREIGN KEY (calendario_evento_id)
REFERENCES evento.calendario_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_evento_padre | type: CONSTRAINT --
-- ALTER TABLE evento.calendario_evento DROP CONSTRAINT IF EXISTS fk_evento_padre CASCADE;
ALTER TABLE evento.calendario_evento ADD CONSTRAINT fk_evento_padre FOREIGN KEY (evento_padre_id)
REFERENCES evento.calendario_evento (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
COMMENT ON CONSTRAINT fk_evento_padre ON evento.calendario_evento  IS 'FK de la evento padre';
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA evento TO test;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA evento TO test;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA evento TO test;
