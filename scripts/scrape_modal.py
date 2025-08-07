#!/usr/bin/env python3
"""
Modal-specific documentation scraper
Optimized for Modal's documentation structure
"""

import json
import os
import sys
import time
from pathlib import Path
from typing import Dict, List, Optional, Set
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

class ModalScraper:
    def __init__(self, output_dir: str):
        self.base_url = "https://modal.com/docs"
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
        print(f"🕷️  Starting Modal docs scrape")
        
        # Create output directory
        self.output_dir.mkdir(parents=True, exist_ok=True)
        
        # Get all documentation pages
        pages = self.discover_modal_pages()
        
        print(f"📄 Found {len(pages)} pages to scrape")
        
        # Metadata structure
        metadata = {
            "title": "Modal",
            "description": "Documentation scraped from https://modal.com/docs",
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
            time.sleep(0.3)
        
        # Save metadata
        with open(self.output_dir / "vd.json", 'w') as f:
            json.dump(metadata, f, indent=2)
        
        print(f"✅ Successfully scraped {len(metadata['files'])} pages!")
        return metadata
    
    def discover_modal_pages(self) -> List[str]:
        """Discover Modal documentation pages"""
        pages = set()
        to_visit = [self.base_url]
        visited = set()
        
        # Main documentation sections
        main_sections = [
            "/docs/guide",
            "/docs/examples", 
            "/docs/reference"
        ]
        
        for section in main_sections:
            to_visit.append(f"https://modal.com{section}")
        
        while to_visit and len(pages) < 300:  # Limit to 300 pages
            url = to_visit.pop(0)
            if url in visited:
                continue
            
            visited.add(url)
            
            # Only process modal.com/docs URLs
            if not url.startswith("https://modal.com/docs"):
                continue
            
            try:
                resp = self.session.get(url, timeout=10)
                if resp.status_code != 200:
                    continue
                
                # Add this page
                pages.add(url)
                
                soup = BeautifulSoup(resp.text, 'html.parser')
                
                # Find all internal doc links
                for link in soup.find_all('a', href=True):
                    href = link['href']
                    
                    # Only follow /docs/ links
                    if href.startswith('/docs/'):
                        full_url = f"https://modal.com{href}"
                        # Remove hash fragments
                        full_url = full_url.split('#')[0]
                        
                        if full_url not in visited and full_url not in to_visit:
                            to_visit.append(full_url)
                            
            except Exception as e:
                print(f"  Error discovering {url}: {e}")
                continue
        
        return sorted(list(pages))
    
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
        # Try h1 first
        h1 = soup.find('h1')
        if h1:
            return h1.get_text(strip=True)
        
        # Try title tag
        title_tag = soup.find('title')
        if title_tag:
            title = title_tag.get_text(strip=True)
            # Remove common suffixes
            for suffix in [' | Modal', ' - Modal', ' | Documentation']:
                if title.endswith(suffix):
                    title = title[:-len(suffix)]
            return title
        
        # Fallback to URL
        path = urlparse(url).path.strip('/')
        return path.split('/')[-1].replace('-', ' ').replace('_', ' ').title()
    
    def extract_content(self, soup: BeautifulSoup) -> Optional:
        """Extract main content from Modal docs page"""
        # Modal-specific content selectors
        selectors = [
            'main',
            'article', 
            '[class*="prose"]',
            '[class*="content"]',
            '[class*="markdown"]',
            '.docs-content',
            '#content'
        ]
        
        for selector in selectors:
            content = soup.select_one(selector)
            if content:
                # Remove navigation, sidebars, etc.
                for elem in content.select('nav, aside, [class*="sidebar"], [class*="navigation"], [class*="toc"]'):
                    elem.decompose()
                # Remove header/footer
                for elem in content.select('header, footer'):
                    elem.decompose()
                return content
        
        # Fallback: try to find the main content area
        # Look for divs with significant text content
        for div in soup.find_all('div'):
            text_length = len(div.get_text(strip=True))
            if text_length > 500:  # Significant content
                return div
        
        return None
    
    def clean_markdown(self, markdown: str, title: str) -> str:
        """Clean up the converted markdown"""
        lines = markdown.split('\n')
        cleaned = []
        
        # Add title if not present
        if not any(line.startswith('# ') for line in lines[:5]):
            cleaned.append(f"# {title}")
            cleaned.append("")
        
        # Clean up lines
        prev_blank = False
        for line in lines:
            # Skip multiple blank lines
            if not line.strip():
                if not prev_blank:
                    cleaned.append("")
                prev_blank = True
            else:
                cleaned.append(line.rstrip())
                prev_blank = False
        
        return '\n'.join(cleaned)
    
    def url_to_filepath(self, url: str) -> str:
        """Convert URL to local file path"""
        parsed = urlparse(url)
        path = parsed.path.strip('/')
        
        # Remove 'docs/' prefix
        if path.startswith('docs/'):
            path = path[5:]
        
        if not path:
            path = 'index'
        
        # Ensure .md extension
        if not path.endswith('.md'):
            path += '.md'
        
        return path


def main():
    if len(sys.argv) != 2:
        print("Usage: python scrape_modal.py <output_dir>")
        sys.exit(1)
    
    output_dir = sys.argv[1]
    
    scraper = ModalScraper(output_dir)
    scraper.scrape()


if __name__ == "__main__":
    main()