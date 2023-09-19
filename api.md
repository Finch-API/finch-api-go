# Top Level

# HRIS

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Income">Income</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Money">Money</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Paging">Paging</a>

## Company

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Company">Company</a>

Methods:

- <code title="get /employer/company">client.HRIS.Company.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISCompanyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Company">Company</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Directory

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualInDirectory">IndividualInDirectory</a>

Methods:

- <code title="get /employer/directory">client.HRIS.Directory.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDirectoryService.ListIndividuals">ListIndividuals</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISDirectoryListIndividualsParams">HRISDirectoryListIndividualsParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualsPage">IndividualsPage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Individuals

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Individual">Individual</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualResponse">IndividualResponse</a>

Methods:

- <code title="post /employer/individual">client.HRIS.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISIndividualService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISIndividualGetManyParams">HRISIndividualGetManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualResponse">IndividualResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Employments

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentData">EmploymentData</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentDataResponse">EmploymentDataResponse</a>

Methods:

- <code title="post /employer/employment">client.HRIS.Employments.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISEmploymentService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISEmploymentGetManyParams">HRISEmploymentGetManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#EmploymentDataResponse">EmploymentDataResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Payment">Payment</a>

Methods:

- <code title="get /employer/payment">client.HRIS.Payments.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPaymentListParams">HRISPaymentListParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Payment">Payment</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PayStatements

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatement">PayStatement</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementResponse">PayStatementResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementResponseBody">PayStatementResponseBody</a>

Methods:

- <code title="post /employer/pay-statement">client.HRIS.PayStatements.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPayStatementService.GetMany">GetMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISPayStatementGetManyParams">HRISPayStatementGetManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ResponsesPage">ResponsesPage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#PayStatementResponse">PayStatementResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Benefits

Params Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitFrequency">BenefitFrequency</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitType">BenefitType</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitFrequency">BenefitFrequency</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenefitType">BenefitType</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#BenfitContribution">BenfitContribution</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyBenefit">CompanyBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CreateCompanyBenefitsResponse">CreateCompanyBenefitsResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SupportedBenefit">SupportedBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UpdateCompanyBenefitResponse">UpdateCompanyBenefitResponse</a>

Methods:

- <code title="post /employer/benefits">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitNewParams">HRISBenefitNewParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CreateCompanyBenefitsResponse">CreateCompanyBenefitsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyBenefit">CompanyBenefit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /employer/benefits/{benefit_id}">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitUpdateParams">HRISBenefitUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UpdateCompanyBenefitResponse">UpdateCompanyBenefitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#CompanyBenefit">CompanyBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/meta">client.HRIS.Benefits.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitService.ListSupportedBenefits">ListSupportedBenefits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SupportedBenefit">SupportedBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Individuals

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualBenefit">IndividualBenefit</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UnenrolledIndividual">UnenrolledIndividual</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualEnrolledIDsResponse">HRISBenefitIndividualEnrolledIDsResponse</a>

Methods:

- <code title="get /employer/benefits/{benefit_id}/enrolled">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualService.EnrolledIDs">EnrolledIDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualEnrolledIDsResponse">HRISBenefitIndividualEnrolledIDsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualService.GetManyBenefits">GetManyBenefits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualGetManyBenefitsParams">HRISBenefitIndividualGetManyBenefitsParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#IndividualBenefit">IndividualBenefit</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /employer/benefits/{benefit_id}/individuals">client.HRIS.Benefits.Individuals.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualService.UnenrollMany">UnenrollMany</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benefitID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#HRISBenefitIndividualUnenrollManyParams">HRISBenefitIndividualUnenrollManyParams</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#UnenrolledIndividual">UnenrolledIndividual</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Provider">Provider</a>

Methods:

- <code title="get /providers">client.Providers.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Provider">Provider</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Account

Response Types:

- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#DisconnectResponse">DisconnectResponse</a>
- <a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Introspection">Introspection</a>

Methods:

- <code title="post /disconnect">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccountService.Disconnect">Disconnect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#DisconnectResponse">DisconnectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /introspect">client.Account.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#AccountService.Introspect">Introspect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go">finchgo</a>.<a href="https://pkg.go.dev/github.com/Finch-API/finch-api-go#Introspection">Introspection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Custom Methods:

- VerifySignature
