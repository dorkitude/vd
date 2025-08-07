# Using MongoDB Atlas Vector and GeoJSON Search with Modal Docs

> Note: This is a placeholder for Modal documentation from https://modal.com/docs/examples/mongodb-search
> The full HTML-to-Markdown conversion will be implemented with a proper HTML parser.

## About This Page

This documentation was scraped from the Modal docs website. Modal is a cloud platform for running Python code in the cloud with minimal configuration.

Visit the original documentation at: https://modal.com/docs/examples/mongodb-search

## Content Preview

The application is a hybrid search engine,
like the retrieval engines that power RAG chatbots,
but for satellite images of the state of California.
Images can be searched based on their
geospatial and temporal metadata or based on their semantic content
as captured by a pre-trained embedding model.

At the center of the application is a MongoDB Atlas instance
that stores metadata for a collection of satellite images.

Modal orchestrates the compute around that database:
retrieving data from elsewhere and storing it in the database,
computing vector embeddings for the data in the database,
and serving both a frontend and a client.

The dataflow looks something like this:

This entire application —
from API queries and frontend UI to GPU inference and hybrid search —
is delivered using nothing but Modal and MongoDB Atlas.
Setting it up for yourself requires only credentials on these platforms
and a few commands, detailed below.

