package gastrofix

import "time"

type BusinessPeriods []BusinessPeriod

type BusinessPeriod struct {
	BusinessDay           string    `json:"businessDay"`
	PeriodID              int       `json:"periodId"`
	StartPeriodTimestamp  time.Time `json:"startPeriodTimestamp"`
	FinishPeriodTimestamp time.Time `json:"finishPeriodTimestamp"`
}

type Transactions []Transaction

type Transaction struct {
	Head struct {
		ReferenceTrUUID  interface{} `json:"referenceTrUuid"`
		StartedTimestamp string      `json:"startedTimestamp"`
		TrainingFlag     bool        `json:"trainingFlag"`
		TypeCode         string      `json:"typeCode"`
		UUID             string      `json:"uuid"`
		VoidedTrUUID     interface{} `json:"voidedTrUuid"`
	} `json:"head"`
	LineItems []struct {
		Amounts struct {
			ActualSubTotal             int64 `json:"actualSubTotal"`
			AppliedToTransactionAmount int64 `json:"appliedToTransactionAmount"`
			DiscountSubTotal           int64 `json:"discountSubTotal"`
			ExchangeRate               int64 `json:"exchangeRate"`
			ForeignCurrencyAmount      int64 `json:"foreignCurrencyAmount"`
			OverAllSubTotal            int64 `json:"overAllSubTotal"`
			Quantity                   int64 `json:"quantity"`
			RegularAmount              int64 `json:"regularAmount"`
			RegularSubTotal            int64 `json:"regularSubTotal"`
			RegularUnitPrice           int64 `json:"regularUnitPrice"`
			TaxAmount                  int64 `json:"taxAmount"`
			TaxPercent                 int64 `json:"taxPercent"`
			TaxSalesGrossAmount        int64 `json:"taxSalesGrossAmount"`
			TaxSalesNetAmount          int64 `json:"taxSalesNetAmount"`
			TenderAmount               int64 `json:"tenderAmount"`
			TipAmount                  int64 `json:"tipAmount"`
			TipForeignCurrencyAmount   int64 `json:"tipForeignCurrencyAmount"`
			TipSubTotal                int64 `json:"tipSubTotal"`
			Units                      int64 `json:"units"`
			NewAmount                  int64 `json:"newamount"`
			PreviousAmount             int64 `json:"previousAmount"`
			AppliedAmount              int64 `json:"appliedAmount"`
			AppliedPercent             int64 `json:"appliedPercent"`
		} `json:"amounts"`
		Extras struct {
			AssociatedLineItemSequenceNumber int64       `json:"associatedLineItemSequenceNumber"`
			AppVersion                       string      `json:"appVersion"`
			AttachmentSubTypeCode            interface{} `json:"attachmentSubTypeCode"`
			AttachmentTypeCode               string      `json:"attachmentTypeCode"`
			Blob                             string      `json:"blob"`
			BusinessCases                    interface{} `json:"businessCases"`
			CardAcquirerReferenceID          interface{} `json:"cardAcquirerReferenceId"`
			CardApplicationID                interface{} `json:"cardApplicationId"`
			CardAuthorizationCode            interface{} `json:"cardAuthorizationCode"`
			CardEntryMethodCode              interface{} `json:"cardEntryMethodCode"`
			CardExpiryDate                   interface{} `json:"cardExpiryDate"`
			CardFingerPrint                  interface{} `json:"cardFingerPrint"`
			CardHolderName                   interface{} `json:"cardHolderName"`
			CardLastFour                     interface{} `json:"cardLastFour"`
			CardMerchantID                   interface{} `json:"cardMerchantId"`
			CardTerminalID                   interface{} `json:"cardTerminalId"`
			CardTerminalReferenceID          interface{} `json:"cardTerminalReferenceId"`
			CardType                         interface{} `json:"cardType"`
			CardVerificationMethod           interface{} `json:"cardVerificationMethod"`
			ClosingBalances                  struct {
				One12 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"112"`
				One14 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"114"`
				One15 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"115"`
				Total []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"total"`
			} `json:"closingBalances"`
			Comment interface{} `json:"comment"`
			Company struct {
				City        interface{} `json:"city"`
				CountryCode interface{} `json:"countryCode"`
				ID          int64       `json:"id"`
				Name        interface{} `json:"name"`
				PostalCode  interface{} `json:"postalCode"`
				Street      interface{} `json:"street"`
				TaxID       interface{} `json:"taxId"`
				TaxNumber   interface{} `json:"taxNumber"`
			} `json:"company"`
			CurrencyCode        string      `json:"currencyCode"`
			CustomerCity        interface{} `json:"customerCity"`
			CustomerCountryCode interface{} `json:"customerCountryCode"`
			CustomerGroupName   string      `json:"customerGroupName"`
			CustomerName        string      `json:"customerName"`
			CustomerPostalCode  interface{} `json:"customerPostalCode"`
			CustomerStreet      interface{} `json:"customerStreet"`
			DeviceUdid          string      `json:"deviceUdid"`
			Devices             struct {
				Master struct {
					CostCenterID int64 `json:"costCenterId"`
					Hardware     struct {
						Brand   string `json:"brand"`
						Model   string `json:"model"`
						Version string `json:"version"`
					} `json:"hardware"`
					IPAddress             string `json:"ipAddress"`
					IsCustomerDisplayFlag bool   `json:"isCustomerDisplayFlag"`
					Name                  string `json:"name"`
					OutletID              int64  `json:"outletId"`
					PriceLevelID          int64  `json:"priceLevelId"`
					Software              struct {
						Brand   string `json:"brand"`
						Version string `json:"version"`
					} `json:"software"`
					Udid string `json:"udid"`
				} `json:"master"`
				SecureElements interface{} `json:"secureElements"`
				Slaves         interface{} `json:"slaves"`
			} `json:"devices"`
			DiscountName               string      `json:"discountName"`
			DivisionName               string      `json:"divisionName"`
			FromLineItemSequenceNumber int         `json:"fromLineItemSequenceNumber"`
			GroupName                  string      `json:"groupName"`
			HotelGuestName             string      `json:"hotelGuestName"`
			HotelPaymentProviderCode   string      `json:"hotelPaymentProviderCode"`
			HotelReservationID         string      `json:"hotelReservationId"`
			HotelRoomID                string      `json:"hotelRoomId"`
			HotelRoomName              string      `json:"hotelRoomName"`
			ItemBarcode                interface{} `json:"itemBarcode"`
			ItemEntryMethod            string      `json:"itemEntryMethod"`
			ItemKind                   int64       `json:"itemKind"`
			ItemName                   string      `json:"itemName"`
			ItemShortName              string      `json:"itemShortName"`
			Location                   struct {
				BaseCurrency string      `json:"baseCurrency"`
				City         interface{} `json:"city"`
				CountryCode  interface{} `json:"countryCode"`
				ID           int64       `json:"id"`
				Name         interface{} `json:"name"`
				PostalCode   interface{} `json:"postalCode"`
				Street       interface{} `json:"street"`
				TaxID        interface{} `json:"taxId"`
				Timezone     string      `json:"timezone"`
			} `json:"location"`
			MobilePaymentProviderCode interface{} `json:"mobilePaymentProviderCode"`
			OpeningBalances           struct {
				One12 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"112"`
				One14 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"114"`
				One15 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"115"`
				Total []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"total"`
			} `json:"openingBalances"`
			Party     int64       `json:"party"`
			PartyName interface{} `json:"partyName"`
			PartySize int64       `json:"partySize"`
			Payments  struct {
				One14 []struct {
					Amount               int64  `json:"amount"`
					BaseCurrencyAmount   int64  `json:"baseCurrencyAmount"`
					CurrencyCode         string `json:"currencyCode"`
					DisbursementAmount   int64  `json:"disbursementAmount"`
					ExpenseAmount        int64  `json:"expenseAmount"`
					FundsReceiptAmount   int64  `json:"fundsReceiptAmount"`
					ID                   int64  `json:"id"`
					LoanAmount           int64  `json:"loanAmount"`
					Name                 string `json:"name"`
					PickupAmount         int64  `json:"pickupAmount"`
					SalesAmount          int64  `json:"salesAmount"`
					TipAmount            int64  `json:"tipAmount"`
					TransferInAmount     int64  `json:"transferInAmount"`
					TransferOutAmount    int64  `json:"transferOutAmount"`
					TypeCode             string `json:"typeCode"`
					UntaxedVoucherAmount int64  `json:"untaxedVoucherAmount"`
				} `json:"114"`
				Total []struct {
					Amount               int64  `json:"amount"`
					BaseCurrencyAmount   int64  `json:"baseCurrencyAmount"`
					CurrencyCode         string `json:"currencyCode"`
					DisbursementAmount   int64  `json:"disbursementAmount"`
					ExpenseAmount        int64  `json:"expenseAmount"`
					FundsReceiptAmount   int64  `json:"fundsReceiptAmount"`
					ID                   int64  `json:"id"`
					LoanAmount           int64  `json:"loanAmount"`
					Name                 string `json:"name"`
					PickupAmount         int64  `json:"pickupAmount"`
					SalesAmount          int64  `json:"salesAmount"`
					TipAmount            int64  `json:"tipAmount"`
					TransferInAmount     int64  `json:"transferInAmount"`
					TransferOutAmount    int64  `json:"transferOutAmount"`
					TypeCode             string `json:"typeCode"`
					UntaxedVoucherAmount int64  `json:"untaxedVoucherAmount"`
				} `json:"total"`
			} `json:"payments"`
			PreviousLineItemSequenceNumber string      `json:"previousLineItemSequenceNumber"`
			PriceEntryMethod               string      `json:"priceEntryMethod"`
			PriceLevelName                 interface{} `json:"priceLevelName"`
			ReasonCategoryCode             interface{} `json:"reasonCategoryCode"`
			ReasonName                     interface{} `json:"reasonName"`
			ResetBalances                  struct {
				One12 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"112"`
				One14 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"114"`
				One15 []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"115"`
				Total []struct {
					Amount             int64  `json:"amount"`
					BaseCurrencyAmount int64  `json:"baseCurrencyAmount"`
					CurrencyCode       string `json:"currencyCode"`
					ID                 int64  `json:"id"`
					Name               string `json:"name"`
					TypeCode           string `json:"typeCode"`
				} `json:"total"`
			} `json:"resetBalances"`
			ReferencedTrUuid string `json:"referencedTrUuid"`
			SalesSummaries   struct {
				Divisions []struct {
					DiscountAmount           int64  `json:"discountAmount"`
					DiscountTransactionCount int64  `json:"discountTransactionCount"`
					DiscountUnitCount        int64  `json:"discountUnitCount"`
					DivisionID               int64  `json:"divisionId"`
					Name                     string `json:"name"`
					SalesAmount              int64  `json:"salesAmount"`
					SalesGuestCount          int64  `json:"salesGuestCount"`
					SalesTaxAmount           int64  `json:"salesTaxAmount"`
					SalesTransactionCount    int64  `json:"salesTransactionCount"`
					SalesUnitCount           int64  `json:"salesUnitCount"`
					Tax                      []struct {
						SalesAmount    int64 `json:"salesAmount"`
						SalesTaxAmount int64 `json:"salesTaxAmount"`
						TaxPercent     int64 `json:"taxPercent"`
						VatID          int64 `json:"vatId"`
						VoidAmount     int64 `json:"voidAmount"`
						VoidTaxAmount  int64 `json:"voidTaxAmount"`
					} `json:"tax"`
					VoidAmount           int64 `json:"voidAmount"`
					VoidTaxAmount        int64 `json:"voidTaxAmount"`
					VoidTransactionCount int64 `json:"voidTransactionCount"`
					VoidUnitCount        int64 `json:"voidUnitCount"`
				} `json:"divisions"`
				Groups []struct {
					DiscountAmount           int64  `json:"discountAmount"`
					DiscountTransactionCount int64  `json:"discountTransactionCount"`
					DiscountUnitCount        int64  `json:"discountUnitCount"`
					GroupID                  int64  `json:"groupId"`
					Name                     string `json:"name"`
					SalesAmount              int64  `json:"salesAmount"`
					SalesGuestCount          int64  `json:"salesGuestCount"`
					SalesTaxAmount           int64  `json:"salesTaxAmount"`
					SalesTransactionCount    int64  `json:"salesTransactionCount"`
					SalesUnitCount           int64  `json:"salesUnitCount"`
					Tax                      []struct {
						SalesAmount    int64 `json:"salesAmount"`
						SalesTaxAmount int64 `json:"salesTaxAmount"`
						TaxPercent     int64 `json:"taxPercent"`
						VatID          int64 `json:"vatId"`
						VoidAmount     int64 `json:"voidAmount"`
						VoidTaxAmount  int64 `json:"voidTaxAmount"`
					} `json:"tax"`
					VoidAmount           int64 `json:"voidAmount"`
					VoidTaxAmount        int64 `json:"voidTaxAmount"`
					VoidTransactionCount int64 `json:"voidTransactionCount"`
					VoidUnitCount        int64 `json:"voidUnitCount"`
				} `json:"groups"`
				Items []struct {
					DiscountAmount           int64  `json:"discountAmount"`
					DiscountTransactionCount int64  `json:"discountTransactionCount"`
					DiscountUnitCount        int64  `json:"discountUnitCount"`
					Name                     string `json:"name"`
					SalesAmount              int64  `json:"salesAmount"`
					SalesGuestCount          int64  `json:"salesGuestCount"`
					SalesTaxAmount           int64  `json:"salesTaxAmount"`
					SalesTransactionCount    int64  `json:"salesTransactionCount"`
					SalesUnitCount           int64  `json:"salesUnitCount"`
					Sku                      int64  `json:"sku"`
					Tax                      []struct {
						SalesAmount    int64 `json:"salesAmount"`
						SalesTaxAmount int64 `json:"salesTaxAmount"`
						TaxPercent     int64 `json:"taxPercent"`
						VatID          int64 `json:"vatId"`
						VoidAmount     int64 `json:"voidAmount"`
						VoidTaxAmount  int64 `json:"voidTaxAmount"`
					} `json:"tax"`
					VoidAmount           int64 `json:"voidAmount"`
					VoidTaxAmount        int64 `json:"voidTaxAmount"`
					VoidTransactionCount int64 `json:"voidTransactionCount"`
					VoidUnitCount        int64 `json:"voidUnitCount"`
				} `json:"items"`
			} `json:"salesSummaries"`
			SignatureAlgorithm               interface{} `json:"signatureAlgorithm"`
			SignatureCounter                 interface{} `json:"signatureCounter"`
			SignatureData                    interface{} `json:"signatureData"`
			SignatureDateTimeFormat          interface{} `json:"signatureDateTimeFormat"`
			SignatureError                   interface{} `json:"signatureError"`
			SignatureProcessData             interface{} `json:"signatureProcessData"`
			SignatureProcessType             interface{} `json:"signatureProcessType"`
			SignaturePublicKey               interface{} `json:"signaturePublicKey"`
			SignatureReferenceData           interface{} `json:"signatureReferenceData"`
			SignatureReferenceNumber         interface{} `json:"signatureReferenceNumber"`
			SignatureSequenceNumber          interface{} `json:"signatureSequenceNumber"`
			SignatureSerial                  interface{} `json:"signatureSerial"`
			SignatureTransactionID           interface{} `json:"signatureTransactionId"`
			TaxedLineItemSequenceNumber      int64       `json:"taxedLineItemSequenceNumber"`
			TenderName                       string      `json:"tenderName"`
			TenderTypeCode                   string      `json:"tenderTypeCode"`
			ToLineItemSequenceNumber         int         `json:"toLineItemSequenceNumber"`
			TransferTypeCode                 string      `json:"transferTypeCode"`
			TriggeringLineItemSequenceNumber int64       `json:"triggeringLineItemSequenceNumber"`
			UnitName                         interface{} `json:"unitName"`
			Vats                             []struct {
				Description string `json:"description"`
				ID          int64  `json:"id"`
				Name        string `json:"name"`
				Percent     int64  `json:"percent"`
			} `json:"vats"`
			VoidedLineItemSequenceNumber interface{} `json:"voidedLineItemSequenceNumber"`
			VoucherProviderCode          interface{} `json:"voucherProviderCode"`
			VoucherSerial                interface{} `json:"voucherSerial"`
			WaiterName                   string      `json:"waiterName"`
			Waiters                      []struct {
				FirstName       string `json:"firstName"`
				ID              int64  `json:"id"`
				IsActiveFlag    bool   `json:"isActiveFlag"`
				IsSafeUserFlag  bool   `json:"isSafeUserFlag"`
				LastName        string `json:"lastName"`
				PersonnelNumber string `json:"personnelNumber"`
			} `json:"waiters"`
		} `json:"extras"`
		Flags struct {
			IsBusinessCaseCompleteFlag bool `json:"isBusinessCaseCompleteFlag"`
			IsCancelFlag               bool `json:"isCancelFlag"`
			IsCarryForwardFlag         bool `json:"isCarryForwardFlag"`
			IsCarryOverFlag            bool `json:"isCarryOverFlag"`
			IsChangeFlag               bool `json:"isChangeFlag"`
			IsInfoFlag                 bool `json:"isInfoFlag"`
			IsNegativeItemFlag         bool `json:"isNegativeItemFlag"`
			IsNonDiscountableFlag      bool `json:"isNonDiscountableFlag"`
			IsNonTurnoverFlag          bool `json:"isNonTurnoverFlag"`
			IsOfflineFlag              bool `json:"isOfflineFlag"`
			IsRestoreFlag              bool `json:"isRestoreFlag"`
			IsSplitFlag                bool `json:"isSplitFlag"`
			IsSuspendedFlag            bool `json:"isSuspendedFlag"`
			IsToGoFlag                 bool `json:"isToGoFlag"`
			IsTransitoryFlag           bool `json:"isTransitoryFlag"`
			IsVoidFlag                 bool `json:"isVoidFlag"`
		} `json:"flags"`
		Primary struct {
			DayDeviceTransactionSequenceNumber int64       `json:"dayDeviceTransactionSequenceNumber"`
			FiscalOperationSequenceNumber      interface{} `json:"fiscalOperationSequenceNumber"`
			InvoiceNumber                      string      `json:"invoiceNumber"`
			OrderSequenceNumber                int64       `json:"orderSequenceNumber"`
			TransactionSequenceNumber          int64       `json:"transactionSequenceNumber"`
		} `json:"primary"`
		Related struct {
			ArticleTypeID   int64       `json:"articleTypeId"`
			CostCenterID    int64       `json:"costCenterId"`
			CourseID        int64       `json:"courseId"`
			CustomerUUID    string      `json:"customerUuid"`
			DeviceID        int64       `json:"deviceId"`
			DivisionID      int64       `json:"divisionId"`
			GroupID         int64       `json:"groupId"`
			ItemSku         int64       `json:"itemSku"`
			PriceLevelID    interface{} `json:"priceLevelId"`
			ReasonClassID   interface{} `json:"reasonClassId"`
			ReasonID        interface{} `json:"reasonId"`
			RoomID          int64       `json:"roomId"`
			SecureElementID interface{} `json:"secureElementId"`
			ServicePeriodID interface{} `json:"servicePeriodId"`
			TableID         int64       `json:"tableId"`
			TenderID        int64       `json:"tenderId"`
			UnitID          interface{} `json:"unitId"`
			VatID           int64       `json:"vatId"`
			WaiterID        int64       `json:"waiterId"`
		} `json:"related"`
		SalesSummaries struct {
			Divisions []struct {
				DiscountAmount           int64  `json:"discountAmount"`
				DiscountTransactionCount int64  `json:"discountTransactionCount"`
				DiscountUnitCount        int64  `json:"discountUnitCount"`
				DivisionID               int64  `json:"divisionId"`
				Name                     string `json:"name"`
				SalesAmount              int64  `json:"salesAmount"`
				SalesGuestCount          int64  `json:"salesGuestCount"`
				SalesTaxAmount           int64  `json:"salesTaxAmount"`
				SalesTransactionCount    int64  `json:"salesTransactionCount"`
				SalesUnitCount           int64  `json:"salesUnitCount"`
				Tax                      []struct {
					SalesAmount    int64 `json:"salesAmount"`
					SalesTaxAmount int64 `json:"salesTaxAmount"`
					TaxPercent     int64 `json:"taxPercent"`
					VatID          int64 `json:"vatId"`
					VoidAmount     int64 `json:"voidAmount"`
					VoidTaxAmount  int64 `json:"voidTaxAmount"`
				} `json:"tax"`
				VoidAmount           int64 `json:"voidAmount"`
				VoidTaxAmount        int64 `json:"voidTaxAmount"`
				VoidTransactionCount int64 `json:"voidTransactionCount"`
				VoidUnitCount        int64 `json:"voidUnitCount"`
			} `json:"divisions"`
			Groups []struct {
				DiscountAmount           int64  `json:"discountAmount"`
				DiscountTransactionCount int64  `json:"discountTransactionCount"`
				DiscountUnitCount        int64  `json:"discountUnitCount"`
				GroupID                  int64  `json:"groupId"`
				Name                     string `json:"name"`
				SalesAmount              int64  `json:"salesAmount"`
				SalesGuestCount          int64  `json:"salesGuestCount"`
				SalesTaxAmount           int64  `json:"salesTaxAmount"`
				SalesTransactionCount    int64  `json:"salesTransactionCount"`
				SalesUnitCount           int64  `json:"salesUnitCount"`
				Tax                      []struct {
					SalesAmount    int64 `json:"salesAmount"`
					SalesTaxAmount int64 `json:"salesTaxAmount"`
					TaxPercent     int64 `json:"taxPercent"`
					VatID          int64 `json:"vatId"`
					VoidAmount     int64 `json:"voidAmount"`
					VoidTaxAmount  int64 `json:"voidTaxAmount"`
				} `json:"tax"`
				VoidAmount           int64 `json:"voidAmount"`
				VoidTaxAmount        int64 `json:"voidTaxAmount"`
				VoidTransactionCount int64 `json:"voidTransactionCount"`
				VoidUnitCount        int64 `json:"voidUnitCount"`
			} `json:"groups"`
			Items []struct {
				DiscountAmount           int64  `json:"discountAmount"`
				DiscountTransactionCount int64  `json:"discountTransactionCount"`
				DiscountUnitCount        int64  `json:"discountUnitCount"`
				Name                     string `json:"name"`
				SalesAmount              int64  `json:"salesAmount"`
				SalesGuestCount          int64  `json:"salesGuestCount"`
				SalesTaxAmount           int64  `json:"salesTaxAmount"`
				SalesTransactionCount    int64  `json:"salesTransactionCount"`
				SalesUnitCount           int64  `json:"salesUnitCount"`
				Sku                      int64  `json:"sku"`
				Tax                      []struct {
					SalesAmount    int64 `json:"salesAmount"`
					SalesTaxAmount int64 `json:"salesTaxAmount"`
					TaxPercent     int64 `json:"taxPercent"`
					VatID          int64 `json:"vatId"`
					VoidAmount     int64 `json:"voidAmount"`
					VoidTaxAmount  int64 `json:"voidTaxAmount"`
				} `json:"tax"`
				VoidAmount           int64 `json:"voidAmount"`
				VoidTaxAmount        int64 `json:"voidTaxAmount"`
				VoidTransactionCount int64 `json:"voidTransactionCount"`
				VoidUnitCount        int64 `json:"voidUnitCount"`
			} `json:"items"`
		} `json:"salesSummaries"`
		SequenceNumber int64 `json:"sequenceNumber"`
		Timestamps     struct {
			RecordedTimestamp       string      `json:"recordedTimestamp"`
			SignatureEndTimestamp   interface{} `json:"signatureEndTimestamp"`
			SignatureStartTimestamp interface{} `json:"signatureStartTimestamp"`
			UpdatedTimestamp        string      `json:"updatedTimestamp"`
		} `json:"timestamps"`
		TypeCode string `json:"typeCode"`
	} `json:"lineItems"`
}
