
# Author: litianyu
# Description: 通过 Python 基于 bs4 的网络爬虫。
# How to use it: 需要修改 headers 中的内容，可以从各自浏览器中输入 about://version 来进行查看 `User-Agent` 内容

from os import write
import urllib.request
from bs4 import BeautifulSoup
from tenacity import retry
import json

BaseUrl = 'https://theme.npm.edu.tw/opendata/DigitImageSets.aspx'


@retry
def handle_request(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) \
        AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36'
    }
    request = urllib.request.Request(url=url, headers=headers)
    return request


# @retry
def wirte_to_file(url):
    request = handle_request(url)
    # print(url)

    # 发送请求等待相应
    content = urllib.request.urlopen(
        request, timeout=10).read().decode('utf-8')
    # 解析内容
    soup = BeautifulSoup(content, 'lxml')
    # 获取对应的部分
    ul = soup.select('.project-detail > ul > li')

    map = {}

    fileName = ul[0].text[7:] + '.json'

    for li in ul:
        map[li.text[:li.text.index('：')]
            ] = li.text[li.text.index('：') + 1:]

    # dict 转化为 json
    j = json.dumps(map, ensure_ascii=False)

    # 写入文件
    file = open(fileName, "w")
    file.write(j)
    file.close()


@retry
def get_link(baseUrl):
    request = handle_request(baseUrl)
    # 发送请求等待相应
    content = urllib.request.urlopen(
        request, timeout=10).read().decode('utf-8')
    # 解析内容
    soup = BeautifulSoup(content, 'lxml')
    # 获取对应的部分
    ul = soup.select('.painting-list > li > a')
    for li in ul:
        childUrl = BaseUrl + li['href'][19:]
        wirte_to_file(childUrl)


# last page url is 2331
if __name__ == '__main__':
    for i in range(93, 500):
        # print(baseUrl + '?pageNo=' + str(i))
        print(BaseUrl + '?pageNo=' + str(i))
        get_link(BaseUrl + '?pageNo=' + str(i))
