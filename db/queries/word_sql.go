package queries

const (
	SQLCreateNewWord = `
	INSERT INTO word (content, status, youglish, created, translation)
	VALUES (?, ?, ?, ?, ?)`

	SQLFetchWordsColumns = `
	w.word_id,
	w.content,
	w.status,
	w.user_id,
	w.youglish,
	w.created,
	w.translation	`

	SQLFetchWords = `
	SELECT
		w.word_id,
		w.content,
		w.status,
		w.youglish,
		w.created,
		w.translation`

	SQLFromWord          = " FROM word w"
	SQLFromWordDomain    = " FROM word_domain wd"
	SQLWhere             = " WHERE 1 = 1 "
	SQLInnerJoinWord     = " INNER JOIN word w ON wd.word_id = w.word_id "
	SQLOrderByContentASC = " ORDER BY w.content ASC"
	SQLSearchWord        = "SELECT word_id, content FROM word WHERE content = ?"

	SQLSearchWordByUser = `
	SELECT
    w.word_id,
    w.content
	FROM
			word w
			INNER JOIN word_domain wd ON w.word_id = wd.word_id
	WHERE
			wd.user_id = ?
			AND w.word_id = ?`

	SQLLinkWordToUser = `
	INSERT INTO word_domain (user_id, status, word_id)
	VALUES (?, 1, ?)`
)
