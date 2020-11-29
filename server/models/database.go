package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetTelemetry(tabletId int) (*TelemetryData, error) {
	rows, err := s.db.Query("SELECT * FROM get_telemetry(?);", tabletId)
	if err != nil {
		return nil, err
	}

	res := &TelemetryData{}
	for rows.Next() {
		if err := scanTelemetry(res, rows); err != nil {
			return nil, err
		}
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Store) AddTelemetry(data *TelemetryData) error {
	rows, err := s.db.Query("SELECT id FROM tablet WHERE id = ?;", data.TabletId)
	if err != nil {
		return errors.New("tablet not found")
	}
	_ = rows.Close()

	query := "INSERT INTO telemetry (battery, device_time, timestamp, current_video) VALUES "
	for _, t := range data.Telemetry {
		var video string
		if t.CurrentVideo != nil {
			video = *t.CurrentVideo
		} else {
			video = "NULL"
		}
		query += fmt.Sprintf("(%d, %s, %s, %s)", t.Battery, t.DeviceTime, t.Timestamp, video)
	}
	query += ";"

	if _, err = s.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func scanTelemetry(tablet *TelemetryData, rows *sql.Rows) error {
	t := &struct {
		tabletId     int
		tabletName   *string
		battery      int
		deviceTime   string
		currentVideo *string
	}{}
	if err := rows.Scan(
		t.tabletId,
		t.tabletName,
		t.battery,
		t.deviceTime,
		t.currentVideo,
	); err != nil {
		return err
	}

	tablet.TabletId = t.tabletId
	tablet.TabletName = t.tabletName
	tablet.Telemetry = append(tablet.Telemetry, &Telemetry{
		Battery:      t.battery,
		DeviceTime:   t.deviceTime,
		Timestamp:    t.deviceTime,
		CurrentVideo: t.currentVideo,
	})

	return nil
}
