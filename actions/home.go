package actions

import (
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/BryanMoslo/QAReport/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
	polyline "github.com/twpayne/go-polyline"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// SendFile is a method that receive a logfile
func SendFile(c buffalo.Context) error {
	c.Request().ParseMultipartForm(0)
	Form := c.Request().MultipartForm
	filename := Form.File["file"][0].Filename

	file, _ := Form.File["file"][0].Open()
	contentFile, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(contentFile), "\n")
	coords := convertLocation(lines)
	urlImage := urlFromPath(coords)

	c.Set("points", coords)
	distance := getDistance(coords)

	report := &models.Report{
		Filename: filename,
		Distance: distance,
		ImageSrc: urlImage,
	}

	tx := c.Get("tx").(*pop.Connection)
	err := tx.Create(report)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(301, "/reports/%v", report.ID)
}

func convertLocation(data []string) [][]float64 {
	var coords = [][]float64{}
	r, _ := regexp.Compile(`<(\S*,\S*)>`)

	for _, l := range data {
		pointLocation := strings.NewReplacer("<", "", ">", "").Replace(r.FindString(l))

		if len(pointLocation) > 1 {
			lat, _ := strconv.ParseFloat(strings.Split(pointLocation, ",")[0], 64)
			lng, _ := strconv.ParseFloat(strings.Split(pointLocation, ",")[1], 64)
			coords = append(coords, []float64{lat, lng})
		}
	}

	return coords
}

func getDistance(coords [][]float64) float64 {
	var distance float64
	var R = 6371.0 // Radius of the earth in km

	for i := 0; i < len(coords)-1; i++ {
		lat1 := coords[i][0]
		lon1 := coords[i][1]
		lat2 := coords[i+1][0]
		lon2 := coords[i+1][1]

		var dLat = deg2Rad(lat2 - lat1)
		var dLon = deg2Rad(lon2 - lon1)
		var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
			math.Cos(deg2Rad(lat1))*math.Cos(deg2Rad(lat2))*
				math.Sin(dLon/2)*math.Sin(dLon/2)

		var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
		// Distance in km
		var d = R * c

		distance += d
	}

	return round((distance*100)/100, 2)
}

func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor((f * shift) + .5)
}

func deg2Rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func urlFromPath(points [][]float64) string {
	var startMarker = "markers=color:blue%7Clabel:I%7C" + strconv.FormatFloat(points[0][0], 'g', -1, 64) + "," + strconv.FormatFloat(points[0][1], 'g', -1, 64)
	var endMarker = "&markers=color:red%7Clabel:F%7C" + strconv.FormatFloat(points[len(points)-1][0], 'g', -1, 64) + "," + strconv.FormatFloat(points[len(points)-1][1], 'g', -1, 64)
	var route = "https://maps.googleapis.com/maps/api/staticmap?" + startMarker + endMarker + "&scale=2&size=800x800&path=weight:6%7Ccolor:red%7Cenc:"

	codec := polyline.Codec{Dim: 2, Scale: 1e5}
	encodecPath := codec.EncodeCoords([]byte{}, points)
	route = route + string(encodecPath)

	return route
}
