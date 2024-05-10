package transaction

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pedrodib/study-expert-session-finance/model/transaction"
	"github.com/pedrodib/study-expert-session-finance/util"
)

// GetTransactions is Used to Get all Transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	transactions := transaction.Transactions{
		transaction.Transaction{Title: "Salário", Amount: 1400.0, Type: 0, CreatedAt: util.StringToTime("2024-05-10T11:10:00")},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateTransactions is used to create transactions
func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

}
