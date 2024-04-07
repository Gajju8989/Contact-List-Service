package mysql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task2/database/mysql/model/authtoken"
	"time"
)

func (m *impl) StoreAuthToken(ctx context.Context, auth *authtoken.AuthToken) error {
	return m.mysqlDb.
		WithContext(ctx).
		Create(auth).
		Error
}

func (m *impl) GetAuthToken(ctx context.Context, phoneNo string) (*authtoken.AuthToken, error) {

	if phoneNo == "" {
		return nil, fmt.Errorf("Unauthorized: Phone number not provided")
	}
	searchPredicate := &authtoken.AuthToken{
		PhoneNo: phoneNo,
	}

	tNow := time.Now().Unix()

	res := &authtoken.AuthToken{}
	err := m.mysqlDb.
		WithContext(ctx).
		Where(searchPredicate).
		Where("expiry > ?", tNow).
		First(&res).
		Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Unauthorized: User Not Found")
		}
		return nil, err
	}

	return res, nil
}
