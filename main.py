
# Author: litianyu
# Description: 通过 Python 基于 bs4 的网络爬虫。

import urllib.request
from bs4 import BeautifulSoup
import json


def handle_request(url):
    headers = {
        'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) \
        AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36'
    }
    request = urllib.request.Request(url=url, headers=headers)
    return request


def wirte_to_file(url):
    request = handle_request(url)

    # 发送请求等待相应
    content = urllib.request.urlopen(request).read().decode('utf-8')
    # 解析内容
    soup = BeautifulSoup(content, 'lxml')
    # 获取对应的部分
    ul = soup.select('.project-detail > ul > li')

    map = {}

    fileName = ul[0].text[7:] + '.json'

    map[ul[0].text[:6]] = ul[0].text[7:]
    map[ul[1].text[:2]] = ul[1].text[3:]
    map[ul[2].text[:2]] = ul[2].text[3:]
    map[ul[3].text[:2]] = ul[3].text[3:]
    map[ul[4].text[:2]] = ul[4].text[3:]
    map[ul[5].text[:3]] = ul[5].text[4:]

    # dict 转化为 json
    j = json.dumps(map, ensure_ascii=False)

    # 写入文件
    file = open(fileName, "w")
    file.write(j)
    file.close()


def get_link(baseUrl):
    request = handle_request(baseUrl)
    # 发送请求等待相应
    content = urllib.request.urlopen(request).read().decode('utf-8')
    # 解析内容
    soup = BeautifulSoup(content, 'lxml')
    # 获取对应的部分
    ul = soup.select('.painting-list > li > a')
    for li in ul:
        childUrl = baseUrl + li['href'][19:]
        wirte_to_file(childUrl)


# last page url is 2331
if __name__ == '__main__':
    baseUrl = 'https://theme.npm.edu.tw/opendata/DigitImageSets.aspx'
    for i in range(1, 5):
        # print(baseUrl + '?pageNo=' + str(i))
        get_link(baseUrl + '?pageNo=' + str(i))
