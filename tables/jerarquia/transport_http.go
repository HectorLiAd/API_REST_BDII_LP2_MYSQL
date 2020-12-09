package jerarquia

import (
	// "context"
	// "encoding/json"
	// "net/http"
	// "strconv"

	"net/http"

	"github.com/go-chi/chi"
	// kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPSHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPSHandler(s Service) http.Handler {
	r := chi.NewRouter()

	//Insertar sucursal
	// addSucursalHandler := kithttp.NewServer(
	// 	makeAddSucursalEndPoint(s),
	// 	addSurcursalRequestDecoder,
	// 	kithttp.EncodeJSONResponse,
	// )
	// r.Method(http.MethodPost, "/", addSucursalHandler)

	return r
}

// func addSurcursalRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
// 	request := addSucursalRequest{}
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	return request, err
// }

// func getSurcursalByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
// 	sucursalID, err := strconv.Atoi(chi.URLParam(r, "id"))
// 	rol := getSucursalByIDRequest{
// 		ID: sucursalID,
// 	}
// 	return rol, err
// }
