#!/usr/bin/env python3
"""
Mintlify documentation scraper for VD
Uses requests and beautifulsoup for robust HTML parsing
"""

import json
import os
import sys
import time
from pathlib import Path
from typing import Dict, List, Optional
from urllib.parse import urljoin, urlparse

try:
    import requests
    from bs4 import BeautifulSoup
    import html2text
except ImportError:
    print("Installing required packages...")
    import subprocess
    subprocess.check_call([sys.executable, "-m", "pip", "install", "requests", "beautifulsoup4", "html2text"])
    import requests
    from bs4 import BeautifulSoup
    import html2text

class MintlifyScraper:
    def __init__(self, base_url: str, output_dir: str):
        self.base_url = base_url.rstrip('/')
        self.output_dir = Path(output_dir)
        self.session = requests.Session()
        self.session.headers.update({
            'User-Agent': 'Mozilla/5.0 (compatible; VD-Scraper/1.0)'
        })
        self.h2t = html2text.HTML2Text()
        self.h2t.ignore_links = False
        self.h2t.ignore_images = False
        self.scraped_urls = set()
        
    def scrape(self) -> Dict:
        """Main scraping function"""
        print(f"🕷️  Starting scrape of {self.base_url}")
        
        # Create output directory
        self.output_dir.mkdir(parents=True, exist_ok=True)
        
        # Try to find the sitemap or navigation
        pages = self.discover_pages()
        
        if not pages:
            print("⚠️  No sitemap found, trying to crawl from homepage...")
            pages = self.crawl_from_homepage()
        
        print(f"📄 Found {len(pages)} pages to scrape")
        
        # Metadata structure
        metadata = {
            "title": self.get_site_title(),
            "description": f"Documentation scraped from {self.base_url}",
            "root_url": self.base_url,
            "scrape_date": time.strftime("%Y-%m-%dT%H:%M:%SZ"),
            "version": "1.0.0",
            "files": [],
            "metadata": {
                "doc_type": "mintlify",
                "status": "scraped"
            }
        }
        
        # Scrape each page
        for i, page_url in enumerate(pages, 1):
            print(f"[{i}/{len(pages)}] Scraping {page_url}")
            try:
                file_info = self.scrape_page(page_url)
                if file_info:
                    metadata["files"].append(file_info)
            except Exception as e:
                print(f"  ⚠️  Error: {e}")
            
            # Rate limiting
            time.sleep(0.5)
        
        # Save metadata
        with open(self.output_dir / "vd.json", 'w') as f:
            json.dump(metadata, f, indent=2)
        
        print(f"✅ Successfully scraped {len(metadata['files'])} pages!")
        return metadata
    
    def discover_pages(self) -> List[str]:
        """Try to discover all pages through sitemap or API"""
        pages = []
        
        # Try sitemap.xml
        sitemap_urls = [
            f"{self.base_url}/sitemap.xml",
            f"{self.base_url}/sitemap_index.xml",
        ]
        
        for url in sitemap_urls:
            try:
                resp = self.session.get(url, timeout=10)
                if resp.status_code == 200:
                    soup = BeautifulSoup(resp.content, 'xml')
                    for loc in soup.find_all('loc'):
                        page_url = loc.text
                        if self.is_doc_page(page_url):
                            pages.append(page_url)
            except:
                continue
        
        return pages
    
    def crawl_from_homepage(self) -> List[str]:
        """Crawl documentation pages starting from homepage"""
        to_visit = [self.base_url]
        visited = set()
        pages = []
        
        while to_visit and len(pages) < 500:  # Limit to 500 pages
            url = to_visit.pop(0)
            if url in visited:
                continue
            
            visited.add(url)
            
            try:
                resp = self.session.get(url, timeout=10)
                if resp.status_code != 200:
                    continue
                
                soup = BeautifulSoup(resp.text, 'html.parser')
                
                # This looks like a documentation page
                if self.is_doc_page(url):
                    pages.append(url)
                
                # Find all links
                for link in soup.find_all('a', href=True):
                    href = link['href']
                    full_url = urljoin(url, href)
                    
                    # Only follow internal links
                    if urlparse(full_url).netloc == urlparse(self.base_url).netloc:
                        if full_url not in visited and full_url not in to_visit:
                            to_visit.append(full_url)
                            
            except Exception as e:
                print(f"  Error crawling {url}: {e}")
                continue
        
        return pages
    
    def is_doc_page(self, url: str) -> bool:
        """Check if URL is likely a documentation page"""
        # Skip non-documentation pages
        skip_patterns = [
            '/api/', '/auth/', '/login', '/signup', '/pricing',
            '.json', '.xml', '.css', '.js', '.png', '.jpg', '.svg',
            '/blog/', '/careers/', '/about/', '/contact/'
        ]
        
        url_lower = url.lower()
        return not any(pattern in url_lower for pattern in skip_patterns)
    
    def scrape_page(self, url: str) -> Optional[Dict]:
        """Scrape a single page"""
        if url in self.scraped_urls:
            return None
        
        self.scraped_urls.add(url)
        
        try:
            resp = self.session.get(url, timeout=10)
            if resp.status_code != 200:
                return None
            
            soup = BeautifulSoup(resp.text, 'html.parser')
            
            # Extract title
            title = self.extract_title(soup, url)
            
            # Extract main content
            content = self.extract_content(soup)
            
            if not content:
                return None
            
            # Convert to markdown
            markdown = self.h2t.handle(str(content))
            
            # Clean up the markdown
            markdown = self.clean_markdown(markdown, title)
            
            # Create file path
            file_path = self.url_to_filepath(url)
            full_path = self.output_dir / file_path
            
            # Create directory if needed
            full_path.parent.mkdir(parents=True, exist_ok=True)
            
            # Save the file
            with open(full_path, 'w', encoding='utf-8') as f:
                f.write(markdown)
            
            return {
                "path": file_path,
                "url": url,
                "title": title,
                "scraped_at": time.strftime("%Y-%m-%dT%H:%M:%SZ")
            }
            
        except Exception as e:
            print(f"  Error scraping {url}: {e}")
            return None
    
    def extract_title(self, soup: BeautifulSoup, url: str) -> str:
        """Extract page title"""
        # Try different title sources
        title = None
        
        # Try h1
        h1 = soup.find('h1')
        if h1:
            title = h1.get_text(strip=True)
        
        # Try title tag
        if not title:
            title_tag = soup.find('title')
            if title_tag:
                title = title_tag.get_text(strip=True)
                # Remove common suffixes
                for suffix in [' | Documentation', ' - Docs', ' | Docs']:
                    if title.endswith(suffix):
                        title = title[:-len(suffix)]
        
        # Fallback to URL
        if not title:
            path = urlparse(url).path.strip('/')
            title = path.split('/')[-1].replace('-', ' ').title()
        
        return title
    
    def extract_content(self, soup: BeautifulSoup) -> Optional:
        """Extract main content from page"""
        # Common content selectors for documentation sites
        selectors = [
            'main',
            'article',
            '[role="main"]',
            '.documentation-content',
            '.markdown-body',
            '.content',
            '#content',
            '.docs-content',
            '.doc-content',
        ]
        
        for selector in selectors:
            content = soup.select_one(selector)
            if content:
                # Remove navigation, sidebars, etc.
                for elem in content.select('nav, aside, .sidebar, .navigation, .toc'):
                    elem.decompose()
                return content
        
        # Fallback: try to find the largest text block
        return soup.find('body')
    
    def clean_markdown(self, markdown: str, title: str) -> str:
        """Clean up the converted markdown"""
        lines = markdown.split('\n')
        cleaned = []
        
        # Add title if not present
        if not any(line.startswith('# ') for line in lines[:5]):
            cleaned.append(f"# {title}")
            cleaned.append("")
        
        # Clean up lines
        for line in lines:
            # Skip empty lines at the start
            if not cleaned and not line.strip():
                continue
            
            # Remove excessive whitespace
            line = line.rstrip()
            
            cleaned.append(line)
        
        # Remove excessive blank lines
        result = []
        prev_blank = False
        for line in cleaned:
            if not line:
                if not prev_blank:
                    result.append(line)
                prev_blank = True
            else:
                result.append(line)
                prev_blank = False
        
        return '\n'.join(result)
    
    def url_to_filepath(self, url: str) -> str:
        """Convert URL to local file path"""
        parsed = urlparse(url)
        path = parsed.path.strip('/')
        
        if not path:
            path = 'index'
        
        # Ensure .md extension
        if not path.endswith('.md'):
            path += '.md'
        
        return path
    
    def get_site_title(self) -> str:
        """Get the site title"""
        try:
            resp = self.session.get(self.base_url, timeout=10)
            soup = BeautifulSoup(resp.text, 'html.parser')
            
            # Try to find site title
            title = soup.find('title')
            if title:
                text = title.get_text(strip=True)
                # Clean up common patterns
                for suffix in [' | Documentation', ' - Documentation', ' Docs']:
                    if text.endswith(suffix):
                        return text[:-len(suffix)]
                return text
        except:
            pass
        
        # Fallback
        return urlparse(self.base_url).netloc


def main():
    if len(sys.argv) != 3:
        print("Usage: python scrape_mintlify.py <url> <output_dir>")
        sys.exit(1)
    
    url = sys.argv[1]
    output_dir = sys.argv[2]
    
    scraper = MintlifyScraper(url, output_dir)
    scraper.scrape()


if __name__ == "__main__":
    main()