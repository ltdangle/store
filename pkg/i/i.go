// Interfaces package.
package i

// AdminEntity interface to work with admin panel.
type AdminEntity interface {
	// Primary key of the table.
	PrimaryKey() string
	// Name of the database table for the entity.
	TableName() string
	// Constructor function.
	New() AdminEntity
}
