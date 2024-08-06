package model

import "github.com/golang-jwt/jwt/v5"

// prepare header structure contains algorithm and token type
type JWSHeader struct {
	Alg string   `json:"alg"`
	X5c []string `json:"x5c"`
}

// Notification signed payload
type NotificationPayload struct {
	jwt.RegisteredClaims
	NotificationType      string                    `json:"notificationType"`
	Subtype               string                    `json:"subtype"`
	NotificationVersion   string                    `json:"notificationVersion"`
	Data                  NotificationData          `json:"data"`
	Summary               SummaryData               `json:"summary"`               //subscription-renewal-date extension notification. type: RENEWAL_EXTENSION + subType:SUMMARY
	ExternalPurchaseToken ExternalPurchaseTokenData `json:"externalPurchaseToken"` //type: EXTERNAL_PURCHASE_TOKEN , This notification type applies to apps that use the External Purchase API to offer alternative payment options.
	Version               string                    `json:"version"`
	SignedDate            int64                     `json:"signedDate"`
	NotificationUUID      string                    `json:"notificationUUID"`
}

type ExternalPurchaseTokenData struct {
	ExternalPurchaseId string `json:"externalPurchaseId"`
	TokenCreationDate  int64  `json:"tokenCreationDate"`
	AppAppleId         string `json:"appAppleId"`
	BundleId           string `json:"bundleId"`
}

type SummaryData struct {
	jwt.RegisteredClaims
	RequestIdentifier      string   `json:"requestIdentifier"`
	Environment            string   `json:"environment"`
	AppAppleID             int      `json:"appAppleId"`
	BundleID               string   `json:"bundleId"`
	ProductId              string   `json:"productId"`
	StorefrontCountryCodes []string `json:"storefrontCountryCodes"`
	FailedCount            int      `json:"failedCount"`
	SucceededCount         int      `json:"succeededCount"`
}

// Notification Data
type NotificationData struct {
	jwt.RegisteredClaims
	AppAppleID            int    `json:"appAppleId"`
	BundleID              string `json:"bundleId"`
	BundleVersion         string `json:"bundleVersion"`
	Environment           string `json:"environment"`
	SignedRenewalInfo     string `json:"signedRenewalInfo"`
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

// Notification Transaction Info
type TransactionPayload struct {
	jwt.RegisteredClaims
	AppAccountToken             string `json:"appAccountToken"`
	BundleID                    string `json:"bundleId"`
	Currency                    string `json:"currency"`
	Environment                 string `json:"environment"`
	ExpiresDate                 int    `json:"expiresDate"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	IsUpgraded                  bool   `json:"isUpgraded"`
	OfferDiscountType           string `json:"offerDiscountType"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferType                   int32  `json:"offerType"`
	OriginalPurchaseDate        int    `json:"originalPurchaseDate"`
	OriginalTransactionID       string `json:"originalTransactionId"`
	Price                       int64  `json:"price"`
	ProductID                   string `json:"productId"`
	PurchaseDate                int    `json:"purchaseDate"`
	Quantity                    int    `json:"quantity"`
	RevocationDate              int    `json:"revocationDate"`
	RevocationReason            int32  `json:"revocationReason"`
	SignedDate                  int    `json:"signedDate"`
	Storefront                  string `json:"storefront"`
	StorefrontId                string `json:"storefrontId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	TransactionId               string `json:"transactionId"`
	TransactionReason           string `json:"transactionReason"`
	Type                        string `json:"type"`
	WebOrderLineItemID          string `json:"webOrderLineItemId"`
}

type transactionClaims struct {
	jwt.RegisteredClaims
	AppAccountToken       string `json:"appAccountToken"`
	TransactionId         string `json:"transactionId"`
	OriginalTransactionId string `json:"originalTransactionId"`
	BundleId              string `json:"bundleId"`
	ProductId             string `json:"productId"`
	PurchaseDate          int64  `json:"purchaseDate"`
	OriginalPurchaseDate  int64  `json:"originalPurchaseDate"`
	Quantity              int    `json:"quantity"`
	Type                  string `json:"type"`
	InAppOwnershipType    string `json:"inAppOwnershipType"`
	SignedDate            int64  `json:"signedDate"`
	Environment           string `json:"environment"`
	TransactionReason     string `json:"transactionReason"`
	Storefront            string `json:"storefront"`
	StorefrontId          string `json:"storefrontId"`
	Price                 int    `json:"price"`
	Currency              string `json:"currency"`
}

// Notification Renewal Info
type RenewalInfoPayload struct {
	jwt.RegisteredClaims
	AutoRenewProductId          string   `json:"autoRenewProductId"`
	AutoRenewStatus             int      `json:"autoRenewStatus"`
	Currency                    string   `json:"currency"`
	EligibleWinBackOfferIds     []string `json:"eligibleWinBackOfferIds"`
	Environment                 string   `json:"environment"`
	ExpirationIntent            int      `json:"expirationIntent"`
	GracePeriodExpiresDate      int      `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool     `json:"isInBillingRetryPeriod"`
	OfferDiscountType           string   `json:"offerDiscountType"`
	OfferIdentifier             string   `json:"offerIdentifier"`
	OfferType                   int32    `json:"offerType"`
	OriginalTransactionID       string   `json:"originalTransactionId"`
	PriceIncreaseStatus         int32    `json:"priceIncreaseStatus"`
	ProductID                   string   `json:"productId"`
	RecentSubscriptionStartDate int64    `json:"recentSubscriptionStartDate"`
	RenewalDate                 int64    `json:"renewalDate"`
	RenewalPrice                int64    `json:"renewalPrice"`
	SignedDate                  int      `json:"signedDate"`
}
