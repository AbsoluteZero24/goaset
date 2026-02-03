package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: AssetKSO{}},
		{Model: MasterBranch{}},
		{Model: MasterDepartment{}},
		{Model: MasterSubDepartment{}},
		{Model: MasterPosition{}},
		{Model: MasterAssetCategory{}},
	}
}
