package dialect

type EQLDialect struct {
	GenericSQLDialect
}

var _ Dialect = &EQLDialect{} // Ensure EQLDialect implements Dialect

func (*EQLDialect) IsIdentifierStart(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || r == '.' // Added '.' as seen in '.name'
}

func (*EQLDialect) IsIdentifierPart(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '.' // Added '.'
}

func (*EQLDialect) IsDelimitedIdentifierStart(r rune) bool {
	return false // EQL spec doesn't mention delimited identifiers, but could be '`' or '"' if needed later
}

func (*EQLDialect) IsPlaceHolderStart(r rune) bool {
	return r == '$' // Functions start with $ in EQL
}

func (*EQLDialect) IsPlaceHolderPart(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == ':' // Added ':' for function names like $my_lib:my_fun
}

func (*EQLDialect) Name() string {
	return "eql"
}

var eqlKeywords = []string{
	"GET", "FROM", "WHERE", "AND", "OR", "ANDNOT", "GROUP BY", "ORDER BY", "PAGE", "LOCK", "FORMAT",
	"SUBSCRIBE", "UNSUBSCRIBE", "SET", "IN", "INSERT", "