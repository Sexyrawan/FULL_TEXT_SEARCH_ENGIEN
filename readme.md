# Simple Text Search Engine

## Overview

This project is a lightweight, conceptual text search engine that leverages an **inverted index** to perform lightning-fast text queries. Instead of linearly scanning every document for a matching word, the engine looks up the word in the index to instantly retrieve all associated document IDs.

## Core Features

- **Text Preprocessing:** Cleans raw text before indexing by removing punctuation and converting all characters to lowercase.
- **Tokenization:** Splits the cleaned text into individual, searchable words (tokens).
- **Inverted Indexing:** Maps individual tokens to the specific documents where they appear, enabling rapid data retrieval.

## How It Works

### 1. Documents

The engine starts with a collection of raw data entries. For example, consider the following three documents:

- **Document 1:** `"a donut on a glass plate"`
- **Document 2:** `"only the donut"`
- **Document 3:** `"listen to the drum machine"`

### 2. Preprocessing & Tokenization

Before building the index, the engine normalizes the text. It strips away formatting and punctuation, ensuring that a search query for "Donut!" successfully matches the base token `"donut"`.

### 3. The Inverted Index

The core data structure maps words (keys) to their document locations (values).

Based on the example documents above, the resulting inverted index looks like this:

| Word        | Document IDs |
| :---------- | :----------- |
| **"a"**     | `[1]`        |
| **"donut"** | `[1, 2]`     |
| **"on"**    | `[1]`        |
| **"glass"** | `[1]`        |
| **"plate"** | `[1]`        |
| **"only"**  | `[2]`        |
| **"the"**   | `[2, 3]`     |

### 4. FLOW

```
Raw Data (CSV/XML/Wikipedia dump)
↓
Parser
↓
Documents Struct
↓
Indexer
↓
Inverted Index
↓
Search Query
↓
Matching Documents
```

## SUMMARY

**TOCKNIZATION** (text is converted into tolen a number) \n

**FILTERING** -->

1. LOWERCASING: SHall ShaLL --> shall small case
2. DROPPING COMMON WORDS: ex. it, al, we, be, ending verb(word)
3. STEMMING: (snow ball stemm -- liberary) fishing+ fished +fisher = fish

### HI

-- if you want suggest me then definately go , i will grateful for you this act.

```

```
