#!/bin/bash
if [ ! -f "./podcast.json" ] ; then
curl "https://radio-t.com/site-api/last/1000?categories=podcast" > ./podcast.json ;
fi
if [ ! -f "./prep.json" ]; then
curl "https://radio-t.com/site-api/last/1000?categories=prep" > ./prep.json
fi
cat ./IntenseDebate_clean.xml | ./rtxml -podcast ./podcast.json -prep ./prep.json > rt.json



