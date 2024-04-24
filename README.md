# Notion CLI Tool

This CLI tool is a convenient command-line interface to interact with a Notion database. It allows users to add entries to a Notion database directly from the terminal, simplifying the process of content management within Notion pages.

## Features

- **Create Entries**: Easily create new entries in a specified Notion database using title and content input from the command line.

## Prerequisites

Before you start using this CLI tool, make sure you have:

- Go installed on your machine.
- An active Notion account and access to the Notion API.
- A Notion integration created to obtain the `NOTION_TOKEN`.
- The database ID of the Notion database you want to interact with (`NOTION_DATABASE_ID`).

## Installation

1. Clone this repository:
   ```bash
   git clone https://your-repository-url
   cd notion_cli

2. Set your environment variables:
    You will need to set these in your .bashrc, .zshrc, or .profile depending on which shell you use, so they are automatically set whenever you open a terminal.

    echo 'export NOTION_TOKEN="your_notion_api_token"' >> ~/.bashrc
    echo 'export NOTION_DATABASE_ID="your_notion_database_id"' >> ~/.bashrc
    source ~/.bashrc



3. Build the Binary
    go build -o ncli (or your preferred command for the program)

4. Verify the Installation
    type ncli in your terminal, a line should appear prompting a title for your new database entry

If you encounter any issues with the tool after installation, make sure that your NOTION_TOKEN and NOTION_DATABASE_ID are correctly configured in your environment variables.
