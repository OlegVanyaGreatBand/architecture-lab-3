CREATE FUNCTION get_telemetry (in in_id int) RETURNS TABLE (
    tablet_id int,
    tablet_name varchar(60),
    battery smallint,
    device_time TIMESTAMP,
    current_video varchar(255)
) AS $BODY$

    SELECT (tablet.tablet_id, tablet.tablet_name, telemetry.battery, telemetry.device_time, telemetry.current_video) FROM tablet
    INNER JOIN telemetry on tablet.tablet_id = telemetry.tablet_id
    WHERE tablet.tablet_id = in_id
    ORDER BY telemetry.device_time DESC
    LIMIT 50;

    $BODY$ LANGUAGE sql;
