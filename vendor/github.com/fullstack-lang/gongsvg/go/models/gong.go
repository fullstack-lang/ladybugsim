// generated by ModelGongFileTemplate
package models

import "sort"

// swagger:ignore
type __void struct{}

// needed for creating set of instances in the stage
var __member __void

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct { // insertion point for definition of arrays registering instances
	Animates           map[*Animate]struct{}
	Animates_mapString map[string]*Animate

	Circles           map[*Circle]struct{}
	Circles_mapString map[string]*Circle

	Ellipses           map[*Ellipse]struct{}
	Ellipses_mapString map[string]*Ellipse

	Lines           map[*Line]struct{}
	Lines_mapString map[string]*Line

	Paths           map[*Path]struct{}
	Paths_mapString map[string]*Path

	Polygones           map[*Polygone]struct{}
	Polygones_mapString map[string]*Polygone

	Polylines           map[*Polyline]struct{}
	Polylines_mapString map[string]*Polyline

	Rects           map[*Rect]struct{}
	Rects_mapString map[string]*Rect

	SVGs           map[*SVG]struct{}
	SVGs_mapString map[string]*SVG

	Texts           map[*Text]struct{}
	Texts_mapString map[string]*Text

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback OnInitCommitInterface
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAnimate(animate *Animate)
	CheckoutAnimate(animate *Animate)
	CommitCircle(circle *Circle)
	CheckoutCircle(circle *Circle)
	CommitEllipse(ellipse *Ellipse)
	CheckoutEllipse(ellipse *Ellipse)
	CommitLine(line *Line)
	CheckoutLine(line *Line)
	CommitPath(path *Path)
	CheckoutPath(path *Path)
	CommitPolygone(polygone *Polygone)
	CheckoutPolygone(polygone *Polygone)
	CommitPolyline(polyline *Polyline)
	CheckoutPolyline(polyline *Polyline)
	CommitRect(rect *Rect)
	CheckoutRect(rect *Rect)
	CommitSVG(svg *SVG)
	CheckoutSVG(svg *SVG)
	CommitText(text *Text)
	CheckoutText(text *Text)
	GetLastCommitNb() uint
	GetLastPushFromFrontNb() uint
}

// swagger:ignore instructs the gong compiler (gongc) to avoid this particular struct
var Stage StageStruct = StageStruct{ // insertion point for array initiatialisation
	Animates:           make(map[*Animate]struct{}),
	Animates_mapString: make(map[string]*Animate),

	Circles:           make(map[*Circle]struct{}),
	Circles_mapString: make(map[string]*Circle),

	Ellipses:           make(map[*Ellipse]struct{}),
	Ellipses_mapString: make(map[string]*Ellipse),

	Lines:           make(map[*Line]struct{}),
	Lines_mapString: make(map[string]*Line),

	Paths:           make(map[*Path]struct{}),
	Paths_mapString: make(map[string]*Path),

	Polygones:           make(map[*Polygone]struct{}),
	Polygones_mapString: make(map[string]*Polygone),

	Polylines:           make(map[*Polyline]struct{}),
	Polylines_mapString: make(map[string]*Polyline),

	Rects:           make(map[*Rect]struct{}),
	Rects_mapString: make(map[string]*Rect),

	SVGs:           make(map[*SVG]struct{}),
	SVGs_mapString: make(map[string]*SVG),

	Texts:           make(map[*Text]struct{}),
	Texts_mapString: make(map[string]*Text),

	// end of insertion point
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
func (stage *StageStruct) getAnimateOrderedStructWithNameField() []*Animate {
	// have alphabetical order generation
	animateOrdered := []*Animate{}
	for animate := range stage.Animates {
		animateOrdered = append(animateOrdered, animate)
	}
	sort.Slice(animateOrdered[:], func(i, j int) bool {
		return animateOrdered[i].Name < animateOrdered[j].Name
	})
	return animateOrdered
}

// Stage puts animate to the model stage
func (animate *Animate) Stage() *Animate {
	Stage.Animates[animate] = __member
	Stage.Animates_mapString[animate.Name] = animate

	return animate
}

// Unstage removes animate off the model stage
func (animate *Animate) Unstage() *Animate {
	delete(Stage.Animates, animate)
	delete(Stage.Animates_mapString, animate.Name)
	return animate
}

// commit animate to the back repo (if it is already staged)
func (animate *Animate) Commit() *Animate {
	if _, ok := Stage.Animates[animate]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitAnimate(animate)
		}
	}
	return animate
}

// Checkout animate to the back repo (if it is already staged)
func (animate *Animate) Checkout() *Animate {
	if _, ok := Stage.Animates[animate]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutAnimate(animate)
		}
	}
	return animate
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of animate to the model stage
func (animate *Animate) StageCopy() *Animate {
	_animate := new(Animate)
	*_animate = *animate
	_animate.Stage()
	return _animate
}

// StageAndCommit appends animate to the model stage and commit to the orm repo
func (animate *Animate) StageAndCommit() *Animate {
	animate.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMAnimate(animate)
	}
	return animate
}

// DeleteStageAndCommit appends animate to the model stage and commit to the orm repo
func (animate *Animate) DeleteStageAndCommit() *Animate {
	animate.Unstage()
	DeleteORMAnimate(animate)
	return animate
}

// StageCopyAndCommit appends a copy of animate to the model stage and commit to the orm repo
func (animate *Animate) StageCopyAndCommit() *Animate {
	_animate := new(Animate)
	*_animate = *animate
	_animate.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMAnimate(animate)
	}
	return _animate
}

// CreateORMAnimate enables dynamic staging of a Animate instance
func CreateORMAnimate(animate *Animate) {
	animate.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMAnimate(animate)
	}
}

// DeleteORMAnimate enables dynamic staging of a Animate instance
func DeleteORMAnimate(animate *Animate) {
	animate.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMAnimate(animate)
	}
}

func (stage *StageStruct) getCircleOrderedStructWithNameField() []*Circle {
	// have alphabetical order generation
	circleOrdered := []*Circle{}
	for circle := range stage.Circles {
		circleOrdered = append(circleOrdered, circle)
	}
	sort.Slice(circleOrdered[:], func(i, j int) bool {
		return circleOrdered[i].Name < circleOrdered[j].Name
	})
	return circleOrdered
}

// Stage puts circle to the model stage
func (circle *Circle) Stage() *Circle {
	Stage.Circles[circle] = __member
	Stage.Circles_mapString[circle.Name] = circle

	return circle
}

// Unstage removes circle off the model stage
func (circle *Circle) Unstage() *Circle {
	delete(Stage.Circles, circle)
	delete(Stage.Circles_mapString, circle.Name)
	return circle
}

// commit circle to the back repo (if it is already staged)
func (circle *Circle) Commit() *Circle {
	if _, ok := Stage.Circles[circle]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitCircle(circle)
		}
	}
	return circle
}

// Checkout circle to the back repo (if it is already staged)
func (circle *Circle) Checkout() *Circle {
	if _, ok := Stage.Circles[circle]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutCircle(circle)
		}
	}
	return circle
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of circle to the model stage
func (circle *Circle) StageCopy() *Circle {
	_circle := new(Circle)
	*_circle = *circle
	_circle.Stage()
	return _circle
}

// StageAndCommit appends circle to the model stage and commit to the orm repo
func (circle *Circle) StageAndCommit() *Circle {
	circle.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMCircle(circle)
	}
	return circle
}

// DeleteStageAndCommit appends circle to the model stage and commit to the orm repo
func (circle *Circle) DeleteStageAndCommit() *Circle {
	circle.Unstage()
	DeleteORMCircle(circle)
	return circle
}

// StageCopyAndCommit appends a copy of circle to the model stage and commit to the orm repo
func (circle *Circle) StageCopyAndCommit() *Circle {
	_circle := new(Circle)
	*_circle = *circle
	_circle.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMCircle(circle)
	}
	return _circle
}

// CreateORMCircle enables dynamic staging of a Circle instance
func CreateORMCircle(circle *Circle) {
	circle.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMCircle(circle)
	}
}

// DeleteORMCircle enables dynamic staging of a Circle instance
func DeleteORMCircle(circle *Circle) {
	circle.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMCircle(circle)
	}
}

func (stage *StageStruct) getEllipseOrderedStructWithNameField() []*Ellipse {
	// have alphabetical order generation
	ellipseOrdered := []*Ellipse{}
	for ellipse := range stage.Ellipses {
		ellipseOrdered = append(ellipseOrdered, ellipse)
	}
	sort.Slice(ellipseOrdered[:], func(i, j int) bool {
		return ellipseOrdered[i].Name < ellipseOrdered[j].Name
	})
	return ellipseOrdered
}

// Stage puts ellipse to the model stage
func (ellipse *Ellipse) Stage() *Ellipse {
	Stage.Ellipses[ellipse] = __member
	Stage.Ellipses_mapString[ellipse.Name] = ellipse

	return ellipse
}

// Unstage removes ellipse off the model stage
func (ellipse *Ellipse) Unstage() *Ellipse {
	delete(Stage.Ellipses, ellipse)
	delete(Stage.Ellipses_mapString, ellipse.Name)
	return ellipse
}

// commit ellipse to the back repo (if it is already staged)
func (ellipse *Ellipse) Commit() *Ellipse {
	if _, ok := Stage.Ellipses[ellipse]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitEllipse(ellipse)
		}
	}
	return ellipse
}

// Checkout ellipse to the back repo (if it is already staged)
func (ellipse *Ellipse) Checkout() *Ellipse {
	if _, ok := Stage.Ellipses[ellipse]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutEllipse(ellipse)
		}
	}
	return ellipse
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of ellipse to the model stage
func (ellipse *Ellipse) StageCopy() *Ellipse {
	_ellipse := new(Ellipse)
	*_ellipse = *ellipse
	_ellipse.Stage()
	return _ellipse
}

// StageAndCommit appends ellipse to the model stage and commit to the orm repo
func (ellipse *Ellipse) StageAndCommit() *Ellipse {
	ellipse.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMEllipse(ellipse)
	}
	return ellipse
}

// DeleteStageAndCommit appends ellipse to the model stage and commit to the orm repo
func (ellipse *Ellipse) DeleteStageAndCommit() *Ellipse {
	ellipse.Unstage()
	DeleteORMEllipse(ellipse)
	return ellipse
}

// StageCopyAndCommit appends a copy of ellipse to the model stage and commit to the orm repo
func (ellipse *Ellipse) StageCopyAndCommit() *Ellipse {
	_ellipse := new(Ellipse)
	*_ellipse = *ellipse
	_ellipse.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMEllipse(ellipse)
	}
	return _ellipse
}

// CreateORMEllipse enables dynamic staging of a Ellipse instance
func CreateORMEllipse(ellipse *Ellipse) {
	ellipse.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMEllipse(ellipse)
	}
}

// DeleteORMEllipse enables dynamic staging of a Ellipse instance
func DeleteORMEllipse(ellipse *Ellipse) {
	ellipse.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMEllipse(ellipse)
	}
}

func (stage *StageStruct) getLineOrderedStructWithNameField() []*Line {
	// have alphabetical order generation
	lineOrdered := []*Line{}
	for line := range stage.Lines {
		lineOrdered = append(lineOrdered, line)
	}
	sort.Slice(lineOrdered[:], func(i, j int) bool {
		return lineOrdered[i].Name < lineOrdered[j].Name
	})
	return lineOrdered
}

// Stage puts line to the model stage
func (line *Line) Stage() *Line {
	Stage.Lines[line] = __member
	Stage.Lines_mapString[line.Name] = line

	return line
}

// Unstage removes line off the model stage
func (line *Line) Unstage() *Line {
	delete(Stage.Lines, line)
	delete(Stage.Lines_mapString, line.Name)
	return line
}

// commit line to the back repo (if it is already staged)
func (line *Line) Commit() *Line {
	if _, ok := Stage.Lines[line]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitLine(line)
		}
	}
	return line
}

// Checkout line to the back repo (if it is already staged)
func (line *Line) Checkout() *Line {
	if _, ok := Stage.Lines[line]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutLine(line)
		}
	}
	return line
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of line to the model stage
func (line *Line) StageCopy() *Line {
	_line := new(Line)
	*_line = *line
	_line.Stage()
	return _line
}

// StageAndCommit appends line to the model stage and commit to the orm repo
func (line *Line) StageAndCommit() *Line {
	line.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMLine(line)
	}
	return line
}

// DeleteStageAndCommit appends line to the model stage and commit to the orm repo
func (line *Line) DeleteStageAndCommit() *Line {
	line.Unstage()
	DeleteORMLine(line)
	return line
}

// StageCopyAndCommit appends a copy of line to the model stage and commit to the orm repo
func (line *Line) StageCopyAndCommit() *Line {
	_line := new(Line)
	*_line = *line
	_line.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMLine(line)
	}
	return _line
}

// CreateORMLine enables dynamic staging of a Line instance
func CreateORMLine(line *Line) {
	line.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMLine(line)
	}
}

// DeleteORMLine enables dynamic staging of a Line instance
func DeleteORMLine(line *Line) {
	line.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMLine(line)
	}
}

func (stage *StageStruct) getPathOrderedStructWithNameField() []*Path {
	// have alphabetical order generation
	pathOrdered := []*Path{}
	for path := range stage.Paths {
		pathOrdered = append(pathOrdered, path)
	}
	sort.Slice(pathOrdered[:], func(i, j int) bool {
		return pathOrdered[i].Name < pathOrdered[j].Name
	})
	return pathOrdered
}

// Stage puts path to the model stage
func (path *Path) Stage() *Path {
	Stage.Paths[path] = __member
	Stage.Paths_mapString[path.Name] = path

	return path
}

// Unstage removes path off the model stage
func (path *Path) Unstage() *Path {
	delete(Stage.Paths, path)
	delete(Stage.Paths_mapString, path.Name)
	return path
}

// commit path to the back repo (if it is already staged)
func (path *Path) Commit() *Path {
	if _, ok := Stage.Paths[path]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitPath(path)
		}
	}
	return path
}

// Checkout path to the back repo (if it is already staged)
func (path *Path) Checkout() *Path {
	if _, ok := Stage.Paths[path]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutPath(path)
		}
	}
	return path
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of path to the model stage
func (path *Path) StageCopy() *Path {
	_path := new(Path)
	*_path = *path
	_path.Stage()
	return _path
}

// StageAndCommit appends path to the model stage and commit to the orm repo
func (path *Path) StageAndCommit() *Path {
	path.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPath(path)
	}
	return path
}

// DeleteStageAndCommit appends path to the model stage and commit to the orm repo
func (path *Path) DeleteStageAndCommit() *Path {
	path.Unstage()
	DeleteORMPath(path)
	return path
}

// StageCopyAndCommit appends a copy of path to the model stage and commit to the orm repo
func (path *Path) StageCopyAndCommit() *Path {
	_path := new(Path)
	*_path = *path
	_path.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPath(path)
	}
	return _path
}

// CreateORMPath enables dynamic staging of a Path instance
func CreateORMPath(path *Path) {
	path.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPath(path)
	}
}

// DeleteORMPath enables dynamic staging of a Path instance
func DeleteORMPath(path *Path) {
	path.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMPath(path)
	}
}

func (stage *StageStruct) getPolygoneOrderedStructWithNameField() []*Polygone {
	// have alphabetical order generation
	polygoneOrdered := []*Polygone{}
	for polygone := range stage.Polygones {
		polygoneOrdered = append(polygoneOrdered, polygone)
	}
	sort.Slice(polygoneOrdered[:], func(i, j int) bool {
		return polygoneOrdered[i].Name < polygoneOrdered[j].Name
	})
	return polygoneOrdered
}

// Stage puts polygone to the model stage
func (polygone *Polygone) Stage() *Polygone {
	Stage.Polygones[polygone] = __member
	Stage.Polygones_mapString[polygone.Name] = polygone

	return polygone
}

// Unstage removes polygone off the model stage
func (polygone *Polygone) Unstage() *Polygone {
	delete(Stage.Polygones, polygone)
	delete(Stage.Polygones_mapString, polygone.Name)
	return polygone
}

// commit polygone to the back repo (if it is already staged)
func (polygone *Polygone) Commit() *Polygone {
	if _, ok := Stage.Polygones[polygone]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitPolygone(polygone)
		}
	}
	return polygone
}

// Checkout polygone to the back repo (if it is already staged)
func (polygone *Polygone) Checkout() *Polygone {
	if _, ok := Stage.Polygones[polygone]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutPolygone(polygone)
		}
	}
	return polygone
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of polygone to the model stage
func (polygone *Polygone) StageCopy() *Polygone {
	_polygone := new(Polygone)
	*_polygone = *polygone
	_polygone.Stage()
	return _polygone
}

// StageAndCommit appends polygone to the model stage and commit to the orm repo
func (polygone *Polygone) StageAndCommit() *Polygone {
	polygone.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPolygone(polygone)
	}
	return polygone
}

// DeleteStageAndCommit appends polygone to the model stage and commit to the orm repo
func (polygone *Polygone) DeleteStageAndCommit() *Polygone {
	polygone.Unstage()
	DeleteORMPolygone(polygone)
	return polygone
}

// StageCopyAndCommit appends a copy of polygone to the model stage and commit to the orm repo
func (polygone *Polygone) StageCopyAndCommit() *Polygone {
	_polygone := new(Polygone)
	*_polygone = *polygone
	_polygone.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPolygone(polygone)
	}
	return _polygone
}

// CreateORMPolygone enables dynamic staging of a Polygone instance
func CreateORMPolygone(polygone *Polygone) {
	polygone.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPolygone(polygone)
	}
}

// DeleteORMPolygone enables dynamic staging of a Polygone instance
func DeleteORMPolygone(polygone *Polygone) {
	polygone.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMPolygone(polygone)
	}
}

func (stage *StageStruct) getPolylineOrderedStructWithNameField() []*Polyline {
	// have alphabetical order generation
	polylineOrdered := []*Polyline{}
	for polyline := range stage.Polylines {
		polylineOrdered = append(polylineOrdered, polyline)
	}
	sort.Slice(polylineOrdered[:], func(i, j int) bool {
		return polylineOrdered[i].Name < polylineOrdered[j].Name
	})
	return polylineOrdered
}

// Stage puts polyline to the model stage
func (polyline *Polyline) Stage() *Polyline {
	Stage.Polylines[polyline] = __member
	Stage.Polylines_mapString[polyline.Name] = polyline

	return polyline
}

// Unstage removes polyline off the model stage
func (polyline *Polyline) Unstage() *Polyline {
	delete(Stage.Polylines, polyline)
	delete(Stage.Polylines_mapString, polyline.Name)
	return polyline
}

// commit polyline to the back repo (if it is already staged)
func (polyline *Polyline) Commit() *Polyline {
	if _, ok := Stage.Polylines[polyline]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitPolyline(polyline)
		}
	}
	return polyline
}

// Checkout polyline to the back repo (if it is already staged)
func (polyline *Polyline) Checkout() *Polyline {
	if _, ok := Stage.Polylines[polyline]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutPolyline(polyline)
		}
	}
	return polyline
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of polyline to the model stage
func (polyline *Polyline) StageCopy() *Polyline {
	_polyline := new(Polyline)
	*_polyline = *polyline
	_polyline.Stage()
	return _polyline
}

// StageAndCommit appends polyline to the model stage and commit to the orm repo
func (polyline *Polyline) StageAndCommit() *Polyline {
	polyline.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPolyline(polyline)
	}
	return polyline
}

// DeleteStageAndCommit appends polyline to the model stage and commit to the orm repo
func (polyline *Polyline) DeleteStageAndCommit() *Polyline {
	polyline.Unstage()
	DeleteORMPolyline(polyline)
	return polyline
}

// StageCopyAndCommit appends a copy of polyline to the model stage and commit to the orm repo
func (polyline *Polyline) StageCopyAndCommit() *Polyline {
	_polyline := new(Polyline)
	*_polyline = *polyline
	_polyline.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPolyline(polyline)
	}
	return _polyline
}

// CreateORMPolyline enables dynamic staging of a Polyline instance
func CreateORMPolyline(polyline *Polyline) {
	polyline.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMPolyline(polyline)
	}
}

// DeleteORMPolyline enables dynamic staging of a Polyline instance
func DeleteORMPolyline(polyline *Polyline) {
	polyline.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMPolyline(polyline)
	}
}

func (stage *StageStruct) getRectOrderedStructWithNameField() []*Rect {
	// have alphabetical order generation
	rectOrdered := []*Rect{}
	for rect := range stage.Rects {
		rectOrdered = append(rectOrdered, rect)
	}
	sort.Slice(rectOrdered[:], func(i, j int) bool {
		return rectOrdered[i].Name < rectOrdered[j].Name
	})
	return rectOrdered
}

// Stage puts rect to the model stage
func (rect *Rect) Stage() *Rect {
	Stage.Rects[rect] = __member
	Stage.Rects_mapString[rect.Name] = rect

	return rect
}

// Unstage removes rect off the model stage
func (rect *Rect) Unstage() *Rect {
	delete(Stage.Rects, rect)
	delete(Stage.Rects_mapString, rect.Name)
	return rect
}

// commit rect to the back repo (if it is already staged)
func (rect *Rect) Commit() *Rect {
	if _, ok := Stage.Rects[rect]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitRect(rect)
		}
	}
	return rect
}

// Checkout rect to the back repo (if it is already staged)
func (rect *Rect) Checkout() *Rect {
	if _, ok := Stage.Rects[rect]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutRect(rect)
		}
	}
	return rect
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of rect to the model stage
func (rect *Rect) StageCopy() *Rect {
	_rect := new(Rect)
	*_rect = *rect
	_rect.Stage()
	return _rect
}

// StageAndCommit appends rect to the model stage and commit to the orm repo
func (rect *Rect) StageAndCommit() *Rect {
	rect.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMRect(rect)
	}
	return rect
}

// DeleteStageAndCommit appends rect to the model stage and commit to the orm repo
func (rect *Rect) DeleteStageAndCommit() *Rect {
	rect.Unstage()
	DeleteORMRect(rect)
	return rect
}

// StageCopyAndCommit appends a copy of rect to the model stage and commit to the orm repo
func (rect *Rect) StageCopyAndCommit() *Rect {
	_rect := new(Rect)
	*_rect = *rect
	_rect.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMRect(rect)
	}
	return _rect
}

// CreateORMRect enables dynamic staging of a Rect instance
func CreateORMRect(rect *Rect) {
	rect.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMRect(rect)
	}
}

// DeleteORMRect enables dynamic staging of a Rect instance
func DeleteORMRect(rect *Rect) {
	rect.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMRect(rect)
	}
}

func (stage *StageStruct) getSVGOrderedStructWithNameField() []*SVG {
	// have alphabetical order generation
	svgOrdered := []*SVG{}
	for svg := range stage.SVGs {
		svgOrdered = append(svgOrdered, svg)
	}
	sort.Slice(svgOrdered[:], func(i, j int) bool {
		return svgOrdered[i].Name < svgOrdered[j].Name
	})
	return svgOrdered
}

// Stage puts svg to the model stage
func (svg *SVG) Stage() *SVG {
	Stage.SVGs[svg] = __member
	Stage.SVGs_mapString[svg.Name] = svg

	return svg
}

// Unstage removes svg off the model stage
func (svg *SVG) Unstage() *SVG {
	delete(Stage.SVGs, svg)
	delete(Stage.SVGs_mapString, svg.Name)
	return svg
}

// commit svg to the back repo (if it is already staged)
func (svg *SVG) Commit() *SVG {
	if _, ok := Stage.SVGs[svg]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitSVG(svg)
		}
	}
	return svg
}

// Checkout svg to the back repo (if it is already staged)
func (svg *SVG) Checkout() *SVG {
	if _, ok := Stage.SVGs[svg]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutSVG(svg)
		}
	}
	return svg
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of svg to the model stage
func (svg *SVG) StageCopy() *SVG {
	_svg := new(SVG)
	*_svg = *svg
	_svg.Stage()
	return _svg
}

// StageAndCommit appends svg to the model stage and commit to the orm repo
func (svg *SVG) StageAndCommit() *SVG {
	svg.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMSVG(svg)
	}
	return svg
}

// DeleteStageAndCommit appends svg to the model stage and commit to the orm repo
func (svg *SVG) DeleteStageAndCommit() *SVG {
	svg.Unstage()
	DeleteORMSVG(svg)
	return svg
}

// StageCopyAndCommit appends a copy of svg to the model stage and commit to the orm repo
func (svg *SVG) StageCopyAndCommit() *SVG {
	_svg := new(SVG)
	*_svg = *svg
	_svg.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMSVG(svg)
	}
	return _svg
}

// CreateORMSVG enables dynamic staging of a SVG instance
func CreateORMSVG(svg *SVG) {
	svg.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMSVG(svg)
	}
}

// DeleteORMSVG enables dynamic staging of a SVG instance
func DeleteORMSVG(svg *SVG) {
	svg.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMSVG(svg)
	}
}

func (stage *StageStruct) getTextOrderedStructWithNameField() []*Text {
	// have alphabetical order generation
	textOrdered := []*Text{}
	for text := range stage.Texts {
		textOrdered = append(textOrdered, text)
	}
	sort.Slice(textOrdered[:], func(i, j int) bool {
		return textOrdered[i].Name < textOrdered[j].Name
	})
	return textOrdered
}

// Stage puts text to the model stage
func (text *Text) Stage() *Text {
	Stage.Texts[text] = __member
	Stage.Texts_mapString[text.Name] = text

	return text
}

// Unstage removes text off the model stage
func (text *Text) Unstage() *Text {
	delete(Stage.Texts, text)
	delete(Stage.Texts_mapString, text.Name)
	return text
}

// commit text to the back repo (if it is already staged)
func (text *Text) Commit() *Text {
	if _, ok := Stage.Texts[text]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CommitText(text)
		}
	}
	return text
}

// Checkout text to the back repo (if it is already staged)
func (text *Text) Checkout() *Text {
	if _, ok := Stage.Texts[text]; ok {
		if Stage.BackRepo != nil {
			Stage.BackRepo.CheckoutText(text)
		}
	}
	return text
}

//
// Legacy, to be deleted
//

// StageCopy appends a copy of text to the model stage
func (text *Text) StageCopy() *Text {
	_text := new(Text)
	*_text = *text
	_text.Stage()
	return _text
}

// StageAndCommit appends text to the model stage and commit to the orm repo
func (text *Text) StageAndCommit() *Text {
	text.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMText(text)
	}
	return text
}

// DeleteStageAndCommit appends text to the model stage and commit to the orm repo
func (text *Text) DeleteStageAndCommit() *Text {
	text.Unstage()
	DeleteORMText(text)
	return text
}

// StageCopyAndCommit appends a copy of text to the model stage and commit to the orm repo
func (text *Text) StageCopyAndCommit() *Text {
	_text := new(Text)
	*_text = *text
	_text.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMText(text)
	}
	return _text
}

// CreateORMText enables dynamic staging of a Text instance
func CreateORMText(text *Text) {
	text.Stage()
	if Stage.AllModelsStructCreateCallback != nil {
		Stage.AllModelsStructCreateCallback.CreateORMText(text)
	}
}

// DeleteORMText enables dynamic staging of a Text instance
func DeleteORMText(text *Text) {
	text.Unstage()
	if Stage.AllModelsStructDeleteCallback != nil {
		Stage.AllModelsStructDeleteCallback.DeleteORMText(text)
	}
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAnimate(Animate *Animate)
	CreateORMCircle(Circle *Circle)
	CreateORMEllipse(Ellipse *Ellipse)
	CreateORMLine(Line *Line)
	CreateORMPath(Path *Path)
	CreateORMPolygone(Polygone *Polygone)
	CreateORMPolyline(Polyline *Polyline)
	CreateORMRect(Rect *Rect)
	CreateORMSVG(SVG *SVG)
	CreateORMText(Text *Text)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnimate(Animate *Animate)
	DeleteORMCircle(Circle *Circle)
	DeleteORMEllipse(Ellipse *Ellipse)
	DeleteORMLine(Line *Line)
	DeleteORMPath(Path *Path)
	DeleteORMPolygone(Polygone *Polygone)
	DeleteORMPolyline(Polyline *Polyline)
	DeleteORMRect(Rect *Rect)
	DeleteORMSVG(SVG *SVG)
	DeleteORMText(Text *Text)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Animates = make(map[*Animate]struct{})
	stage.Animates_mapString = make(map[string]*Animate)

	stage.Circles = make(map[*Circle]struct{})
	stage.Circles_mapString = make(map[string]*Circle)

	stage.Ellipses = make(map[*Ellipse]struct{})
	stage.Ellipses_mapString = make(map[string]*Ellipse)

	stage.Lines = make(map[*Line]struct{})
	stage.Lines_mapString = make(map[string]*Line)

	stage.Paths = make(map[*Path]struct{})
	stage.Paths_mapString = make(map[string]*Path)

	stage.Polygones = make(map[*Polygone]struct{})
	stage.Polygones_mapString = make(map[string]*Polygone)

	stage.Polylines = make(map[*Polyline]struct{})
	stage.Polylines_mapString = make(map[string]*Polyline)

	stage.Rects = make(map[*Rect]struct{})
	stage.Rects_mapString = make(map[string]*Rect)

	stage.SVGs = make(map[*SVG]struct{})
	stage.SVGs_mapString = make(map[string]*SVG)

	stage.Texts = make(map[*Text]struct{})
	stage.Texts_mapString = make(map[string]*Text)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Animates = nil
	stage.Animates_mapString = nil

	stage.Circles = nil
	stage.Circles_mapString = nil

	stage.Ellipses = nil
	stage.Ellipses_mapString = nil

	stage.Lines = nil
	stage.Lines_mapString = nil

	stage.Paths = nil
	stage.Paths_mapString = nil

	stage.Polygones = nil
	stage.Polygones_mapString = nil

	stage.Polylines = nil
	stage.Polylines_mapString = nil

	stage.Rects = nil
	stage.Rects_mapString = nil

	stage.SVGs = nil
	stage.SVGs_mapString = nil

	stage.Texts = nil
	stage.Texts_mapString = nil

}
