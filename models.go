package main

import (
  "gopkg.in/guregu/null.v3"
)

// app/models/business_category.rb
// Table name: business_categories
type BusinessCategory struct {
	Id               null.Int    `db:"id"`
	Name             null.String `db:"name"`
	BusinessTypeId   null.Int    `db:"business_type_id"`
	CreatedAt        null.Time   `db:"created_at"`
	UpdatedAt        null.Time   `db:"updated_at"`
}

// app/models/business_type.rb
// Table name: business_types
type BusinessType struct {
	Id         null.Int    `db:"id"`
	Name       null.String `db:"name"`
	CreatedAt  null.Time   `db:"created_at"`
	UpdatedAt  null.Time   `db:"updated_at"`
}

type BusinessCategoryTypeView struct {
  BusinessCategoryId   null.Int
  BusinessTypeId       null.Int
  BusinessCategoryName null.String
  BusinessTypeName     null.String
}

type BusinessOutletEmailListView struct {
  BusinessId    null.Int
  BusinessName  null.String
  OutletIds     *[]uint8
}
// ====================================================================



// LOW INVENTORY
// ===============
// select * from item_variants where
//
// select business_id from businesses
//
// select setting_id from setting where inventory_alerts = true and bsiness_id in ()
//
// select email from setting_alert_recipients where setting_id in () and is_deleted = false
//
// select * from outlets where business_id in ()
//
//
// select * from item_variants where alert = true and stock_alert < in_stock and outlet in ()


// DAILY SUMMARY
// ================
