package main

import (
  "gopkg.in/guregu/null.v3"
)

// Tested
type AccessToken struct {
  Id        null.Int    `db:"id"`
  Token     null.String `db:token`
  DeviceId  null.Int    `db:device_id`
  UserId    null.Int    `db:user_id`
  OutletId  null.Int    `db:outlet_id`
  UpdatedAt null.String `db:updated_at`
  CreatedAt null.String `db:created_at`
  ExpiredAt null.String `db:expired_at`
}

type Activity struct {
  Id             null.Int      `db:id`
  TrackableId    null.Int     `db:trackable_id`
  TrackableType  null.String  `db:trackable_type`
  OwnerId        null.Int     `db:owner_id`
  OwnerType      null.String  `db:owner_key`
  Key            null.String  `db:key`
  Parameters     null.String  `db:parameters`
  RecipientId    null.Int     `db:recipient_id`
  RecipientType  null.String  `db:recipient_type`
  CreatedAt      null.Time    `db:created_at`
  UpdatedAt      null.Time    `db:updated_at`
}

type BankAccount struct {
  Id                null.Int    `db:id`
  UserId            null.Int    `db:user_id`
  BankType          null.String `db:bank_type`
  AccountHolderName null.String `db:account_holder_name`
  SwiftCode         null.String `db:swift_code`
  AccountNumber     null.String `db:account_number`
  CreatedAt         null.Time   `db:created_at`
  UpdatedAt         null.Time   `db:updated_at`
}

type Billing struct {
  Id                null.Int    `db:id`
  InvoiceNumber     null.String `db:invoice_number`
  BusinessId        null.Int    `db:business_id`
  status            null.String `db:status`
  OutletIds         null.String `db:outlet_ids`
  outlet_count      null.Int    `db:outlet_count`
  Amount            null.Int    `db:amount`
  Discount          null.Int    `db:discount`
  TotalDiscount     null.Int    `db:total_discount`
  TotalAmount       null.Int    `db:total_amount`
  PaidAt            null.Time   `db:paid_at`
  SubscriptionPlan  null.Int    `db:subscription_plan`
  StartPeriod       null.Time   `db:start_period`
  EndPeriod         null.Time   `db:end_period`
  Type              null.String `db:type`
  CreatedAt         null.Time   `db:created_at`
  UpdatedAt         null.Time   `db:updated_at`
}

// Tested
type BusinessCategory struct {
  Id             null.Int     `db:id`
  Name           null.String  `db:name`
  BusinessTypeId null.Int     `db:business_type_id`
  CreatedAt      null.Time    `db:created_at`
  UpdatedAt      null.Time    `db:updated_at`
}

// Tested
type BusinessType struct {
  Id             null.Int    `db:id`
  Name           null.String `db:name`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
}

//Tested
type Business struct {
  Id                 null.Int    `db:id`
  Name               null.String `db:name`
  Address            null.String `db:address`
  Suite              null.String `db:suite`
  City               null.String `db:city`
  Province           null.String `db:province`
  PostalCode         null.String `db:postal_code`
  Description        null.String `db:description`
  Email              null.String `db:email`
  Phone              null.String `db:phone`
  ClientKey          null.String `db:client_key`
  Logo               null.String `db:logo`
  NotifyPerDeposit   null.Bool   `db:notify_per_deposit`
  NotifyPerTrans     null.Bool   `db:notify_per_trans`
  UserId             null.Int    `db:user_id`
  IsDeleted          null.Bool   `db:is_deleted`
  CreatedAt          null.Time   `db:created_at`
  UpdatedAt          null.Time   `db:updated_at`
  BusinessTypeId     null.Int    `db:business_type_id`
  BusinessCategoryId null.Int    `db:business_category_id`
  Website            null.String `db:website`
  Twitter            null.String `db:twitter`
  Facebook           null.String `db:facebook`
  Instagram          null.String `db:instagram`
  Notes              null.String `db:notes`
  Latitude           null.Float  `db:latitude`
  Longitude          null.Float  `db:longitude`
  SynchronizedAt     null.Time   `db:synchronized_at`
  OutletSlot         null.Int    `db:outlet_slot`
  IsTrial            null.Bool   `db:is_trial`
  CurrentPlan        null.Int    `db:current_plan`
  ExpiredAt          null.Time   `db:expired_at`
  LastPaymentDate    null.Time   `db:last_payment_date`
  NextBillingDate    null.Time   `db:next_billing_date`
  OldUserOutletSlot  null.Int    `db:old_user_outlet_slot`
}

//Tested
type Category struct {
  Id             null.Int    `db:id`
  Name           null.String `db:name`
  Description    null.String `db:description`
  BusinessId     null.Int    `db:business_id`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  ColorCode      null.String `db:color_code`
  OutletId       null.Int    `db:outlet_id`
  Guid           null.String `db:guid`
  UniqId         null.String `db:uniq_id`
  SynchronizedAt null.Time   `db:synchronized_at`
}

//Tested
type CheckoutDiscount struct {
  Id                 null.Int    `db:id`
  DiscountId         null.Int    `db:discount_id`
  IsDeleted          null.Bool   `db:is_deleted`
  CreatedAt          null.Time   `db:created_at`
  UpdatedAt          null.Time   `db:updated_at`
  CheckoutId         null.Int    `db:checkout_id`
  DiscountAmount     null.Int    `db:discount_amount`
  DiscountName       null.String `db:discount_name`
  DiscountType       null.String `db:discount_type`
  DiscountPercentage null.Float  `db:discount_percentage`
}

//Tested
type CheckoutModifier struct {
  Id                  null.Int    `db:id`
  CheckoutId          null.Int    `db:checkout_id`
  ModifierId          null.Int    `db:modifier_id`
  ModifierOptionId    null.Int    `db:modifier_option_id`
  ModifierName        null.String `db:modifier_name`
  ModifierOptionName  null.String `db:modifier_option_name`
  Price               null.Int    `db:price`
  CreatedAt           null.Time   `db:created_at`
  UpdatedAt           null.Time   `db:updated_at`
  Cogs                null.Int    `db:cogs`
}

//Tested
type Checkout struct {
  Id                        null.Int    `db:id`
  CustomAmount              null.Int    `db:custom_amount`
  ItemVariantId             null.Int    `db:item_variant_id`
  Quantity                  null.Int    `db:quantity`
  DiscountAmount            null.Int    `db:discount_amount`
  TaxAmount                 null.Int    `db:tax_amount`
  BusinessId                null.Int    `db:business_id`
  PaymentId                 null.Int    `db:payment_id`
  IsDeleted                 null.Bool   `db:is_deleted`
  CreatedAt                 null.Time   `db:created_at`
  UpdatedAt                 null.Time   `db:updated_at`
  ItemId                    null.Int    `db:item_id`
  ItemDiscount              null.Int    `db:item_discount`
  ItemPriceLibrary          null.Int    `db:item_price_library`
  ItemPrice                 null.Int    `db:item_price`
  ItemPriceDiscount         null.Int    `db:item_price_discount`
  GratuityAmount            null.Int    `db:gratuity_amount`
  ItemPriceDiscountGratuity null.Int    `db:item_price_discount_gratuity`
  TotalPrice                null.Int    `db:total_price`
  ItemPriceQuantity         null.Int    `db:item_price_quantity`
  CategoryName              null.String `db:category_name`
  CategoryId                null.Int    `db:category_id`
  ItemName                  null.String `db:item_name`
  ItemVariantName           null.String `db:item_variant_name`
  Sku                       null.String `db:sku`
  Note                      null.String `db:note`
  Cogs                      null.Int    `db:cogs`
  GrossSales                null.Int    `db:gross_sales`
  OutletId                  null.Int    `db:outlet_id`
}

//Tested
type CustomerFeedbackDetail struct {
  Id                 null.Int    `db:id`
  Comment            null.String `db:comment`
  CustomerFeedbackId null.Int    `db:customer_feedback_id`
  CreatedAt          null.Time   `db:created_at`
  UpdatedAt          null.Time   `db:updated_at`
  FeedbackById       null.Int    `db:feedback_by_id`
  FeedbackByType     null.String `db:feedback_by_type`
}

//Tested
type CustomerFeedback struct {
  Id            null.Int    `db:id`
  Mode          null.String `db:mode`
  FeedbackType  null.String `db:feedback_type`
  CustomerId    null.Int    `db:customer_id`
  CreatedAt     null.Time   `db:created_at`
  UpdatedAt     null.Time   `db:updated_at`
  BusinessId    null.Int    `db:business_id`
  FeedbackToken null.String `db:feedback_token`
  PaymentId     null.Int    `db:payment_id`
  OutletId      null.Int    `db:outlet_id`
}

//Tested
type Customer struct {
  Id             null.Int    `db:id`
  Email          null.String `db:email`
  Name           null.String `db:name`
  Phone          null.String `db:phone`
  Address        null.String `db:address`
  City           null.String `db:city`
  State          null.String `db:state`
  PostalCode     null.String `db:postal_code`
  BusinessId     null.Int    `db:business_id`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  image          null.String `db:image`
  Birthday       null.Time   `db:birthday`
  SynchronizedAt null.Time   `db:synchronized_at`
}

//Tested
type Discount struct {
  Id             null.Int    `db:id`
  Name           null.String `db:name`
  Amount         null.Float  `db:amount`
  Type           null.String `db:type`
  BusinessId     null.Int    `db:business_id`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  OutletId       null.Int    `db:outlet_id`
  Guid           null.String `db:guid`
  UniqId         null.String `db:uniq_id`
  SynchronizedAt null.Time   `db:synchronized_at`
}

// type DokuPayment struct {
//
// }
//
// type EmployeeBusiness struct {
//
// }

//Tested
type EmployeeOutlet struct {
  Id             null.Int `db:id`
  EmployeeId     null.Int `db:employee_id`
  OutletId       null.Int `db:outlet_id`
  IsDeleted      null.Bool `db:is_deleted`
  CreatedAt      null.Time `db:created_at`
  UpdatedAt      null.Time `db:updated_at`
  SynchronizedAt null.Time `db:synchronized_at`
  UserType       null.Int  `db:user_type`
}

// Tested by the location always printed as null
type Favorite struct {
  Id             null.Int    `db:id`
  BusinessId     null.Int    `db:business_id`
  ItemId         null.Int    `db:item_id`
  Locationx      null.Int    `db:location_x`
  Locationy      null.Int    `db:location_y`
  FavoriteType   null.String `db:favorite_type`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  OutletId       null.Int    `db:outlet_id`
  Guid           null.String `db:guid`
  UniqId         null.String `db:uniq_id`
  SynchronizedAt null.Time   `db:synchronized_at`
}

//Tested
type Gratuity struct {
  Id             null.Int    `db:id`
  Name           null.String `db:name`
  Amount         null.Float  `db:amount`
  BusinessId     null.Int    `db:business_id`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  OutletId       null.Int    `db:outlet_id`
  Guid           null.String `db:guid`
  UniqId         null.String `db:uniq_id`
  SynchronizedAt null.Time   `db:synchronized_at`
}

//Tested
type History struct {
  Id             null.Int    `db:id`
  OutletId       null.Int    `db:outlet_id`
  Note           null.String `db:note`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  Type           null.String `db:type`
  SupplierId     null.Int `db:supplier_id`
  OrderNo        null.String `db:order_no`
  Total              null.Int    `db:total`
  Status             null.String `db:status`
  CreatedBy          null.Int    `db:created_by`
  ToOutletId         null.Int    `db:to_outlet_id`
  SupplierName       null.String `db:supplier_name`
  SupplierPhone      null.String `db:supplier_phone`
  SupplierEmail      null.String `db:supplier_email`
  SupplierAddress    null.String `db:supplier_address`
  SupplierCity       null.String `db:supplier_city`
  SupplierState      null.String `db:supplier_state`
  SupplierPostalCode null.String `db:supplier_postal_code`
  IsSender           null.Bool   `db:is_sender`
  IsReceiver         null.Bool   `db:is_receiver`
  FromOutletId       null.Int    `db;from_outlet_id`
  UpdatedBy          null.Int    `db:updated_by`
  PaymentId          null.Int    `db:payment_id`
  SynchronizedAt     null.Time   `db:synchronized_at`
}

//Tested
type HistoryDetail struct {
  Id              null.Int    `db:id`
  BeginningStock  null.Int    `db:beginning_stock`
  EndStock        null.Int    `db:end_stock`
  Quantity        null.Int    `db:quantity`
  ItemId          null.Int    `db:item_id`
  ItemVariantId   null.Int    `db:item_variant_id`
  HistoryId       null.Int    `db:history_id`
  CreatedAt       null.Time   `db:created_at`
  UpdatedAt       null.Time   `db:updated_at`
  Price           null.Int    `db:int`
  Adjustment      null.Int    `db:adjustment`
  ItemName        null.String `db:item_name`
  ItemVariantName null.String `db:item_variant_name`
  CategoryName    null.String `db:category_name`
  OutletName      null.String `db:outlet_name`
  IsItem          null.Bool   `db:is_item`
  SynchronizedAt  null.Time   `db:synchronized_at`
}

//Tested
type ItemModifier struct {
  Id              null.Int  `db:id`
  ItemId          null.Int  `db:item_id`
  ModifierId      null.Int  `db:modifier_id`
  IsDeleted       null.Bool `db:is_deleted`
  CreatedAt       null.Time `db:created_at`
  UpdatedAt       null.Time `db:updated_at`
  SynchronizedAt  null.Time `db:synchronized_at`
}

//Tested
type ItemVariant struct {
  Id              null.Int    `db:id`
  Name            null.String `db:name`
  Sku             null.String `db:sku`
  Price           null.Int    `db:price`
  InStock         null.Int    `db:in_stock`
  StockAlert      null.Int    `db:stock_alert`
  Position        null.Int    `db:position`
  ItemId          null.Int    `db:item_id`
  IsDeleted       null.Bool   `db:is_deleted`
  CreatedAt       null.Time   `db:created_at`
  UpdatedAt       null.Time   `db:updated_at`
  AddInventory    null.Int    `db:add_inventory`
  TrackStock      null.Bool   `db:track_stock`
  Alert           null.Bool   `db:alert`
  Cogs            null.Int    `db:cogs`
  SynchronizedAt  null.Time   `db:synchronized_at`
  LastModified    null.Time   `db:last_modified`
}

//Tested
type Item struct {
  Id              null.Int    `db:id`
  Name            null.String `db:name`
  Description     null.String `db:description`
  Image           null.String `db:image`
  BusinessId      null.Int    `db:business_id`
  CategoryId      null.Int    `db:category_id`
  IsDeleted       null.Bool   `db:is_deleted`
  CreatedAt       null.Time   `db:created_at`
  UpdatedAt       null.Time   `db:updated_at`
  BackgroundColor null.String `db:background_color`
  OutletId        null.Int    `db:outlet_id`
  Guid            null.String `db:guid`
  UniqId          null.String `db:uniq_id`
  SynchronizedAt  null.Time   `db:synchronized_at`
}

//Tested
type ModifierOption struct {
  Id              null.Int    `db:id`
  ModifierId      null.Int    `db:modifier_id`
  Name            null.String `db:name`
  Price           null.Int    `db:price`
  IsDeleted       null.Bool   `db:is_deleted`
  CreatedAt       null.Time   `db:created_at`
  UpdatedAt       null.Time   `db:updated_at`
  Position        null.Int    `db:position`
  Cogs            null.Int    `db:cogs`
  OutletId        null.Int    `db:outlet_id`
  SynchronizedAt  null.Time   `db:synchronized_at`
}

//Tested
type Modifier struct {
  Id              null.Int    `db:id`
  BusinessId      null.Int    `db:business_id`
  Name            null.String `db:name`
  IsDeleted       null.Bool   `db:is_deleted`
  CreatedAt       null.Time   `db:created_at`
  UpdatedAt       null.Time   `db:updated_at`
  OutletId        null.Int    `db:outlet_id`
  Guid            null.String `db:guid`
  UniqId          null.String `db:uniq_id`
  SynchronizedAt  null.Time   `db:synchronized_at`
}

// type OauthAccessGrants
// type OauthAccessTokens
// type OauthApplications

//Tested
type Outlet struct {
  Id             null.Int    `db:id`
  BusinessId     null.Int    `db:business_id`
  Name           null.String `db:name`
  Address        null.String `db:address`
  PhoneNumber    null.String `db:phone_number`
  City           null.String `db:city`
  Province       null.String `db:province`
  PostalCode     null.String `db:postal_code`
  Latitude       null.Float  `db:latitude`
  Longitude      null.Float  `db:longitude`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  SynchronizedAt null.Time   `db:synchronized_at`
  Importing      null.Bool   `db:importing`
  Exporting      null.Bool   `db:exporting`
  Notes          null.String `db:notes`
}

//Tested
type PaymentDiscount struct {
  Id             null.Int    `db:id`
  PaymentId      null.Int    `db:payment_id`
  DiscountId     null.Int    `db:discount_id`
  DiscountAmount null.Int    `db:discount_amount`
  DiscountName   null.String `db:discount_name`
  discount_type  null.String `db:discount_type`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
}

//Tested
type PaymentGratuity struct {
  Id          null.Int    `db:id`
  Name        null.String `db:name`
  amount      null.Float  `db:amount`
  Total       null.Int    `db:total`
  PaymentId   null.Int    `db:payment_id`
  CreatedAt   null.Time   `db:created_at`
  UpdatedAt   null.Time   `db:updated_at`
  GratuityId  null.Int    `db:gratuity_id`
}

//Tested
type PaymentRecord struct {
  Id             null.Int    `db:id`
  AmountReceived null.Int    `db:amount_received`
  PaymentType    null.String `db:payment_type`
  Note           null.String `db:note`
  PaymentId      null.Int    `db:payment_id`
  Guid           null.String `db:guid`
  UniqId         null.String `db:uniq_id`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  PaymentDate    null.Time   `db:payment_date`
  SynchronizedAt null.Time   `db:synchronized_at`
  ShiftId        null.Int    `db:shift_id`
}

//Tested
type PaymentTax struct {
  Id            null.Int    `db:id`
  Name          null.String `db:name`
  Amount        null.Float  `db:amount`
  Total         null.Int    `db:total`
  TaxableAmount null.Int    `db:taxable_amount`
  PaymentId     null.Int    `db:payment_id`
  CreatedAt     null.Time   `db:created_at`
  UpdatedAt     null.Time   `db:updated_at`
  TaxId         null.Int    `db:tax_id`
}

//Tested
type Payment struct {
  Id                     null.Int    `db:id`
  Status                 null.String `db:status`
  PaymentNo              null.String `db:payment_no`
  PaymentType            null.String `db:payment_type`
  AmountPay              null.Int    `db:amount_pay`
  AmountChange           null.Int    `db:amount_change`
  TotalDiscountAmount    null.Int    `db:total_discount_amount`
  TotalGratuityAmount    null.Int    `db:total_gratuity_amount`
  TotalItemPriceAmount   null.Int    `db:total_item_price_amount`
  TotalTaxAmount         null.Int    `db:total_tax_amount`
  TotalCheckoutsAmount   null.Int    `db:total_checkouts_amount`
  IsReceiptPrinted       null.Bool   `db:is_receipt_printed`
  IsReceiptEmailed       null.Bool   `db:is_receipt_emailed`
  IsRefunded             null.Bool   `db:is_refunded`
  RefundedReason         null.String `db:refunded_reason`
  RefundedDate           null.Time   `db:refunded_date`
  CustomerFeedback       null.Bool   `db:customer_feedback`
  BusinessId             null.Int    `db:business_id`
  CustomerId             null.Int    `db:customer_id`
  DiscountId             null.Int    `db:discount_id`
  IsDeleted              null.Bool   `db:is_deleted`
  CreatedAt              null.Time   `db:created_at`
  UpdatedAt              null.Time   `db:updated_at`
  TotalCustomPriceAmount null.Int    `db:total_custom_price_amount`
  GratuityId             null.Int    `db:gratuity_id`
  TaxId                  null.Int    `db:tax_id`
  IncludeGratuityTax     null.Bool   `db:include_gratuity_tax`
  ParentPaymentId        null.Int    `db:parent_payment_id`
  RefundAmount           null.Int    `db:refund_amount`
  RefundType             null.String `db:refund_type`
  CardNo                 null.String `db:card_no`
  PartialRefundCompleted null.Bool   `db:partial_refund_completed`
  FullRefundCompleted    null.Bool   `db:full_refund_completed`
  CreatedBy              null.Int    `db:created_by`
  Note                   null.String `db:note`
  TotalCollectedAmount   null.Int    `db:total_collected_amount`
  ServerId               null.Int    `db:server_id`
  ServerName             null.String `db:server_name`
  ServerTitle            null.String `db:server_title`
  AuthCode               null.String `db:auth_code`
  CcName                 null.String `db:cc_name`
  CardType               null.String `db:card_type`
  TransactionNumber      null.String `db:transaction_number`
  TransactionReference   null.String `db:transaction_reference`
  OutletId               null.Int    `db:outlet_id`
  Guid                   null.String `db:guid`
  SynchronizedAt         null.Time   `db:synchronized_at`
  UniqId                 null.String `db:uniq_id`
  TransactionCertificate null.String `db:transaction_certificate`
  TransactionStatusInfo  null.String `db:transaction_status_info`
  MerchantId             null.String `db:merchant_id`
  MposDeviceId           null.String `db:mpos_device_id`
  PgMid                  null.String `db:pg_mid`
  PgSetting              null.String `db:pg_setting`
  OrderInfo              null.String  `db:order_info`
  Tvr                    null.String `db:tvr`
  CvmResult              null.String `db:cvm_result`
  Aid                    null.String `db:aid`
  TransactionDate        null.Time   `db:transaction_date`
  Pii                    null.String `db:pii`
  CollectedBy            null.String `db:collected_by`
  ShiftId                null.Int    `db:shift_id`
  InvoiceNo              null.String `db:invoice_no`
  InvoiceDepositAmount   null.Int    `db:invoice_deposit_amount`
  InvoiceDueDate         null.Time   `db:invoice_due_date`
  InvoiceStatus          null.String `db:invoice_status`
}

//Tested
type Referral struct {
  Id        null.Int    `db:id`
  Name      null.String `db:name`
  Code      null.String `db:code`
  Emails    null.String `db:emails`
  Logo      null.String `db:logo`
  CreatedAt null.Time   `db:created_at`
  UpdatedAt null.Time   `db:updated_at`
}

//Tested
type SettingAlertRecipient struct {
  Id        null.Int    `db:id`
  SettingId null.Int    `db:setting_id`
  Email     null.String `db:email`
  CreatedAt null.Time   `db:created_at`
  UpdatedAt null.Time   `db:updated_at`
  IsDeleted null.Bool   `db:is_deleted`
}

//Tested
type Setting struct {
  Id                 null.Int    `db:id`
  EnableTax          null.Bool   `db:enable_tax`
  CustomerSignOnIpad null.Bool   `db:customer_sign_on_ipad`
  EnableGratuity     null.Bool   `db:enable_gratuity`
  LogoOnReceipt      null.Bool   `db:logo_on_receipt`
  BusinessId         null.Int    `db:business_id`
  IsDeleted          null.Bool   `db:is_deleted`
  CreatedAt          null.Time   `db:created_at`
  UpdatedAt          null.Time   `db:updated_at`
  IncludeGratuityTax null.Bool   `db:include_gratuity_tax`
  StartFilteredDate  null.Time   `db:start_filtered_date`
  EndFilteredDate    null.Time   `db:end_filtered_date`
  AllowStaffRefund   null.Bool   `db:allow_staff_refund`
  AuthCode           null.String `db:auth_code`
  DailySalesSummary  null.Bool   `db:daily_sales_summary`
  ProductUpdates     null.Bool   `db:product_updates`
  InventoryAlerts    null.Bool   `db:inventory_alerts`
  CustomerFeedback   null.Bool   `db:customer_feedback`
  TrackServer        null.Bool   `db:track_server`
  StaffDeleteBill    null.Bool   `db:staff_delete_bill`
  SynchronizedAt     null.Time   `db:synchronized_at`
}

//Tested
type ShiftDetail struct {
  Id             null.Int    `db:id`
  ShiftId        null.Int    `db:shift_id`
  Amount         null.Int    `db:amount`
  Description    null.String `db:description`
  Type           null.String `db:type`
  SynchronizedAt null.Time   `db:synchronized_at`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  UniqId         null.String `db:uniq_id`
  Guid           null.String `db:guid`
}

//Tested
type Shift struct {
  Id                     null.Int    `db:id`
  OutletId               null.Int    `db:outlet_id`
  EmployeeId             null.Int    `db:employee_id`
  StartingCash           null.Int    `db:starting_cash`
  CashSales              null.Int    `db:cash_sales`
  CashRefund             null.Int    `db:cash_refund`
  TotalExpected          null.Int    `db:total_expected`
  TotalActual            null.Int    `db:total_actual`
  TotalDifference        null.Int    `db:total_difference`
  EndedAt                null.Int    `db:ended_at`
  SynchronizedAt         null.Time   `db:synchronized_at`
  CreatedAt              null.Time   `db:created_at`
  UpdatedAt              null.Time   `db:updated_at`
  ExpectedEndingCash     null.Int    `db:expected_ending_cash`
  ActualEndingCash       null.Int    `db:actual_ending_cash`
  TotalDebitCredit       null.Int    `db:total_debit_credit`
  TotalVoided            null.Int    `db:total_voided`
  ExpectedCardPayment    null.Int    `db:expected_card_payment`
  TotalEdc               null.Int    `db:total_edc`
  TotalGiftCard          null.Int    `db:total_gift_card`
  TotalOther             null.Int    `db:total_other`
  OutletName             null.String `db:outlet_name`
  EmployeeName           null.String `db:employee_name`
  TotalExpense           null.Int    `db:total_expense`
  TotalIncome            null.Int    `db:total_income`
  UniqId                 null.String `db:uniq_id`
  Guid                   null.String `db:guid`
  OtherRefunds           null.Int    `db:other_refunds`
  ExpectedOtherPayment   null.Int    `db:expected_other_payment`
  TotalInvoice           null.Int    `db:total_invoice`
  TotalCancelledInvoice  null.Int    `db:total_cancelled_invoice`
  ExpectedInvoicePayment null.Int    `db:expected_invoice_payment`
  CashInvoice            null.Int    `db:cash_invoice`
}

//Tested
type Supplier struct {
  Id         null.Int    `db:id`
  Name       null.String `db:name`
  Phone      null.String `db:phone`
  Email      null.String `db:email`
  Address    null.String `db:address`
  City       null.String `db:city`
  State      null.String `db:state`
  PostalCode null.String `db:postal_code`
  IsDeleted  null.String `db:is_deleted`
  BusinessId null.Int    `db:business_id`
  CreatedAt  null.Time   `db:created_at`
  UpdatedAt  null.Time   `db:updated_at`
}

//Tested
type Tax struct {
  Id             null.Int    `db:id`
  Name           null.String `db:name`
  Amount         null.Float  `db:amount`
  BusinessId     null.Int    `db:business_id`
  IsDeleted      null.Bool   `db:is_deleted`
  CreatedAt      null.Time   `db:created_at`
  UpdatedAt      null.Time   `db:updated_at`
  OutletId       null.Int    `db:outlet_id`
  Guid           null.String `db:guid`
  UniqId         null.String `db:uniq_id`
  SynchronizedAt null.Time   `db:synchronized_at`
}

//Tested
type Unit struct {
  Id        null.Int    `db:id`
  Name      null.String `db:name`
  CreatedAt null.Time   `db:created_at`
  UpdatedAt null.Time   `db:updated_at`
}

//Tested
type UserDevice struct {
  Id                  null.Int    `db:id`
  BusinessId          null.Int    `db:business_id`
  DeviceId            null.String `db:device_id`
  UniqId              null.String `db:uniq_id`
  CreatedAt           null.Time   `db:created_at`
  UpdatedAt           null.Time   `db:updated_at`
  ReceiptCountre      null.Int    `db:receipt_counter`
  AppVersion          null.String `db:app_version`
  UserId              null.Int    `db:user_id`
  OutletId            null.Int    `db:outlet_id`
  IsShiftAuto         null.Bool   `db:is_shift_auto`
  DefaultStartingCash null.Int    `db:default_starting_cash`
  InvoiceCounter      null.Int    `db:invoice_counter`
}

//Tested
type User struct {
  Id                      null.Int    `db:id`
  Email                   null.String `db:email`
  PasswordDigest          null.String `db:password_digest`
  FirstName               null.String `db:first_name`
  LastName                null.String `db:last_name`
  Picture                 null.String `db:picture`
  Thumbnail               null.String `db:thumbnail`
  RememberToken           null.String `db:remember_token`
  SecurityQuestion        null.String `db:security_question`
  SecurityQuestionAnswer  null.String `db:security_question_answer`
  IsDeleted               null.Bool   `db:is_deleted`
  CreatedAt               null.Time   `db:created_at`
  UpdatedAt               null.Time   `db:updated_at`
  IsOwner                 null.Bool   `db:is_owner`
  IsConfirmed             null.Bool   `db:is_confirmed`
  Phone                   null.String `db:phone`
  ForgotPasswordToken     null.String `db:forgot_password_token`
  ForgotPasswordExpiredAt null.Time   `db:forgot_password_expired_at`
  ImportItemsFile         null.String `db:import_items_file`
  ImportInventoryFile     null.String `db:import_inventory_file`
  OutletId                null.String `db:outlet_id`
  SynchronizedAt          null.Time   `db:synchronized_at`
  ReferralId              null.Int    `db:referral_id`
  ImportIncomingFile      null.String `db:import_incoming_file`
  ImportCustomersFile     null.String `db:import_customers_file`
  IsManager               null.Bool   `db:is_manager`
  ImportAdjustmentFile    null.String `db:import_adjustment_file`
  ExportItemsFile         null.String `db:export_items_file`
  Importing               null.Bool   `db:importing`
  Exporting               null.Bool   `db:exporting`
  ImportSuppliersFile     null.String `db:import_suppliers_file`
  EncryptedPassword       null.String `db:encrypted_password`
  ResetPasswordToken      null.String `db:reset_password_token`
  ResetPasswordSentAt     null.Time   `db:reset_password_sent_at`
  SignInCount             null.Int    `db:sign_in_count`
  CurrentSignInAt         null.Time   `db:current_sign_in_at`
  LastSignInAt            null.Time   `db:last_sign_in_at`
  CurrentSignInIp         null.String `db:last_sign_in_ip;sql:"type:inet;"`
  LastSignInIp            null.String `db:last_sign_in_ip;sql:"type:inet;"`
  RememberCreatedAt       null.Time   `db:remember_created_at`
}
