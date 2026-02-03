package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/administration/employee", server.ListEmployees).Methods("GET")
	server.Router.HandleFunc("/administration/employee/create", server.CreateEmployeeForm).Methods("GET")
	server.Router.HandleFunc("/administration/employee", server.StoreEmployee).Methods("POST")
	server.Router.HandleFunc("/administration/employee/edit/{id}", server.EditEmployeeForm).Methods("GET")
	server.Router.HandleFunc("/administration/employee/update/{id}", server.UpdateEmployee).Methods("POST")
	server.Router.HandleFunc("/administration/employee/delete/{id}", server.DeleteEmployee).Methods("GET")

	// Master Data Employee
	server.Router.HandleFunc("/administration/master-data/branch", server.ListMasterBranch).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/branch/store", server.StoreMasterBranch).Methods("POST")
	server.Router.HandleFunc("/administration/master-data/branch/delete/{id}", server.DeleteMasterBranch).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/branch/edit/{id}", server.EditMasterBranch).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/branch/update/{id}", server.UpdateMasterBranch).Methods("POST")

	server.Router.HandleFunc("/administration/master-data/department", server.ListMasterDepartment).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/department/store", server.StoreMasterDepartment).Methods("POST")
	server.Router.HandleFunc("/administration/master-data/department/delete/{id}", server.DeleteMasterDepartment).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/department/edit/{id}", server.EditMasterDepartment).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/department/update/{id}", server.UpdateMasterDepartment).Methods("POST")

	server.Router.HandleFunc("/administration/master-data/sub-department", server.ListMasterSubDepartment).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/sub-department/store", server.StoreMasterSubDepartment).Methods("POST")
	server.Router.HandleFunc("/administration/master-data/sub-department/delete/{id}", server.DeleteMasterSubDepartment).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/sub-department/edit/{id}", server.EditMasterSubDepartment).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/sub-department/update/{id}", server.UpdateMasterSubDepartment).Methods("POST")

	server.Router.HandleFunc("/administration/master-data/position", server.ListMasterPosition).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/position/store", server.StoreMasterPosition).Methods("POST")
	server.Router.HandleFunc("/administration/master-data/position/delete/{id}", server.DeleteMasterPosition).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/position/edit/{id}", server.EditMasterPosition).Methods("GET")
	server.Router.HandleFunc("/administration/master-data/position/update/{id}", server.UpdateMasterPosition).Methods("POST")

	// Master Data Asset Category
	server.Router.HandleFunc("/inventori/master-data/asset-category", server.ListMasterAssetCategory).Methods("GET")
	server.Router.HandleFunc("/inventori/master-data/asset-category/store", server.StoreMasterAssetCategory).Methods("POST")
	server.Router.HandleFunc("/inventori/master-data/asset-category/delete/{id}", server.DeleteMasterAssetCategory).Methods("GET")
	server.Router.HandleFunc("/inventori/master-data/asset-category/edit/{id}", server.EditMasterAssetCategory).Methods("GET")
	server.Router.HandleFunc("/inventori/master-data/asset-category/update/{id}", server.UpdateMasterAssetCategory).Methods("POST")

	server.Router.HandleFunc("/inventori/aset-laptop", server.ListAssetKSO).Methods("GET")
	server.Router.HandleFunc("/inventori/aset-laptop/create", server.CreateAssetKSOForm).Methods("GET")
	server.Router.HandleFunc("/inventori/aset-laptop/bulk-create", server.CreateAssetKSOBulkForm).Methods("GET")
	server.Router.HandleFunc("/inventori/aset-laptop", server.StoreAssetKSO).Methods("POST")
	server.Router.HandleFunc("/inventori/aset-laptop/bulk-store", server.StoreAssetKSOBulk).Methods("POST")
	server.Router.HandleFunc("/inventori/aset-laptop/edit/{id}", server.EditAssetKSOForm).Methods("GET")
	server.Router.HandleFunc("/inventori/aset-laptop/update/{id}", server.UpdateAssetKSO).Methods("POST")
	server.Router.HandleFunc("/inventori/aset-laptop/delete/{id}", server.DeleteAssetKSO).Methods("GET")

	server.Router.HandleFunc("/asset-management/laptop", server.ListAssetLaptop).Methods("GET")
	server.Router.HandleFunc("/asset-management/laptop/create", server.CreateAssetLaptopForm).Methods("GET")
	server.Router.HandleFunc("/asset-management/laptop/edit/{id}", server.EditAssetLaptopForm).Methods("GET")
	server.Router.HandleFunc("/asset-management/laptop/delete/{id}", server.DeleteAssetLaptop).Methods("GET")
	server.Router.HandleFunc("/asset-management/laptop/assign", server.AssignAssetLaptop).Methods("POST")

	server.Router.HandleFunc("/asset-management/komputer", server.ListAssetKomputer).Methods("GET")
	server.Router.HandleFunc("/asset-management/komputer/create", server.CreateAssetKomputerForm).Methods("GET")
	server.Router.HandleFunc("/asset-management/komputer/edit/{id}", server.EditAssetKomputerForm).Methods("GET")
	server.Router.HandleFunc("/asset-management/komputer/delete/{id}", server.DeleteAssetKomputer).Methods("GET")
	server.Router.HandleFunc("/asset-management/komputer/assign", server.AssignAssetKomputer).Methods("POST")

	// Static files from dist (renamed to public in suggestion or keep dist?)
	// I'll keep it as dist for now but use the standard name "public" if preferred.
	// Actually, I'll use "public" and rename the folder "dist" to "public".
	staticFileDirectory := http.Dir("./public")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler)
}
