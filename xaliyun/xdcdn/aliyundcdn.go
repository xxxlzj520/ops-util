package xdcdn

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dcdn20180115 "github.com/alibabacloud-go/dcdn-20180115/client"
	"github.com/alibabacloud-go/tea/tea"
)

func authDCdnPoint(accessKeyId, accessKeySecret, endpoint string) (_result *dcdn20180115.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(accessKeyId),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String(endpoint)
	_result, _err = dcdn20180115.NewClient(config)
	return _result, _err
}
func RefreshDCdnObject(accessKeyId, accessKeySecret, endpoint, objectUrl, objectType string) error {
	client, err := authDCdnPoint(accessKeyId, accessKeySecret, endpoint)
	if err != nil {
		return err
	}
	refreshDcdnObjectCachesRequest := &dcdn20180115.RefreshDcdnObjectCachesRequest{
		ObjectType: tea.String(objectType),
		ObjectPath: tea.String(objectUrl),
	}
	_, err1 := client.RefreshDcdnObjectCaches(refreshDcdnObjectCachesRequest)
	if err1 != nil {
		return err1
	}
	return nil
}
func UpdateDCdnDomain(accessKeyId, accessKeySecret, endpoint, objectUrl, sources string) error {
	client, err := authDCdnPoint(accessKeyId, accessKeySecret, endpoint)
	if err != nil {
		return err
	}
	updateDcdnDomainRequest := &dcdn20180115.UpdateDcdnDomainRequest{
		DomainName: tea.String(objectUrl),
		Sources:    tea.String(sources),
	}
	_, err1 := client.UpdateDcdnDomain(updateDcdnDomainRequest)
	if err1 != nil {
		return err1
	}
	return nil
}
func DescribeDcdnDomainDetail(accessKeyId, accessKeySecret, endpoint, objectUrl string) (_result *dcdn20180115.DescribeDcdnDomainDetailResponse, _err error) {
	client, err := authDCdnPoint(accessKeyId, accessKeySecret, endpoint)
	if err != nil {
		return nil, err
	}
	describeDcdnDomainDetailRequest := &dcdn20180115.DescribeDcdnDomainDetailRequest{
		DomainName: tea.String(objectUrl),
	}
	response, _err := client.DescribeDcdnDomainDetail(describeDcdnDomainDetailRequest)
	if _err != nil {
		return nil, _err
	}
	//fmt.Println(response.Body.DomainDetail.Sources.Source)
	return response, nil
}
