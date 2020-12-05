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
	query, err := s.db.Prepare("select * from get_telemetry($1)")
	if err != nil {
		return nil, err
	}

	rows, err := query.Query(tabletId)
	if err != nil {
		return nil, err
	}

	res := &TelemetryData{
		Telemetry: []*Telemetry{},
	}
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
	q, err := s.db.Prepare("SELECT tablet_id FROM tablet WHERE tablet_id=$1")
	if err != nil {
		return errors.New(fmt.Sprintf("tablet %d not found", data.TabletId))
	}
	rows, err := q.Query(data.TabletId)
	if err != nil {
		return errors.New(fmt.Sprintf("tablet %d not found", data.TabletId))
	}

	_ = rows.Close()

	query := "INSERT INTO telemetry (tablet_id, battery, device_time, server_time, current_video) VALUES "
	for _, t := range data.Telemetry {
		var video string
		if t.CurrentVideo != nil {
			video = *t.CurrentVideo
		} else {
			video = "NULL"
		}
		query += fmt.Sprintf("(%d, %d, '%s', '%s', '%s'),", data.TabletId, t.Battery, t.DeviceTime, t.ServerTime, video)
	}
	query = query[:len(query) - 1]
	query += ";"

	if _, err = s.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func scanTelemetry(tablet *TelemetryData, rows *sql.Rows) error {
	t := &struct {
		tabletId     int
		tabletName   string
		battery      *int
		deviceTime   *string
		serverTime   *string
		currentVideo *string
	}{}
	if err := rows.Scan(
		&t.tabletId,
		&t.tabletName,
		&t.battery,
		&t.deviceTime,
		&t.serverTime,
		&t.currentVideo,
	); err != nil {
		return err
	}

	tablet.TabletId = t.tabletId
	tablet.TabletName = &t.tabletName
	if t.battery != nil && t.deviceTime != nil && t.serverTime != nil {
		tablet.Telemetry = append(tablet.Telemetry, &Telemetry{
			Battery:      *t.battery,
			DeviceTime:   *t.deviceTime,
			ServerTime:   *t.serverTime,
			CurrentVideo: t.currentVideo,
		})
	}

	return nil
}
