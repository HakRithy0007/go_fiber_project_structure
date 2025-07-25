package custom_models

import "time"

type PlayerContext struct {
	PlayerID     float64   `json:"player_id"`
	UserName     string    `json:"user_name"`
	LoginSession string    `json:"login_session"`
	Exp          time.Time `json:"exp"`
	UserAgent    string    `json:"user_agent"`
	Ip           string    `json:"ip"`
	MembershipId float64   `json:"membership_id"`
	RoleID       int       `json:"role_id"`
}

// Deadline implements context.Context.
func (p *PlayerContext) Deadline() (deadline time.Time, ok bool) {
	panic("unimplemented")
}

// Done implements context.Context.
func (p *PlayerContext) Done() <-chan struct{} {
	panic("unimplemented")
}

// Err implements context.Context.
func (p *PlayerContext) Err() error {
	panic("unimplemented")
}

// Value implements context.Context.
func (p *PlayerContext) Value(key any) any {
	panic("unimplemented")
}

type Token struct {
	Id       float64 `json:"id"`
	Username string  `json:"user_name"`
	// MembershipId   float64 `json:"membership_id"`
	// MembershipRole string  `json:"membership_role"`
	// RoleId         float64 `json:"role_id"`
}

type Paging struct {
	PerPage int `json:"perpage" query:"per_page" validate:"required"`
	Page    int `json:"page" query:"page" validate:"required"`
}

type Filter struct {
	Property string      `json:"property" query:"property"`
	Value    interface{} `json:"value" query:"value"`
}

type Sort struct {
	Property  string `json:"property" query:"property"`
	Direction string `json:"direction" query:"direction"`
}
