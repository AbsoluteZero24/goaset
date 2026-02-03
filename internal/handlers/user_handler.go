package handlers

import (
	"net/http"

	"github.com/AbsoluteZero24/goaset/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (server *Server) ListEmployees(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "administration/employee", map[string]interface{}{
		"title": "Daftar Karyawan",
		"users": users,
	})
}

func (server *Server) CreateEmployeeForm(w http.ResponseWriter, r *http.Request) {
	var branches []models.MasterBranch
	var positions []models.MasterPosition

	server.DB.Preload("Departments.SubDepartments").Find(&branches)
	server.DB.Find(&positions)

	_ = server.Renderer.HTML(w, http.StatusOK, "administration/employee_form", map[string]interface{}{
		"title":     "Tambah Karyawan",
		"branches":  branches,
		"positions": positions,
	})
}

func (server *Server) StoreEmployee(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:             uuid.New().String(),
		NIK:            r.FormValue("nik"),
		Name:           r.FormValue("name"),
		Email:          r.FormValue("email"),
		Branch:         r.FormValue("branch"),
		Department:     r.FormValue("department"),
		SubDepartment:  r.FormValue("sub_department"),
		Position:       r.FormValue("position"),
		StatusKaryawan: r.FormValue("status_karyawan"),
		Password:       "password123", // Default password
	}

	server.DB.Create(&user)
	http.Redirect(w, r, "/administration/employee", http.StatusSeeOther)
}

func (server *Server) EditEmployeeForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	if err := server.DB.Where("id = ?", id).First(&user).Error; err != nil {
		http.Redirect(w, r, "/administration/employee", http.StatusSeeOther)
		return
	}

	var branches []models.MasterBranch
	var positions []models.MasterPosition

	server.DB.Preload("Departments.SubDepartments").Find(&branches)
	server.DB.Find(&positions)

	_ = server.Renderer.HTML(w, http.StatusOK, "administration/employee_form", map[string]interface{}{
		"title":     "Edit Karyawan",
		"user":      user,
		"branches":  branches,
		"positions": positions,
	})
}

func (server *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := server.DB.Where("id = ?", id).First(&user).Error; err != nil {
		http.Redirect(w, r, "/administration/employee", http.StatusSeeOther)
		return
	}

	user.NIK = r.FormValue("nik")
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Branch = r.FormValue("branch")
	user.Department = r.FormValue("department")
	user.SubDepartment = r.FormValue("sub_department")
	user.Position = r.FormValue("position")
	user.StatusKaryawan = r.FormValue("status_karyawan")

	server.DB.Save(&user)
	http.Redirect(w, r, "/administration/employee", http.StatusSeeOther)
}

func (server *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	server.DB.Where("id = ?", id).Delete(&models.User{})
	http.Redirect(w, r, "/administration/employee", http.StatusSeeOther)
}
