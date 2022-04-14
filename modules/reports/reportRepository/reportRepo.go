package reportRepository

import (
	"database/sql"
	"errors"

	"github.com/muhammadnurbasari/test-be-majoo/models/reportModel"
	"github.com/muhammadnurbasari/test-be-majoo/modules/reports"
)

type sqlRepository struct {
	Conn *sql.DB
}

//NewReportRepository - will create object that represent that reports.ReportRepository interface
func NewReportRepository(Conn *sql.DB) reports.ReportRepository {
	return &sqlRepository{Conn}
}

// PaginationOmzetMerchants -  get info omzet merchant with pagination
func (config *sqlRepository) PaginationOmzetMerchants(userID int64, start, rowPerPage int) (*reportModel.ResOmzetMerchant, error) {
	rows, err := config.Conn.Query(`SELECT
	CASE WHEN mysubquery.userid IS NULL THEN ? ELSE mysubquery.userid END userid,
	CASE WHEN mysubquery.name IS NULL THEN (SELECT name FROM users WHERE id = ?) ELSE mysubquery.name END name,
	CASE WHEN mysubquery.merchant_name IS NULL THEN (SELECT merchant_name FROM merchants WHERE id = (SELECT id FROM merchants WHERE user_id = ?)) ELSE mysubquery.merchant_name END merchant_name,
	CASE WHEN mysubquery.tanggal = DATE(tm.periodic) THEN mysubquery.tanggal ELSE DATE(tm.periodic) END tanggal,
	CASE WHEN mysubquery.omzet IS NULL THEN 0 ELSE mysubquery.omzet END omzet
	FROM temporary_periodics tm
	LEFT JOIN (
	SELECT u.id userid , u.name, m.id merchantid, m.merchant_name, DATE(t.created_at) tanggal, SUM(t.bill_total) omzet FROM users u
	LEFT JOIN merchants m ON m.user_id = u.id
	LEFT JOIN transactions t ON t.merchant_id = m.id
	WHERE u.id = ? AND DATE(t.created_at) BETWEEN '2021-11-01' AND '2021-11-30'
	GROUP BY DATE(t.created_at)
	ORDER BY DATE(t.created_at) ASC
	) mysubquery ON mysubquery.tanggal = DATE(tm.periodic)
	LIMIT ? OFFSET ?`, userID, userID, userID, userID, rowPerPage, start)

	if err != nil {
		return nil, errors.New("PaginationOmzetMerchants err = " + err.Error())
	}
	defer rows.Close()

	var data []reportModel.DataOmzetMerchants
	for rows.Next() {
		var each reportModel.DataOmzetMerchants

		err = rows.Scan(&each.UserID, &each.Name, &each.MerchantName, &each.Tanggal, &each.Omzet)
		if err != nil {
			return nil, errors.New("PaginationOmzetMerchants err = " + err.Error())
		}

		data = append(data, each)
	}

	resp := reportModel.ResOmzetMerchant{
		Start:      start,
		RowPerPage: rowPerPage,
		Data:       data,
	}

	return &resp, nil

}

// PaginationOmzetOutlet -  get info omzet outlet with pagination
func (config *sqlRepository) PaginationOmzetOutlet(userID int64, start, rowPerPage int) (*reportModel.ResOmzetOutlet, error) {
	rows, err := config.Conn.Query(`SELECT
	CASE WHEN mysubquery.userid IS NULL THEN ? ELSE mysubquery.userid END userid,
	CASE WHEN mysubquery.name IS NULL THEN (SELECT name FROM users WHERE id = ?) ELSE mysubquery.name END name,
	CASE WHEN mysubquery.merchant_name IS NULL THEN (SELECT merchant_name FROM merchants WHERE id = (SELECT id FROM merchants WHERE user_id = ?)) ELSE mysubquery.merchant_name END merchant_name,
	CASE WHEN mysubquery.outlet_name IS NULL THEN '' ELSE mysubquery.outlet_name END outlet_name,
	CASE WHEN mysubquery.tanggal = DATE(tm.periodic) THEN mysubquery.tanggal ELSE DATE(tm.periodic) END tanggal,
	CASE WHEN mysubquery.omzet IS NULL THEN 0 ELSE mysubquery.omzet END omzet
	FROM temporary_periodics tm
	LEFT JOIN (
	SELECT u.id userid , u.name, m.id merchantid, m.merchant_name, o.outlet_name, DATE(t.created_at) tanggal, SUM(t.bill_total) omzet FROM users u
	LEFT JOIN merchants m ON m.user_id = u.id
	LEFT JOIN transactions t ON t.merchant_id = m.id
	LEFT JOIN outlets o ON o.id = t.outlet_id
	WHERE u.id = ? AND DATE(t.created_at) BETWEEN '2021-11-01' AND '2021-11-30'
	GROUP BY DATE(t.created_at), t.outlet_id
	ORDER BY DATE(t.created_at) ASC
	) mysubquery ON mysubquery.tanggal = DATE(tm.periodic)
	LIMIT ? OFFSET ?;`, userID, userID, userID, userID, rowPerPage, start)

	if err != nil {
		return nil, errors.New("PaginationOmzetOutlet err = " + err.Error())
	}
	defer rows.Close()

	var data []reportModel.DataOmzetOutlets
	for rows.Next() {
		var each reportModel.DataOmzetOutlets

		err = rows.Scan(&each.UserID, &each.Name, &each.MerchantName, &each.OutletName, &each.Tanggal, &each.Omzet)
		if err != nil {
			return nil, errors.New("PaginationOmzetOutlet err = " + err.Error())
		}

		data = append(data, each)
	}

	resp := reportModel.ResOmzetOutlet{
		Start:      start,
		RowPerPage: rowPerPage,
		Data:       data,
	}

	return &resp, nil

}
