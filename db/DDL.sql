CREATE TABLE tablet(
tablet_id serial primary key,
tablet_name varchar(60) NOT NULL UNIQUE
    );
CREATE TABLE telemetry(
    telemetry_id serial primary key,
    battery smallint NOT NULL,
    device_time TIMESTAMP NOT NULL,
     current_video varchar(255),
    tablet_id int NOT NULL ,
    CONSTRAINT percent CHECK ( battery BETWEEN 0 and 100)
);
CREATE INDEX telemetry$tablet_id_idx ON telemetry (tablet_id);
ALTER TABLE telemetry
ADD CONSTRAINT fk_telemetry$tablet_id FOREIGN KEY (tablet_id) REFERENCES tablet (tablet_id);


