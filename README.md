# nena-manipu-latina
Nena Manipulatina: High Concurrency Ad Server

## Table of Contents
[Overview](#overview)  
[Goals](#goals)  
[Phase 1 Scope](#phase-1-scope)  
[Components](#components)  
[Technical Requirements](#technical-requirements)  
[Request Flow](#basic-request-auction-flow)  
[Future Roadmap](#future-roadmap)  
[Open Questions](#open-questions)    
[Appendix](#appendix)  

## Overview
This outlines the initial design and goals for building a lightweight, modular ad server capable fo conducting real-time auctions for ad impressions. This part will serve as a blueprint before implementation. Requirements will include: high concurrency, scalability and demand-side integrations. 

## Goals
* Serve display ads via real-time bidding (RTB) auctions or simplified waterfall
* Modular architecture to allow for plug-in bidding strategies
* Handle requests with low latency (sub-100 ms target)
* Be horizontally scalable
* Allow integration with multiple DSPs or internal bidders
* Future-proofed for high-concurrency traffic
* Minimal external dependencies initially
* Easily testable and observable
* Collect impression click, and event data for analytics

## Phase 1 Scope
The initial version will support:
* A single ad format (e.g., 300x250 display banner)
* Direct connections to a small set of mock bidders
* A basic auction mechanism (e.g., highest bid wins)
* Basic logging and request tracing

## Components
### Request Handler
* Entry point for HTTP requests to the ad server
* Accepts ad requests from clients (e.g., Javascript tags, SDKs)
* Parses ad request parameters (site ID, user ID, placement ID, etc)
* Enforces any basic validation and filtering
### Auction Engine
* Coordinates the bidding logic
* Responsible for managing the auction lifecycle
* Receives valid bid requests from the handler
* Calls out to mock bidders
* Evaluates bids and determines the winner
* Supports:
    * First-price or second-price auctions
    * Static CPM floors
    * Timeout handling for slow bidders
### Bidder Interface (Mock)
* Simulated endpoints that represent demand partners
* Return bid responses based on simple logic or randomness
* Eventually replaced or extended with real DSPs and/or SSPs
* Timeout-sensitive response with bid metadata
### Response Builder
* Constructs a final ad response to the client (HTML, VAST, etc)
* Includes ad markup, tracking URLs and bid metadata
### Metrics
* Tracks
    * Impressions
    * Clicks
    * Bid win/loss
    * Errors/timeouts
### Logging and Monitoring
* Structured logging for auction events
* Basic metrics (e.g., request count, avg response time, win rate)
* Hooks for future integration with observability tools like Prometheus or OpenTelemetry

## Technical Requirements
* Language: Go for performance and native concurrency support
* Data stores: TDB, Redis caching, PgSQL or NoSQL for event logs
* Deployment: Docker for containerization, Kubernetes for orchestration
* Networking: HTTP-based APIs, eventual gRPC support between internal modules
* Concurrency Model: Will leverage goroutines, rate limiters and sync primitives. 

## Basic Request-Auction Flow
1. Client makes a request to `/ad` endpoint with parameters
2. Request handler validates and constructs an internal bid request object, enriches context (e.g., geo)
3. Auction engine calls out to all bidders and mock bidders concurrently within a timeout.
4. Bids are collected, evaluated and the winner is selected
5. Response is built and returned to the client
6. Impression and bid data are logged asynchronously. Logs are written and metrics are updated.

## Future Roadmap
* Support for multiple formats (video, native)
* Support user and context targeting and segmentation (geo, time, device, etc)
* Header bidding support
* Integration with external DSPs and ad exchanges
* Real-time logging dashboard
* Security and abuse mitigation 
    * IP rate limiting and throttling
    * Signature validation on incoming ad requests
    * Filtering invalid traffic (IVT) and bots

## Open Questions
* What is the latency budget we want to guarantee at scale?
* Will we support cookie-based or server-side user IDs?
* Do we want to simulate auctions with floors and timeouts in Phase 1?
* How will inventory be defined and matched (placement ID vs site ID)?
* What types of creatives are we serving first (banner, native, video)?
* What is our initial traffic expectation?
* Do we need frequency capping or user session management?

## Appendix
* Placeholder bidder response JSON format
* Sample HTTP request and response structures