package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongsvg/go/orm"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongsvg/go")
	{ // insertion point for registrations
		v1.GET("/v1/animates", GetAnimates)
		v1.GET("/v1/animates/:id", GetAnimate)
		v1.POST("/v1/animates", PostAnimate)
		v1.PATCH("/v1/animates/:id", UpdateAnimate)
		v1.PUT("/v1/animates/:id", UpdateAnimate)
		v1.DELETE("/v1/animates/:id", DeleteAnimate)

		v1.GET("/v1/circles", GetCircles)
		v1.GET("/v1/circles/:id", GetCircle)
		v1.POST("/v1/circles", PostCircle)
		v1.PATCH("/v1/circles/:id", UpdateCircle)
		v1.PUT("/v1/circles/:id", UpdateCircle)
		v1.DELETE("/v1/circles/:id", DeleteCircle)

		v1.GET("/v1/ellipses", GetEllipses)
		v1.GET("/v1/ellipses/:id", GetEllipse)
		v1.POST("/v1/ellipses", PostEllipse)
		v1.PATCH("/v1/ellipses/:id", UpdateEllipse)
		v1.PUT("/v1/ellipses/:id", UpdateEllipse)
		v1.DELETE("/v1/ellipses/:id", DeleteEllipse)

		v1.GET("/v1/lines", GetLines)
		v1.GET("/v1/lines/:id", GetLine)
		v1.POST("/v1/lines", PostLine)
		v1.PATCH("/v1/lines/:id", UpdateLine)
		v1.PUT("/v1/lines/:id", UpdateLine)
		v1.DELETE("/v1/lines/:id", DeleteLine)

		v1.GET("/v1/paths", GetPaths)
		v1.GET("/v1/paths/:id", GetPath)
		v1.POST("/v1/paths", PostPath)
		v1.PATCH("/v1/paths/:id", UpdatePath)
		v1.PUT("/v1/paths/:id", UpdatePath)
		v1.DELETE("/v1/paths/:id", DeletePath)

		v1.GET("/v1/polygones", GetPolygones)
		v1.GET("/v1/polygones/:id", GetPolygone)
		v1.POST("/v1/polygones", PostPolygone)
		v1.PATCH("/v1/polygones/:id", UpdatePolygone)
		v1.PUT("/v1/polygones/:id", UpdatePolygone)
		v1.DELETE("/v1/polygones/:id", DeletePolygone)

		v1.GET("/v1/polylines", GetPolylines)
		v1.GET("/v1/polylines/:id", GetPolyline)
		v1.POST("/v1/polylines", PostPolyline)
		v1.PATCH("/v1/polylines/:id", UpdatePolyline)
		v1.PUT("/v1/polylines/:id", UpdatePolyline)
		v1.DELETE("/v1/polylines/:id", DeletePolyline)

		v1.GET("/v1/rects", GetRects)
		v1.GET("/v1/rects/:id", GetRect)
		v1.POST("/v1/rects", PostRect)
		v1.PATCH("/v1/rects/:id", UpdateRect)
		v1.PUT("/v1/rects/:id", UpdateRect)
		v1.DELETE("/v1/rects/:id", DeleteRect)

		v1.GET("/v1/svgs", GetSVGs)
		v1.GET("/v1/svgs/:id", GetSVG)
		v1.POST("/v1/svgs", PostSVG)
		v1.PATCH("/v1/svgs/:id", UpdateSVG)
		v1.PUT("/v1/svgs/:id", UpdateSVG)
		v1.DELETE("/v1/svgs/:id", DeleteSVG)

		v1.GET("/v1/texts", GetTexts)
		v1.GET("/v1/texts/:id", GetText)
		v1.POST("/v1/texts", PostText)
		v1.PATCH("/v1/texts/:id", UpdateText)
		v1.PUT("/v1/texts/:id", UpdateText)
		v1.DELETE("/v1/texts/:id", DeleteText)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
