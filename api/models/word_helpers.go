package models

import "github.com/words-reminder-api/db/queries"

func fetchQueryBuilder(f WordFilter) (string, []interface{}) {
	q := queries.SQLFetchWords
	params := []interface{}{}

	if f.UserID > 0 {
		q = q + queries.SQLFromWordDomain
		q = q + queries.SQLInnerJoinWord + queries.SQLWhere
		q = q + " AND wd.user_id = ? "
		params = append(params, f.UserID)
	} else {
		q = q + queries.SQLFromWord + queries.SQLWhere
	}

	q = q + queries.SQLOrderByContentASC

	return q, params
}
