package models

import (
	"fmt"

	"github.com/words-reminder-api/db/queries"
)

func userQueryBuilder(filter UserFilter) (string, []interface{}) {
	q := ""
	params := make([]interface{}, 0)
	offset := filter.Offset
	limit := filter.Limit

	if filter.Count {
		q = queries.SQLUserCount
	} else {
		q = queries.SQLFetchUsers
	}

	if filter.Email != "" {
		q = q + " AND email LIKE ? "
		params = append(params, fmt.Sprintf("%v%v%v", "%", filter.Email, "%"))
	}

	if !filter.Count {
		q = q + queries.SQLLimitOffset
		params = append(params, limit)
		params = append(params, offset)
	}

	return q, params
}
