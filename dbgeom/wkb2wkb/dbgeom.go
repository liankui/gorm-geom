package wkb2wkb

import (
	"database/sql/driver"
	"errors"
	"log"

	"github.com/spatial-go/geoos/encoding/wkb"
	"github.com/spatial-go/geoos/space"
)

// DbGeom ...
type DbGeom struct {
	Geometry space.Geometry
}

// GormDataType ...
func (g DbGeom) GormDataType() string {
	return "geometry"
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 DbGeom
func (g *DbGeom) Scan(value interface{}) error {
	cgeo, err := wkb.GeomFromWKBHexStr(value.(string))
	if err != nil {
		return errors.New("cannot convert database value to geometry")
	}
	g.Geometry = cgeo
	return nil
}

// Value 实现 driver.Valuer 接口，Value 返回 string value
func (g DbGeom) Value() (driver.Value, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic err=%v\n", err)
		}
	}()
	//str := wkt.MarshalString(g.Geometry.Geom())
	str, err := wkb.GeomToWKBHexStr(g.Geometry.Geom())
	if err != nil {
		log.Printf("wkb.GeomToWKBHexStr err=%v\n", err)
		return "", errors.New("cannot convert geometry to string")
	}
	return "SRID=4326;" + str, nil
}
