import scrapy
import json
import time
from scrapy.crawler import CrawlerProcess

class WikipediaSpider(scrapy.Spider):
    name = "wikipedia"
    start_urls = [
        "https://en.wikipedia.org/wiki/Robotics",
        "https://en.wikipedia.org/wiki/Robot",
        "https://en.wikipedia.org/wiki/Reinforcement_learning",
        "https://en.wikipedia.org/wiki/Robot_Operating_System",
        "https://en.wikipedia.org/wiki/Intelligent_agent",
        "https://en.wikipedia.org/wiki/Software_agent",
        "https://en.wikipedia.org/wiki/Robotic_process_automation",
        "https://en.wikipedia.org/wiki/Chatbot",
        "https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
        "https://en.wikipedia.org/wiki/Android_(robot)",
    ]

    def parse(self, response):
        text = ''.join(response.css('div.mw-parser-output *::text').getall())
        text = text.strip()
        if text:
            data = {
                'url': response.url,
                'text': text,
            }
            print(json.dumps(data))

if __name__ == "__main__":
    start_time = time.time()
    
    process = CrawlerProcess(settings={
        'LOG_LEVEL': 'ERROR',  # Set to 'DEBUG' to view log details
    })
    process.crawl(WikipediaSpider)
    process.start()  # the script will block here until the crawling is finished
    
    elapsed_time = time.time() - start_time
    print("Crawling took %s seconds" % elapsed_time)
