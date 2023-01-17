package whendy

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	ses "sevian.com/whendy/session"
	"sevian.com/whendy/tool"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
// var session map[string]interface{}

type Store struct {
	Session ses.Session
	vform   map[string]interface{}
	vexp    map[string]interface{}
	User    User
	DB      *DB
}

func (store *Store) Start(w http.ResponseWriter, req *http.Request) {
	//store.User = User{}
	store.User.Init(w, req)

	store.vform = make(map[string]interface{})
	store.vexp = make(map[string]interface{})
	//UploadHandler(w, req)

	fmt.Println("el usuario es ", store.User.Get())

	mode := req.Header.Get("Content-Type")

	if mode == "application/json" {

		defer req.Body.Close()
		b, _ := io.ReadAll(req.Body)

		json.Unmarshal(b, &store.vform)

	} else if req.Method == "POST" {

		req.ParseForm()
		for key := range req.Form {
			store.vform[key] = req.FormValue(key)
		}

	} else if req.Method == "GET" {

		values := req.URL.Query()
		for key := range values {
			store.vform[key] = req.URL.Query().Get(key)
		}
	}

}

func (s *Store) SetForm(key string, value interface{}) {

	s.vform[key] = value
}

func (s *Store) GetForm(key string) interface{} {
	return s.vform[key]

}

func (s *Store) LoadExp(data map[string]interface{}) {
	for key, value := range data {
		s.vexp[key] = value
	}

}

func (s *Store) SetExp(key string, value interface{}) {

	s.vexp[key] = value
}

func (s *Store) GetExp(key string) interface{} {
	return s.vexp[key]

}

func (s *Store) SetSes(key string, value interface{}) {
	s.Session.Set(key, value)

}

func (s *Store) GetSes(key string) interface{} {
	return s.Session.Get(key)

}

func (s *Store) GetDB(name string) *sql.DB {

	return s.DB.dbs[name]
}

func (s *Store) SetSessionData(data map[string]interface{}) {

	for k, v := range data {
		s.SetSes(k, v)
	}

}

func (s *Store) LoadJson(path string) {
	data := tool.LoadJsonFile(path)

	for k, v := range data {
		//println(k, v)
		/*if _, ok := v.(string); ok {
			ss[k] = v
		}*/

		s.SetSes(k, v)
	}

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// truncated for brevity

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")
}

func UploadHandlerMultiple(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 32 MB is the default used by FormFile()
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := r.MultipartForm.File["file"]

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Fprintf(w, "Upload successful")
}

type Progress struct {
	TotalSize int64
	BytesRead int64
}

// Write is used to satisfy the io.Writer interface.
// Instead of writing somewhere, it simply aggregates
// the total bytes on each read
func (pr *Progress) Write(p []byte) (n int, err error) {
	n, err = len(p), nil
	pr.BytesRead += int64(n)
	pr.Print()
	return
}

// Print displays the current progress of the file upload
// each time Write is called
func (pr *Progress) Print() {
	if pr.BytesRead == pr.TotalSize {
		fmt.Println("DONE!")
		return
	}

	fmt.Printf("File upload in progress: %d\n", pr.BytesRead)
}

func KuploadHandler1(w http.ResponseWriter, req *http.Request) {
	// truncated for brevity
	/*
		files := req.MultipartForm.File["file"]
		for _, fileHeader := range files {
			// [..]

			pr := &Progress{
				TotalSize: fileHeader.Size,
			}

			_, err := io.Copy(f, io.TeeReader(file, pr))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	*/

	fmt.Fprintf(w, "Upload successful")
}
