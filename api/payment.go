package api
 
import (
	service "cesjb/domain/service/payment"
	handler "cesjb/handlers/payment"
	payment "cesjb/repository/postgresql/payment"
 
	"github.com/jackc/pgx/v5/pgxpool"
)
 
func InitPayment(pool *pgxpool.Pool) handler.Handler {
	r := payment.NewPostgresRepository(pool)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	