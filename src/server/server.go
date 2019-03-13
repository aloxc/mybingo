package server
import(
	"mybingo1/config"
)

type MybingoServer struct {
	config *config.MySqlConfig
}
func NewMybingoServer(config *config.MySqlConfig)(mybingoServer *MybingoServer,err){
	
	
	fmt.Print("a")
	
}