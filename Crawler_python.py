import requests
import re
from bs4 import BeautifulSoup

headers = {
    'User-Agent' : 'Mozilla/s.o (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36' 
}

# r = requests.get("https://cn.bing.com/",timeout = 1)
# r = requests.get("http://stockpage.10jqka.com.cn/000001/",headers = headers,timeout = 1)



# print(r.text)
# soup.findAll('div')
# t = soup.find_all('div',class_= 'sub_cont_3') :


def getStockInfo(url):
    r = requests.get(url,headers = headers,timeout = 1)
    print(r.status_code) 
    infoDict = {}
    soup = BeautifulSoup(r.text,'html.parser')
    name = soup.find('strong').text
    print(name)
    # infoDict.update({'股票名称': name.text.split()[0]})
    infoDict.update({'股票名称': name})
    stockInfo = soup.find('div',attrs={'class':'sub_cont_3'})
    keyList = stockInfo.find_all('dt')
    valueList = stockInfo.find_all('dd')
    for i in range(len(keyList)) :
        key = keyList[i].text
        value = valueList[i].text
        infoDict[key] = value
    with open("douban.txt",'w',encoding='utf-8') as f:
        f.write(str(infoDict))
        f.write("\n")

def getData(js) :
    r = requests.get(js,headers = headers,timeout = 1)
    print(r.status_code) 
    pat = "items:\[(.*?)\]"
    data = re.compile(pat,re.S).findall(r.text)
    print(data)



def main() :
    url = "http://stockpage.10jqka.com.cn/000001/"
    # js = "http://d.10jqka.com.cn/v2/realhead/hs_000001/last.js"
    # getStockInfo(url)
    # getData(js)
    
main()