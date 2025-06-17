# nena-manipu-latina
Nena Manipulatina: High Concurrency Ad Server

## Overview
This outlines the initial design and goals for building a lightweight, modular ad server capable fo conducting real-time auctions for ad impressions. This part will serve as a blueprint before implementation. Requirements will include: high concurrency, scalability and demand-side integrations. 

## Goals
* Serve display ads via real-time bidding (RTB) auctions
* Modular architecture to allow for plug-in bidding strategies
* Handle requests with low latency (sub-100 ms target)
* Future-proofed for high-concurrency traffic
* Minimal external dependencies initially
* Easily testable and observable


