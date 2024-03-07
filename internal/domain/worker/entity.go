package worker

type Entity struct {
	ID          string  `db:"id" bson:"_id"`
	FullName    *string `db:"full_name" bson:"full_name"`
	Pseudonym   *string `db:"pseudonym" bson:"pseudonym"`
	Description *string `db:"description" bson:"description"`
	Position    *string `db:"position" bson:"position"`
}
