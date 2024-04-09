package view

type View struct {
	Initial  *InitialView
	Register *RegisterView
	Login    *LoginView
	Card     *CardView
}

func NewView(r *RegisterView, in *InitialView, lv *LoginView, c *CardView) *View {
	return &View{
		Register: r,
		Initial:  in,
		Login:    lv,
		Card:     c,
	}
}
