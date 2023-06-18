package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentsHandler struct{}

func (sh studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize data: %q", err)
	}
	return buf.Bytes(), nil
}

// /students - entire class
// /students/{id} - single student record
// /students/{id}/grades - single student's grades
func (sh studentsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO(moosch): Needs validation and verification. Look at using 3rd party server like fasthttp/gin.
	pathSegments := strings.Split(req.URL.Path, "/")
	switch len(pathSegments) {
	case 2: // /students
		sh.getAll(w, req)
	case 3: // /students/{:id}
		id, err := strconv.Atoi(pathSegments[2]) // ["", "student", "{id}"]
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.getOne(w, req, id)
	case 4: // /students/{:id}/grades
		id, err := strconv.Atoi(pathSegments[2]) // ["", "student", "{id}"]
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sh.addGrade(w, req, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (sh studentsHandler) getAll(w http.ResponseWriter, req *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data, err := sh.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("failed to serialize students: %q", err))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
func (sh studentsHandler) getOne(w http.ResponseWriter, req *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}
	data, err := sh.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(fmt.Errorf("failed to serialize student: %q", err))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
func (sh studentsHandler) addGrade(w http.ResponseWriter, req *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	var grade Grade
	dec := json.NewDecoder(req.Body)
	err = dec.Decode(&grade)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	student.Grades = append(student.Grades, grade)

	w.WriteHeader(http.StatusCreated)
	data, err := sh.toJSON(grade)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
