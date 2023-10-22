# WebScrapingProject

Web Scraping: Go vs Python
This project aims to develop and compare web scraping tools using Go and Python. The primary focus is on scraping Wikipedia pages related to intelligent systems and robotics. The results are intended to help the technology firm in creating its own online knowledge base.

Go Web Scraper
Overview
The Go web scraper is built using the Colly framework, which is a popular choice for web scraping in Go due to its efficiency and ease of use. The program is designed to concurrently scrape a predefined list of Wikipedia URLs, extract text content from specific sections of the web pages, and output the results in a JSON lines format.

Implementation
The Go scraper is implemented in a single Go file. It initializes a Colly collector, sets up error handling and scraping logic, and then iterates over a list of target URLs to start the crawling process. The program takes advantage of Go's concurrency model by initiating multiple goroutines to handle the web requests concurrently.

The text content is extracted from the <div class="mw-parser-output"> element of each Wikipedia page, as this div typically contains the main content of the article.

The results are printed to the standard output in a JSON lines format, where each line is a JSON object containing the URL of the scraped page and the extracted text content.

Performance
The execution time of the Go web scraper was measured using the time package in Go. The program took approximately 558.009292ms to complete the scraping of all the target URLs.

Python Web Scraper
Overview
The Python web scraper is built using Scrapy, a widely-used web crawling and scraping framework in Python. Similar to the Go scraper, this program is designed to scrape a list of Wikipedia URLs, extract text content, and output the results in a JSON lines format.

Implementation
The Python scraper is implemented as a Scrapy spider in a single Python script. The spider defines the target URLs, the logic for sending web requests, and the method for parsing the responses to extract text content.

The text content is extracted using CSS selectors to target the <div class="mw-parser-output"> element and its child text nodes. The results are printed to the standard output in a JSON lines format.

Performance
The execution time of the Python web scraper was measured using the time module in Python. The program took approximately 0.931437 seconds to complete the scraping of all the target URLs.

Comparison and Conclusion
When comparing the execution times of the two programs, the Go web scraper proved to be faster, completing the task in 558.009292ms compared to the Python scraper's 0.931437 seconds (or 931.437ms). This performance difference aligns with the general expectation that Go, with its ability to efficiently handle concurrent processes, would outperform Python in web scraping tasks involving multiple network requests.

Both programs successfully scraped the target Wikipedia pages and extracted the required text content, demonstrating that Go can be a viable alternative to Python for web scraping tasks, especially when performance is a critical concern.
