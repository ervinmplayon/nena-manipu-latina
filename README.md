# nena-manipu-latina
Nena Manipulatina: High Concurrency Ad Server

## Table of Contents
[Overview](#overview)  
[Goals](#goals)  
[Phase 1 Scope](#phase-1-scope)  
[Components](#components)  
[Request Flow](#request-flow)  
[Future Roadmap](#future-roadmap)  
[Open Questions](#open-questions)    
[Appendix](#appendix)  

## Overview
This outlines the initial design and goals for building a lightweight, modular ad server capable fo conducting real-time auctions for ad impressions. This part will serve as a blueprint before implementation. Requirements will include: high concurrency, scalability and demand-side integrations. 

## Goals
* Serve display ads via real-time bidding (RTB) auctions
* Modular architecture to allow for plug-in bidding strategies
* Handle requests with low latency (sub-100 ms target)
* Future-proofed for high-concurrency traffic
* Minimal external dependencies initially
* Easily testable and observable

## Phase 1 Scope
The initial version will support:
* A single ad format (e.g., 300x250 display banner)
* Direct connections to a small set of mock bidders
* A basic auction mechanism (e.g., highest bid wins)
* Basic logging and request tracing

## Components
### Request Handler
* Entry point for HTTP requests to the ad server
* Parses ad request parameters (site ID, user ID, placement ID, etc)
* Enforces any basic validation and filtering
### Auction Engine
* Responsible for managing the auction lifecycle
* Receives valid bid requests from the handler
* Calls out to mock bidders
* Evaluates bids and determines the winner
### Bidder Interface (Mock)
* Simulated endpoints that represent demand partners
* Return bid responses based on simple logic or randomness
* Eventuall replaced or extended with real DSPs and/or SSPs
### Response Builder
* Constructs a final ad response to the client
* Includes ad markup, tracking URLs and bid metadata
### Logging and Monitoring
* Structured logging for auction events
* Basic metrics (e.g., request count, avg response time, win rate)
* Hooks for future integration with observability tools like Prometheus or OpenTelemetry

## Request Flow
1. Client makes a request to `/ad` endpoint with parameters
2. Request handler validates and constructs an internal bid request object
3. Auction engine calls mock bidders concurrently
4. Bids are collected, evaluated and the winner is selected
5. Response is built and returned to the client
6. Logs are written and metrics are updated.

## Future Roadmap
* Support for multiple formats (video, native)
* User targeting & segmentation
* Header bidding support
* Integration with external DSPs and ad exchanges
* Real-time logging dashboard
* Throttling and rate limiting

## Open Questions
* What is the latency budget we want to guarantee at scale?
* Will we support cookie-based or server-side user IDs?
* Do we want to simulate auctions with floors and timeouts in Phase 1?
* How will inventory be defined and matched (placement ID vs site ID)?

## Appendix
* Placeholder bidder response JSON format
* Sample HTTP request and response structures