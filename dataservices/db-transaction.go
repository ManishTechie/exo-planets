package dataservices

import "gas-giant/api/v1"

// BeginTransaction --
func (ms *DBClient) BeginTransaction() (txn BackendServiceDBInterface, err *api.APIError) {
	var result DBClient
	result.DB = ms.DB.Begin()
	if result.DB.Error != nil {
		return nil, api.NewAPIError(api.InternalServerError, "error occurred while begin DB Transaction")
	}
	return &result, nil
}

// CommitTransaction --
func (ms *DBClient) CommitTransaction() (err *api.APIError) {
	if err := ms.DB.Commit().Error; err != nil {
		if rollbackErr := ms.DB.Rollback().Error; rollbackErr != nil {
			return api.NewAPIError(api.InternalServerError, "rollback failed after unsuccessful commit")
		}
		return api.NewAPIError(api.InternalServerError, "commit failed for DB transaction")
	}
	return nil
}

// RollbackTransaction -
func (ms *DBClient) RollbackTransaction() (err *api.APIError) {
	if rollbackErr := ms.DB.Rollback().Error; rollbackErr != nil {
		return api.NewAPIError(api.InternalServerError, "rollback failed")
	}
	return nil
}
