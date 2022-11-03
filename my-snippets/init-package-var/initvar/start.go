package initvar

var RefA = &VarA

var (
	VarA string
	VarB string
	VarC int
)

func init() {

	VarB = "Initial value"

}