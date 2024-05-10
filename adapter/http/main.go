package http

import (
	"net/http"

	"github.com/pedrodib/study-expert-session-finance/adapter/http/actuator"
	"github.com/pedrodib/study-expert-session-finance/adapter/http/transaction"
)

/*
Init the application
*/
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
