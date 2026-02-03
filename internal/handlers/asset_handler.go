package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/AbsoluteZero24/goaset/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (server *Server) ListAssetKSO(w http.ResponseWriter, r *http.Request) {
	var assets []models.AssetKSO
	server.DB.Preload("User").Find(&assets)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/asetkso", map[string]interface{}{
		"title":  "Daftar Aset KSO",
		"assets": assets,
	})
}

func (server *Server) CreateAssetKSOForm(w http.ResponseWriter, r *http.Request) {
	var categories []models.MasterAssetCategory
	server.DB.Find(&categories)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/asetkso_form", map[string]interface{}{
		"title":      "Tambah Aset KSO",
		"categories": categories,
	})
}

func (server *Server) StoreAssetKSO(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	purchaseDate, _ := time.Parse("2006-01-02", r.FormValue("purchase_date"))
	userID := r.FormValue("user_id")
	var userIDPtr *string
	if userID != "" {
		userIDPtr = &userID
	}

	asset := models.AssetKSO{
		ID:              uuid.New().String(),
		InventoryNumber: r.FormValue("inventory_number"),
		AssetName:       r.FormValue("asset_name"),
		DeviceName:      r.FormValue("device_name"),
		Category:        r.FormValue("category"),
		Brand:           r.FormValue("brand"),
		Specification:   r.FormValue("specification"),
		Color:           r.FormValue("color"),
		Location:        r.FormValue("location"),
		UserID:          userIDPtr,
		PurchaseDate:    purchaseDate,
		Status:          r.FormValue("status"),
	}

	server.DB.Create(&asset)

	redirectPath := r.FormValue("redirect_to")
	if redirectPath == "" {
		redirectPath = "/inventori/aset-laptop"
		if asset.Category == "Laptop" {
			redirectPath = "/asset-management/laptop"
		} else if asset.Category == "Komputer" {
			redirectPath = "/asset-management/komputer"
		}
	}

	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func (server *Server) CreateAssetKSOBulkForm(w http.ResponseWriter, r *http.Request) {
	var categories []models.MasterAssetCategory
	server.DB.Find(&categories)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/asetkso_bulk_form", map[string]interface{}{
		"title":      "Sisipan Masal Aset KSO",
		"categories": categories,
	})
}

func (server *Server) StoreAssetKSOBulk(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	qtyStr := r.FormValue("quantity")
	quantity, _ := strconv.Atoi(qtyStr)
	if quantity < 1 {
		quantity = 1
	}

	invStart := r.FormValue("inventory_number_start")
	purchaseDate, _ := time.Parse("2006-01-02", r.FormValue("purchase_date"))

	// Helper to increment inventory number
	// It looks for digits at the end of the string
	re := regexp.MustCompile(`(\d+)$`)
	matches := re.FindStringSubmatch(invStart)

	var prefix string
	var currentNum int
	var padding int

	if len(matches) > 0 {
		numStr := matches[1]
		padding = len(numStr)
		currentNum, _ = strconv.Atoi(numStr)
		prefix = invStart[:len(invStart)-padding]
	} else {
		// If no digits at end, we just append numbers
		prefix = invStart + "-"
		currentNum = 1
		padding = 1
	}

	for i := 0; i < quantity; i++ {
		newInvNum := ""
		if len(matches) > 0 {
			newInvNum = fmt.Sprintf("%s%0*d", prefix, padding, currentNum+i)
		} else {
			if i == 0 {
				newInvNum = invStart
			} else {
				newInvNum = fmt.Sprintf("%s%d", prefix, currentNum+i)
			}
		}

		asset := models.AssetKSO{
			ID:              uuid.New().String(),
			InventoryNumber: newInvNum,
			AssetName:       r.FormValue("asset_name"),
			Category:        r.FormValue("category"),
			Brand:           r.FormValue("brand"),
			Specification:   r.FormValue("specification"),
			Color:           r.FormValue("color"),
			Location:        r.FormValue("location"),
			PurchaseDate:    purchaseDate,
			Status:          r.FormValue("status"),
		}
		server.DB.Create(&asset)
	}

	redirectPath := r.FormValue("redirect_to")
	if redirectPath == "" {
		redirectPath = "/inventori/aset-laptop"
	}

	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func (server *Server) EditAssetKSOForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var asset models.AssetKSO
	if err := server.DB.Preload("User").Where("id = ?", id).First(&asset).Error; err != nil {
		http.Redirect(w, r, "/inventori/aset-laptop", http.StatusSeeOther)
		return
	}

	var categories []models.MasterAssetCategory
	server.DB.Find(&categories)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/asetkso_form", map[string]interface{}{
		"title":      "Edit Aset",
		"asset":      asset,
		"categories": categories,
	})
}

func (server *Server) UpdateAssetKSO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var asset models.AssetKSO
	if err := server.DB.Where("id = ?", id).First(&asset).Error; err != nil {
		http.Redirect(w, r, "/inventori/aset-laptop", http.StatusSeeOther)
		return
	}

	purchaseDate, _ := time.Parse("2006-01-02", r.FormValue("purchase_date"))
	userID := r.FormValue("user_id")
	var userIDPtr *string
	if userID != "" {
		userIDPtr = &userID
	}

	asset.InventoryNumber = r.FormValue("inventory_number")
	asset.AssetName = r.FormValue("asset_name")
	asset.DeviceName = r.FormValue("device_name")
	asset.Category = r.FormValue("category")
	asset.Brand = r.FormValue("brand")
	asset.Specification = r.FormValue("specification")
	asset.Color = r.FormValue("color")
	asset.Location = r.FormValue("location")
	asset.UserID = userIDPtr
	asset.PurchaseDate = purchaseDate
	asset.Status = r.FormValue("status")

	server.DB.Save(&asset)

	redirectPath := r.FormValue("redirect_to")
	if redirectPath == "" {
		redirectPath = "/inventori/aset-laptop"
		if asset.Category == "Laptop" {
			redirectPath = "/asset-management/laptop"
		} else if asset.Category == "Komputer" {
			redirectPath = "/asset-management/komputer"
		}
	}

	http.Redirect(w, r, redirectPath, http.StatusSeeOther)
}

func (server *Server) DeleteAssetKSO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	server.DB.Unscoped().Where("id = ?", id).Delete(&models.AssetKSO{})
	http.Redirect(w, r, "/inventori/aset-laptop", http.StatusSeeOther)
}

func (server *Server) ListAssetLaptop(w http.ResponseWriter, r *http.Request) {
	var assets []models.AssetKSO
	server.DB.Preload("User").Where("category = ?", "Laptop").Order("inventory_number asc").Find(&assets)

	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/laptop_management", map[string]interface{}{
		"title":  "Asset Management - Laptop",
		"assets": assets,
		"users":  users,
	})
}

func (server *Server) CreateAssetLaptopForm(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/laptop_mgmt_form", map[string]interface{}{
		"title":    "Tambah Laptop",
		"category": "Laptop",
		"users":    users,
	})
}

func (server *Server) EditAssetLaptopForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var asset models.AssetKSO
	if err := server.DB.Preload("User").Where("id = ?", id).First(&asset).Error; err != nil {
		http.Redirect(w, r, "/asset-management/laptop", http.StatusSeeOther)
		return
	}

	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/laptop_mgmt_form", map[string]interface{}{
		"title": "Edit Laptop",
		"asset": asset,
		"users": users,
	})
}

func (server *Server) DeleteAssetLaptop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	server.DB.Unscoped().Where("id = ?", id).Delete(&models.AssetKSO{})
	http.Redirect(w, r, "/asset-management/laptop", http.StatusSeeOther)
}

func (server *Server) AssignAssetLaptop(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	assetID := r.FormValue("asset_id")
	userID := r.FormValue("user_id")

	if assetID == "" {
		http.Redirect(w, r, "/asset-management/laptop", http.StatusSeeOther)
		return
	}

	var asset models.AssetKSO
	if err := server.DB.Where("id = ?", assetID).First(&asset).Error; err == nil {
		if userID == "" {
			asset.UserID = nil
		} else {
			asset.UserID = &userID
		}
		server.DB.Save(&asset)
	}

	http.Redirect(w, r, "/asset-management/laptop", http.StatusSeeOther)
}

func (server *Server) ListAssetKomputer(w http.ResponseWriter, r *http.Request) {
	var assets []models.AssetKSO
	server.DB.Preload("User").Where("category = ?", "Komputer").Order("inventory_number asc").Find(&assets)

	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/komputer_management", map[string]interface{}{
		"title":  "Asset Management - Komputer",
		"assets": assets,
		"users":  users,
	})
}

func (server *Server) CreateAssetKomputerForm(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/komputer_mgmt_form", map[string]interface{}{
		"title":    "Tambah Komputer",
		"category": "Komputer",
		"users":    users,
	})
}

func (server *Server) EditAssetKomputerForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var asset models.AssetKSO
	if err := server.DB.Preload("User").Where("id = ?", id).First(&asset).Error; err != nil {
		http.Redirect(w, r, "/asset-management/komputer", http.StatusSeeOther)
		return
	}

	var users []models.User
	server.DB.Find(&users)

	_ = server.Renderer.HTML(w, http.StatusOK, "assets/komputer_mgmt_form", map[string]interface{}{
		"title": "Edit Komputer",
		"asset": asset,
		"users": users,
	})
}

func (server *Server) DeleteAssetKomputer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	server.DB.Unscoped().Where("id = ?", id).Delete(&models.AssetKSO{})
	http.Redirect(w, r, "/asset-management/komputer", http.StatusSeeOther)
}

func (server *Server) AssignAssetKomputer(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	assetID := r.FormValue("asset_id")
	userID := r.FormValue("user_id")

	if assetID == "" {
		http.Redirect(w, r, "/asset-management/komputer", http.StatusSeeOther)
		return
	}

	var asset models.AssetKSO
	if err := server.DB.Where("id = ?", assetID).First(&asset).Error; err == nil {
		if userID == "" {
			asset.UserID = nil
		} else {
			asset.UserID = &userID
		}
		server.DB.Save(&asset)
	}

	http.Redirect(w, r, "/asset-management/komputer", http.StatusSeeOther)
}
