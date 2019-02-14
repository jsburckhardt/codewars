package main

import (
	"fmt"
	"strings"
)

func main() {
	a := `<?xml version="1.0" encoding="UTF-8"?><soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:partner.soap.sforce.com" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><soapenv:Body><loginResponse><result><metadataServerUrl>https://na52.salesforce.com/services/Soap/m/40.0/00D90000000JJYK</metadataServerUrl><passwordExpired>false</passwordExpired><sandbox>false</sandbox><serverUrl>https://na52.salesforce.com/services/Soap/u/40.0/00D90000000JJYK</serverUrl><sessionId>00D90000000JJYK!ARoAQCEgmYrlZVwfMeE5DlCbEam5ViHw61.pfzJfCsP_EsikRWvXotzxMQRmTKJlpdSJPi.MXTI6JGJ2f4a9tFHRASX.Npzg</sessionId><userId>00590000000FoOrAAK</userId><userInfo><accessibilityMode>false</accessibilityMode><chatterExternal>false</chatterExternal><currencySymbol>$</currencySymbol><orgAttachmentFileSizeLimit>20971520</orgAttachmentFileSizeLimit><orgDefaultCurrencyIsoCode>AUD</orgDefaultCurrencyIsoCode><orgDefaultCurrencyLocale>en_AU</orgDefaultCurrencyLocale><orgDisallowHtmlAttachments>false</orgDisallowHtmlAttachments><orgHasPersonAccounts>true</orgHasPersonAccounts><organizationId>00D90000000JJYKEA4</organizationId><organizationMultiCurrency>false</organizationMultiCurrency><organizationName>Advantedge PVT</organizationName><profileId>00e90000000cGpdAAE</profileId><roleId>00E90000000WgTOEA0</roleId><sessionSecondsValid>14400</sessionSecondsValid><userDefaultCurrencyIsoCode xsi:nil="true"/><userEmail>podium.broker.crm.admin@podium.net.au</userEmail><userFullName>Admin User</userFullName><userId>00590000000FoOrAAK</userId><userLanguage>en_US</userLanguage><userLocale>en_AU</userLocale><userName>sadmin@advantedge.choice.pvt</userName><userTimeZone>Australia/Sydney</userTimeZone><userType>Standard</userType><userUiSkin>Theme3</userUiSkin></userInfo></result></loginResponse></soapenv:Body></soapenv:Envelope>`
	//out := strings.TrimLeft(strings.TrimRight(a, "</organizationName>"), "<organizationName>")
	out := GetStringInBetween(a, "<organizationId>", "</organizationId>")
	fmt.Println(out)
}

//GetStringInBetween does something
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	return str[s:e]
}
