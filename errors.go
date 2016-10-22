package tengo

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

// IsDatabaseError returns true if err came from a database server, typically
// as a response to a query or connection attempt.
func IsDatabaseError(err error) bool {
	_, result := err.(*mysql.MySQLError)
	return result
}

// IsSyntaxError returns true if err is a SQL syntax error, or false otherwise.
func IsSyntaxError(err error) bool {
	if merr, ok := err.(*mysql.MySQLError); ok {
		return merr.Number == mysqlerr.ER_PARSE_ERROR || merr.Number == mysqlerr.ER_SYNTAX_ERROR
	}
	return false
}

// IsAccessError returns true if err indicates an authentication or authorization
// problem, at connection time or query time. Can be a problem with credentials,
// client host, no access to requested default database, missing privilege, etc.
// There is no sense in immediately retrying the connection or query when
// encountering this type of error.
func IsAccessError(err error) bool {
	if merr, ok := err.(*mysql.MySQLError); ok {
		authErrors := map[uint16]bool{
			mysqlerr.ER_ACCESS_DENIED_ERROR:          true,
			mysqlerr.ER_BAD_HOST_ERROR:               true,
			mysqlerr.ER_DBACCESS_DENIED_ERROR:        true,
			mysqlerr.ER_BAD_DB_ERROR:                 true,
			mysqlerr.ER_HOST_NOT_PRIVILEGED:          true,
			mysqlerr.ER_HOST_IS_BLOCKED:              true,
			mysqlerr.ER_SPECIFIC_ACCESS_DENIED_ERROR: true,
		}
		return authErrors[merr.Number]
	}
	return false
}
