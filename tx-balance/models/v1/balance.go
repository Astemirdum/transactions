package v1

type Balance struct {
	ID              int    `db:"id"`
	UserID          int    `db:"user_id"`
	Cash            uint64 `db:"cash"`
	LastTransaction int64  `db:"last_transaction"`
}

type CashOutMsg struct {
	Cash uint64 `json:"cash"`
}

type CashOut struct {
	UserID int
	Cash   uint64
}
