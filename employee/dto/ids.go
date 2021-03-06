package dto

type MonthsColumns struct {
	Month,
	WorkDays,
	Year,
	ID,
	PreviousMonthLink,
	PreviousMonth,
	IsCurrent string
}

type InvoicesColumns struct {
	Id                   string
	InvoiceNo            string
	Month                string
	Employee             string
	PreviousInvoice      string
	RequestedSubtotalRub string
	EurRubExpected       string
	RequestedSubtotalEur string
	RoundingPrevMonEur   string
	Rounding             string
	AmountRequestedEur   string
	Hours                string
	EurRubActual         string
	AmountRubActual      string
	RateErrorRub         string
	CostOfDay            string
	BankTotalFees        string
	OpeningDateIp        string
	CorrectionRefs       string
	CorrectionRub        string
	SalaryRub            string
	PatentRub            string
	TaxesRub             string
	PatentRefs           string
	TaxesRefs            string
	BaseSalaryRub        string
	BaseSalaryEur        string
	BankFees             string
	RateErrorPrevMon     string
	PaymentChecksPassed  string
}

type CorrectionsColumns struct {
	Comment                  string
	AbsoluteCorrectionEur    string
	AbsoluteCorrectionRub    string
	EurRubExpected           string
	AbsCorrectionEurInRub    string
	PerDayType               string
	NumberOfDays             string
	CostOfDay                string
	PerDay                   string
	TotalCorrectionRub       string
	PaymentInvoice           string
	PerDayCoefficient        string
	PerDayCalculationInvoice string
	Display                  string
	Category                 string
}

type TaxesColumns struct {
	Invoice            string
	OpeningDateIp      string
	PeriodStart        string
	PeriodEnd          string
	AmountIpDays       string
	MedicalFund        string
	PensionFundFixed   string
	PensionFundPercent string
	Amount             string
}

type PatentColumns struct {
	Invoice       string
	OpeningPatent string
	PeriodEnd     string
	FullMonths    string
	AnnualCost    string
	PeriodCost    string
	Period        string
}

type EmployeesColumns struct {
	Location                string
	Name                    string
	StartDate               string
	ProbationEnd            string
	AnnualLeaveFrom         string
	SalaryBeforeIp          string
	SalaryAfterIp           string
	NetSalaryAfterProbation string
	EndDate                 string
	HourRate                string
	BankCurrencyControl     string
	BankService             string
	BankTotalFees           string
	Address                 string
	OpeningDateIp           string
	StartMonth              string
	MattermostLogin         string
	RussianFullName         string
	Position                string
	INN                     string
	IsWorkingNow            string
	AllSalaries             string
	PatentDocuments         string
	PatentFee               string
	LastSalary              string
	EnglishFullName         string
	BankRequisites          string
	BillTo                  string
	GeneralSdLink           string
	PersonnelSdLink         string
	FinanceSdLink           string
	AclSdLink               string
	ContractDate            string
	ContractNumber          string
}

type IdOnly struct {
	Id string
}

type InvoicesDecl struct {
	IdOnly
	Cols InvoicesColumns
}

type MonthsDecl struct {
	IdOnly
	Cols MonthsColumns
}

type CorrectionsDecl struct {
	IdOnly
	Cols CorrectionsColumns
}

type TaxesDecl struct {
	IdOnly
	Cols TaxesColumns
}

type PatentDecl struct {
	IdOnly
	Cols PatentColumns
}

type EmployeesDecl struct {
	IdOnly
	Cols EmployeesColumns
}

type CodaFormulasDecl struct {
	CurrentMonth string
}

type CodaViewsDecl struct {
	ThisMonthInvoices    string // table-_k4m2krdyp
	ThisMonthCorrections string // table-29Y7W0xzVA
	ThisMonthTaxes       string // table-X3zgg4jJ8-
	ThisMonthPatents     string // table-n2aIqJxO80
}

type IdStruct struct {
	Invoices     InvoicesDecl
	Corrections  CorrectionsDecl
	Taxes        TaxesDecl
	Patent       PatentDecl
	Months       MonthsDecl
	Employees    EmployeesDecl
	CodaFormulas CodaFormulasDecl
}

var Ids IdStruct

func init() {
	Ids = IdStruct{
		Invoices: InvoicesDecl{
			IdOnly{"grid-Wdy6Agpxou"},
			InvoicesColumns{
				Id:                   "c-bZ_nLfZufG",
				InvoiceNo:            "c-eJ2e_cRCaM",
				Month:                "c-wR0IONcxGH",
				Employee:             "c-bbHUhqlbfN",
				PreviousInvoice:      "c-FQ7rKmbXr6",
				RequestedSubtotalRub: "c-WLMdyl4WqI",
				EurRubExpected:       "c-tvtGu9juVL",
				RequestedSubtotalEur: "c-9rnJJZ6gA7",
				RoundingPrevMonEur:   "c-hLrmDsk89g",
				Rounding:             "c-Tri-EGUP_n",
				AmountRequestedEur:   "c-bJpHVxywXD",
				Hours:                "c-KtVV9if8P7",
				EurRubActual:         "c-kLIyv9EvyH",
				AmountRubActual:      "c-AxLSgrt7e3",
				RateErrorRub:         "c-SsHRhKa_uC",
				CostOfDay:            "c-yJnq9stsgi",
				BankTotalFees:        "c-IYFe5lD4td",
				OpeningDateIp:        "c-hI1iZG3xzY",
				CorrectionRefs:       "c-tpeCMU21_I",
				CorrectionRub:        "c-jNcl4nZe_h",
				PatentRub:            "c-qA_pPM9kuZ",
				TaxesRub:             "c-ug709va8_K",
				PatentRefs:           "c-bkVlencAZt",
				TaxesRefs:            "c-aYkpi97eXt",
				BaseSalaryRub:        "c-wqNhZf9EQY",
				BaseSalaryEur:        "c-B0nU6ZI_Z4",
				BankFees:             "c-sRGR6jYC7g",
				RateErrorPrevMon:     "c-_9tuuG4RIN",
				PaymentChecksPassed:  "c-DRPGK3XTmD",
			},
		},
		Corrections: CorrectionsDecl{
			IdOnly{"grid-wBmvgFgaGi"},
			CorrectionsColumns{
				Comment:                  "c--_r48PQnSn",
				AbsoluteCorrectionEur:    "c-pRmEece9pf",
				AbsoluteCorrectionRub:    "c-P2x5IJuMXN",
				EurRubExpected:           "c-8LD34cnmCh",
				AbsCorrectionEurInRub:    "c-_GvG9w3Qs7",
				PerDayType:               "c-3Ivn-M1j7-",
				NumberOfDays:             "c-gDOyigH1cm",
				CostOfDay:                "c-K_Iy0iERKR",
				PerDay:                   "c-Y2E1Vwe2_-",
				TotalCorrectionRub:       "c-0arkfr4qXv",
				PaymentInvoice:           "c-7SU0iOBY9J",
				PerDayCoefficient:        "c-pz6W2IRzFR",
				PerDayCalculationInvoice: "c-bK4qXZUCqs",
				Display:                  "c-FVW_9PPzZ2",
				Category:                 "c-zDY58PF0P6",
			},
		},
		Taxes: TaxesDecl{
			IdOnly{"grid-057HtXfvYH"},
			TaxesColumns{
				Invoice:            "c-sYkd0N-9Ef",
				OpeningDateIp:      "c-Q83cwaB-vf",
				PeriodStart:        "c-UgsMCGn1oX",
				PeriodEnd:          "c-7Bt0dg_odm",
				AmountIpDays:       "c-wraB-EUIPu",
				MedicalFund:        "c-hRpESBxMnx",
				PensionFundFixed:   "c-5Gm9BIf7sa",
				PensionFundPercent: "c-nO-EnVUZSb",
				Amount:             "c-FlQQoKxoau",
			},
		},
		Patent: PatentDecl{
			IdOnly{"grid-_IJllxLQCt"},
			PatentColumns{
				Invoice:       "c-sYkd0N-9Ef",
				OpeningPatent: "c-GfXRAu21-_",
				PeriodEnd:     "c-7Bt0dg_odm",
				FullMonths:    "c-gARgn0s4h-",
				AnnualCost:    "c-ayAtt8TYPI",
				PeriodCost:    "c-FlQQoKxoau",
				Period:        "c-gtV-Qz9osQ",
			},
		},
		Months: MonthsDecl{
			IdOnly: IdOnly{"grid-laH8qsdDyP"},
			Cols: MonthsColumns{
				Month:             "c-u_Pgrgevw7",
				WorkDays:          "c-2eD8ott5yA",
				Year:              "c-NgPwXZshJM",
				ID:                "c-iRCjZ0JBcM",
				PreviousMonthLink: "c-3Cc_lYdvmW",
				PreviousMonth:     "c-vuW159vf-o",
				IsCurrent:         "c-OLxcAJLQoW",
			},
		},
		Employees: EmployeesDecl{
			IdOnly: IdOnly{"grid-TGESBHJkVA"},
			Cols: EmployeesColumns{
				Location:                "c-WcmDQXPChx",
				Name:                    "c-tCDt6yt4Ix",
				StartDate:               "c-Zs7oQbj-_J",
				ProbationEnd:            "c-35cUDxDnAo",
				AnnualLeaveFrom:         "c-uzP6UrJs3a",
				SalaryBeforeIp:          "c-MZLKSGSxFn",
				SalaryAfterIp:           "c-5asgm6tqZs",
				NetSalaryAfterProbation: "c-fvHtrLntTN",
				EndDate:                 "c-7OoHuXqt8n",
				HourRate:                "c-SItGvyE3ie",
				BankCurrencyControl:     "c-HbuzOpvaPf",
				BankService:             "c-214wDdnGtE",
				BankTotalFees:           "c-156kDMsJzb",
				Address:                 "c-lEdEHoijqv",
				OpeningDateIp:           "c-uU3-6piESs",
				StartMonth:              "c-pnvQkCAJ7U",
				MattermostLogin:         "c-BmjWOCSXHd",
				RussianFullName:         "c-7N8qU2h-Zv",
				Position:                "c-4nDWkuySVp",
				INN:                     "c-Gsx7-a_kGU",
				IsWorkingNow:            "c-QfV5QzjuJP",
				AllSalaries:             "c-P0cTAu016r",
				PatentDocuments:         "c-FKDG1g17Su",
				PatentFee:               "c-kwYGeFl40p",
				LastSalary:              "c-Is6fozVp1a",
				EnglishFullName:         "c-TAlfzDcFzQ",
				BankRequisites:          "c-OlCoWd7n4S",
				BillTo:                  "c-XnSDzWgAgQ",
				GeneralSdLink:           "c-EMbqnhOkkr",
				PersonnelSdLink:         "c-6L5oS7LZ4a",
				FinanceSdLink:           "c-NeQmHu-raB",
				AclSdLink:               "c-aKn3TiCF_A",
				ContractDate:            "c-fDIRo1JHHX",
				ContractNumber:          "c-XX2OkCkOSR",
			},
		},
		CodaFormulas: CodaFormulasDecl{
			CurrentMonth: "f-rnJn4-MytN",
		},
	}
}
