package wkb

import (
	"github.com/twpayne/gogeom/geom"
	"reflect"
	"testing"
)

func TestWKB(t *testing.T) {

	var testCases = []struct {
		g   geom.T
		xdr []byte
		ndr []byte
	}{
		{
			g: geom.Point{X: 1, Y: 2},
			xdr: []byte("\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@"),
		},
		{
			g: geom.PointZ{X: 1, Y: 2, Z: 3},
			xdr: []byte("\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g: geom.PointM{X: 1, Y: 2, M: 3},
			xdr: []byte("\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g: geom.PointZM{X: 1, Y: 2, Z: 3, M: 4},
			xdr: []byte("\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g: geom.LineString{Points: []geom.Point{{1, 2}, {3, 4}}},
			xdr: []byte("\x00\x00\x00\x00\x02\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x02\x00\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g: geom.LineStringZ{Points: []geom.PointZ{{1, 2, 3}, {4, 5, 6}}},
			xdr: []byte("\x00\x00\x00\x03\xea\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xea\x03\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g: geom.LineStringM{Points: []geom.PointM{{1, 2, 3}, {4, 5, 6}}},
			xdr: []byte("\x00\x00\x00\x07\xd2\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd2\x07\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g: geom.LineStringZM{Points: []geom.PointZM{{1, 2, 3, 4}, {5, 6, 7, 8}}},
			xdr: []byte("\x00\x00\x00\x0b\xba\x00\x00\x00\x02?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xba\x0b\x00\x00\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @"),
		},
		{
			g: geom.Polygon{Rings: [][]geom.Point{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}},
			xdr: []byte("\x00\x00\x00\x00\x03\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x03\x00\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@"),
		},
		{
			g: geom.PolygonZ{Rings: [][]geom.PointZ{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}},
			xdr: []byte("\x00\x00\x00\x03\xeb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xeb\x03\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g: geom.PolygonM{Rings: [][]geom.PointM{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}},
			xdr: []byte("\x00\x00\x00\x07\xd3\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd3\x07\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@"),
		},
		{
			g: geom.PolygonZM{Rings: [][]geom.PointZM{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}}},
			xdr: []byte("\x00\x00\x00\x0b\xbb\x00\x00\x00\x01\x00\x00\x00\x04?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00@\"\x00\x00\x00\x00\x00\x00@$\x00\x00\x00\x00\x00\x00@&\x00\x00\x00\x00\x00\x00@(\x00\x00\x00\x00\x00\x00?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xbb\x0b\x00\x00\x01\x00\x00\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @\x00\x00\x00\x00\x00\x00\"@\x00\x00\x00\x00\x00\x00$@\x00\x00\x00\x00\x00\x00&@\x00\x00\x00\x00\x00\x00(@\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g: geom.MultiPoint{Points: []geom.Point{{1, 2}, {3, 4}}},
			xdr: []byte("\x00\x00\x00\x00\x04\x00\x00\x00\x02\x00\x00\x00\x00\x01?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\x04\x00\x00\x00\x02\x00\x00\x00\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x01\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@"),
		},
		{
			g: geom.MultiPointZ{Points: []geom.PointZ{{1, 2, 3}, {4, 5, 6}}},
			xdr: []byte("\x00\x00\x00\x03\xec\x00\x00\x00\x02\x00\x00\x00\x03\xe9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x03\xe9@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xec\x03\x00\x00\x02\x00\x00\x00\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x01\xe9\x03\x00\x00\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g: geom.MultiPointM{Points: []geom.PointM{{1, 2, 3}, {4, 5, 6}}},
			xdr: []byte("\x00\x00\x00\x07\xd4\x00\x00\x00\x02\x00\x00\x00\x07\xd1?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00\x00\x00\x00\x07\xd1@\x10\x00\x00\x00\x00\x00\x00@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xd4\x07\x00\x00\x02\x00\x00\x00\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x01\xd1\x07\x00\x00\x00\x00\x00\x00\x00\x00\x10@\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@"),
		},
		{
			g: geom.MultiPointZM{Points: []geom.PointZM{{1, 2, 3, 4}, {5, 6, 7, 8}}},
			xdr: []byte("\x00\x00\x00\x0b\xbc\x00\x00\x00\x02\x00\x00\x00\x0b\xb9?\xf0\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x00@\x08\x00\x00\x00\x00\x00\x00@\x10\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0b\xb9@\x14\x00\x00\x00\x00\x00\x00@\x18\x00\x00\x00\x00\x00\x00@\x1c\x00\x00\x00\x00\x00\x00@ \x00\x00\x00\x00\x00\x00"),
			ndr: []byte("\x01\xbc\x0b\x00\x00\x02\x00\x00\x00\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\xf0?\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x00\x00\x00\x08@\x00\x00\x00\x00\x00\x00\x10@\x01\xb9\x0b\x00\x00\x00\x00\x00\x00\x00\x00\x14@\x00\x00\x00\x00\x00\x00\x18@\x00\x00\x00\x00\x00\x00\x1c@\x00\x00\x00\x00\x00\x00 @"),
		},
	}

	for _, tc := range testCases {

		// test XDR decoding
		if got, err := Decode(tc.xdr); err != nil || !reflect.DeepEqual(got, tc.g) {
			t.Errorf("Decode(%#v) == %#v, %s, want %#v, nil", tc.xdr, got, err, tc.g)
		}

		// test XDR encoding
		if got, err := Encode(tc.g, XDR); err != nil || !reflect.DeepEqual(got, tc.xdr) {
			t.Errorf("Encode(%#v, %#v) == %#v, %#v, want %#v, nil", tc.g, XDR, got, err, tc.xdr)
		}

		// test NDR decoding
		if got, err := Decode(tc.ndr); err != nil || !reflect.DeepEqual(got, tc.g) {
			t.Errorf("Decode(%#v) == %#v, %s, want %#v, nil", tc.ndr, got, err, tc.g)
		}

		// test NDR encoding
		if got, err := Encode(tc.g, NDR); err != nil || !reflect.DeepEqual(got, tc.ndr) {
			t.Errorf("Encode(%#v, %#v) == %#v, %#v, want %#v, nil", tc.g, NDR, got, err, tc.ndr)
		}

	}

}
