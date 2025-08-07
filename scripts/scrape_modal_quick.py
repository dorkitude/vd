#!/usr/bin/env python3
"""
Quick Modal documentation scraper - focused approach
"""

import json
import sys
import time
from pathlib import Path
from urllib.parse import urlparse

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

def scrape_modal_quick(output_dir: str):
    """Quick scrape of essential Modal docs"""
    output_path = Path(output_dir)
    output_path.mkdir(parents=True, exist_ok=True)
    
    session = requests.Session()
    session.headers.update({'User-Agent': 'Mozilla/5.0'})
    
    h2t = html2text.HTML2Text()
    h2t.ignore_links = False
    
    # Essential pages to scrape
    essential_pages = [
        "https://modal.com/docs",
        "https://modal.com/docs/guide",
        "https://modal.com/docs/guide/environments",
        "https://modal.com/docs/guide/cloud-storage",
        "https://modal.com/docs/guide/functions",
        "https://modal.com/docs/guide/images",
        "https://modal.com/docs/guide/webhooks",
        "https://modal.com/docs/examples",
        "https://modal.com/docs/reference",
        "https://modal.com/docs/reference/modal.App",
        "https://modal.com/docs/reference/modal.Function",
        "https://modal.com/docs/reference/modal.Image",
    ]
    
    metadata = {
        "title": "Modal",
        "description": "Documentation scraped from https://modal.com/docs",
        "root_url": "https://modal.com/docs",
        "scrape_date": time.strftime("%Y-%m-%dT%H:%M:%SZ"),
        "version": "1.0.0",
        "files": [],
        "metadata": {
            "doc_type": "modal",
            "status": "scraped"
        }
    }
    
    print(f"🕷️  Scraping {len(essential_pages)} essential Modal docs pages...")
    
    for i, url in enumerate(essential_pages, 1):
        print(f"[{i}/{len(essential_pages)}] {url}")
        try:
            resp = session.get(url, timeout=10)
            if resp.status_code != 200:
                print(f"  ⚠️  Got status {resp.status_code}")
                continue
            
            soup = BeautifulSoup(resp.text, 'html.parser')
            
            # Find title
            title = "Modal Docs"
            h1 = soup.find('h1')
            if h1:
                title = h1.get_text(strip=True)
            elif soup.title:
                title = soup.title.get_text(strip=True).replace(' | Modal', '')
            
            # Find main content
            content = None
            for selector in ['main', 'article', '[class*="prose"]', '[class*="content"]']:
                content = soup.select_one(selector)
                if content:
                    break
            
            if not content:
                print(f"  ⚠️  No content found")
                continue
            
            # Clean content
            for elem in content.select('nav, aside, header, footer'):
                elem.decompose()
            
            # Convert to markdown
            markdown = h2t.handle(str(content))
            
            # Create file path
            path = urlparse(url).path.strip('/').replace('docs/', '')
            if not path:
                path = 'index'
            if not path.endswith('.md'):
                path += '.md'
            
            # Save file
            file_path = output_path / path
            file_path.parent.mkdir(parents=True, exist_ok=True)
            
            # Add title if missing
            if not markdown.startswith('# '):
                markdown = f"# {title}\n\n{markdown}"
            
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(markdown)
            
            metadata["files"].append({
                "path": path,
                "url": url,
                "title": title,
                "scraped_at": time.strftime("%Y-%m-%dT%H:%M:%SZ")
            })
            
            print(f"  ✅ Saved to {path}")
            time.sleep(0.2)  # Rate limit
            
        except Exception as e:
            print(f"  ❌ Error: {e}")
    
    # Save metadata
    with open(output_path / "vd.json", 'w') as f:
        json.dump(metadata, f, indent=2)
    
    print(f"\n✅ Successfully scraped {len(metadata['files'])} pages!")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python scrape_modal_quick.py <output_dir>")
        sys.exit(1)
    
    scrape_modal_quick(sys.argv[1])