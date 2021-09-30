# Author: litianyu
# Description: 通过 Python 基于 bs4 的网络爬虫。
# How to use it: 需要修改 headers 中的内容，可以从各自浏览器中输入 about://version 来进行查看 `User-Agent` 内容

import requests
from bs4 import BeautifulSoup
from tenacity import retry, stop_after_attempt
import concurrent
import json

BaseUrl = 'https://theme.npm.edu.tw/opendata/DigitImageSets.aspx'
Headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) \
    AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36'
}
Proxies = {
    "http": "http://localhost:8889",
    "https": "http://localhost:8889"
}


# session = requests.Session()


@retry(stop=stop_after_attempt(3))
def get_html_content(url: str, headers: dict) -> str:
    session = requests.session()
    reponse = session.get(url=url, headers=headers, proxies=Proxies)

    session.close()
    return reponse.text.encode("utf-8")


def wirte_to_file(url: str, referer: str) -> None:
    dataheaders = Headers
    dataheaders.update({"Referer": referer})
    # print(url)

    content = get_html_content(url=url, headers=dataheaders)
    # 解析内容
    soup = BeautifulSoup(content, 'lxml')
    # 获取对应的部分
    ul = soup.select('.project-detail > ul > li')

    map = {}

    document_number = ul[0].text[7:]

    fileName = document_number + '.json'

    for li in ul:
        map[li.text[:li.text.index('：')]
            ] = li.text[li.text.index('：') + 1:]

    # dict 转化为 json
    j = json.dumps(map, ensure_ascii=False)

    # 写入文件
    print("文物圖檔編號: {}".format(document_number), end=",")
    destfile = open(fileName, "w")
    destfile.write(j)
    destfile.close()
    print("已保存")


def get_link(baseurl: str) -> None:
    content = get_html_content(url=baseurl, headers=Headers)
    # 解析内容
    print("Analyzing html content of {}".format(baseurl))
    soup = BeautifulSoup(content, 'lxml')
    # 获取对应的部分
    ul = soup.select('.painting-list > li > a')
    for li in ul:
        childUrl = BaseUrl + li['href'][19:]
        wirte_to_file(childUrl, baseurl)


# last page url is 2331
if __name__ == '__main__':
    try:
        with concurrent.futures.ThreadPoolExecutor(5) as executor:
            for pageno in range(501, 2332):
                # print(baseurl + '?pageNo=' + str(i))
                print(BaseUrl + '?pageNo=' + str(pageno))
                args = BaseUrl + '?pageNo=' + str(pageno)
                executor.submit(get_link, args)
    #        session.close()
    except KeyboardInterrupt:
        #        session.close()
        print("Aborted")
        exit(0)

