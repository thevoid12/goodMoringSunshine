package dbpkg

const (
	SCHEMA = "CREATE TABLE IF NOT EXISTS email_record (id TEXT NOT NULL,email_id TEXT NOT NULL,expiry_date TEXT NOT NULL,created_on TIME NOT NULL,is_deleted BOOL NOT NULL);"

	//QUERY
	CREATE_EMAIL_RECORD_QUERY        = "INSERT OR REPLACE INTO email_record(id,email_id,expiry_date,created_on,is_deleted) VALUES(?,?,?,?,?);"
	HARD_DELETE_RECORD_QUERY         = "DELETE FROM email_record WHERE expiry_date< ?;"
	SOFT_DELETE_EXPIRED_RECORD_QUERY = "UPDATE TABLE email_record SET is_deleted = true WHERE WHERE expiry_date< ?"
	LIST_ACTIVE_EMAIL_RECORD_QUERY   = "SELECT * FROM email_record WHERE expiry_date< ? AND is_deleted=false;"
)