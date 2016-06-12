Newshound: The Breaking News Email Aggregator
=========

Newshound is a tool to analyze, visualize and share breaking news email alerts. This project was mirrored from [jprobinson's newshound](https://github.com/jprobinson/newshound).

This repository contains a [service to pull and parse breaking news alerts from an email inbox](https://github.com/news-ai/newshound/tree/master/fetch) and a [fast noun-phrase extracting 'microservice'](https://github.com/news-ai/newshound/tree/master/lib/np_extractor) to extract important phrases and help [detect any News Events](https://github.com/news-ai/newshound/tree/master/common.go#L124) that may have occurred. That News Event data is then used to [generate historic reports for each news source.](https://github.com/news-ai/newshound/blob/master/fetch/mapreduce.go)

To emit alert notifications to Slack, Twitter or WebSocket connections, [fetchd](https://github.com/news-ai/newshound/tree/master/fetch/fetchd) can pass information to [barkd](https://github.com/news-ai/newshound/tree/master/bark/barkd) via [NSQ.](http://nsq.io/)

There is also a [web server](https://github.com/news-ai/newshound/tree/master/web/webserver) that can host a [UI](https://github.com/news-ai/newshound/tree/master/web/frontend) and an [API](https://github.com/news-ai/newshound/tree/master/web/webserver/api) for displaying and sharing Newshound information.
