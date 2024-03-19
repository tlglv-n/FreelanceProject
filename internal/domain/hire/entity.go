package hire

type Entity struct {
	ID          string  `db:"id" bson:"_id"`
	JobName     *string `db:"job_name" bson:"job_name"`
	Amount      *int    `db:"amount" bson:"amount"`
	Description *string `db:"description" bson:"description"`
	Position    *string `db:"position" bson:"position"`
	CustomerID  string  `db:"customer_id" bson:"customer_id"`
}
