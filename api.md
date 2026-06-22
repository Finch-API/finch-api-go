# Shared Params Types

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared#ConnectionStatusType">ConnectionStatusType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared#OperationSupport">OperationSupport</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared#ConnectionStatusType">ConnectionStatusType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared#OperationSupport">OperationSupport</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared#OperationSupportMatrix">OperationSupportMatrix</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/shared#Paging">Paging</a>

# finchgo

Methods:

- <code>client.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#FinchgoService.GetAuthURL">GetAuthURL</a>(products string, redirectUri string, sandbox bool, opts ...option.RequestOption) (string, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code>client.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#FinchgoService.WithAccessToken">WithAccessToken</a>(accessToken string) (Client, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccessTokens

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CreateAccessTokenResponse">CreateAccessTokenResponse</a>

Methods:

- <code title="post /auth/token">client.AccessTokens.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccessTokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccessTokenNewParams">AccessTokenNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CreateAccessTokenResponse">CreateAccessTokenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# HRIS

Params Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IncomeParam">IncomeParam</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#LocationParam">LocationParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Income">Income</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Money">Money</a>

## Company

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Company">Company</a>

Methods:

- <code title="get /employer/company">client.HRIS.Company.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISCompanyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISCompanyGetParams">HRISCompanyGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Company">Company</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PayStatementItem

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemListResponse">HRISPayStatementItemListResponse</a>

Methods:

- <code title="get /employer/pay-statement-item">client.HRIS.PayStatementItem.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemListParams">HRISPayStatementItemListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemListResponse">HRISPayStatementItemListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleNewResponse">HRISPayStatementItemRuleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleUpdateResponse">HRISPayStatementItemRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleListResponse">HRISPayStatementItemRuleListResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleDeleteResponse">HRISPayStatementItemRuleDeleteResponse</a>

Methods:

- <code title="post /employer/pay-statement-item/rule">client.HRIS.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleNewParams">HRISPayStatementItemRuleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleNewResponse">HRISPayStatementItemRuleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /employer/pay-statement-item/rule/{rule_id}">client.HRIS.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleUpdateParams">HRISPayStatementItemRuleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleUpdateResponse">HRISPayStatementItemRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/pay-statement-item/rule">client.HRIS.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleListParams">HRISPayStatementItemRuleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleListResponse">HRISPayStatementItemRuleListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /employer/pay-statement-item/rule/{rule_id}">client.HRIS.PayStatementItem.Rules.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleDeleteParams">HRISPayStatementItemRuleDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementItemRuleDeleteResponse">HRISPayStatementItemRuleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Directory

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualInDirectory">IndividualInDirectory</a>

Methods:

- <code title="get /employer/directory">client.HRIS.Directory.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDirectoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDirectoryListParams">HRISDirectoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualsPage">IndividualsPage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Individuals

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Individual">Individual</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualResponse">IndividualResponse</a>

Methods:

- <code title="post /employer/individual">client.HRIS.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISIndividualService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISIndividualGetManyParams">HRISIndividualGetManyParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualResponse">IndividualResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Employments

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#EmploymentData">EmploymentData</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#EmploymentDataResponse">EmploymentDataResponse</a>

Methods:

- <code title="post /employer/employment">client.HRIS.Employments.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISEmploymentService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISEmploymentGetManyParams">HRISEmploymentGetManyParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#EmploymentDataResponse">EmploymentDataResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Payment">Payment</a>

Methods:

- <code title="get /employer/payment">client.HRIS.Payments.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPaymentListParams">HRISPaymentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Payment">Payment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PayStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayStatement">PayStatement</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayStatementData">PayStatementData</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayStatementDataSyncInProgress">PayStatementDataSyncInProgress</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayStatementResponse">PayStatementResponse</a>

Methods:

- <code title="post /employer/pay-statement">client.HRIS.PayStatements.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISPayStatementGetManyParams">HRISPayStatementGetManyParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayStatementResponse">PayStatementResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Documents

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#DocumentResponse">DocumentResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#W42005">W42005</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#W42020">W42020</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentListResponse">HRISDocumentListResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentRetreiveResponse">HRISDocumentRetreiveResponse</a>

Methods:

- <code title="get /employer/documents">client.HRIS.Documents.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentListParams">HRISDocumentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentListResponse">HRISDocumentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/documents/{document_id}">client.HRIS.Documents.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentService.Retreive">Retreive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, documentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentRetreiveParams">HRISDocumentRetreiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISDocumentRetreiveResponse">HRISDocumentRetreiveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Benefits

Params Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BenefitFrequency">BenefitFrequency</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BenefitType">BenefitType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BenefitFeaturesAndOperations">BenefitFeaturesAndOperations</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BenefitFrequency">BenefitFrequency</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BenefitType">BenefitType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BenefitsSupport">BenefitsSupport</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CompanyBenefit">CompanyBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CreateCompanyBenefitsResponse">CreateCompanyBenefitsResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#RegisterCompanyBenefitResponse">RegisterCompanyBenefitResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SupportPerBenefitType">SupportPerBenefitType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SupportedBenefit">SupportedBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#UpdateCompanyBenefitResponse">UpdateCompanyBenefitResponse</a>

Methods:

- <code title="post /employer/benefits">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitNewParams">HRISBenefitNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CreateCompanyBenefitsResponse">CreateCompanyBenefitsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitGetParams">HRISBenefitGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CompanyBenefit">CompanyBenefit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /employer/benefits/{benefit_id}">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitUpdateParams">HRISBenefitUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#UpdateCompanyBenefitResponse">UpdateCompanyBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitListParams">HRISBenefitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CompanyBenefit">CompanyBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/meta">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitService.ListSupportedBenefits">ListSupportedBenefits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitListSupportedBenefitsParams">HRISBenefitListSupportedBenefitsParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SupportedBenefit">SupportedBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /employer/benefits/register">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitRegisterParams">HRISBenefitRegisterParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#RegisterCompanyBenefitResponse">RegisterCompanyBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Individuals

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#EnrolledIndividualBenefitResponse">EnrolledIndividualBenefitResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualBenefit">IndividualBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#UnenrolledIndividualBenefitResponse">UnenrolledIndividualBenefitResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualEnrolledIDsResponse">HRISBenefitIndividualEnrolledIDsResponse</a>

Methods:

- <code title="post /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualService.EnrollMany">EnrollMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualEnrollManyParams">HRISBenefitIndividualEnrollManyParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#EnrolledIndividualBenefitResponse">EnrolledIndividualBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}/enrolled">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualService.EnrolledIDs">EnrolledIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualEnrolledIDsParams">HRISBenefitIndividualEnrolledIDsParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualEnrolledIDsResponse">HRISBenefitIndividualEnrolledIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualService.GetManyBenefits">GetManyBenefits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualGetManyBenefitsParams">HRISBenefitIndividualGetManyBenefitsParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualBenefit">IndividualBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualService.UnenrollMany">UnenrollMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#HRISBenefitIndividualUnenrollManyParams">HRISBenefitIndividualUnenrollManyParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#UnenrolledIndividualBenefitResponse">UnenrolledIndividualBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ProviderListResponse">ProviderListResponse</a>

Methods:

- <code title="get /providers">client.Providers.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ProviderListResponse">ProviderListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#DisconnectEntityResponse">DisconnectEntityResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#DisconnectResponse">DisconnectResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Introspection">Introspection</a>

Methods:

- <code title="post /disconnect">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccountService.Disconnect">Disconnect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#DisconnectResponse">DisconnectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /disconnect-entity">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccountService.DisconnectEntity">DisconnectEntity</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccountDisconnectEntityParams">AccountDisconnectEntityParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#DisconnectEntityResponse">DisconnectEntityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /introspect">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccountService.Introspect">Introspect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#Introspection">Introspection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AccountUpdateEvent">AccountUpdateEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#BaseWebhookEvent">BaseWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#CompanyEvent">CompanyEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#DirectoryEvent">DirectoryEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#EmploymentEvent">EmploymentEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#IndividualEvent">IndividualEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobCompletionEvent">JobCompletionEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayStatementEvent">PayStatementEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PaymentEvent">PaymentEvent</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#WebhookEventUnion">WebhookEventUnion</a>

Methods:

- <code>client.Webhooks.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#WebhookService.Unwrap">Unwrap</a>(payload []byte, headers http.Header, secret string, now time.Time) (WebhookEvent, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code>client.Webhooks.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#WebhookService.VerifySignature">VerifySignature</a>(payload []byte, headers http.Header, secret string, now time.Time) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# RequestForwarding

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#RequestForwardingForwardResponse">RequestForwardingForwardResponse</a>

Methods:

- <code title="post /forward">client.RequestForwarding.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#RequestForwardingService.Forward">Forward</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#RequestForwardingForwardParams">RequestForwardingForwardParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#RequestForwardingForwardResponse">RequestForwardingForwardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

## Automated

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AutomatedAsyncJob">AutomatedAsyncJob</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedNewResponse">JobAutomatedNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedListResponse">JobAutomatedListResponse</a>

Methods:

- <code title="post /jobs/automated">client.Jobs.Automated.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedNewParams">JobAutomatedNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedNewResponse">JobAutomatedNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/automated/{job_id}">client.Jobs.Automated.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#AutomatedAsyncJob">AutomatedAsyncJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/automated">client.Jobs.Automated.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedListParams">JobAutomatedListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobAutomatedListResponse">JobAutomatedListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Manual

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ManualAsyncJob">ManualAsyncJob</a>

Methods:

- <code title="get /jobs/manual/{job_id}">client.Jobs.Manual.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#JobManualService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ManualAsyncJob">ManualAsyncJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sandbox

## Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionNewResponse">SandboxConnectionNewResponse</a>

Methods:

- <code title="post /sandbox/connections">client.Sandbox.Connections.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionNewParams">SandboxConnectionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionNewResponse">SandboxConnectionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountNewResponse">SandboxConnectionAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountUpdateResponse">SandboxConnectionAccountUpdateResponse</a>

Methods:

- <code title="post /sandbox/connections/accounts">client.Sandbox.Connections.Accounts.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountNewParams">SandboxConnectionAccountNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountNewResponse">SandboxConnectionAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandbox/connections/accounts">client.Sandbox.Connections.Accounts.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountUpdateParams">SandboxConnectionAccountUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxConnectionAccountUpdateResponse">SandboxConnectionAccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Company

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxCompanyUpdateResponse">SandboxCompanyUpdateResponse</a>

Methods:

- <code title="put /sandbox/company">client.Sandbox.Company.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxCompanyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxCompanyUpdateParams">SandboxCompanyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxCompanyUpdateResponse">SandboxCompanyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Directory

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxDirectoryNewResponse">SandboxDirectoryNewResponse</a>

Methods:

- <code title="post /sandbox/directory">client.Sandbox.Directory.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxDirectoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxDirectoryNewParams">SandboxDirectoryNewParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxDirectoryNewResponse">SandboxDirectoryNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Individual

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxIndividualUpdateResponse">SandboxIndividualUpdateResponse</a>

Methods:

- <code title="put /sandbox/individual/{individual_id}">client.Sandbox.Individual.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxIndividualService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, individualID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxIndividualUpdateParams">SandboxIndividualUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxIndividualUpdateResponse">SandboxIndividualUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Employment

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxEmploymentUpdateResponse">SandboxEmploymentUpdateResponse</a>

Methods:

- <code title="put /sandbox/employment/{individual_id}">client.Sandbox.Employment.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxEmploymentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, individualID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxEmploymentUpdateParams">SandboxEmploymentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxEmploymentUpdateResponse">SandboxEmploymentUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payment

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxPaymentNewResponse">SandboxPaymentNewResponse</a>

Methods:

- <code title="post /sandbox/payment">client.Sandbox.Payment.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxPaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxPaymentNewParams">SandboxPaymentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxPaymentNewResponse">SandboxPaymentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobNewResponse">SandboxJobNewResponse</a>

Methods:

- <code title="post /sandbox/jobs">client.Sandbox.Jobs.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobNewParams">SandboxJobNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobNewResponse">SandboxJobNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Configuration

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobConfiguration">SandboxJobConfiguration</a>

Methods:

- <code title="get /sandbox/jobs/configuration">client.Sandbox.Jobs.Configuration.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobConfigurationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobConfiguration">SandboxJobConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandbox/jobs/configuration">client.Sandbox.Jobs.Configuration.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobConfigurationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobConfigurationUpdateParams">SandboxJobConfigurationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#SandboxJobConfiguration">SandboxJobConfiguration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payroll

## PayGroups

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupGetResponse">PayrollPayGroupGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupListResponse">PayrollPayGroupListResponse</a>

Methods:

- <code title="get /employer/pay-groups/{pay_group_id}">client.Payroll.PayGroups.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, payGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupGetParams">PayrollPayGroupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupGetResponse">PayrollPayGroupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/pay-groups">client.Payroll.PayGroups.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupListParams">PayrollPayGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#PayrollPayGroupListResponse">PayrollPayGroupListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Connect

## Sessions

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionNewResponse">ConnectSessionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionReauthenticateResponse">ConnectSessionReauthenticateResponse</a>

Methods:

- <code title="post /connect/sessions">client.Connect.Sessions.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionNewParams">ConnectSessionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionNewResponse">ConnectSessionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /connect/sessions/reauthenticate">client.Connect.Sessions.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionService.Reauthenticate">Reauthenticate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionReauthenticateParams">ConnectSessionReauthenticateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go/v2#ConnectSessionReauthenticateResponse">ConnectSessionReauthenticateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
