# kiteapi

The source is taken from project OTAMIA which is private. 

# UNOFFICIAL

This api is unofficial although, it source codes are borrowed from official API. This is **incompatible** with official api of Zerodha. 

# Major Differences 

1. Supports OI flag in historical data.
2. The error is displayed in debug mode and you will not receive error = nil when there is one. - __FIX__
3. NewConnect and NewTicker are the functions instead of two packages with New. - __Improvement__
4. The structure used for Ticker data and Connect data is now common.  - __Improvement__ (Major) 
5. int64 is used everywhere while offical discrepancy is: uint32 in ticker VS and int64 in connect that in official API. __Improvement__ which avoids various conversions.
6. Fixed the minute historical data date is not parsed in known utility when doing unmarshal - Converting the JSON to Struct was making dates as 0000-00-00 __FIX__ (Major)

# Status 

This is used in the production environment and I will improve it wherever needed!


# Why not a fork?

The official API is simply incompatible with this one. The core is changed and it doesn't make sense to keep a fork.
