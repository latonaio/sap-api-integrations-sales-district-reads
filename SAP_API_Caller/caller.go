package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-sales-district-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesDistrict(salesDistrict, language, salesDistrictName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "SalesDistrict":
			func() {
				c.SalesDistrict(salesDistrict)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(language, salesDistrictName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) SalesDistrict(salesDistrict string) {
	salesDistrictData, err := c.callSalesDistrictSrvAPIRequirementSalesDistrict("A_SalesDistrict", salesDistrict)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(salesDistrictData)

	textData, err := c.callToText(salesDistrictData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callSalesDistrictSrvAPIRequirementSalesDistrict(api, salesDistrict string) ([]sap_api_output_formatter.SalesDistrict, error) {
	url := strings.Join([]string{c.baseURL, "API_SALESDISTRICT_SRV", api}, "/")
	param := c.getQueryWithSalesDistrict(map[string]string{}, salesDistrict)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSalesDistrict(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.ToText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(language, salesDistrictName string) {
	data, err := c.callSalesDistrictSrvAPIRequirementText("A_SalesDistrictText", language, salesDistrictName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callSalesDistrictSrvAPIRequirementText(api, language, salesDistrictName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_SALESDISTRICT_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, language, salesDistrictName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithSalesDistrict(params map[string]string, salesDistrict string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SalesDistrict eq '%s'", salesDistrict)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, language, salesDistrictName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', SalesDistrictName)", language, salesDistrictName)
	return params
}
