package Server

/**
 * The entity interface regroups methods needed by a structure to be
 * saved and initialized from the key value data store.
 */
type Entity interface {

	/**
	 * Create a new entity.
	 */
	SaveEntity()

	/**
	 * Create and initialize an entity with a given id.
	 * if lazy is set at true sub-entity are not initialised
	 */
	InitEntity(id string, lazy bool) error

	/**
	 * Delete the entity
	 */
	DeleteEntity()

	/**
	 * Return the type name of an entity
	 */
	GetTypeName() string

	/**
	 * Return the size of an object in memory
	 */
	GetSize() uint

	/**
	 * Get an entity's uuid
	 * Each entity must have one uuid.
	 */
	GetUuid() string

	/**
	 * If the entity is created by a parent entity.
	 */
	GetParentPtr() Entity
	SetParentPtr(Entity)

	/**
	 * Return the list of child entity (pointers)
	 */
	GetChildsPtr() []Entity

	/**
	 * Set the array of childs pointers
	 */
	SetChildsPtr(childsPtr []Entity)

	/**
	 * Remove child with a given uuid.
	 */
	RemoveChild(name string, uuid string)

	/**
	 * Append a child.
	 */
	AppendChild(attributeName string, child Entity) error

	/**
	 * Get the childs uuid.
	 */
	GetChildsUuid() []string

	/**
	 * Set the array of childs uuid.
	 */
	SetChildsUuid(uuids []string)

	/**
	 * Return the list of references uuid of an entity
	 */
	GetReferencesUuid() []string

	/**
	 * Set reference uuid
	 */
	SetReferencesUuid(uuid []string)

	/**
	 * Return the list of reference of an entity
	 */
	GetReferencesPtr() []Entity

	/**
	 * Set reference uuid
	 */
	SetReferencesPtr(ref []Entity)

	/**
	 * Append a reference.
	 */
	AppendReference(reference Entity)

	/**
	 * Remove the reference.
	 */
	RemoveReference(name string, reference Entity)

	/**
	 * Append an entity that references this entity.
	 */
	AppendReferenced(name string, owner Entity)

	/**
	 * Remove the referenced.
	 */
	RemoveReferenced(name string, owner Entity)

	/**
	 * Return the list of references.
	 */
	GetReferenced() []EntityRef

	/**
	 * Return the object wrapped by this entity.
	 */
	GetObject() interface{}

	/**
	 * Evaluate if an entity needs to be saved.
	 */
	NeedSave() bool

	/**
	 * Set the need save value.
	 */
	SetNeedSave(bool)

	/**
	 * Test if the entity is initialyse.
	 */
	IsInit() bool

	/**
	 * If the entity is lazy loaded.
	 */
	IsLazy() bool

	/**
	 * Set if an entity must be inityalyse.
	 */
	SetInit(isInit bool)

	/**
	 * Get the checksum value
	 */
	GetChecksum() string

	/**
	 * Determine if an entity exists.
	 */
	Exist() bool

	/**
	 * Return the entity prototype.
	 */
	GetPrototype() *EntityPrototype
}
