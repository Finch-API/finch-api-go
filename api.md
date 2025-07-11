# Shared Params Types

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared#ConnectionStatusType">ConnectionStatusType</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared#ConnectionStatusType">ConnectionStatusType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared#OperationSupport">OperationSupport</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared#OperationSupportMatrix">OperationSupportMatrix</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/shared#Paging">Paging</a>

# finchgo

Methods:

- <code>client.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#FinchgoService.GetAuthURL">GetAuthURL</a>(products string, redirectUri string, sandbox bool, opts ...option.RequestOption) (string, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code>client.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#FinchgoService.WithAccessToken">WithAccessToken</a>(accessToken string) (Client, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccessTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CreateAccessTokenResponse">CreateAccessTokenResponse</a>

Methods:

- <code title="post /auth/token">client.AccessTokens.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccessTokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccessTokenNewParams">AccessTokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CreateAccessTokenResponse">CreateAccessTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# HRIS

Params Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IncomeParam">IncomeParam</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#LocationParam">LocationParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Income">Income</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Money">Money</a>

## Company

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Company">Company</a>

Methods:

- <code title="get /employer/company">client.HRIS.Company.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Company">Company</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### PayStatementItem

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemListResponse">HRISCompanyPayStatementItemListResponse</a>

Methods:

- <code title="get /employer/pay-statement-item">client.HRIS.Company.PayStatementItem.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemListParams">HRISCompanyPayStatementItemListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemListResponse">HRISCompanyPayStatementItemListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleNewResponse">HRISCompanyPayStatementItemRuleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleUpdateResponse">HRISCompanyPayStatementItemRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleListResponse">HRISCompanyPayStatementItemRuleListResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleDeleteResponse">HRISCompanyPayStatementItemRuleDeleteResponse</a>

Methods:

- <code title="post /employer/pay-statement-item/rule">client.HRIS.Company.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleNewParams">HRISCompanyPayStatementItemRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleNewResponse">HRISCompanyPayStatementItemRuleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /employer/pay-statement-item/rule/{rule_id}">client.HRIS.Company.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleUpdateParams">HRISCompanyPayStatementItemRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleUpdateResponse">HRISCompanyPayStatementItemRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/pay-statement-item/rule">client.HRIS.Company.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleListResponse">HRISCompanyPayStatementItemRuleListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /employer/pay-statement-item/rule/{rule_id}">client.HRIS.Company.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyPayStatementItemRuleDeleteResponse">HRISCompanyPayStatementItemRuleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Directory

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualInDirectory">IndividualInDirectory</a>

Methods:

- <code title="get /employer/directory">client.HRIS.Directory.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDirectoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDirectoryListParams">HRISDirectoryListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualsPage">IndividualsPage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Individuals

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Individual">Individual</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualResponse">IndividualResponse</a>

Methods:

- <code title="post /employer/individual">client.HRIS.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISIndividualService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISIndividualGetManyParams">HRISIndividualGetManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualResponse">IndividualResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Employments

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentData">EmploymentData</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentDataResponse">EmploymentDataResponse</a>

Methods:

- <code title="post /employer/employment">client.HRIS.Employments.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISEmploymentService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISEmploymentGetManyParams">HRISEmploymentGetManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentDataResponse">EmploymentDataResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Payment">Payment</a>

Methods:

- <code title="get /employer/payment">client.HRIS.Payments.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPaymentListParams">HRISPaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Payment">Payment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PayStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatement">PayStatement</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementDataSyncInProgress">PayStatementDataSyncInProgress</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementResponse">PayStatementResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementResponseBody">PayStatementResponseBody</a>

Methods:

- <code title="post /employer/pay-statement">client.HRIS.PayStatements.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPayStatementService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPayStatementGetManyParams">HRISPayStatementGetManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementResponse">PayStatementResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Documents

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#DocumentResponse">DocumentResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#W42005">W42005</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#W42020">W42020</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentListResponse">HRISDocumentListResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentRetreiveResponse">HRISDocumentRetreiveResponse</a>

Methods:

- <code title="get /employer/documents">client.HRIS.Documents.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentListParams">HRISDocumentListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentListResponse">HRISDocumentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/documents/{document_id}">client.HRIS.Documents.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentService.Retreive">Retreive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDocumentRetreiveResponse">HRISDocumentRetreiveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Benefits

Params Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitFrequency">BenefitFrequency</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitType">BenefitType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitContribution">BenefitContribution</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitFeaturesAndOperations">BenefitFeaturesAndOperations</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitFrequency">BenefitFrequency</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitType">BenefitType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitsSupport">BenefitsSupport</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyBenefit">CompanyBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CreateCompanyBenefitsResponse">CreateCompanyBenefitsResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SupportPerBenefitType">SupportPerBenefitType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SupportedBenefit">SupportedBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UpdateCompanyBenefitResponse">UpdateCompanyBenefitResponse</a>

Methods:

- <code title="post /employer/benefits">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitNewParams">HRISBenefitNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CreateCompanyBenefitsResponse">CreateCompanyBenefitsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyBenefit">CompanyBenefit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /employer/benefits/{benefit_id}">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitUpdateParams">HRISBenefitUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UpdateCompanyBenefitResponse">UpdateCompanyBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyBenefit">CompanyBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/meta">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.ListSupportedBenefits">ListSupportedBenefits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SupportedBenefit">SupportedBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Individuals

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualBenefit">IndividualBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UnenrolledIndividualBenefitResponse">UnenrolledIndividualBenefitResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualEnrolledIDsResponse">HRISBenefitIndividualEnrolledIDsResponse</a>

Methods:

- <code title="get /employer/benefits/{benefit_id}/enrolled">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualService.EnrolledIDs">EnrolledIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualEnrolledIDsResponse">HRISBenefitIndividualEnrolledIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualService.GetManyBenefits">GetManyBenefits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualGetManyBenefitsParams">HRISBenefitIndividualGetManyBenefitsParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualBenefit">IndividualBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualService.UnenrollMany">UnenrollMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualUnenrollManyParams">HRISBenefitIndividualUnenrollManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UnenrolledIndividualBenefitResponse">UnenrolledIndividualBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Provider">Provider</a>

Methods:

- <code title="get /providers">client.Providers.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Provider">Provider</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#DisconnectResponse">DisconnectResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Introspection">Introspection</a>

Methods:

- <code title="post /disconnect">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccountService.Disconnect">Disconnect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#DisconnectResponse">DisconnectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /introspect">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccountService.Introspect">Introspect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Introspection">Introspection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccountUpdateEvent">AccountUpdateEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BaseWebhookEvent">BaseWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyEvent">CompanyEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#DirectoryEvent">DirectoryEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentEvent">EmploymentEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualEvent">IndividualEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobCompletionEvent">JobCompletionEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementEvent">PayStatementEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PaymentEvent">PaymentEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#WebhookEventUnion">WebhookEventUnion</a>

Methods:

- <code>client.Webhooks.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#WebhookService.Unwrap">Unwrap</a>(payload []byte, headers http.Header, secret string, now time.Time) (WebhookEvent, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code>client.Webhooks.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#WebhookService.VerifySignature">VerifySignature</a>(payload []byte, headers http.Header, secret string, now time.Time) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# RequestForwarding

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#RequestForwardingForwardResponse">RequestForwardingForwardResponse</a>

Methods:

- <code title="post /forward">client.RequestForwarding.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#RequestForwardingService.Forward">Forward</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#RequestForwardingForwardParams">RequestForwardingForwardParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#RequestForwardingForwardResponse">RequestForwardingForwardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

## Automated

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AutomatedAsyncJob">AutomatedAsyncJob</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedNewResponse">JobAutomatedNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedListResponse">JobAutomatedListResponse</a>

Methods:

- <code title="post /jobs/automated">client.Jobs.Automated.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedNewParams">JobAutomatedNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedNewResponse">JobAutomatedNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/automated/{job_id}">client.Jobs.Automated.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AutomatedAsyncJob">AutomatedAsyncJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/automated">client.Jobs.Automated.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedListParams">JobAutomatedListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobAutomatedListResponse">JobAutomatedListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Manual

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ManualAsyncJob">ManualAsyncJob</a>

Methods:

- <code title="get /jobs/manual/{job_id}">client.Jobs.Manual.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#JobManualService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ManualAsyncJob">ManualAsyncJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sandbox

## Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionNewResponse">SandboxConnectionNewResponse</a>

Methods:

- <code title="post /sandbox/connections">client.Sandbox.Connections.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionNewParams">SandboxConnectionNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionNewResponse">SandboxConnectionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountNewResponse">SandboxConnectionAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountUpdateResponse">SandboxConnectionAccountUpdateResponse</a>

Methods:

- <code title="post /sandbox/connections/accounts">client.Sandbox.Connections.Accounts.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountNewParams">SandboxConnectionAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountNewResponse">SandboxConnectionAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandbox/connections/accounts">client.Sandbox.Connections.Accounts.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountUpdateParams">SandboxConnectionAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxConnectionAccountUpdateResponse">SandboxConnectionAccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Company

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxCompanyUpdateResponse">SandboxCompanyUpdateResponse</a>

Methods:

- <code title="put /sandbox/company">client.Sandbox.Company.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxCompanyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxCompanyUpdateParams">SandboxCompanyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxCompanyUpdateResponse">SandboxCompanyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Directory

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxDirectoryNewResponse">SandboxDirectoryNewResponse</a>

Methods:

- <code title="post /sandbox/directory">client.Sandbox.Directory.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxDirectoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxDirectoryNewParams">SandboxDirectoryNewParams</a>) ([]<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxDirectoryNewResponse">SandboxDirectoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Individual

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxIndividualUpdateResponse">SandboxIndividualUpdateResponse</a>

Methods:

- <code title="put /sandbox/individual/{individual_id}">client.Sandbox.Individual.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxIndividualService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, individualID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxIndividualUpdateParams">SandboxIndividualUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxIndividualUpdateResponse">SandboxIndividualUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Employment

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxEmploymentUpdateResponse">SandboxEmploymentUpdateResponse</a>

Methods:

- <code title="put /sandbox/employment/{individual_id}">client.Sandbox.Employment.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxEmploymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, individualID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxEmploymentUpdateParams">SandboxEmploymentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxEmploymentUpdateResponse">SandboxEmploymentUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payment

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxPaymentNewResponse">SandboxPaymentNewResponse</a>

Methods:

- <code title="post /sandbox/payment">client.Sandbox.Payment.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxPaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxPaymentNewParams">SandboxPaymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxPaymentNewResponse">SandboxPaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobNewResponse">SandboxJobNewResponse</a>

Methods:

- <code title="post /sandbox/jobs">client.Sandbox.Jobs.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobNewParams">SandboxJobNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobNewResponse">SandboxJobNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Configuration

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobConfiguration">SandboxJobConfiguration</a>

Methods:

- <code title="get /sandbox/jobs/configuration">client.Sandbox.Jobs.Configuration.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobConfigurationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobConfiguration">SandboxJobConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandbox/jobs/configuration">client.Sandbox.Jobs.Configuration.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobConfigurationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobConfigurationUpdateParams">SandboxJobConfigurationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SandboxJobConfiguration">SandboxJobConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payroll

## PayGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupGetResponse">PayrollPayGroupGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupListResponse">PayrollPayGroupListResponse</a>

Methods:

- <code title="get /employer/pay-groups/{pay_group_id}">client.Payroll.PayGroups.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, payGroupID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupGetResponse">PayrollPayGroupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/pay-groups">client.Payroll.PayGroups.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupListParams">PayrollPayGroupListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayrollPayGroupListResponse">PayrollPayGroupListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Connect

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionNewResponse">ConnectSessionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionReauthenticateResponse">ConnectSessionReauthenticateResponse</a>

Methods:

- <code title="post /connect/sessions">client.Connect.Sessions.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionNewParams">ConnectSessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionNewResponse">ConnectSessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /connect/sessions/reauthenticate">client.Connect.Sessions.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionService.Reauthenticate">Reauthenticate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionReauthenticateParams">ConnectSessionReauthenticateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ConnectSessionReauthenticateResponse">ConnectSessionReauthenticateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
