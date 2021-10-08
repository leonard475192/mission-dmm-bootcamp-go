package object

type (
	StatusID = int64

	Status struct {
		ID               StatusID  `json:"id,omitempty"`
		AccountID        int64     `json:"account_id,omitempty"  db:"account_id"`
		Account          *Account  `json:"account,omitempty"`
		Content          string    `json:"content,omitempty"`
		MediaAttachments []int64   `json:"media_attachments,omitempty"`
		CreateAt         *DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
