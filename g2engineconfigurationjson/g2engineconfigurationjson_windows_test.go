//go:build windows

package g2engineconfigurationjson

var testCasesForOsArch = []testCaseMetadata{
	{
		name:        "sqlite3-001",
		databaseUrl: "sqlite3://na:na@nowhere/C:\\Temp\\sqlite\\G2C.db",
		databaseFile: "C:\\Temp\\sqlite\\G2C.db",
	},
	{
		name:        "sqlite3-002",
		databaseUrl: `sqlite3://na:na@nowhere/C:\Temp\sqlite\G2C.db`,
		databaseFile: "C:\Temp\sqlite\G2C.db",
	},
}
