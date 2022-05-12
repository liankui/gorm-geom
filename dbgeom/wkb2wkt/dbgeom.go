package wkb2wkt

import (
	"database/sql/driver"
	"errors"
	"github.com/spatial-go/geoos/encoding/wkb"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
	"log"
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
			log.Println(err)
		}
	}()
	str := wkt.MarshalString(g.Geometry.Geom())
	// return "SRID=4326;" + str, nil
	return str, nil
}
