from urllib2 import Request, urlopen
from urllib import urlencode, quote_plus

url = 'http://openapi.data.go.kr/openapi/service/rest/Covid19/getCovid19InfStateJson'

serviceKey = 'l99dHcvNzvph1PhVP0%252BJbsSo4Ron616lka2GaUxUPVCUWZG3YMaxo9PRD6nO8n2fDOOlAOcryhpHokGZR52iig%253D%253D'
serviceKey = "l99dHcvNzvph1PhVP0+JbsSo4Ron616lka2GaUxUPVCUWZG3YMaxo9PRD6nO8n2fDOOlAOcryhpHokGZR52iig=="

queryParams = '?' + urlencode({ quote_plus('ServiceKey') : serviceKey , quote_plus('pageNo') : '1', quote_plus('numOfRows') : '10', quote_plus('startCreateDt') : '20200310', quote_plus('endCreateDt') : '20200315' })

request = Request(url + queryParams)
request.get_method = lambda: 'GET'
response_body = urlopen(request).read()
print(response_body)
