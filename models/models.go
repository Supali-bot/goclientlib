package models

type Data struct {
        DataStruct struct {
                Type           string `json:"type"`
                ID             string `json:"id"`
                Version        int    `json:"version"`
                OrganisationID string `json:"organisation_id"`
                Attributes     struct {
                        Country       string   `json:"country"`
                        BaseCurrency  string   `json:"base_currency"`
                        AccountNumber string   `json:"account_number"`
                        BankID        string   `json:"bank_id"`
                        BankIDCode    string   `json:"bank_id_code"`
                        Bic           string   `json:"bic"`
                        Iban          string   `json:"iban"`
                        Name          []string `json:"name"`
                        Status        string   `json:"status"`
                } `json:"attributes"`
        } `json:"data"`
}

